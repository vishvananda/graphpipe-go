/*
** Copyright © 2018, Oracle and/or its affiliates. All rights reserved.
** Licensed under the Universal Permissive License v 1.0 as shown at http://oss.oracle.com/licenses/upl.
 */

package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	graphpipe "github.com/vishvananda/graphpipe-go"
	graphpipefb "github.com/vishvananda/graphpipe-go/graphpipefb"
)

// #cgo CXXFLAGS: -DONNX_NAMESPACE=onnx_c2 -D_GNU_SOURCE -D_REENTRANT -D CAFFE2_USE_GFLAGS -D CAFFE2_USE_GOOGLE_GLOG -D NDEBUG -I /usr/local/cuda/include -std=c++11 -O2
// #cgo LDFLAGS: -ldl /usr/local/lib/libcaffe2.so -lprotobuf -lglog -lgflags
// #include <stdlib.h>
// #include <c2_api.h>
import "C"

var (
	ver string
	sha string
)

func version() string {
	if ver == "" {
		ver = "dev"
		sha = "unknown"
	}
	return fmt.Sprintf("version %s (built from sha %s)", ver, sha)
}

type options struct {
	model       string
	listen      string
	cacheDir    string
	verbose     bool
	version     bool
	cache       bool
	disableCuda bool
	valueInputs string
	initNet     string
	predictNet  string
	profile     string
	engineCount int
}

func loadFile(uri string) ([]byte, error) {
	logrus.Infof("Loading file ", uri)
	if strings.HasPrefix(uri, "http://") ||
		strings.HasPrefix(uri, "https://") {
		var transport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			Proxy:               http.ProxyFromEnvironment,
			TLSHandshakeTimeout: 5 * time.Second,
		}
		var client = &http.Client{
			Timeout:   time.Second * 60,
			Transport: transport,
		}
		response, err := client.Get(uri)
		if err != nil {
			logrus.Errorf("Failed to get '%s': %v", uri, err)
			return nil, err
		}
		if response.StatusCode != 200 {
			logrus.Fatalf("Failed to load file: %s, status code: %v", uri, response.StatusCode)
		}
		return ioutil.ReadAll(response.Body)
	}
	return ioutil.ReadFile(uri)
}

