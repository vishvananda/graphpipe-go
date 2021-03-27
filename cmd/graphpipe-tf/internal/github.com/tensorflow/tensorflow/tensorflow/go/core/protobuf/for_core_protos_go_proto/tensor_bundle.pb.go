// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/protobuf/tensor_bundle.proto

package for_core_protos_go_proto

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	tensor_shape_go_proto "github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf/internal/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	tensor_slice_go_proto "github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf/internal/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_slice_go_proto"
	types_go_proto "github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf/internal/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	versions_go_proto "github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf/internal/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/versions_go_proto"
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

// An enum indicating the endianness of the platform that produced this
// bundle.  A bundle can only be read by a platform with matching endianness.
// Defaults to LITTLE, as most modern platforms are little-endian.
//
// Affects the binary tensor data bytes only, not the metadata in protobufs.
type BundleHeaderProto_Endianness int32

const (
	BundleHeaderProto_LITTLE BundleHeaderProto_Endianness = 0
	BundleHeaderProto_BIG    BundleHeaderProto_Endianness = 1
)

var BundleHeaderProto_Endianness_name = map[int32]string{
	0: "LITTLE",
	1: "BIG",
}

var BundleHeaderProto_Endianness_value = map[string]int32{
	"LITTLE": 0,
	"BIG":    1,
}

func (x BundleHeaderProto_Endianness) String() string {
	return proto.EnumName(BundleHeaderProto_Endianness_name, int32(x))
}

func (BundleHeaderProto_Endianness) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9ab648e6509929dc, []int{0, 0}
}

// Special header that is associated with a bundle.
//
// TODO(zongheng,zhifengc): maybe in the future, we can add information about
// which binary produced this checkpoint, timestamp, etc. Sometime, these can be
// valuable debugging information. And if needed, these can be used as defensive
// information ensuring reader (binary version) of the checkpoint and the writer
// (binary version) must match within certain range, etc.
type BundleHeaderProto struct {
	// Number of data files in the bundle.
	NumShards  int32                        `protobuf:"varint,1,opt,name=num_shards,json=numShards,proto3" json:"num_shards,omitempty"`
	Endianness BundleHeaderProto_Endianness `protobuf:"varint,2,opt,name=endianness,proto3,enum=tensorflow.BundleHeaderProto_Endianness" json:"endianness,omitempty"`
	// Versioning of the tensor bundle format.
	Version              *versions_go_proto.VersionDef `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *BundleHeaderProto) Reset()         { *m = BundleHeaderProto{} }
func (m *BundleHeaderProto) String() string { return proto.CompactTextString(m) }
func (*BundleHeaderProto) ProtoMessage()    {}
func (*BundleHeaderProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab648e6509929dc, []int{0}
}
func (m *BundleHeaderProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BundleHeaderProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *BundleHeaderProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleHeaderProto.Merge(m, src)
}
func (m *BundleHeaderProto) XXX_Size() int {
	return m.Size()
}
func (m *BundleHeaderProto) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleHeaderProto.DiscardUnknown(m)
}

var xxx_messageInfo_BundleHeaderProto proto.InternalMessageInfo

func (m *BundleHeaderProto) GetNumShards() int32 {
	if m != nil {
		return m.NumShards
	}
	return 0
}

func (m *BundleHeaderProto) GetEndianness() BundleHeaderProto_Endianness {
	if m != nil {
		return m.Endianness
	}
	return BundleHeaderProto_LITTLE
}

func (m *BundleHeaderProto) GetVersion() *versions_go_proto.VersionDef {
	if m != nil {
		return m.Version
	}
	return nil
}

// Describes the metadata related to a checkpointed tensor.
type BundleEntryProto struct {
	// The tensor dtype and shape.
	Dtype types_go_proto.DataType                 `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	// The binary content of the tensor lies in:
	//   File "shard_id": bytes [offset, offset + size).
	ShardId int32 `protobuf:"varint,3,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	Offset  int64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Size_   int64 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	// The CRC32C checksum of the tensor bytes.
	Crc32C uint32 `protobuf:"fixed32,6,opt,name=crc32c,proto3" json:"crc32c,omitempty"`
	// Iff present, this entry represents a partitioned tensor.  The previous
	// fields are interpreted as follows:
	//
	//   "dtype", "shape": describe the full tensor.
	//   "shard_id", "offset", "size", "crc32c": all IGNORED.
	//      These information for each slice can be looked up in their own
	//      BundleEntryProto, keyed by each "slice_name".
	Slices               []*tensor_slice_go_proto.TensorSliceProto `protobuf:"bytes,7,rep,name=slices,proto3" json:"slices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *BundleEntryProto) Reset()         { *m = BundleEntryProto{} }
func (m *BundleEntryProto) String() string { return proto.CompactTextString(m) }
func (*BundleEntryProto) ProtoMessage()    {}
func (*BundleEntryProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab648e6509929dc, []int{1}
}
func (m *BundleEntryProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BundleEntryProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *BundleEntryProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleEntryProto.Merge(m, src)
}
func (m *BundleEntryProto) XXX_Size() int {
	return m.Size()
}
func (m *BundleEntryProto) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleEntryProto.DiscardUnknown(m)
}

