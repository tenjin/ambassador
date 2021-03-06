// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/lua/v3alpha/lua.proto

package envoy_config_filter_http_lua_v3alpha

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Lua struct {
	// The Lua code that Envoy will execute. This can be a very small script that
	// further loads code from disk if desired. Note that if JSON configuration is used, the code must
	// be properly escaped. YAML configuration may be easier to read since YAML supports multi-line
	// strings so complex scripts can be easily expressed inline in the configuration.
	InlineCode           string   `protobuf:"bytes,1,opt,name=inline_code,json=inlineCode,proto3" json:"inline_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lua) Reset()         { *m = Lua{} }
func (m *Lua) String() string { return proto.CompactTextString(m) }
func (*Lua) ProtoMessage()    {}
func (*Lua) Descriptor() ([]byte, []int) {
	return fileDescriptor_d72dd20a329954ba, []int{0}
}
func (m *Lua) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lua) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lua.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lua) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lua.Merge(m, src)
}
func (m *Lua) XXX_Size() int {
	return m.Size()
}
func (m *Lua) XXX_DiscardUnknown() {
	xxx_messageInfo_Lua.DiscardUnknown(m)
}

var xxx_messageInfo_Lua proto.InternalMessageInfo

func (m *Lua) GetInlineCode() string {
	if m != nil {
		return m.InlineCode
	}
	return ""
}

func init() {
	proto.RegisterType((*Lua)(nil), "envoy.config.filter.http.lua.v3alpha.Lua")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/lua/v3alpha/lua.proto", fileDescriptor_d72dd20a329954ba)
}

var fileDescriptor_d72dd20a329954ba = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0xcf, 0x29, 0x4d, 0xd4, 0x2f, 0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48,
	0x04, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0xc0, 0xea, 0xf5, 0x20, 0xea, 0xf5,
	0x20, 0xea, 0xf5, 0x40, 0xea, 0xf5, 0x40, 0x6a, 0xa0, 0xea, 0xa5, 0xc4, 0xcb, 0x12, 0x73, 0x32,
	0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x76, 0x25, 0x43, 0x2e, 0x66, 0x9f, 0xd2, 0x44,
	0x21, 0x2d, 0x2e, 0xee, 0xcc, 0xbc, 0x9c, 0xcc, 0xbc, 0xd4, 0xf8, 0xe4, 0xfc, 0x94, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xce, 0x5d, 0x2f, 0x0f, 0x30, 0xb3, 0x14, 0x31, 0x29, 0x30,
	0x06, 0x71, 0x41, 0x64, 0x9d, 0xf3, 0x53, 0x52, 0x9d, 0x7c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0x46, 0x2e, 0xa3, 0xcc, 0x7c, 0x88, 0x93, 0x0b, 0x8a, 0xf2,
	0x2b, 0x2a, 0xf5, 0x88, 0x71, 0x8d, 0x13, 0x87, 0x4f, 0x69, 0x62, 0x00, 0xc8, 0xfa, 0x00, 0xc6,
	0x24, 0x36, 0xb0, 0x3b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x66, 0x38, 0x7d, 0xf8,
	0x00, 0x00, 0x00,
}

func (m *Lua) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lua) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lua) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.InlineCode) > 0 {
		i -= len(m.InlineCode)
		copy(dAtA[i:], m.InlineCode)
		i = encodeVarintLua(dAtA, i, uint64(len(m.InlineCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLua(dAtA []byte, offset int, v uint64) int {
	offset -= sovLua(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lua) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InlineCode)
	if l > 0 {
		n += 1 + l + sovLua(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLua(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLua(x uint64) (n int) {
	return sovLua(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lua) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLua
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
			return fmt.Errorf("proto: Lua: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lua: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InlineCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLua
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
				return ErrInvalidLengthLua
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLua
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InlineCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLua(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLua
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLua
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
func skipLua(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLua
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
					return 0, ErrIntOverflowLua
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
					return 0, ErrIntOverflowLua
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
				return 0, ErrInvalidLengthLua
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLua
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLua
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
				next, err := skipLua(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLua
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
	ErrInvalidLengthLua = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLua   = fmt.Errorf("proto: integer overflow")
)