func main() {
	var cmdExitCode int
	var opts options
	cmd := cobra.Command{
		Use:   "graphpipe-caffe2",
		Short: "graphpipe-caffe2 - serving up delectable caffe2 and onnx models",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if opts.verbose {
				logrus.SetLevel(logrus.DebugLevel)
			} else {
				logrus.SetLevel(logrus.InfoLevel)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if opts.disableCuda {
				logrus.Infof("CUDA is disabled")
			}
			if opts.version {
				fmt.Printf("%s\n", version())
				return
			}
			if len(args) != 0 {
				cmdExitCode = 1
				cmd.Usage()
				return
			}
			if opts.model == "" {
				if opts.initNet == "" && opts.predictNet == "" {
					logrus.Errorf("either --model/env(\"GP_MODEL\") must be set or --init_net/env(GP_INIT_NET) AND --predict_net/env(GP_PREDICT_NET) must be set.")
					cmdExitCode = 1
					return
				}
			}
			if opts.engineCount < 0 || opts.engineCount > 32 {
				logrus.Errorf("--engine_count must be >= 1 && <=32")
				cmdExitCode = 1
				return
			}
			if opts.valueInputs == "" {
				logrus.Errorf("--value_inputs or env(\"GP_VALUE_INPUTS\") must be set")
				cmdExitCode = 1
				return
			}
			if opts.profile != "" {
				f, err := os.Create(opts.profile)
				if err != nil {
					logrus.Fatal(err)
				}
				pprof.StartCPUProfile(f)
				c := make(chan os.Signal, 1)
				signal.Notify(c, os.Interrupt)
				go func() {
					for _ = range c {
						pprof.StopCPUProfile()
						os.Exit(0)
					}
				}()
			}

			logrus.Infof("Starting graphpipe-caffe2 %s", version())

			if err := serve(opts); err != nil {
				logrus.Errorf("Failed to serve: %v", err)
				cmdExitCode = 1
			}
			if opts.profile != "" {
			}
		},
	}
	f := cmd.Flags()
	f.StringVarP(&opts.cacheDir, "cache-dir", "", "~/.graphpipe", "directory for local cache state")
	f.StringVarP(&opts.listen, "listen", "l", "127.0.0.1:9000", "listen string")
	f.StringVarP(&opts.model, "model", "m", "", "ONNX model to load.   Accepts local file or http(s) url.")
	f.StringVarP(&opts.initNet, "init-net", "", "", "init_net file to load. Accepts local file or http(s) url.")
	f.StringVarP(&opts.predictNet, "predict-net", "", "", "predict_net file to load.  Accepts local file or http(s) url.")
	f.StringVarP(&opts.valueInputs, "value-inputs", "", "", "value_inputs.json for the model.  Accepts local file or http(s) url.")
	f.BoolVarP(&opts.cache, "cache", "", false, "enable results caching")
	f.BoolVarP(&opts.disableCuda, "disable-cuda", "", false, "disable Cuda")
	f.StringVarP(&opts.profile, "profile", "", "", "profile and write profiling output to this file")
	f.IntVarP(&opts.engineCount, "engine-count", "", 1, "number of caffe2 graph engines to create")

	f = cmd.PersistentFlags()
	f.BoolVarP(&opts.verbose, "verbose", "v", false, "enable verbose output")
	f.BoolVarP(&opts.version, "version", "V", false, "show version")

	opts.cacheDir = strings.Replace(opts.cacheDir, "~", os.Getenv("HOME"), -1)

	if opts.model == "" {
		opts.model = os.Getenv("GP_MODEL")
	}
	if opts.valueInputs == "" {
		opts.valueInputs = os.Getenv("GP_VALUE_INPUTS")
	}
	if opts.initNet == "" {
		opts.initNet = os.Getenv("GP_INIT_NET")
	}
	if opts.predictNet == "" {
		opts.predictNet = os.Getenv("GP_PREDICT_NET")
	}

	if os.Getenv("GP_CACHE") != "" {
		val := strings.ToLower(os.Getenv("GP_CACHE"))
		if val == "1" || val == "true" {
			logrus.Infof("Caching is enabled")
			opts.cache = true
		}
	}

	if os.Getenv("GP_ENGINE_COUNT") != "" {
		count, err := strconv.Atoi(os.Getenv("GP_ENGINE_COUNT"))
		if err != nil {
			logrus.Errorf("Could not parse GP_ENGINE_COUNT")
			os.Exit(1)
		}
		opts.engineCount = count
	}

	if os.Getenv("MKL_NUM_THREADS") == "" {
		logrus.Infof("Setting MKL_NUM_THREADS=4.  You can override this variable in your environment.")
		os.Setenv("MKL_NUM_THREADS", "4")
	}
	cmd.Execute()
	os.Exit(cmdExitCode)
}

type c2Context struct {
	EngineCount    int
	CEngineCtxs    []*C.c2_engine_ctx
	InputDims      map[string][]int64
	OutputDims     map[string][]int64
	Inputs         []string
	Outputs        []string
	modelHash      []byte
	meta           *graphpipe.NativeMetadataResponse
	engineChannels chan *C.c2_engine_ctx
}

