// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/protobuf/worker_service.proto

package for_core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/worker_service.proto", fileDescriptor_cf631a047edac54b)
}

var fileDescriptor_cf631a047edac54b = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xf1, 0x85, 0xd0, 0x11, 0x55, 0xd0, 0xf6, 0x96, 0xb6, 0xa6, 0xb4, 0xfc, 0x39, 0x91,
	0x48, 0xf0, 0x04, 0x75, 0x82, 0x02, 0x02, 0xa4, 0x92, 0x04, 0x21, 0x45, 0x42, 0x91, 0x63, 0xc6,
	0x5b, 0x0b, 0x67, 0xd7, 0xec, 0x9f, 0xf6, 0x5d, 0x78, 0x22, 0x8e, 0x1c, 0x39, 0xa2, 0x3c, 0x01,
	0x8f, 0x80, 0x6c, 0xef, 0x26, 0xde, 0x74, 0x93, 0x9c, 0x62, 0x7d, 0xdf, 0x37, 0xbf, 0x99, 0xc8,
	0x9e, 0x81, 0x97, 0x0a, 0x99, 0xe4, 0x22, 0xcd, 0xf9, 0x6d, 0x2f, 0xe1, 0x02, 0x7b, 0x85, 0xe0,
	0x8a, 0xcf, 0x75, 0xda, 0xbb, 0xe5, 0xe2, 0x3b, 0x8a, 0x99, 0x44, 0x71, 0x93, 0x25, 0xd8, 0xad,
	0x74, 0xd2, 0x5e, 0xc7, 0xbb, 0x54, 0x14, 0x49, 0xe7, 0xd9, 0x9e, 0xfa, 0xba, 0xee, 0xd5, 0xbf,
	0x03, 0x38, 0xfc, 0x52, 0x09, 0xe3, 0x9a, 0x47, 0xde, 0xc2, 0xc1, 0x10, 0xd5, 0x58, 0xc5, 0x4a,
	0x4b, 0x72, 0xd2, 0x6d, 0x70, 0x57, 0xf2, 0x08, 0x7f, 0x68, 0x94, 0xaa, 0x73, 0xba, 0xc5, 0x95,
	0x05, 0x67, 0x12, 0x49, 0x0a, 0x47, 0x7d, 0x81, 0xb1, 0x42, 0xdb, 0x40, 0xca, 0x8c, 0x33, 0xf2,
	0xbc, 0x59, 0xe5, 0x09, 0x58, 0xfa, 0x8b, 0xbd, 0xb9, 0x75, 0x9f, 0x01, 0xe6, 0xb8, 0xb3, 0x8f,
	0x27, 0xe0, 0xed, 0xe3, 0xcd, 0x99, 0x3e, 0x13, 0x38, 0x1c, 0x21, 0xcd, 0xa4, 0x42, 0x31, 0x14,
	0x71, 0x71, 0x4d, 0xce, 0x9a, 0x95, 0x8e, 0x65, 0xd9, 0x4f, 0x76, 0x24, 0x0c, 0x75, 0x0a, 0xed,
	0x01, 0x0a, 0x87, 0x7b, 0xee, 0x4e, 0x24, 0x7c, 0xe4, 0x8b, 0x9d, 0x19, 0xc3, 0x7e, 0x03, 0x0f,
	0x46, 0x9a, 0xd5, 0xd0, 0x63, 0x67, 0x14, 0xa3, 0x5a, 0xda, 0x89, 0xdf, 0x34, 0x98, 0x4f, 0xf0,
	0xb0, 0x9f, 0x63, 0xcc, 0x74, 0x51, 0xa3, 0x1e, 0x3b, 0x6f, 0xa6, 0xe1, 0x58, 0xdc, 0xd9, 0xf6,
	0x80, 0x41, 0xbe, 0x07, 0x30, 0xfa, 0x65, 0x9e, 0x93, 0x53, 0x4f, 0xfe, 0x32, 0xcf, 0x2d, 0x2e,
	0xdc, 0x66, 0x1b, 0xd8, 0x47, 0x80, 0x11, 0x26, 0x37, 0x93, 0x2a, 0xe4, 0xc2, 0xd6, 0xba, 0x17,
	0xd6, 0xb4, 0x6b, 0xd8, 0xf9, 0x3d, 0x12, 0x41, 0xeb, 0x03, 0xa7, 0x34, 0x63, 0x94, 0x74, 0x9a,
	0x61, 0x23, 0x5a, 0xd0, 0xb1, 0xd7, 0x33, 0x23, 0x45, 0xd0, 0x9a, 0x88, 0x38, 0xb9, 0xc3, 0x30,
	0xa2, 0x97, 0xb1, 0xf2, 0x0c, 0x63, 0x00, 0xad, 0x72, 0xbe, 0x48, 0xa7, 0x2e, 0xc3, 0x88, 0x5e,
	0xc6, 0xca, 0x5b, 0xfd, 0x9b, 0x29, 0xb4, 0xab, 0xd5, 0xc4, 0x62, 0x5c, 0x16, 0xb0, 0x04, 0xdd,
	0xef, 0x6b, 0xc3, 0xf4, 0x7e, 0x5f, 0x77, 0x32, 0xeb, 0x8d, 0xe8, 0xf3, 0x45, 0x51, 0xae, 0xcc,
	0x50, 0x70, 0x5d, 0xb8, 0x1b, 0xe1, 0x58, 0xde, 0x8d, 0xd8, 0x48, 0x18, 0xea, 0x57, 0x78, 0x64,
	0x8d, 0x77, 0x4c, 0xaa, 0xb8, 0x1c, 0xf9, 0xc2, 0x57, 0x66, 0x5d, 0xcb, 0x7e, 0xba, 0x3b, 0x54,
	0xe3, 0xa3, 0x9f, 0xc1, 0xaf, 0x65, 0x18, 0xfc, 0x5e, 0x86, 0xc1, 0x9f, 0x65, 0x18, 0xfc, 0x5d,
	0x86, 0x01, 0x74, 0xb8, 0xa0, 0xcd, 0xda, 0x6f, 0x99, 0x54, 0x42, 0x33, 0x95, 0x2d, 0x30, 0x3a,
	0x72, 0xce, 0xe3, 0x55, 0x79, 0x35, 0xe5, 0x55, 0x30, 0xfd, 0x4c, 0x33, 0x75, 0xad, 0xe7, 0xdd,
	0x84, 0x2f, 0x7a, 0x8d, 0x53, 0xeb, 0x7f, 0xa4, 0x7c, 0xe3, 0x06, 0xa7, 0x5c, 0xcc, 0x4a, 0x65,
	0x56, 0x29, 0x72, 0x46, 0x79, 0xfd, 0x34, 0xbf, 0x5f, 0xfd, 0xbc, 0xfe, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0x04, 0xc4, 0x82, 0x57, 0xff, 0x05, 0x00, 0x00,
}
