// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/protobuf/cluster.proto

package for_core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
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

// Defines a single job in a TensorFlow cluster.
type JobDef struct {
	// The name of this job.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Mapping from task ID to "hostname:port" string.
	//
	// If the `name` field contains "worker", and the `tasks` map contains a
	// mapping from 7 to "example.org:2222", then the device prefix
	// "/job:worker/task:7" will be assigned to "example.org:2222".
	Tasks                map[int32]string `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *JobDef) Reset()         { *m = JobDef{} }
func (m *JobDef) String() string { return proto.CompactTextString(m) }
func (*JobDef) ProtoMessage()    {}
func (*JobDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ea47a9615190cff, []int{0}
}
func (m *JobDef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JobDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *JobDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobDef.Merge(m, src)
}
func (m *JobDef) XXX_Size() int {
	return m.Size()
}
func (m *JobDef) XXX_DiscardUnknown() {
	xxx_messageInfo_JobDef.DiscardUnknown(m)
}

var xxx_messageInfo_JobDef proto.InternalMessageInfo

func (m *JobDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobDef) GetTasks() map[int32]string {
	if m != nil {
		return m.Tasks
	}
	return nil
}

// Defines a TensorFlow cluster as a set of jobs.
type ClusterDef struct {
	// The jobs that comprise the cluster.
	Job                  []*JobDef `protobuf:"bytes,1,rep,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ClusterDef) Reset()         { *m = ClusterDef{} }
func (m *ClusterDef) String() string { return proto.CompactTextString(m) }
func (*ClusterDef) ProtoMessage()    {}
func (*ClusterDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ea47a9615190cff, []int{1}
}
func (m *ClusterDef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ClusterDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterDef.Merge(m, src)
}
func (m *ClusterDef) XXX_Size() int {
	return m.Size()
}
func (m *ClusterDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterDef.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterDef proto.InternalMessageInfo

func (m *ClusterDef) GetJob() []*JobDef {
	if m != nil {
		return m.Job
	}
	return nil
}

func init() {
	proto.RegisterType((*JobDef)(nil), "tensorflow.JobDef")
	proto.RegisterMapType((map[int32]string)(nil), "tensorflow.JobDef.TasksEntry")
	proto.RegisterType((*ClusterDef)(nil), "tensorflow.ClusterDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/cluster.proto", fileDescriptor_8ea47a9615190cff)
}

var fileDescriptor_8ea47a9615190cff = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0xd2, 0x03, 0x0b, 0x08,
	0x71, 0x21, 0xd4, 0x29, 0x75, 0x33, 0x72, 0xb1, 0x79, 0xe5, 0x27, 0xb9, 0xa4, 0xa6, 0x09, 0x09,
	0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42,
	0xc6, 0x5c, 0xac, 0x25, 0x89, 0xc5, 0xd9, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xb2,
	0x7a, 0x08, 0xad, 0x7a, 0x10, 0x6d, 0x7a, 0x21, 0x20, 0x79, 0xd7, 0xbc, 0x92, 0xa2, 0xca, 0x20,
	0x88, 0x5a, 0x29, 0x0b, 0x2e, 0x2e, 0x84, 0xa0, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0x25, 0xd8,
	0x54, 0xd6, 0x20, 0x10, 0x53, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09,
	0x6c, 0x13, 0x84, 0x63, 0xc5, 0x64, 0xc1, 0xa8, 0x64, 0xc4, 0xc5, 0xe5, 0x0c, 0x71, 0x2a, 0xc8,
	0x41, 0x2a, 0x5c, 0xcc, 0x59, 0xf9, 0x49, 0x12, 0x8c, 0x60, 0xab, 0x85, 0x30, 0xad, 0x0e, 0x02,
	0x49, 0x3b, 0x4d, 0x67, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x1b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0xe4, 0x92, 0xca, 0x2f, 0x4a, 0x47, 0x56, 0x9e, 0x92, 0x59, 0x5c,
	0x52, 0x54, 0x9a, 0x57, 0x92, 0x99, 0x9b, 0xea, 0xc4, 0x0b, 0xb5, 0x20, 0x00, 0x14, 0x14, 0xc5,
	0x01, 0x8c, 0x51, 0xa1, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x48,
	0x01, 0x88, 0x9d, 0x99, 0x9e, 0x8f, 0x16, 0xb2, 0x69, 0xf9, 0x45, 0xf1, 0x20, 0x91, 0x78, 0xb0,
	0x48, 0x71, 0x7c, 0x7a, 0x3e, 0x84, 0xf5, 0x83, 0x91, 0x31, 0x89, 0x0d, 0xcc, 0x32, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xeb, 0x74, 0x4a, 0x7f, 0x98, 0x01, 0x00, 0x00,
}

func (m *JobDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobDef) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JobDef) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tasks) > 0 {
		keysForTasks := make([]int32, 0, len(m.Tasks))
		for k := range m.Tasks {
			keysForTasks = append(keysForTasks, int32(k))
		}
		github_com_gogo_protobuf_sortkeys.Int32s(keysForTasks)
		for iNdEx := len(keysForTasks) - 1; iNdEx >= 0; iNdEx-- {
			v := m.Tasks[int32(keysForTasks[iNdEx])]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintCluster(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i = encodeVarintCluster(dAtA, i, uint64(keysForTasks[iNdEx]))
			i--
			dAtA[i] = 0x8
			i = encodeVarintCluster(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCluster(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClusterDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterDef) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterDef) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Job) > 0 {
		for iNdEx := len(m.Job) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Job[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCluster(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCluster(dAtA []byte, offset int, v uint64) int {
	offset -= sovCluster(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *JobDef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCluster(uint64(l))
	}
	if len(m.Tasks) > 0 {
		for k, v := range m.Tasks {
			_ = k
			_ = v
			mapEntrySize := 1 + sovCluster(uint64(k)) + 1 + len(v) + sovCluster(uint64(len(v)))
			n += mapEntrySize + 1 + sovCluster(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClusterDef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Job) > 0 {
		for _, e := range m.Job {
			l = e.Size()
			n += 1 + l + sovCluster(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCluster(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCluster(x uint64) (n int) {
	return sovCluster(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *JobDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
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
			return fmt.Errorf("proto: JobDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JobDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tasks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tasks == nil {
				m.Tasks = make(map[int32]string)
			}
			var mapkey int32
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCluster
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCluster
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCluster
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthCluster
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthCluster
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCluster(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthCluster
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Tasks[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCluster
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
func (m *ClusterDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
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
			return fmt.Errorf("proto: ClusterDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Job", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Job = append(m.Job, &JobDef{})
			if err := m.Job[len(m.Job)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCluster
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
func skipCluster(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCluster
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
					return 0, ErrIntOverflowCluster
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
					return 0, ErrIntOverflowCluster
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
				return 0, ErrInvalidLengthCluster
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCluster
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCluster
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCluster        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCluster          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCluster = fmt.Errorf("proto: unexpected end of group")
)
