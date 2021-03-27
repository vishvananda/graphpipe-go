// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/protobuf/device_properties.proto

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

type DeviceProperties struct {
	// Device type (CPU, GPU, ...)
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Vendor (Intel, nvidia, ...)
	Vendor string `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Model (Haswell, K40, ...)
	Model string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	// Core Frequency in Mhz
	Frequency int64 `protobuf:"varint,4,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Number of cores
	NumCores int64 `protobuf:"varint,5,opt,name=num_cores,json=numCores,proto3" json:"num_cores,omitempty"`
	// Version of the tools and libraries used with this device (e.g. gcc 4.9,
	// cudnn 5.1)
	Environment map[string]string `protobuf:"bytes,6,rep,name=environment,proto3" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Number of registers per core.
	NumRegisters int64 `protobuf:"varint,7,opt,name=num_registers,json=numRegisters,proto3" json:"num_registers,omitempty"`
	// L1 cache size in bytes
	L1CacheSize int64 `protobuf:"varint,8,opt,name=l1_cache_size,json=l1CacheSize,proto3" json:"l1_cache_size,omitempty"`
	// L2 cache size in bytes
	L2CacheSize int64 `protobuf:"varint,9,opt,name=l2_cache_size,json=l2CacheSize,proto3" json:"l2_cache_size,omitempty"`
	// L3 cache size in bytes
	L3CacheSize int64 `protobuf:"varint,10,opt,name=l3_cache_size,json=l3CacheSize,proto3" json:"l3_cache_size,omitempty"`
	// Shared memory size per multiprocessor in bytes. This field is
	// applicable to GPUs only.
	SharedMemorySizePerMultiprocessor int64 `protobuf:"varint,11,opt,name=shared_memory_size_per_multiprocessor,json=sharedMemorySizePerMultiprocessor,proto3" json:"shared_memory_size_per_multiprocessor,omitempty"`
	// Memory size in bytes
	MemorySize int64 `protobuf:"varint,12,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	// Memory bandwidth in KB/s
	Bandwidth            int64    `protobuf:"varint,13,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceProperties) Reset()         { *m = DeviceProperties{} }
func (m *DeviceProperties) String() string { return proto.CompactTextString(m) }
func (*DeviceProperties) ProtoMessage()    {}
func (*DeviceProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c4fdb3c62f9732, []int{0}
}
func (m *DeviceProperties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DeviceProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProperties.Merge(m, src)
}
func (m *DeviceProperties) XXX_Size() int {
	return m.Size()
}
func (m *DeviceProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProperties.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProperties proto.InternalMessageInfo

func (m *DeviceProperties) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DeviceProperties) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *DeviceProperties) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *DeviceProperties) GetFrequency() int64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *DeviceProperties) GetNumCores() int64 {
	if m != nil {
		return m.NumCores
	}
	return 0
}

func (m *DeviceProperties) GetEnvironment() map[string]string {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *DeviceProperties) GetNumRegisters() int64 {
	if m != nil {
		return m.NumRegisters
	}
	return 0
}

func (m *DeviceProperties) GetL1CacheSize() int64 {
	if m != nil {
		return m.L1CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetL2CacheSize() int64 {
	if m != nil {
		return m.L2CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetL3CacheSize() int64 {
	if m != nil {
		return m.L3CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetSharedMemorySizePerMultiprocessor() int64 {
	if m != nil {
		return m.SharedMemorySizePerMultiprocessor
	}
	return 0
}

func (m *DeviceProperties) GetMemorySize() int64 {
	if m != nil {
		return m.MemorySize
	}
	return 0
}

func (m *DeviceProperties) GetBandwidth() int64 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

type NamedDevice struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Properties           *DeviceProperties `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NamedDevice) Reset()         { *m = NamedDevice{} }
func (m *NamedDevice) String() string { return proto.CompactTextString(m) }
func (*NamedDevice) ProtoMessage()    {}
func (*NamedDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c4fdb3c62f9732, []int{1}
}
func (m *NamedDevice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NamedDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *NamedDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedDevice.Merge(m, src)
}
func (m *NamedDevice) XXX_Size() int {
	return m.Size()
}
func (m *NamedDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedDevice.DiscardUnknown(m)
}

var xxx_messageInfo_NamedDevice proto.InternalMessageInfo

func (m *NamedDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedDevice) GetProperties() *DeviceProperties {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceProperties)(nil), "tensorflow.DeviceProperties")
	proto.RegisterMapType((map[string]string)(nil), "tensorflow.DeviceProperties.EnvironmentEntry")
	proto.RegisterType((*NamedDevice)(nil), "tensorflow.NamedDevice")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/device_properties.proto", fileDescriptor_07c4fdb3c62f9732)
}