var xxx_messageInfo_BundleEntryProto proto.InternalMessageInfo

func (m *BundleEntryProto) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *BundleEntryProto) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *BundleEntryProto) GetShardId() int32 {
	if m != nil {
		return m.ShardId
	}
	return 0
}

func (m *BundleEntryProto) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *BundleEntryProto) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *BundleEntryProto) GetCrc32C() uint32 {
	if m != nil {
		return m.Crc32C
	}
	return 0
}

func (m *BundleEntryProto) GetSlices() []*tensor_slice_go_proto.TensorSliceProto {
	if m != nil {
		return m.Slices
	}
	return nil
}

func init() {
	proto.RegisterEnum("tensorflow.BundleHeaderProto_Endianness", BundleHeaderProto_Endianness_name, BundleHeaderProto_Endianness_value)
	proto.RegisterType((*BundleHeaderProto)(nil), "tensorflow.BundleHeaderProto")
	proto.RegisterType((*BundleEntryProto)(nil), "tensorflow.BundleEntryProto")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/tensor_bundle.proto", fileDescriptor_9ab648e6509929dc)
}

var fileDescriptor_9ab648e6509929dc = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x99, 0xa6, 0xb6, 0x61, 0x2a, 0x55, 0x61, 0x41, 0x95, 0x41, 0x10, 0x99, 0x48, 0x48,
	0x16, 0x42, 0x0e, 0x72, 0x79, 0x82, 0xa8, 0x11, 0x8d, 0xd4, 0x43, 0xe5, 0x06, 0x0e, 0x5c, 0x2c,
	0xc7, 0x5e, 0xa7, 0x16, 0xc9, 0x6e, 0xb4, 0x6b, 0x53, 0x85, 0x17, 0xe0, 0x09, 0x78, 0x27, 0x8e,
	0x70, 0xe3, 0x88, 0xfc, 0x14, 0x1c, 0x91, 0x67, 0x1d, 0x62, 0x51, 0x8a, 0x7a, 0x9b, 0x9d, 0xf9,
	0xe6, 0xdf, 0xf9, 0xc7, 0x6b, 0x7c, 0x59, 0x72, 0xa1, 0xa5, 0xca, 0x97, 0xf2, 0x6a, 0x94, 0x4a,
	0xc5, 0x47, 0x6b, 0x25, 0x4b, 0x39, 0xaf, 0xf2, 0x91, 0x29, 0xc4, 0xf3, 0x4a, 0x64, 0x4b, 0x1e,
	0x50, 0x9a, 0xe1, 0x8e, 0x7e, 0x7c, 0xad, 0x33, 0x57, 0xc9, 0x8a, 0x5f, 0x49, 0xf5, 0x61, 0xdb,
	0xaa, 0x2f, 0x93, 0x75, 0xdb, 0x79, 0x1b, 0x7a, 0x59, 0xa4, 0x5b, 0xfa, 0xf9, 0x7f, 0xe8, 0xcd,
	0x9a, 0xeb, 0x16, 0xf3, 0x6f, 0xc6, 0x3e, 0x72, 0xa5, 0x0b, 0x29, 0x5a, 0x72, 0xf8, 0x1d, 0xf0,
	0xfe, 0x98, 0x9c, 0x9c, 0xf2, 0x24, 0xe3, 0xea, 0x9c, 0xec, 0x3c, 0x45, 0x14, 0xd5, 0xaa, 0x99,
	0x53, 0x65, 0xda, 0x05, 0x0f, 0x7c, 0x2b, 0xba, 0x27, 0xaa, 0xd5, 0x05, 0x25, 0xd8, 0x29, 0x22,
	0x17, 0x59, 0x91, 0x08, 0xc1, 0xb5, 0x76, 0xf7, 0x3c, 0xf0, 0x0f, 0x43, 0x3f, 0xd8, 0xdd, 0x19,
	0x5c, 0x53, 0x0c, 0x26, 0x7f, 0xf8, 0xa8, 0xd3, 0xcb, 0x5e, 0xa1, 0xd3, 0x0e, 0xe4, 0xf6, 0x3c,
	0xf0, 0x0f, 0xc2, 0xa3, 0xae, 0xcc, 0x3b, 0x53, 0x3a, 0xe1, 0x79, 0xb4, 0xc5, 0x86, 0xcf, 0x10,
	0x77, 0x5a, 0x0c, 0xd1, 0x3e, 0x9b, 0xce, 0x66, 0x67, 0x93, 0xfe, 0x1d, 0xe6, 0x60, 0x6f, 0x3c,
	0x7d, 0xd3, 0x87, 0xe1, 0xe7, 0x3d, 0xec, 0x9b, 0x09, 0x26, 0xa2, 0x54, 0x1b, 0x63, 0xe9, 0x05,
	0x5a, 0x59, 0xb3, 0x22, 0x72, 0x73, 0x18, 0x3e, 0xec, 0xde, 0x73, 0x92, 0x94, 0xc9, 0x6c, 0xb3,
	0xe6, 0x91, 0x41, 0x58, 0x88, 0x16, 0x7d, 0x22, 0xb2, 0x76, 0x10, 0x3e, 0xe9, 0xb2, 0x33, 0x0a,
	0x2f, 0x9a, 0x32, 0x09, 0x47, 0x06, 0x65, 0x8f, 0xf0, 0x2e, 0xad, 0x2b, 0x2e, 0x32, 0xb2, 0x62,
	0x45, 0x0e, 0x9d, 0xa7, 0x19, 0x3b, 0x42, 0x5b, 0xe6, 0xb9, 0xe6, 0xa5, 0xbb, 0xef, 0x81, 0xdf,
	0x8b, 0xda, 0x13, 0x63, 0xb8, 0xaf, 0x8b, 0x4f, 0xdc, 0xb5, 0x28, 0x4b, 0x71, 0xc3, 0xa6, 0x2a,
	0x3d, 0x0e, 0x53, 0xd7, 0xf6, 0xc0, 0x77, 0xa2, 0xf6, 0xc4, 0x5e, 0xa3, 0x4d, 0xef, 0x40, 0xbb,
	0x8e, 0xd7, 0xbb, 0x61, 0xa6, 0xa6, 0x6e, 0x66, 0x6a, 0xd9, 0xf1, 0x17, 0xf8, 0x5a, 0x0f, 0xe0,
	0x5b, 0x3d, 0x80, 0x1f, 0xf5, 0x00, 0x7e, 0xd6, 0x03, 0xc0, 0x07, 0x52, 0x2d, 0xba, 0xbd, 0x55,
	0x59, 0x2c, 0xc7, 0xcc, 0x28, 0x98, 0xa5, 0x91, 0x84, 0x3e, 0x87, 0xf7, 0x6f, 0x17, 0x45, 0x79,
	0x59, 0xcd, 0x83, 0x54, 0xae, 0x46, 0x9d, 0x07, 0xf5, 0xef, 0x70, 0x21, 0xff, 0xfa, 0x4d, 0x72,
	0xa9, 0xe2, 0x26, 0x13, 0x53, 0x46, 0xc7, 0x0b, 0x69, 0xa2, 0x5f, 0x00, 0x73, 0x9b, 0xa2, 0xe3,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xed, 0x6b, 0x94, 0x65, 0x03, 0x00, 0x00,
}