func serve(opts options) error {
	c2c := &c2Context{}

	tryToUseCuda := 1
	if opts.disableCuda {
		tryToUseCuda = 0
	}

	c2c.EngineCount = opts.engineCount
	c2c.engineChannels = make(chan *C.c2_engine_ctx, c2c.EngineCount)
	for i := 0; i < c2c.EngineCount; i++ {
		engine_ctx := C.c2_engine_create(C.int(tryToUseCuda))
		if engine_ctx == nil {
			logrus.Fatalf("Could not create engine\n")
		}
		c2c.CEngineCtxs = append(c2c.CEngineCtxs, engine_ctx)
		c2c.engineChannels <- engine_ctx
	}

	valueInputData := make(map[string]interface{})
	valueInputJson, err := loadFile(opts.valueInputs)
	if err != nil {
		logrus.Fatalf("Could not load value_input: %v", err)
	}

	err = json.Unmarshal([]byte(valueInputJson), &valueInputData)
	if err != nil {
		logrus.Fatalf("Could not unmarshall value_input: %v", err)
	}

	for k, v := range valueInputData {
		if reflect.ValueOf(v).Len() == 2 {
			dims := []int64{}
			dtype := reflect.ValueOf(v).Index(0).Interface().(float64)
			shape := reflect.ValueOf(v).Index(1).Elem()
			for i := 0; i < shape.Len(); i++ {
				d := int64(shape.Index(i).Interface().(float64))
				dims = append(dims, d)
			}
			for _, engine_ctx := range c2c.CEngineCtxs {
				C.c2_engine_register_input(engine_ctx, C.CString(k), (*C.int64_t)(&dims[0]), (C.int(len(dims))), C.int(dtype))
			}
		} else {
			logrus.Fatalf("Invalid value for value_input with key: %s.  Format should be {\"k\": [dtype, [dim1, dim2, ... ]]}, eg {\"k\": [1, [1, 3, 217, 217]]}", k)
		}
	}

	if opts.model != "" {
		modelData, err := loadFile(opts.model)
		if err != nil {
			logrus.Errorln("Could not read file ", opts.model)
			return err
		}

		for _, engine_ctx := range c2c.CEngineCtxs {
			status := C.c2_engine_initialize_onnx(engine_ctx, (*C.char)(unsafe.Pointer((&modelData[0]))), C.size_t(len(modelData)))
			if status != 0 {
				logrus.Fatalf("Could not initialize engine\n")
			}
		}

		h := sha256.New()
		h.Write(modelData)
		c2c.modelHash = h.Sum(nil)
		logrus.Infof("Model hash is '%x'", c2c.modelHash)
	} else {
		initData, err := loadFile(opts.initNet)
		if err != nil {
			logrus.Errorln("Could not read init file ", opts.initNet)
			return err
		}

		predData, err := loadFile(opts.predictNet)
		if err != nil {
			logrus.Errorln("Could not read predict file ", opts.predictNet)
			return err
		}

		for _, engine_ctx := range c2c.CEngineCtxs {
			status := C.c2_engine_initialize_caffe2(engine_ctx, (*C.char)(unsafe.Pointer((&initData[0]))), C.size_t(len(initData)), (*C.char)(unsafe.Pointer((&predData[0]))), C.size_t(len(predData)))
			if status != 0 {
				logrus.Fatalf("Could not initialize engine\n")
			}
		}

		h := sha256.New()
		h.Write(initData)
		h.Write(predData)
		c2c.modelHash = h.Sum(nil)
		logrus.Infof("Model hash is '%x'", c2c.modelHash)
	}

	inputCount := C.c2_engine_get_input_count(c2c.CEngineCtxs[0])
	c2c.Inputs = make([]string, inputCount)
	c2c.InputDims = make(map[string][]int64)
	c2c.OutputDims = make(map[string][]int64)
	meta := &graphpipe.NativeMetadataResponse{}

	for i := 0; i < int(inputCount); i++ {
		cname := C.c2_engine_get_input_name(c2c.CEngineCtxs[0], C.int(i))
		name := C.GoString(cname)
		c2c.Inputs[i] = name
		dims := make([]int64, 32)
		dimCount := C.c2_engine_get_dimensions(c2c.CEngineCtxs[0], cname, (*C.int64_t)(unsafe.Pointer(&dims[0])))
		if dimCount < 0 {
			logrus.Fatalf("Could not find dimensions for: %s", name)
		}
		c2c.InputDims[name] = dims[:dimCount]
		dtype := C.c2_engine_get_dtype(c2c.CEngineCtxs[0], cname)

		logrus.Debugf("Adding input '%s'", name)

		io := graphpipe.NativeIOMetadata{}
		io.Name = name
		io.Shape = dims[:dimCount]
		io.Type = ctype2gptype[dtype]
		meta.Inputs = append(meta.Inputs, io)

	}

	outputCount := C.c2_engine_get_output_count(c2c.CEngineCtxs[0])
	c2c.Outputs = make([]string, outputCount)
	for i := 0; i < int(outputCount); i++ {
		cname := C.c2_engine_get_output_name(c2c.CEngineCtxs[0], C.int(i))
		name := C.GoString(cname)
		c2c.Outputs[i] = name
		dims := make([]int64, 32)
		dimCount := C.c2_engine_get_dimensions(c2c.CEngineCtxs[0], cname, (*C.int64_t)(unsafe.Pointer(&dims[0])))
		if dimCount < 0 {
			logrus.Fatalf("Could not find dimensions for: %s", name)
		}
		c2c.OutputDims[name] = dims[:dimCount]
		dtype := C.c2_engine_get_dtype(c2c.CEngineCtxs[0], cname)

		logrus.Debugf("Adding output '%s'", name)
		io := graphpipe.NativeIOMetadata{}
		io.Name = name
		io.Shape = dims[:dimCount]
		io.Type = ctype2gptype[dtype]
		meta.Outputs = append(meta.Outputs, io)
	}

	meta.Name = opts.model
	meta.Description = "Implementation of onnx/caffe2 model server using graphpipe.  Use a graphpipe client to make requests to this server."
	meta.Server = "graphpipe-tf"
	meta.Version = version()

	c2c.meta = meta

	cachePath := ""
	if opts.cache {
		cachePath = filepath.Join(opts.cacheDir, fmt.Sprintf("%x.db", c2c.modelHash))
	}

	serveOpts := &graphpipe.ServeRawOptions{
		Listen:         opts.listen,
		CacheFile:      cachePath,
		Meta:           c2c.meta,
		Apply:          c2c.apply,
		GetHandler:     c2c.getHandler,
		DefaultInputs:  c2c.Inputs,
		DefaultOutputs: c2c.Outputs,
	}
	if err := graphpipe.ServeRaw(serveOpts); err != nil {
		return err
	}
	return nil
}