var fileDescriptor_07c4fdb3c62f9732 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0x6f, 0xd3, 0x3e,
	0x14, 0xc7, 0xe5, 0xb5, 0xeb, 0x6f, 0x7d, 0x59, 0xa5, 0xca, 0xfa, 0x69, 0xb2, 0x60, 0x2a, 0xa5,
	0x08, 0xa9, 0x17, 0x5a, 0xd6, 0x5e, 0x10, 0x42, 0x1c, 0x36, 0x76, 0x1c, 0x54, 0x45, 0x5c, 0xb8,
	0x58, 0x69, 0xf2, 0xda, 0x46, 0xc4, 0x76, 0xb0, 0x9d, 0x4e, 0xdd, 0x91, 0x3f, 0x8d, 0x13, 0x47,
	0x8e, 0x1c, 0x51, 0xff, 0x0a, 0x8e, 0xc8, 0x4e, 0x68, 0xb2, 0x0a, 0x71, 0xfb, 0xbe, 0x6f, 0x3e,
	0xef, 0xd9, 0xca, 0xfb, 0x1a, 0x9e, 0x5b, 0x94, 0x46, 0xe9, 0x65, 0xaa, 0x6e, 0xc7, 0x91, 0xd2,
	0x38, 0xce, 0xb4, 0xb2, 0x6a, 0x91, 0x2f, 0xc7, 0x31, 0x6e, 0x92, 0x08, 0x79, 0xa6, 0x55, 0x86,
	0xda, 0x26, 0x68, 0x46, 0xfe, 0x13, 0x85, 0xaa, 0x63, 0xf0, 0xb5, 0x09, 0xdd, 0x37, 0x9e, 0x9b,
	0xed, 0x31, 0x4a, 0xa1, 0x69, 0xb7, 0x19, 0x32, 0xd2, 0x27, 0xc3, 0xf6, 0xdc, 0x6b, 0x7a, 0x06,
	0xad, 0x0d, 0xca, 0x58, 0x69, 0x76, 0xe4, 0xdd, 0xb2, 0xa2, 0xff, 0xc3, 0xb1, 0x50, 0x31, 0xa6,
	0xac, 0xe1, 0xed, 0xa2, 0xa0, 0xe7, 0xd0, 0x5e, 0x6a, 0xfc, 0x9c, 0xa3, 0x8c, 0xb6, 0xac, 0xd9,
	0x27, 0xc3, 0xc6, 0xbc, 0x32, 0xe8, 0x43, 0x68, 0xcb, 0x5c, 0x70, 0x77, 0x5b, 0xc3, 0x8e, 0xfd,
	0xd7, 0x13, 0x99, 0x8b, 0x2b, 0x57, 0xd3, 0x77, 0x10, 0xa0, 0xdc, 0x24, 0x5a, 0x49, 0x81, 0xd2,
	0xb2, 0x56, 0xbf, 0x31, 0x0c, 0x26, 0xcf, 0x46, 0xd5, 0x9d, 0x47, 0x87, 0xf7, 0x1d, 0x5d, 0x57,
	0xfc, 0xb5, 0xb4, 0x7a, 0x3b, 0xaf, 0x4f, 0xa0, 0x4f, 0xa0, 0xe3, 0x4e, 0xd3, 0xb8, 0x4a, 0x8c,
	0x45, 0x6d, 0xd8, 0x7f, 0xfe, 0xc4, 0x53, 0x99, 0x8b, 0xf9, 0x1f, 0x8f, 0x0e, 0xa0, 0x93, 0x5e,
	0xf0, 0x28, 0x8c, 0xd6, 0xc8, 0x4d, 0x72, 0x87, 0xec, 0xc4, 0x43, 0x41, 0x7a, 0x71, 0xe5, 0xbc,
	0xf7, 0xc9, 0x1d, 0x7a, 0x66, 0x52, 0x67, 0xda, 0x25, 0x33, 0xb9, 0xcf, 0x4c, 0xeb, 0x0c, 0x94,
	0xcc, 0xb4, 0x62, 0x66, 0xf0, 0xd4, 0xac, 0x43, 0x8d, 0x31, 0x17, 0x28, 0x94, 0xde, 0x7a, 0x90,
	0x67, 0xa8, 0xb9, 0xc8, 0x53, 0x9b, 0x64, 0x5a, 0x45, 0x68, 0x8c, 0xd2, 0x2c, 0xf0, 0xbd, 0x8f,
	0x0b, 0xf8, 0xc6, 0xb3, 0x6e, 0xc0, 0x0c, 0xf5, 0xcd, 0x3d, 0x90, 0x3e, 0x82, 0xa0, 0x36, 0x8a,
	0x9d, 0xfa, 0x3e, 0x10, 0xfb, 0x0e, 0xb7, 0x8f, 0x45, 0x28, 0xe3, 0xdb, 0x24, 0xb6, 0x6b, 0xd6,
	0x29, 0xf6, 0xb1, 0x37, 0x1e, 0xbc, 0x86, 0xee, 0xe1, 0x2f, 0xa4, 0x5d, 0x68, 0x7c, 0xc2, 0x6d,
	0x19, 0x01, 0x27, 0xdd, 0xa6, 0x37, 0x61, 0x9a, 0x63, 0x19, 0x80, 0xa2, 0x78, 0x79, 0xf4, 0x82,
	0x0c, 0x38, 0x04, 0x6f, 0x43, 0x81, 0x71, 0xb1, 0x18, 0x17, 0x1f, 0x19, 0x8a, 0x7d, 0x7c, 0x9c,
	0xa6, 0xaf, 0x00, 0xaa, 0x1c, 0xfa, 0x09, 0xc1, 0xe4, 0xfc, 0x5f, 0x4b, 0x9d, 0xd7, 0xf8, 0xcb,
	0x2f, 0xe4, 0xdb, 0xae, 0x47, 0xbe, 0xef, 0x7a, 0xe4, 0xc7, 0xae, 0x47, 0x7e, 0xee, 0x7a, 0xe4,
	0xf2, 0xec, 0xb0, 0x61, 0xe6, 0xb2, 0x6d, 0x3e, 0x7e, 0x58, 0x25, 0x76, 0x9d, 0x2f, 0x46, 0x91,
	0x12, 0xe3, 0xda, 0xcb, 0xf8, 0xbb, 0x5c, 0xa9, 0x83, 0x27, 0xb3, 0x54, 0xda, 0xc7, 0x92, 0x7b,
	0xc7, 0xf0, 0x95, 0x2a, 0xd4, 0x2f, 0x42, 0x16, 0x2d, 0xaf, 0xa6, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x8f, 0xb1, 0x1a, 0xc5, 0x71, 0x03, 0x00, 0x00,
}

