// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/protobuf/debug.proto

package for_core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Option for watching a node in TensorFlow Debugger (tfdbg).
type DebugTensorWatch struct {
	// Name of the node to watch.
	// Use "*" for wildcard. But note: currently, regex is not supported in
	// general.
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Output slot to watch.
	// The semantics of output_slot == -1 is that all outputs of the node
	// will be watched (i.e., a wildcard).
	// Other negative values of output_slot are invalid and will lead to
	// errors currently.
	OutputSlot int32 `protobuf:"varint,2,opt,name=output_slot,json=outputSlot,proto3" json:"output_slot,omitempty"`
	// Name(s) of the debugging op(s).
	// One or more than one probes on a tensor.
	// e.g., {"DebugIdentity", "DebugNanCount"}
	DebugOps []string `protobuf:"bytes,3,rep,name=debug_ops,json=debugOps,proto3" json:"debug_ops,omitempty"`
	// URL(s) for debug targets(s).
	//
	// Supported URL formats are:
	//   - file:///foo/tfdbg_dump: Writes out Event content to file
	//     /foo/tfdbg_dump.  Assumes all directories can be created if they don't
	//     already exist.
	//   - grpc://localhost:11011: Sends an RPC request to an EventListener
	//     service running at localhost:11011 with the event.
	//   - memcbk:///event_key: Routes tensors to clients using the
	//     callback registered with the DebugCallbackRegistry for event_key.
	//
	// Each debug op listed in debug_ops will publish its output tensor (debug
	// signal) to all URLs in debug_urls.
	//
	// N.B. Session::Run() supports concurrent invocations of the same inputs
	// (feed keys), outputs and target nodes. If such concurrent invocations
	// are to be debugged, the callers of Session::Run() must use distinct
	// debug_urls to make sure that the streamed or dumped events do not overlap
	// among the invocations.
	// TODO(cais): More visible documentation of this in g3docs.
	DebugUrls []string `protobuf:"bytes,4,rep,name=debug_urls,json=debugUrls,proto3" json:"debug_urls,omitempty"`
	// Do not error out if debug op creation fails (e.g., due to dtype
	// incompatibility). Instead, just log the failure.
	TolerateDebugOpCreationFailures bool     `protobuf:"varint,5,opt,name=tolerate_debug_op_creation_failures,json=tolerateDebugOpCreationFailures,proto3" json:"tolerate_debug_op_creation_failures,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *DebugTensorWatch) Reset()         { *m = DebugTensorWatch{} }
func (m *DebugTensorWatch) String() string { return proto.CompactTextString(m) }
func (*DebugTensorWatch) ProtoMessage()    {}
func (*DebugTensorWatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fbf764b7c91eef6, []int{0}
}
func (m *DebugTensorWatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DebugTensorWatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DebugTensorWatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugTensorWatch.Merge(m, src)
}
func (m *DebugTensorWatch) XXX_Size() int {
	return m.Size()
}
func (m *DebugTensorWatch) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugTensorWatch.DiscardUnknown(m)
}

var xxx_messageInfo_DebugTensorWatch proto.InternalMessageInfo

func (m *DebugTensorWatch) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *DebugTensorWatch) GetOutputSlot() int32 {
	if m != nil {
		return m.OutputSlot
	}
	return 0
}

func (m *DebugTensorWatch) GetDebugOps() []string {
	if m != nil {
		return m.DebugOps
	}
	return nil
}

func (m *DebugTensorWatch) GetDebugUrls() []string {
	if m != nil {
		return m.DebugUrls
	}
	return nil
}

func (m *DebugTensorWatch) GetTolerateDebugOpCreationFailures() bool {
	if m != nil {
		return m.TolerateDebugOpCreationFailures
	}
	return false
}

// Options for initializing DebuggerState in TensorFlow Debugger (tfdbg).
type DebugOptions struct {
	// Debugging options
	DebugTensorWatchOpts []*DebugTensorWatch `protobuf:"bytes,4,rep,name=debug_tensor_watch_opts,json=debugTensorWatchOpts,proto3" json:"debug_tensor_watch_opts,omitempty"`
	// Caller-specified global step count.
	// Note that this is distinct from the session run count and the executor
	// step count.
	GlobalStep int64 `protobuf:"varint,10,opt,name=global_step,json=globalStep,proto3" json:"global_step,omitempty"`
	// Whether the total disk usage of tfdbg is to be reset to zero
	// in this Session.run call. This is used by wrappers and hooks
	// such as the local CLI ones to indicate that the dumped tensors
	// are cleaned up from the disk after each Session.run.
	ResetDiskByteUsage   bool     `protobuf:"varint,11,opt,name=reset_disk_byte_usage,json=resetDiskByteUsage,proto3" json:"reset_disk_byte_usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugOptions) Reset()         { *m = DebugOptions{} }
func (m *DebugOptions) String() string { return proto.CompactTextString(m) }
func (*DebugOptions) ProtoMessage()    {}
func (*DebugOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fbf764b7c91eef6, []int{1}
}
func (m *DebugOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DebugOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DebugOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugOptions.Merge(m, src)
}
func (m *DebugOptions) XXX_Size() int {
	return m.Size()
}
func (m *DebugOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugOptions.DiscardUnknown(m)
}

var xxx_messageInfo_DebugOptions proto.InternalMessageInfo

func (m *DebugOptions) GetDebugTensorWatchOpts() []*DebugTensorWatch {
	if m != nil {
		return m.DebugTensorWatchOpts
	}
	return nil
}

func (m *DebugOptions) GetGlobalStep() int64 {
	if m != nil {
		return m.GlobalStep
	}
	return 0
}

func (m *DebugOptions) GetResetDiskByteUsage() bool {
	if m != nil {
		return m.ResetDiskByteUsage
	}
	return false
}

type DebuggedSourceFile struct {
	// The host name on which a source code file is located.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Path to the source code file.
	FilePath string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// The timestamp at which the source code file is last modified.
	LastModified int64 `protobuf:"varint,3,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	// Byte size of the file.
	Bytes int64 `protobuf:"varint,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
	// Line-by-line content of the source code file.
	Lines                []string `protobuf:"bytes,5,rep,name=lines,proto3" json:"lines,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebuggedSourceFile) Reset()         { *m = DebuggedSourceFile{} }
func (m *DebuggedSourceFile) String() string { return proto.CompactTextString(m) }
func (*DebuggedSourceFile) ProtoMessage()    {}
func (*DebuggedSourceFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fbf764b7c91eef6, []int{2}
}
func (m *DebuggedSourceFile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DebuggedSourceFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DebuggedSourceFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebuggedSourceFile.Merge(m, src)
}
func (m *DebuggedSourceFile) XXX_Size() int {
	return m.Size()
}
func (m *DebuggedSourceFile) XXX_DiscardUnknown() {
	xxx_messageInfo_DebuggedSourceFile.DiscardUnknown(m)
}

var xxx_messageInfo_DebuggedSourceFile proto.InternalMessageInfo

func (m *DebuggedSourceFile) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *DebuggedSourceFile) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *DebuggedSourceFile) GetLastModified() int64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

func (m *DebuggedSourceFile) GetBytes() int64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *DebuggedSourceFile) GetLines() []string {
	if m != nil {
		return m.Lines
	}
	return nil
}

type DebuggedSourceFiles struct {
	// A collection of source code files.
	SourceFiles          []*DebuggedSourceFile `protobuf:"bytes,1,rep,name=source_files,json=sourceFiles,proto3" json:"source_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DebuggedSourceFiles) Reset()         { *m = DebuggedSourceFiles{} }
func (m *DebuggedSourceFiles) String() string { return proto.CompactTextString(m) }
func (*DebuggedSourceFiles) ProtoMessage()    {}
func (*DebuggedSourceFiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fbf764b7c91eef6, []int{3}
}
func (m *DebuggedSourceFiles) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DebuggedSourceFiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DebuggedSourceFiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebuggedSourceFiles.Merge(m, src)
}
func (m *DebuggedSourceFiles) XXX_Size() int {
	return m.Size()
}
func (m *DebuggedSourceFiles) XXX_DiscardUnknown() {
	xxx_messageInfo_DebuggedSourceFiles.DiscardUnknown(m)
}

var xxx_messageInfo_DebuggedSourceFiles proto.InternalMessageInfo

func (m *DebuggedSourceFiles) GetSourceFiles() []*DebuggedSourceFile {
	if m != nil {
		return m.SourceFiles
	}
	return nil
}

func init() {
	proto.RegisterType((*DebugTensorWatch)(nil), "tensorflow.DebugTensorWatch")
	proto.RegisterType((*DebugOptions)(nil), "tensorflow.DebugOptions")
	proto.RegisterType((*DebuggedSourceFile)(nil), "tensorflow.DebuggedSourceFile")
	proto.RegisterType((*DebuggedSourceFiles)(nil), "tensorflow.DebuggedSourceFiles")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/debug.proto", fileDescriptor_4fbf764b7c91eef6)
}

var fileDescriptor_4fbf764b7c91eef6 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xc1, 0x8e, 0x12, 0x4d,
	0x10, 0xc7, 0xd3, 0x1f, 0xbb, 0x5f, 0xa0, 0xc1, 0xc4, 0xb4, 0x6b, 0x9c, 0x44, 0x65, 0x09, 0xeb,
	0x81, 0x13, 0x44, 0x7d, 0x02, 0x91, 0xec, 0x49, 0x5d, 0x32, 0x48, 0x34, 0x5e, 0x3a, 0x3d, 0x4c,
	0xcd, 0x30, 0xa1, 0xa1, 0x26, 0x5d, 0x35, 0x21, 0xfb, 0x1c, 0x26, 0xbe, 0x8a, 0xaf, 0xe0, 0xd1,
	0xa3, 0xf1, 0x64, 0x78, 0x0a, 0x8f, 0xa6, 0x7b, 0xd8, 0xb0, 0xb2, 0xde, 0xaa, 0x7e, 0xff, 0xea,
	0xea, 0x9a, 0x7f, 0x4d, 0xcb, 0x67, 0x0c, 0x1b, 0x42, 0x97, 0x59, 0xdc, 0x8e, 0x16, 0xe8, 0x60,
	0x54, 0x3a, 0x64, 0x4c, 0xaa, 0x6c, 0x94, 0x42, 0x52, 0xe5, 0xc3, 0x90, 0x2a, 0x79, 0xa8, 0xea,
	0xff, 0x14, 0xf2, 0xfe, 0xc4, 0x6b, 0xef, 0x03, 0xfb, 0x60, 0x78, 0xb1, 0x54, 0x8f, 0x65, 0x6b,
	0x83, 0x29, 0xe8, 0x8d, 0x59, 0x43, 0x24, 0x7a, 0x62, 0xd0, 0x8a, 0x9b, 0x1e, 0xbc, 0x33, 0x6b,
	0x50, 0xe7, 0xb2, 0x8d, 0x15, 0x97, 0x15, 0x6b, 0xb2, 0xc8, 0xd1, 0x7f, 0x3d, 0x31, 0x38, 0x8d,
	0x65, 0x8d, 0x66, 0x16, 0xd9, 0x9f, 0x0e, 0xb7, 0x69, 0x2c, 0x29, 0x6a, 0xf4, 0x1a, 0xfe, 0x74,
	0x00, 0x57, 0x25, 0xa9, 0xa7, 0x52, 0xd6, 0x62, 0xe5, 0x2c, 0x45, 0x27, 0x41, 0xad, 0xcb, 0xe7,
	0xce, 0x92, 0x7a, 0x23, 0x2f, 0x18, 0x2d, 0x38, 0xc3, 0xa0, 0x6f, 0x9a, 0xe8, 0x85, 0x03, 0xc3,
	0x05, 0x6e, 0x74, 0x66, 0x0a, 0x5b, 0x39, 0xa0, 0xe8, 0xb4, 0x27, 0x06, 0xcd, 0xf8, 0xfc, 0xa6,
	0x74, 0x52, 0x77, 0x7f, 0xbd, 0xaf, 0xbb, 0xdc, 0x97, 0xf5, 0xbf, 0x0a, 0xd9, 0xd9, 0x6b, 0x9e,
	0x93, 0x9a, 0xc9, 0x47, 0x75, 0xd7, 0xda, 0x01, 0xbd, 0xf5, 0x9f, 0xab, 0xb1, 0xe4, 0x7a, 0x94,
	0xf6, 0x8b, 0x27, 0xc3, 0x83, 0x37, 0xc3, 0x63, 0x5f, 0xe2, 0xb3, 0xf4, 0x88, 0x5c, 0x95, 0x4c,
	0xde, 0x90, 0xdc, 0x62, 0x62, 0xac, 0x26, 0x86, 0x32, 0x92, 0x3d, 0x31, 0x68, 0xc4, 0xb2, 0x46,
	0x33, 0x86, 0x52, 0x3d, 0x97, 0x0f, 0x1d, 0x10, 0xb0, 0x4e, 0x0b, 0x5a, 0xe9, 0xe4, 0x9a, 0x41,
	0x57, 0x64, 0x72, 0x88, 0xda, 0xe1, 0x33, 0x54, 0x10, 0x27, 0x05, 0xad, 0xc6, 0xd7, 0x0c, 0x73,
	0xaf, 0xf4, 0xbf, 0x08, 0xa9, 0xc2, 0xf5, 0x39, 0xa4, 0x33, 0xac, 0xdc, 0x02, 0x2e, 0x0b, 0x0b,
	0x4a, 0xc9, 0x93, 0x25, 0x12, 0xef, 0x77, 0x12, 0x62, 0x6f, 0x77, 0x56, 0x58, 0xd0, 0xa5, 0xe1,
	0x65, 0xd8, 0x46, 0x2b, 0x6e, 0x7a, 0x30, 0x35, 0xbc, 0x54, 0x17, 0xf2, 0x9e, 0x35, 0xc4, 0x7a,
	0x8d, 0x69, 0x91, 0x15, 0x90, 0x46, 0x8d, 0x30, 0x5d, 0xc7, 0xc3, 0xb7, 0x7b, 0xa6, 0xce, 0xe4,
	0xa9, 0x1f, 0xca, 0x7b, 0xe0, 0xc5, 0x3a, 0xf1, 0xd4, 0x16, 0x9b, 0x60, 0xb6, 0x5f, 0x52, 0x9d,
	0xf4, 0x3f, 0xca, 0x07, 0x77, 0xe7, 0x22, 0xf5, 0x4a, 0x76, 0x28, 0xa4, 0xda, 0x5f, 0x4d, 0x91,
	0x08, 0x6e, 0x76, 0xef, 0xb8, 0xf9, 0xd7, 0xb1, 0xb8, 0x4d, 0x87, 0x16, 0xe3, 0xcf, 0xe2, 0xdb,
	0xae, 0x2b, 0xbe, 0xef, 0xba, 0xe2, 0xc7, 0xae, 0x2b, 0x7e, 0xed, 0xba, 0x42, 0x46, 0xe8, 0xf2,
	0xdb, 0x2d, 0x32, 0x67, 0xd6, 0xb0, 0x45, 0xb7, 0x1a, 0xb7, 0x43, 0xb7, 0xa9, 0xff, 0x9d, 0x69,
	0x2a, 0x3e, 0xcd, 0xf3, 0x82, 0x97, 0x55, 0x32, 0x5c, 0xe0, 0x7a, 0x74, 0xeb, 0x09, 0xfc, 0x3b,
	0xcc, 0xf1, 0xe8, 0x6d, 0x64, 0xe8, 0xb4, 0x27, 0x3a, 0x10, 0xd2, 0x39, 0xd6, 0xd1, 0x6f, 0x21,
	0x92, 0xff, 0x43, 0xf4, 0xf2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xea, 0x79, 0x67, 0x5a,
	0x03, 0x00, 0x00,
}