func (c2c *c2Context) getHandler(w http.ResponseWriter, r *http.Request, body []byte) error {
	js, err := json.MarshalIndent(c2c.meta, "", "    ")
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
	return err
}

var gptype2ctype = []int{
	C.TensorProto_DataType_UNDEFINED, // Type_Null = 0,
	C.TensorProto_DataType_UINT8,     // Type_Uint8 = 1,
	C.TensorProto_DataType_INT8,      // Type_Int8 = 2,
	C.TensorProto_DataType_UINT16,    // Type_Uint16 = 3,
	C.TensorProto_DataType_INT16,     // Type_Int16 = 4,
	C.TensorProto_DataType_UNDEFINED, // Type_Uint32 = 5,
	C.TensorProto_DataType_INT32,     // Type_Int32 = 6,
	C.TensorProto_DataType_UNDEFINED, // Type_Uint64 = 7,
	C.TensorProto_DataType_INT64,     // Type_Int64 = 8,
	C.TensorProto_DataType_FLOAT16,   // Type_Float16 = 9,
	C.TensorProto_DataType_FLOAT,     // Type_Float32 = 10,
	C.TensorProto_DataType_DOUBLE,    // Type_Float64 = 11,
	C.TensorProto_DataType_STRING,    // Type_String = 12,
}