func (m *DeviceProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceProperties) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceProperties) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Bandwidth != 0 {
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.Bandwidth))
		i--
		dAtA[i] = 0x68
	}
	if m.MemorySize != 0 {
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.MemorySize))
		i--
		dAtA[i] = 0x60
	}
	if m.SharedMemorySizePerMultiprocessor != 0 {
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.SharedMemorySizePerMultiprocessor))
		i--
		dAtA[i] = 0x58
	}
	if m.L3CacheSize != 0 {
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.L3CacheSize))
		i--
		dAtA[i] = 0x50
	}
	if m.L2CacheSize != 0 {
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.L2CacheSize))
		i--
		dAtA[i] = 0x48
	}
	if m.L1CacheSize != 0 {
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.L1CacheSize))
		i--
		dAtA[i] = 0x40
	}
	if m.NumRegisters != 0 {
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.NumRegisters))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Environment) > 0 {
		keysForEnvironment := make([]string, 0, len(m.Environment))
		for k := range m.Environment {
			keysForEnvironment = append(keysForEnvironment, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForEnvironment)
		for iNdEx := len(keysForEnvironment) - 1; iNdEx >= 0; iNdEx-- {
			v := m.Environment[string(keysForEnvironment[iNdEx])]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintDeviceProperties(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(keysForEnvironment[iNdEx])
			copy(dAtA[i:], keysForEnvironment[iNdEx])
			i = encodeVarintDeviceProperties(dAtA, i, uint64(len(keysForEnvironment[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintDeviceProperties(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.NumCores != 0 {
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.NumCores))
		i--
		dAtA[i] = 0x28
	}
	if m.Frequency != 0 {
		i = encodeVarintDeviceProperties(dAtA, i, uint64(m.Frequency))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Model) > 0 {
		i -= len(m.Model)
		copy(dAtA[i:], m.Model)
		i = encodeVarintDeviceProperties(dAtA, i, uint64(len(m.Model)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Vendor) > 0 {
		i -= len(m.Vendor)
		copy(dAtA[i:], m.Vendor)
		i = encodeVarintDeviceProperties(dAtA, i, uint64(len(m.Vendor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintDeviceProperties(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NamedDevice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamedDevice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NamedDevice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Properties != nil {
		{
			size, err := m.Properties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDeviceProperties(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDeviceProperties(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeviceProperties(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeviceProperties(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceProperties) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovDeviceProperties(uint64(l))
	}
	l = len(m.Vendor)
	if l > 0 {
		n += 1 + l + sovDeviceProperties(uint64(l))
	}
	l = len(m.Model)
	if l > 0 {
		n += 1 + l + sovDeviceProperties(uint64(l))
	}
	if m.Frequency != 0 {
		n += 1 + sovDeviceProperties(uint64(m.Frequency))
	}
	if m.NumCores != 0 {
		n += 1 + sovDeviceProperties(uint64(m.NumCores))
	}
	if len(m.Environment) > 0 {
		for k, v := range m.Environment {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDeviceProperties(uint64(len(k))) + 1 + len(v) + sovDeviceProperties(uint64(len(v)))
			n += mapEntrySize + 1 + sovDeviceProperties(uint64(mapEntrySize))
		}
	}
	if m.NumRegisters != 0 {
		n += 1 + sovDeviceProperties(uint64(m.NumRegisters))
	}
	if m.L1CacheSize != 0 {
		n += 1 + sovDeviceProperties(uint64(m.L1CacheSize))
	}
	if m.L2CacheSize != 0 {
		n += 1 + sovDeviceProperties(uint64(m.L2CacheSize))
	}
	if m.L3CacheSize != 0 {
		n += 1 + sovDeviceProperties(uint64(m.L3CacheSize))
	}
	if m.SharedMemorySizePerMultiprocessor != 0 {
		n += 1 + sovDeviceProperties(uint64(m.SharedMemorySizePerMultiprocessor))
	}
	if m.MemorySize != 0 {
		n += 1 + sovDeviceProperties(uint64(m.MemorySize))
	}
	if m.Bandwidth != 0 {
		n += 1 + sovDeviceProperties(uint64(m.Bandwidth))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NamedDevice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDeviceProperties(uint64(l))
	}
	if m.Properties != nil {
		l = m.Properties.Size()
		n += 1 + l + sovDeviceProperties(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDeviceProperties(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeviceProperties(x uint64) (n int) {
	return sovDeviceProperties(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceProperties
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
			return fmt.Errorf("proto: DeviceProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
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
				return ErrInvalidLengthDeviceProperties
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vendor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
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
				return ErrInvalidLengthDeviceProperties
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vendor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Model", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
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
				return ErrInvalidLengthDeviceProperties
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Model = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frequency", wireType)
			}
			m.Frequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Frequency |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumCores", wireType)
			}
			m.NumCores = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumCores |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Environment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
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
				return ErrInvalidLengthDeviceProperties
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Environment == nil {
				m.Environment = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDeviceProperties
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
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDeviceProperties
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthDeviceProperties
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthDeviceProperties
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDeviceProperties
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
						return ErrInvalidLengthDeviceProperties
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthDeviceProperties
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDeviceProperties(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthDeviceProperties
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Environment[mapkey] = mapvalue
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRegisters", wireType)
			}
			m.NumRegisters = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRegisters |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field L1CacheSize", wireType)
			}
			m.L1CacheSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.L1CacheSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field L2CacheSize", wireType)
			}
			m.L2CacheSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.L2CacheSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field L3CacheSize", wireType)
			}
			m.L3CacheSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.L3CacheSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharedMemorySizePerMultiprocessor", wireType)
			}
			m.SharedMemorySizePerMultiprocessor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SharedMemorySizePerMultiprocessor |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemorySize", wireType)
			}
			m.MemorySize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemorySize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bandwidth", wireType)
			}
			m.Bandwidth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bandwidth |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceProperties(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeviceProperties
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
func (m *NamedDevice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceProperties
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
			return fmt.Errorf("proto: NamedDevice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamedDevice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
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
				return ErrInvalidLengthDeviceProperties
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceProperties
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
				return ErrInvalidLengthDeviceProperties
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Properties == nil {
				m.Properties = &DeviceProperties{}
			}
			if err := m.Properties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceProperties(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeviceProperties
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
func skipDeviceProperties(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceProperties
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
					return 0, ErrIntOverflowDeviceProperties
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
					return 0, ErrIntOverflowDeviceProperties
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
				return 0, ErrInvalidLengthDeviceProperties
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeviceProperties
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeviceProperties
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeviceProperties        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceProperties          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeviceProperties = fmt.Errorf("proto: unexpected end of group")
)
