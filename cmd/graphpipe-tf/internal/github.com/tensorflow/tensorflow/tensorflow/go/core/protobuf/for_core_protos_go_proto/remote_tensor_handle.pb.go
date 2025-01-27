// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/protobuf/remote_tensor_handle.proto

package for_core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	tensor_shape_go_proto "github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf/internal/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/vishvananda/graphpipe-go/cmd/graphpipe-tf/internal/github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
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

type ResourceDtypeAndShape struct {
	Dtype                types_go_proto.DataType                 `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ResourceDtypeAndShape) Reset()         { *m = ResourceDtypeAndShape{} }
func (m *ResourceDtypeAndShape) String() string { return proto.CompactTextString(m) }
func (*ResourceDtypeAndShape) ProtoMessage()    {}
func (*ResourceDtypeAndShape) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7497b828cba9b15, []int{0}
}
func (m *ResourceDtypeAndShape) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceDtypeAndShape) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ResourceDtypeAndShape) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDtypeAndShape.Merge(m, src)
}
func (m *ResourceDtypeAndShape) XXX_Size() int {
	return m.Size()
}
func (m *ResourceDtypeAndShape) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDtypeAndShape.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDtypeAndShape proto.InternalMessageInfo

func (m *ResourceDtypeAndShape) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *ResourceDtypeAndShape) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

type RemoteTensorHandle struct {
	// The ID of the operation that produced this tensor.
	OpId int64 `protobuf:"varint,1,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	// The index into the outputs of the operation that produced this tensor.
	OutputNum int32 `protobuf:"varint,2,opt,name=output_num,json=outputNum,proto3" json:"output_num,omitempty"`
	// Device where the tensor is located. Cannot be empty.
	// For multi-device functions, it's the default device passed to placer.
	Device string `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	// Device of the operation producing this tensor. Can be empty if the
	// operation producing this tensor is a multi-device function.
	OpDevice string `protobuf:"bytes,4,opt,name=op_device,json=opDevice,proto3" json:"op_device,omitempty"`
	// Tensor type.
	Dtype types_go_proto.DataType `protobuf:"varint,5,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// Optional data types and shapes of a remote resource variable.
	ResourceDtypesAndShapes []*ResourceDtypeAndShape `protobuf:"bytes,6,rep,name=resource_dtypes_and_shapes,json=resourceDtypesAndShapes,proto3" json:"resource_dtypes_and_shapes,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *RemoteTensorHandle) Reset()         { *m = RemoteTensorHandle{} }
func (m *RemoteTensorHandle) String() string { return proto.CompactTextString(m) }
func (*RemoteTensorHandle) ProtoMessage()    {}
func (*RemoteTensorHandle) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7497b828cba9b15, []int{1}
}
func (m *RemoteTensorHandle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoteTensorHandle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RemoteTensorHandle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteTensorHandle.Merge(m, src)
}
func (m *RemoteTensorHandle) XXX_Size() int {
	return m.Size()
}
func (m *RemoteTensorHandle) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteTensorHandle.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteTensorHandle proto.InternalMessageInfo

func (m *RemoteTensorHandle) GetOpId() int64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *RemoteTensorHandle) GetOutputNum() int32 {
	if m != nil {
		return m.OutputNum
	}
	return 0
}

func (m *RemoteTensorHandle) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *RemoteTensorHandle) GetOpDevice() string {
	if m != nil {
		return m.OpDevice
	}
	return ""
}

func (m *RemoteTensorHandle) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *RemoteTensorHandle) GetResourceDtypesAndShapes() []*ResourceDtypeAndShape {
	if m != nil {
		return m.ResourceDtypesAndShapes
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceDtypeAndShape)(nil), "tensorflow.eager.ResourceDtypeAndShape")
	proto.RegisterType((*RemoteTensorHandle)(nil), "tensorflow.eager.RemoteTensorHandle")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/remote_tensor_handle.proto", fileDescriptor_c7497b828cba9b15)
}

var fileDescriptor_c7497b828cba9b15 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0xaa, 0xd3, 0x40,
	0x18, 0x65, 0xda, 0xa6, 0xd8, 0x29, 0x88, 0x8c, 0x7f, 0xa1, 0x6a, 0x08, 0x05, 0x31, 0x88, 0x24,
	0x90, 0x3e, 0x81, 0xa5, 0x0b, 0xdd, 0x48, 0x19, 0xeb, 0xc6, 0xcd, 0x90, 0x66, 0x26, 0x69, 0xb1,
	0xc9, 0x37, 0xcc, 0x24, 0x96, 0x3e, 0x85, 0xaf, 0xe0, 0xe3, 0xb8, 0x74, 0xe9, 0x52, 0xf2, 0x14,
	0x2e, 0x25, 0x33, 0xb1, 0x86, 0xde, 0x72, 0xef, 0xee, 0xcb, 0x39, 0x67, 0xce, 0x97, 0x73, 0x66,
	0xf0, 0xa2, 0x12, 0xa5, 0x06, 0x95, 0x1d, 0xe0, 0x18, 0xa5, 0xa0, 0x44, 0x24, 0x15, 0x54, 0xb0,
	0xad, 0xb3, 0x48, 0x89, 0x02, 0x2a, 0xc1, 0x2c, 0xcf, 0x76, 0x49, 0xc9, 0x0f, 0x22, 0x34, 0x2c,
	0x79, 0xf0, 0xff, 0x50, 0x28, 0x92, 0x5c, 0xa8, 0xd9, 0x9b, 0x4b, 0x9b, 0x4c, 0x25, 0x85, 0x38,
	0x82, 0xfa, 0x12, 0x75, 0x06, 0x7a, 0x97, 0xc8, 0xee, 0xfc, 0xec, 0xe5, 0x2d, 0xea, 0x93, 0x14,
	0xda, 0xca, 0xe6, 0x47, 0xfc, 0x98, 0x0a, 0x0d, 0xb5, 0x4a, 0xc5, 0xaa, 0xc5, 0xdf, 0x96, 0xfc,
	0x63, 0xeb, 0x42, 0x5e, 0x63, 0x87, 0xb7, 0x80, 0x8b, 0x7c, 0x14, 0xdc, 0x8f, 0x1f, 0x85, 0xbd,
	0xff, 0x59, 0x25, 0x55, 0xb2, 0x39, 0x49, 0x41, 0xad, 0x84, 0xc4, 0xd8, 0x31, 0xab, 0xdd, 0x81,
	0x8f, 0x82, 0x69, 0xfc, 0xbc, 0xaf, 0xdd, 0x98, 0xd1, 0x78, 0xae, 0xdb, 0x8d, 0xd4, 0x4a, 0xe7,
	0xdf, 0x06, 0x98, 0x50, 0x13, 0xdf, 0x2a, 0xde, 0x99, 0xf0, 0xe4, 0x21, 0x76, 0x40, 0xb2, 0x3d,
	0x37, 0x6b, 0x87, 0x74, 0x04, 0xf2, 0x3d, 0x27, 0x2f, 0x30, 0x86, 0xba, 0x92, 0x75, 0xc5, 0xca,
	0xba, 0x30, 0x4b, 0x1c, 0x3a, 0xb1, 0xc8, 0x87, 0xba, 0x20, 0x4f, 0xf0, 0x98, 0x8b, 0xaf, 0xfb,
	0x54, 0xb8, 0x43, 0x1f, 0x05, 0x13, 0xda, 0x7d, 0x91, 0x67, 0x78, 0x02, 0x92, 0x75, 0xd4, 0xc8,
	0x50, 0xf7, 0x40, 0xae, 0x2c, 0x79, 0xce, 0xe7, 0xdc, 0x9d, 0x8f, 0xe3, 0x99, 0xea, 0x4a, 0x62,
	0x06, 0xd1, 0x2c, 0x29, 0xb9, 0xad, 0x5b, 0xbb, 0x63, 0x7f, 0x18, 0x4c, 0xe3, 0x57, 0xe1, 0xe5,
	0x85, 0x85, 0x57, 0x8b, 0xa5, 0x4f, 0x55, 0x1f, 0xd6, 0xff, 0x70, 0xbd, 0xfc, 0x8e, 0x7e, 0x34,
	0x1e, 0xfa, 0xd9, 0x78, 0xe8, 0x57, 0xe3, 0xa1, 0xdf, 0x8d, 0x87, 0xb0, 0x0b, 0x2a, 0xef, 0xfb,
	0x9e, 0xef, 0x70, 0xe9, 0xde, 0xec, 0xce, 0xd4, 0xab, 0xd7, 0xe8, 0xf3, 0xa7, 0x7c, 0x5f, 0xed,
	0xea, 0x6d, 0x98, 0x42, 0x11, 0xf5, 0x5e, 0xc1, 0xf5, 0x31, 0x87, 0x8b, 0x37, 0x99, 0x81, 0x62,
	0x2d, 0xc2, 0x0c, 0xa2, 0x59, 0x0e, 0x76, 0xfa, 0x83, 0xd0, 0x76, 0x6c, 0xa6, 0xc5, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x36, 0xf0, 0x25, 0x84, 0xd2, 0x02, 0x00, 0x00,
}

func (m *ResourceDtypeAndShape) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceDtypeAndShape) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceDtypeAndShape) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Shape != nil {
		{
			size, err := m.Shape.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRemoteTensorHandle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Dtype != 0 {
		i = encodeVarintRemoteTensorHandle(dAtA, i, uint64(m.Dtype))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RemoteTensorHandle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteTensorHandle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoteTensorHandle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ResourceDtypesAndShapes) > 0 {
		for iNdEx := len(m.ResourceDtypesAndShapes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResourceDtypesAndShapes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRemoteTensorHandle(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Dtype != 0 {
		i = encodeVarintRemoteTensorHandle(dAtA, i, uint64(m.Dtype))
		i--
		dAtA[i] = 0x28
	}
	if len(m.OpDevice) > 0 {
		i -= len(m.OpDevice)
		copy(dAtA[i:], m.OpDevice)
		i = encodeVarintRemoteTensorHandle(dAtA, i, uint64(len(m.OpDevice)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Device) > 0 {
		i -= len(m.Device)
		copy(dAtA[i:], m.Device)
		i = encodeVarintRemoteTensorHandle(dAtA, i, uint64(len(m.Device)))
		i--
		dAtA[i] = 0x1a
	}
	if m.OutputNum != 0 {
		i = encodeVarintRemoteTensorHandle(dAtA, i, uint64(m.OutputNum))
		i--
		dAtA[i] = 0x10
	}
	if m.OpId != 0 {
		i = encodeVarintRemoteTensorHandle(dAtA, i, uint64(m.OpId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRemoteTensorHandle(dAtA []byte, offset int, v uint64) int {
	offset -= sovRemoteTensorHandle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResourceDtypeAndShape) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dtype != 0 {
		n += 1 + sovRemoteTensorHandle(uint64(m.Dtype))
	}
	if m.Shape != nil {
		l = m.Shape.Size()
		n += 1 + l + sovRemoteTensorHandle(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RemoteTensorHandle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OpId != 0 {
		n += 1 + sovRemoteTensorHandle(uint64(m.OpId))
	}
	if m.OutputNum != 0 {
		n += 1 + sovRemoteTensorHandle(uint64(m.OutputNum))
	}
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovRemoteTensorHandle(uint64(l))
	}
	l = len(m.OpDevice)
	if l > 0 {
		n += 1 + l + sovRemoteTensorHandle(uint64(l))
	}
	if m.Dtype != 0 {
		n += 1 + sovRemoteTensorHandle(uint64(m.Dtype))
	}
	if len(m.ResourceDtypesAndShapes) > 0 {
		for _, e := range m.ResourceDtypesAndShapes {
			l = e.Size()
			n += 1 + l + sovRemoteTensorHandle(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRemoteTensorHandle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRemoteTensorHandle(x uint64) (n int) {
	return sovRemoteTensorHandle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceDtypeAndShape) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemoteTensorHandle
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
			return fmt.Errorf("proto: ResourceDtypeAndShape: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceDtypeAndShape: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtype", wireType)
			}
			m.Dtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteTensorHandle
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
					return ErrIntOverflowRemoteTensorHandle
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
				return ErrInvalidLengthRemoteTensorHandle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRemoteTensorHandle
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
		default:
			iNdEx = preIndex
			skippy, err := skipRemoteTensorHandle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRemoteTensorHandle
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
func (m *RemoteTensorHandle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemoteTensorHandle
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
			return fmt.Errorf("proto: RemoteTensorHandle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteTensorHandle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpId", wireType)
			}
			m.OpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteTensorHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OpId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputNum", wireType)
			}
			m.OutputNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteTensorHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OutputNum |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteTensorHandle
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
				return ErrInvalidLengthRemoteTensorHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRemoteTensorHandle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpDevice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteTensorHandle
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
				return ErrInvalidLengthRemoteTensorHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRemoteTensorHandle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OpDevice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtype", wireType)
			}
			m.Dtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteTensorHandle
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceDtypesAndShapes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteTensorHandle
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
				return ErrInvalidLengthRemoteTensorHandle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRemoteTensorHandle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceDtypesAndShapes = append(m.ResourceDtypesAndShapes, &ResourceDtypeAndShape{})
			if err := m.ResourceDtypesAndShapes[len(m.ResourceDtypesAndShapes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemoteTensorHandle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRemoteTensorHandle
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
func skipRemoteTensorHandle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRemoteTensorHandle
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
					return 0, ErrIntOverflowRemoteTensorHandle
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
					return 0, ErrIntOverflowRemoteTensorHandle
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
				return 0, ErrInvalidLengthRemoteTensorHandle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRemoteTensorHandle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRemoteTensorHandle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRemoteTensorHandle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRemoteTensorHandle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRemoteTensorHandle = fmt.Errorf("proto: unexpected end of group")
)