var ctype2gptype = []uint8{
	graphpipefb.TypeNull,    // TensorProto_DataType_UNDEFINED = 0,
	graphpipefb.TypeFloat32, // TensorProto_DataType_FLOAT = 1,
	graphpipefb.TypeInt32,   // TensorProto_DataType_INT32 = 2,
	graphpipefb.TypeInt8,    // TensorProto_DataType_BYTE = 3,
	graphpipefb.TypeString,  // TensorProto_DataType_STRING = 4,
	graphpipefb.TypeNull,    // TensorProto_DataType_BOOL = 5,
	graphpipefb.TypeInt8,    // TensorProto_DataType_UINT8 = 6,
	graphpipefb.TypeInt8,    // TensorProto_DataType_INT8 = 7,
	graphpipefb.TypeUint16,  // TensorProto_DataType_UINT16 = 8,
	graphpipefb.TypeNull,    // TensorProto_DataType_INT16 = 9,
	graphpipefb.TypeInt64,   // TensorProto_DataType_INT64 = 10,
	graphpipefb.TypeFloat16, // TensorProto_DataType_FLOAT16 = 12,
	graphpipefb.TypeFloat64, // TensorProto_DataType_DOUBLE = 13
}

func (c2c *c2Context) apply(requestContext *graphpipe.RequestContext, config string, inputs map[string]*graphpipe.NativeTensor, outputNames []string) ([]*graphpipe.NativeTensor, error) {
	engine_ctx := <-c2c.engineChannels
	defer func() {
		c2c.engineChannels <- engine_ctx
	}()

	outputNts := make([]*graphpipe.NativeTensor, len(outputNames))
	for name, input := range inputs {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))
		dtype := int(C.c2_engine_get_dtype(engine_ctx, cname))
		if dtype < 0 {
			return nil, fmt.Errorf("Could not find input: %s", name)
		}

		if gptype2ctype[input.Type] != dtype {
			return nil, fmt.Errorf("Input type mismatch.  Got %d expected %d", gptype2ctype[input.Type], C.c2_engine_get_dtype(engine_ctx, cname))
		}

		itemSize := int(C.c2_engine_get_itemsize(engine_ctx, cname))
		size := len(input.Data)
		if 0 != int(C.c2_set_input_batch(engine_ctx, cname, unsafe.Pointer(&input.Data[0]), C.int(size/itemSize),
			(*C.int64_t)(&input.Shape[0]), C.int(len(input.Shape)))) {
			return nil, fmt.Errorf("Could not set input batch for: %s", name)
		}
	}

	C.c2_execute_batch(engine_ctx)
	for i, name := range outputNames {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		idx := C.c2_engine_get_output_index(engine_ctx, cname)
		if idx < 0 {
			return nil, fmt.Errorf("Could not find requested output: %s", name)
		}

		itemSize := int64(C.c2_engine_get_itemsize(engine_ctx, cname))
		if itemSize < 0 {
			return nil, fmt.Errorf("Could not find itemSize for requested output: %s", name)
		}

		dtype := int(C.c2_engine_get_dtype(engine_ctx, cname))
		if dtype < 0 {
			return nil, fmt.Errorf("Could not find dtype for requested output: %s", name)
		}

		size := C.c2_engine_get_output_size(engine_ctx, idx)
		if size < 0 {
			return nil, fmt.Errorf("Could not find size for requested output: %s", name)
		}
		buf := make([]byte, size)

		shape := make([]int64, len(c2c.OutputDims[name]))
		rsize := C.c2_engine_get_output(engine_ctx, idx, unsafe.Pointer(&buf[0]), (*C.int64_t)(&shape[0]), C.int(len(shape)))
		if rsize != size {
			return nil, fmt.Errorf("C.c2_engine_get_output size mismatch: %d != %d", size, rsize)
		}

		gptype := -1
		for gp, dt := range gptype2ctype {
			if dt == dtype {
				gptype = gp
				break
			}
		}

		if gptype < 0 {
			return nil, fmt.Errorf("Unhandled type %d", dtype)
		}

		nt := &graphpipe.NativeTensor{}
		nt.InitWithData(buf, shape, uint8(gptype))
		outputNts[i] = nt
	}
	return outputNts, nil
}