func (m *DebugTensorWatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DebugTensorWatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DebugTensorWatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TolerateDebugOpCreationFailures {
		i--
		if m.TolerateDebugOpCreationFailures {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.DebugUrls) > 0 {
		for iNdEx := len(m.DebugUrls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DebugUrls[iNdEx])
			copy(dAtA[i:], m.DebugUrls[iNdEx])
			i = encodeVarintDebug(dAtA, i, uint64(len(m.DebugUrls[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DebugOps) > 0 {
		for iNdEx := len(m.DebugOps) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DebugOps[iNdEx])
			copy(dAtA[i:], m.DebugOps[iNdEx])
			i = encodeVarintDebug(dAtA, i, uint64(len(m.DebugOps[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.OutputSlot != 0 {
		i = encodeVarintDebug(dAtA, i, uint64(m.OutputSlot))
		i--
		dAtA[i] = 0x10
	}
	if len(m.NodeName) > 0 {
		i -= len(m.NodeName)
		copy(dAtA[i:], m.NodeName)
		i = encodeVarintDebug(dAtA, i, uint64(len(m.NodeName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DebugOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DebugOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DebugOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ResetDiskByteUsage {
		i--
		if m.ResetDiskByteUsage {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.GlobalStep != 0 {
		i = encodeVarintDebug(dAtA, i, uint64(m.GlobalStep))
		i--
		dAtA[i] = 0x50
	}
	if len(m.DebugTensorWatchOpts) > 0 {
		for iNdEx := len(m.DebugTensorWatchOpts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DebugTensorWatchOpts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDebug(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	return len(dAtA) - i, nil
}

func (m *DebuggedSourceFile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DebuggedSourceFile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DebuggedSourceFile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Lines) > 0 {
		for iNdEx := len(m.Lines) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Lines[iNdEx])
			copy(dAtA[i:], m.Lines[iNdEx])
			i = encodeVarintDebug(dAtA, i, uint64(len(m.Lines[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Bytes != 0 {
		i = encodeVarintDebug(dAtA, i, uint64(m.Bytes))
		i--
		dAtA[i] = 0x20
	}
	if m.LastModified != 0 {
		i = encodeVarintDebug(dAtA, i, uint64(m.LastModified))
		i--
		dAtA[i] = 0x18
	}
	if len(m.FilePath) > 0 {
		i -= len(m.FilePath)
		copy(dAtA[i:], m.FilePath)
		i = encodeVarintDebug(dAtA, i, uint64(len(m.FilePath)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarintDebug(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DebuggedSourceFiles) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DebuggedSourceFiles) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DebuggedSourceFiles) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.SourceFiles) > 0 {
		for iNdEx := len(m.SourceFiles) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SourceFiles[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDebug(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDebug(dAtA []byte, offset int, v uint64) int {
	offset -= sovDebug(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DebugTensorWatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NodeName)
	if l > 0 {
		n += 1 + l + sovDebug(uint64(l))
	}
	if m.OutputSlot != 0 {
		n += 1 + sovDebug(uint64(m.OutputSlot))
	}
	if len(m.DebugOps) > 0 {
		for _, s := range m.DebugOps {
			l = len(s)
			n += 1 + l + sovDebug(uint64(l))
		}
	}
	if len(m.DebugUrls) > 0 {
		for _, s := range m.DebugUrls {
			l = len(s)
			n += 1 + l + sovDebug(uint64(l))
		}
	}
	if m.TolerateDebugOpCreationFailures {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DebugOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DebugTensorWatchOpts) > 0 {
		for _, e := range m.DebugTensorWatchOpts {
			l = e.Size()
			n += 1 + l + sovDebug(uint64(l))
		}
	}
	if m.GlobalStep != 0 {
		n += 1 + sovDebug(uint64(m.GlobalStep))
	}
	if m.ResetDiskByteUsage {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DebuggedSourceFile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovDebug(uint64(l))
	}
	l = len(m.FilePath)
	if l > 0 {
		n += 1 + l + sovDebug(uint64(l))
	}
	if m.LastModified != 0 {
		n += 1 + sovDebug(uint64(m.LastModified))
	}
	if m.Bytes != 0 {
		n += 1 + sovDebug(uint64(m.Bytes))
	}
	if len(m.Lines) > 0 {
		for _, s := range m.Lines {
			l = len(s)
			n += 1 + l + sovDebug(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DebuggedSourceFiles) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SourceFiles) > 0 {
		for _, e := range m.SourceFiles {
			l = e.Size()
			n += 1 + l + sovDebug(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDebug(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDebug(x uint64) (n int) {
	return sovDebug(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DebugTensorWatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DebugTensorWatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DebugTensorWatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputSlot", wireType)
			}
			m.OutputSlot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OutputSlot |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebugOps", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DebugOps = append(m.DebugOps, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebugUrls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DebugUrls = append(m.DebugUrls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TolerateDebugOpCreationFailures", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TolerateDebugOpCreationFailures = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DebugOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DebugOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DebugOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebugTensorWatchOpts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DebugTensorWatchOpts = append(m.DebugTensorWatchOpts, &DebugTensorWatch{})
			if err := m.DebugTensorWatchOpts[len(m.DebugTensorWatchOpts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalStep", wireType)
			}
			m.GlobalStep = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GlobalStep |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResetDiskByteUsage", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ResetDiskByteUsage = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DebuggedSourceFile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DebuggedSourceFile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DebuggedSourceFile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastModified", wireType)
			}
			m.LastModified = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastModified |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			m.Bytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lines", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lines = append(m.Lines, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DebuggedSourceFiles) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DebuggedSourceFiles: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DebuggedSourceFiles: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceFiles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceFiles = append(m.SourceFiles, &DebuggedSourceFile{})
			if err := m.SourceFiles[len(m.SourceFiles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDebug(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDebug
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDebug
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDebug
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDebug        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDebug          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDebug = fmt.Errorf("proto: unexpected end of group")
)
