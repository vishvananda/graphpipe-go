// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/protobuf/saved_model.proto

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

// SavedModel is the high level serialization format for TensorFlow Models.
// See [todo: doc links, similar to session_bundle] for more information.
type SavedModel struct {
	// The schema version of the SavedModel instance. Used for versioning when
	// making future changes to the specification/implementation. Initial value
	// at release will be 1.
	SavedModelSchemaVersion int64 `protobuf:"varint,1,opt,name=saved_model_schema_version,json=savedModelSchemaVersion,proto3" json:"saved_model_schema_version,omitempty"`
	// One or more MetaGraphs.
	MetaGraphs           []*MetaGraphDef `protobuf:"bytes,2,rep,name=meta_graphs,json=metaGraphs,proto3" json:"meta_graphs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SavedModel) Reset()         { *m = SavedModel{} }
func (m *SavedModel) String() string { return proto.CompactTextString(m) }
func (*SavedModel) ProtoMessage()    {}
func (*SavedModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_537826d0bcc2f334, []int{0}
}
func (m *SavedModel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SavedModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SavedModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedModel.Merge(m, src)
}
func (m *SavedModel) XXX_Size() int {
	return m.Size()
}
func (m *SavedModel) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedModel.DiscardUnknown(m)
}

var xxx_messageInfo_SavedModel proto.InternalMessageInfo

func (m *SavedModel) GetSavedModelSchemaVersion() int64 {
	if m != nil {
		return m.SavedModelSchemaVersion
	}
	return 0
}

func (m *SavedModel) GetMetaGraphs() []*MetaGraphDef {
	if m != nil {
		return m.MetaGraphs
	}
	return nil
}

func init() {
	proto.RegisterType((*SavedModel)(nil), "tensorflow.SavedModel")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/saved_model.proto", fileDescriptor_537826d0bcc2f334)
}

var fileDescriptor_537826d0bcc2f334 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x4e, 0x2c, 0x4b, 0x4d, 0x89, 0xcf, 0xcd, 0x4f, 0x49, 0xcd,
	0xd1, 0x03, 0x0b, 0x0a, 0x71, 0x21, 0xd4, 0x4a, 0x69, 0xe2, 0xd4, 0x97, 0x9b, 0x5a, 0x92, 0x18,
	0x9f, 0x5e, 0x94, 0x58, 0x90, 0x01, 0xd1, 0xa6, 0xd4, 0xc2, 0xc8, 0xc5, 0x15, 0x0c, 0x32, 0xcc,
	0x17, 0x64, 0x96, 0x90, 0x35, 0x97, 0x14, 0x92, 0xd1, 0xf1, 0xc5, 0xc9, 0x19, 0xa9, 0xb9, 0x89,
	0xf1, 0x65, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0xe2,
	0xc5, 0x70, 0xf5, 0xc1, 0x60, 0xf9, 0x30, 0x88, 0xb4, 0x90, 0x25, 0x17, 0x37, 0xc2, 0xfc, 0x62,
	0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x09, 0x3d, 0x84, 0x63, 0xf4, 0x7c, 0x53, 0x4b, 0x12,
	0xdd, 0x41, 0xb2, 0x2e, 0xa9, 0x69, 0x41, 0x5c, 0xb9, 0x30, 0x5e, 0xb1, 0xd3, 0x0c, 0xc6, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0xbc, 0xf1, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0x46,
	0x2e, 0x89, 0xfc, 0xa2, 0x74, 0x64, 0xbd, 0x69, 0x45, 0x89, 0xb9, 0xa9, 0xe5, 0xf9, 0x45, 0xd9,
	0x4e, 0x02, 0x08, 0x07, 0x07, 0x80, 0x3c, 0x51, 0x1c, 0xc0, 0x18, 0x15, 0x9a, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0xe4, 0x7b, 0xec, 0xcc, 0xf4, 0x7c, 0xb4, 0x60,
	0x49, 0xcb, 0x2f, 0x8a, 0x07, 0x89, 0xc4, 0x83, 0x45, 0x8a, 0xe3, 0xd3, 0xf3, 0x21, 0xac, 0x1f,
	0x8c, 0x8c, 0x49, 0x6c, 0x60, 0x96, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xc8, 0x47, 0x2f,
	0x8d, 0x01, 0x00, 0x00,
}

func (m *SavedModel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SavedModel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SavedModel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.MetaGraphs) > 0 {
		for iNdEx := len(m.MetaGraphs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MetaGraphs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSavedModel(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.SavedModelSchemaVersion != 0 {
		i = encodeVarintSavedModel(dAtA, i, uint64(m.SavedModelSchemaVersion))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSavedModel(dAtA []byte, offset int, v uint64) int {
	offset -= sovSavedModel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SavedModel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SavedModelSchemaVersion != 0 {
		n += 1 + sovSavedModel(uint64(m.SavedModelSchemaVersion))
	}
	if len(m.MetaGraphs) > 0 {
		for _, e := range m.MetaGraphs {
			l = e.Size()
			n += 1 + l + sovSavedModel(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSavedModel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSavedModel(x uint64) (n int) {
	return sovSavedModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SavedModel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSavedModel
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
			return fmt.Errorf("proto: SavedModel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SavedModel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SavedModelSchemaVersion", wireType)
			}
			m.SavedModelSchemaVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSavedModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SavedModelSchemaVersion |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaGraphs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSavedModel
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
				return ErrInvalidLengthSavedModel
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSavedModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetaGraphs = append(m.MetaGraphs, &MetaGraphDef{})
			if err := m.MetaGraphs[len(m.MetaGraphs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSavedModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSavedModel
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
func skipSavedModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSavedModel
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
					return 0, ErrIntOverflowSavedModel
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
					return 0, ErrIntOverflowSavedModel
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
				return 0, ErrInvalidLengthSavedModel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSavedModel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSavedModel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSavedModel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSavedModel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSavedModel = fmt.Errorf("proto: unexpected end of group")
)
