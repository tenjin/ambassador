// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/buffer/v2/buffer.proto

package envoy_config_filter_http_buffer_v2

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type Buffer struct {
	// The maximum request size that the filter will buffer before the connection
	// manager will stop buffering and return a 413 response.
	MaxRequestBytes      *types.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Buffer) Reset()         { *m = Buffer{} }
func (m *Buffer) String() string { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()    {}
func (*Buffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e782fc75ce4c789f, []int{0}
}
func (m *Buffer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Buffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Buffer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Buffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buffer.Merge(m, src)
}
func (m *Buffer) XXX_Size() int {
	return m.Size()
}
func (m *Buffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Buffer.DiscardUnknown(m)
}

var xxx_messageInfo_Buffer proto.InternalMessageInfo

func (m *Buffer) GetMaxRequestBytes() *types.UInt32Value {
	if m != nil {
		return m.MaxRequestBytes
	}
	return nil
}

type BufferPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*BufferPerRoute_Disabled
	//	*BufferPerRoute_Buffer
	Override             isBufferPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BufferPerRoute) Reset()         { *m = BufferPerRoute{} }
func (m *BufferPerRoute) String() string { return proto.CompactTextString(m) }
func (*BufferPerRoute) ProtoMessage()    {}
func (*BufferPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_e782fc75ce4c789f, []int{1}
}
func (m *BufferPerRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BufferPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BufferPerRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BufferPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferPerRoute.Merge(m, src)
}
func (m *BufferPerRoute) XXX_Size() int {
	return m.Size()
}
func (m *BufferPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_BufferPerRoute proto.InternalMessageInfo

type isBufferPerRoute_Override interface {
	isBufferPerRoute_Override()
	MarshalTo([]byte) (int, error)
	Size() int
}

type BufferPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}
type BufferPerRoute_Buffer struct {
	Buffer *Buffer `protobuf:"bytes,2,opt,name=buffer,proto3,oneof"`
}

func (*BufferPerRoute_Disabled) isBufferPerRoute_Override() {}
func (*BufferPerRoute_Buffer) isBufferPerRoute_Override()   {}

func (m *BufferPerRoute) GetOverride() isBufferPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *BufferPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*BufferPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *BufferPerRoute) GetBuffer() *Buffer {
	if x, ok := m.GetOverride().(*BufferPerRoute_Buffer); ok {
		return x.Buffer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BufferPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BufferPerRoute_Disabled)(nil),
		(*BufferPerRoute_Buffer)(nil),
	}
}

func init() {
	proto.RegisterType((*Buffer)(nil), "envoy.config.filter.http.buffer.v2.Buffer")
	proto.RegisterType((*BufferPerRoute)(nil), "envoy.config.filter.http.buffer.v2.BufferPerRoute")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/buffer/v2/buffer.proto", fileDescriptor_e782fc75ce4c789f)
}