func (m *BundleHeaderProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BundleHeaderProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BundleHeaderProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Version != nil {
		{
			size, err := m.Version.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTensorBundle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Endianness != 0 {
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.Endianness))
		i--
		dAtA[i] = 0x10
	}
	if m.NumShards != 0 {
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.NumShards))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BundleEntryProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BundleEntryProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BundleEntryProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Slices) > 0 {
		for iNdEx := len(m.Slices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Slices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTensorBundle(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Crc32C != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Crc32C))
		i--
		dAtA[i] = 0x35
	}
	if m.Size_ != 0 {
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x28
	}
	if m.Offset != 0 {
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x20
	}
	if m.ShardId != 0 {
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.ShardId))
		i--
		dAtA[i] = 0x18
	}
	if m.Shape != nil {
		{
			size, err := m.Shape.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTensorBundle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Dtype != 0 {
		i = encodeVarintTensorBundle(dAtA, i, uint64(m.Dtype))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTensorBundle(dAtA []byte, offset int, v uint64) int {
	offset -= sovTensorBundle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BundleHeaderProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumShards != 0 {
		n += 1 + sovTensorBundle(uint64(m.NumShards))
	}
	if m.Endianness != 0 {
		n += 1 + sovTensorBundle(uint64(m.Endianness))
	}
	if m.Version != nil {
		l = m.Version.Size()
		n += 1 + l + sovTensorBundle(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BundleEntryProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dtype != 0 {
		n += 1 + sovTensorBundle(uint64(m.Dtype))
	}
	if m.Shape != nil {
		l = m.Shape.Size()
		n += 1 + l + sovTensorBundle(uint64(l))
	}
	if m.ShardId != 0 {
		n += 1 + sovTensorBundle(uint64(m.ShardId))
	}
	if m.Offset != 0 {
		n += 1 + sovTensorBundle(uint64(m.Offset))
	}
	if m.Size_ != 0 {
		n += 1 + sovTensorBundle(uint64(m.Size_))
	}
	if m.Crc32C != 0 {
		n += 5
	}
	if len(m.Slices) > 0 {
		for _, e := range m.Slices {
			l = e.Size()
			n += 1 + l + sovTensorBundle(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTensorBundle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTensorBundle(x uint64) (n int) {
	return sovTensorBundle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BundleHeaderProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensorBundle
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
			return fmt.Errorf("proto: BundleHeaderProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BundleHeaderProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumShards", wireType)
			}
			m.NumShards = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumShards |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endianness", wireType)
			}
			m.Endianness = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Endianness |= BundleHeaderProto_Endianness(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
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
				return ErrInvalidLengthTensorBundle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTensorBundle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Version == nil {
				m.Version = &versions_go_proto.VersionDef{}
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTensorBundle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTensorBundle
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
func (m *BundleEntryProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensorBundle
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
			return fmt.Errorf("proto: BundleEntryProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BundleEntryProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtype", wireType)
			}
			m.Dtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dtype |= types_go_proto.DataType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shape", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
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
				return ErrInvalidLengthTensorBundle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTensorBundle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Shape == nil {
				m.Shape = &tensor_shape_go_proto.TensorShapeProto{}
			}
			if err := m.Shape.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardId", wireType)
			}
			m.ShardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShardId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Crc32C", wireType)
			}
			m.Crc32C = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Crc32C = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorBundle
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
				return ErrInvalidLengthTensorBundle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTensorBundle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slices = append(m.Slices, &tensor_slice_go_proto.TensorSliceProto{})
			if err := m.Slices[len(m.Slices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTensorBundle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTensorBundle
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
func skipTensorBundle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTensorBundle
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
					return 0, ErrIntOverflowTensorBundle
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
					return 0, ErrIntOverflowTensorBundle
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
				return 0, ErrInvalidLengthTensorBundle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTensorBundle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTensorBundle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTensorBundle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTensorBundle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTensorBundle = fmt.Errorf("proto: unexpected end of group")
)
