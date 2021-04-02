NAME=graphpipe-tf
RUN_TYPE?=cpu
TF_DIR=/usr/local
TF_VER=2.4.1
SUDO=sudo
LIBTF_IMAGE=libtensorflow
GRAPHPIPE_TF_IMAGE=graphpipe-tf
DOCKER=docker
MODEL_URL?=https://objectstorage.us-phoenix-1.oraclecloud.com/n/bmcskeppareuser/b/c4/o/squeezenet.pb 

ifeq ('$(USE_SUDO)', '0')
	SUDO=
endif

ifeq ('$(RUN_TYPE)', 'gpu')
	DOCKER=nvidia-docker
	LD_STUBS=LD_LIBRARY_PATH=/usr/local/cuda/lib64/stubs/:$$LD_LIBRARY_PATH
endif

sha = $(shell git rev-parse --short HEAD | tr -d ' \n')
ifeq ($(VERSION),)
VERSION = $(shell git describe --tags --match 'v*.*.*' 2> /dev/null  | tr -d 'v \n')
realv = $(shell echo $(VERSION) | cut -d- -f1)
ifneq ($(VERSION),$(realv))
commits = $(shell echo $(VERSION) | cut -d- -f2)
VERSION := $(realv).$(commits).$(sha)
endif
endif
dirty = $(shell git diff --shortstat 2> /dev/null | tail -n1 | tr -d ' \n')
ifneq ($(dirty),)
VERSION := $(VERSION).dev
endif

.PHONY: $(NAME)

in-docker:
	docker inspect --type=image $(LIBTF_IMAGE):$(RUN_TYPE)-dev > /dev/null ||  make dev-container
	docker run -it --rm \
		-v $(PWD)/../../:/go/src/github.com/vishvananda/graphpipe-go \
		-w /go/src/github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf \
		-e http_proxy=$(http_proxy) \
		-e https_proxy=$(https_proxy) \
		-e USE_SUDO=0 \
		-e GOPATH=/go \
		-e CHOWN=$$(id -u):$$(id -g) \
		$(LIBTF_IMAGE):$(RUN_TYPE)-dev \
		/usr/bin/make $(RUN_TYPE) tf-generate graphpipe-tf

check-go-version:
	$(eval min_version := go1.8)
	$(eval lesser_version=$(shell echo "`go version | cut -d " " -f3` ${min_version}" | tr " " "\n" | sort -s -t- -k 2,2nr |  sort -t. -s -k 1,1nr -k 2,2nr -k 3,3nr -k 4,4nr | tail -1))
	@if [ "${min_version}" != "${lesser_version}" ]; then \
		echo "minimum go version is ${min_version}.  You have ${lesser_version}"; \
		exit 1;\
	fi

$(NAME): check-go-version
	$(LD_STUBS) go build -ldflags '-X "main.ver=$(VERSION)" -X "main.sha=$(sha)"'

local: $(NAME)

cpu:
	# pass

gpu:
	ln -s /usr/local/cuda/lib64/stubs/libcuda.so /usr/local/cuda/lib64/libcuda.so.1

devshell:
	docker run -it --rm \
		-v $(PWD)/../../:/go/src/github.com/vishvananda/graphpipe-go \
        -w /go/src/github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf \
        -p 9000:9000 \
        -e http_proxy=$(http_proxy) \
        -e https_proxy=$(https_proxy) \
        -e GOPATH=/go \
        $(LIBTF_IMAGE):$(RUN_TYPE)-dev \
        /bin/bash

dev-container:
	docker build \
		-f Dockerfile.$(RUN_TYPE)-dev \
		-t $(LIBTF_IMAGE):$(RUN_TYPE)-dev \
		--build-arg http_proxy=$(http_proxy) \
		--build-arg https_proxy=$(https_proxy) \
		.