var fileDescriptor_e782fc75ce4c789f = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0xeb, 0xf4, 0x8f, 0x82, 0x2b, 0x41, 0x9b, 0x85, 0xaa, 0x42, 0x51, 0xd5, 0x05, 0xd4,
	0xc1, 0x46, 0xe9, 0x0d, 0x3c, 0x01, 0x12, 0x52, 0x15, 0x04, 0x6b, 0xe5, 0x34, 0x2f, 0x25, 0x28,
	0xad, 0x83, 0xe3, 0x84, 0xf6, 0x0a, 0x1c, 0x82, 0x73, 0x20, 0xa6, 0x8e, 0x8c, 0x1c, 0x01, 0x65,
	0xeb, 0x2d, 0x50, 0xec, 0x94, 0x15, 0xb6, 0x27, 0x3d, 0xff, 0x7e, 0x9f, 0xdf, 0x87, 0x29, 0xac,
	0x0b, 0xb1, 0xa5, 0x0b, 0xb1, 0x8e, 0xe2, 0x25, 0x8d, 0xe2, 0x44, 0x81, 0xa4, 0x8f, 0x4a, 0xa5,
	0x34, 0xc8, 0xa3, 0x08, 0x24, 0x2d, 0xbc, 0x7a, 0x22, 0xa9, 0x14, 0x4a, 0x38, 0x63, 0x0d, 0x10,
	0x03, 0x10, 0x03, 0x90, 0x0a, 0x20, 0xf5, 0xb3, 0xc2, 0x1b, 0xba, 0x4b, 0x21, 0x96, 0x09, 0x50,
	0x4d, 0x04, 0x79, 0x44, 0x5f, 0x24, 0x4f, 0x53, 0x90, 0x99, 0x71, 0x0c, 0x4f, 0x0b, 0x9e, 0xc4,
	0x21, 0x57, 0x40, 0x0f, 0x83, 0x59, 0x8c, 0x17, 0xb8, 0xc3, 0xb4, 0xc5, 0xb9, 0xc3, 0xfd, 0x15,
	0xdf, 0xcc, 0x25, 0x3c, 0xe7, 0x90, 0xa9, 0x79, 0xb0, 0x55, 0x90, 0x0d, 0xd0, 0x08, 0x5d, 0x74,
	0xbd, 0x33, 0x62, 0xf4, 0xe4, 0xa0, 0x27, 0xf7, 0xd7, 0x6b, 0x35, 0xf5, 0x1e, 0x78, 0x92, 0x03,
	0x3b, 0xfa, 0xd8, 0xef, 0x9a, 0xad, 0x89, 0x35, 0x6a, 0xf8, 0x27, 0x2b, 0xbe, 0xf1, 0x8d, 0x80,
	0x55, 0xfc, 0x4d, 0xcb, 0xb6, 0x7a, 0xcd, 0xf1, 0x1b, 0xc2, 0xc7, 0x26, 0x65, 0x06, 0xd2, 0x17,
	0xb9, 0x02, 0xe7, 0x1c, 0xdb, 0x61, 0x9c, 0xf1, 0x20, 0x81, 0x50, 0x87, 0xd8, 0xb5, 0xe6, 0xc9,
	0xb2, 0xd1, 0x55, 0xc3, 0xff, 0x5d, 0x3a, 0x33, 0xdc, 0x31, 0x67, 0x0e, 0x2c, 0xfd, 0x97, 0x09,
	0xf9, 0xbb, 0x0e, 0x62, 0xc2, 0x18, 0xae, 0x94, 0xed, 0x57, 0x64, 0xf5, 0x2a, 0x67, 0xed, 0x61,
	0x7d, 0x6c, 0x8b, 0x02, 0xa4, 0x8c, 0x43, 0x70, 0xda, 0xef, 0xfb, 0x5d, 0x13, 0xb1, 0xdb, 0xcf,
	0xd2, 0x45, 0x5f, 0xa5, 0x8b, 0xbe, 0x4b, 0x17, 0xe1, 0xcb, 0x58, 0x98, 0x90, 0x54, 0x8a, 0xcd,
	0xf6, 0x1f, 0x79, 0xac, 0x5b, 0x5f, 0x57, 0xd5, 0x33, 0x43, 0x41, 0x47, 0xf7, 0x34, 0xfd, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x74, 0xb8, 0xee, 0xc8, 0xeb, 0x01, 0x00, 0x00,
}

func (m *Buffer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Buffer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Buffer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxRequestBytes != nil {
		{
			size, err := m.MaxRequestBytes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBuffer(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BufferPerRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BufferPerRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BufferPerRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Override != nil {
		{
			size := m.Override.Size()
			i -= size
			if _, err := m.Override.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *BufferPerRoute_Disabled) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *BufferPerRoute_Disabled) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.Disabled {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}
func (m *BufferPerRoute_Buffer) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *BufferPerRoute_Buffer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Buffer != nil {
		{
			size, err := m.Buffer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBuffer(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintBuffer(dAtA []byte, offset int, v uint64) int {
	offset -= sovBuffer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Buffer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxRequestBytes != nil {
		l = m.MaxRequestBytes.Size()
		n += 1 + l + sovBuffer(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BufferPerRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Override != nil {
		n += m.Override.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BufferPerRoute_Disabled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *BufferPerRoute_Buffer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Buffer != nil {
		l = m.Buffer.Size()
		n += 1 + l + sovBuffer(uint64(l))
	}
	return n
}

func sovBuffer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBuffer(x uint64) (n int) {
	return sovBuffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Buffer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuffer
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
			return fmt.Errorf("proto: Buffer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Buffer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequestBytes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBuffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRequestBytes == nil {
				m.MaxRequestBytes = &types.UInt32Value{}
			}
			if err := m.MaxRequestBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBuffer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBuffer
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
func (m *BufferPerRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuffer
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
			return fmt.Errorf("proto: BufferPerRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BufferPerRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
			b := bool(v != 0)
			m.Override = &BufferPerRoute_Disabled{b}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buffer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBuffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Buffer{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Override = &BufferPerRoute_Buffer{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBuffer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBuffer
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
func skipBuffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBuffer
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
					return 0, ErrIntOverflowBuffer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBuffer
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
				return 0, ErrInvalidLengthBuffer
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthBuffer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBuffer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBuffer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthBuffer
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBuffer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBuffer   = fmt.Errorf("proto: integer overflow")
)
