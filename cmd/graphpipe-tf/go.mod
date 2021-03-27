module github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf

go 1.16

replace github.com/vishvananda/graphpipe-go v1.0.0 => ../../

require (
	github.com/gogo/protobuf v1.2.1
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/tensorflow/tensorflow v2.4.1+incompatible
	github.com/vishvananda/graphpipe-go v1.0.0
	google.golang.org/protobuf v1.26.0 // indirect
)