devserver:
	$(DOCKER) run -it --rm \
        -v $(PWD)/../../:/go/src/github.com/vishvananda/graphpipe-go \
        -w /go/src/github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf \
        -e http_proxy=$(http_proxy) \
        -e https_proxy=$(https_proxy) \
        -p 9000:9000 \
        -e GOPATH=/go \
        $(LIBTF_IMAGE):$(RUN_TYPE)-dev \
        ./graphpipe-tf --model=$(MODEL_URL) \
		--listen=0.0.0.0:9000

runtime-container:
	cd ../../ && docker build \
		-f cmd/graphpipe-tf/Dockerfile.$(RUN_TYPE) \
		-t $(GRAPHPIPE_TF_IMAGE):$(RUN_TYPE) \
		--build-arg http_proxy=$(http_proxy) \
		--build-arg https_proxy=$(https_proxy) \
		.

runserver:
	$(DOCKER) run -it --rm \
        -e http_proxy=$(http_proxy) \
        -e https_proxy=$(https_proxy) \
        -p 9000:9000 \
        $(GRAPHPIPE_TF_IMAGE):$(RUN_TYPE) \
        --model=$(MODEL_URL) \
		--listen=0.0.0.0:9000

install-tensorflow:
	@if [ ! -e /usr/local/include/tensorflow/c/c_api.h ]; then \
		curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-$(RUN_TYPE)-linux-x86_64-$(TF_VER).tar.gz" | \
			$(SUDO) tar -C $(TF_DIR) -xz; \
	fi

tensorflow:
	@if [ ! -e /usr/local/include/tensorflow/c/c_api.h ]; then \
		echo "You need tensorflow installed: https://www.tensorflow.org/install/install_go" && exit 1; \
	fi

PROTOC=protoc
TF=$(GOPATH)/src/github.com/tensorflow/tensorflow
PROTOS = $(TF)/tensorflow/core/protobuf/config.proto \
	$(TF)/tensorflow/core/protobuf/debug.proto \
	$(TF)/tensorflow/core/protobuf/cluster.proto \
	$(TF)/tensorflow/core/protobuf/rewriter_config.proto \
	$(TF)/tensorflow/core/protobuf/verifier_config.proto

protobuf:
	go get github.com/gogo/protobuf/proto
	go get github.com/gogo/protobuf/gogoproto
	cd protoc-gen-gogotf && go install
	for f in $(TF)/tensorflow/core/framework/*.proto; do $(PROTOC) -I$(TF)/ $$f --gogotf_out=internal/; done
	$(PROTOC) -I$(TF)/ $(TF)/tensorflow/stream_executor/*.proto --gogotf_out=internal/
	$(PROTOC) -I$(TF)/ $(TF)/tensorflow/core/protobuf/*.proto --gogotf_out=\
	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types:internal/
	sed -i 's;github.com/tensorflow;github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf/internal/github.com/tensorflow;g' internal/github.com/tensorflow/tensorflow/tensorflow/go/*/*/*/*.pb.go

VER=tensorflow@v$(TF_VER)+incompatible

tf-generate:
	go mod download
	mkdir -p /go/src/github.com/tensorflow
	ln -s $(GOPATH)/pkg/mod/github.com/tensorflow/$(VER) $(GOPATH)/src/github.com/tensorflow/tensorflow
	cd $(GOPATH)/pkg/mod/github.com/tensorflow/$(VER); \
		chmod -R +w .; \
		ln -s vendor/github.com/tensorflow/tensorflow/tensorflow/go/core tensorflow/go; \
		ln -s vendor/github.com/tensorflow/tensorflow/tensorflow/go/stream_executor tensorflow/go; \
		go mod init github.com/tensorflow/tensorflow; \
		go get google.golang.org/protobuf/runtime/protoimpl@v1.26.0; \
		go get google.golang.org/protobuf/reflect/protoreflect@v1.26.0; \
		go generate github.com/tensorflow/tensorflow/tensorflow/go/op

all: dev-container in-docker runtime-container
