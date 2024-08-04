// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine_chain/jklmint/minted_block.proto

package types

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

type MintedBlock struct {
	Height int64  `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Minted int64  `protobuf:"varint,2,opt,name=minted,proto3" json:"minted,omitempty"`
	Denom  string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *MintedBlock) Reset()         { *m = MintedBlock{} }
func (m *MintedBlock) String() string { return proto.CompactTextString(m) }
func (*MintedBlock) ProtoMessage()    {}
func (*MintedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_2708e6d79ccb3843, []int{0}
}
func (m *MintedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintedBlock.Merge(m, src)
}
func (m *MintedBlock) XXX_Size() int {
	return m.Size()
}
func (m *MintedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_MintedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_MintedBlock proto.InternalMessageInfo

func (m *MintedBlock) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *MintedBlock) GetMinted() int64 {
	if m != nil {
		return m.Minted
	}
	return 0
}

func (m *MintedBlock) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*MintedBlock)(nil), "canine_chain.jklmint.MintedBlock")
}

func init() {
	proto.RegisterFile("canine_chain/jklmint/minted_block.proto", fileDescriptor_2708e6d79ccb3843)
}

var fileDescriptor_2708e6d79ccb3843 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0x4e, 0xcc, 0xcb,
	0xcc, 0x4b, 0x8d, 0x4f, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0xca, 0xce, 0xc9, 0xcd, 0xcc, 0x2b,
	0xd1, 0x07, 0x11, 0xa9, 0x29, 0xf1, 0x49, 0x39, 0xf9, 0xc9, 0xd9, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x22, 0xc8, 0x0a, 0xf5, 0xa0, 0x0a, 0x95, 0x82, 0xb9, 0xb8, 0x7d, 0xc1, 0x6a, 0x9d,
	0x40, 0x4a, 0x85, 0xc4, 0xb8, 0xd8, 0x32, 0x52, 0x33, 0xd3, 0x33, 0x4a, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x98, 0x83, 0xa0, 0x3c, 0x90, 0x38, 0xc4, 0x48, 0x09, 0x26, 0x88, 0x38, 0x84, 0x27, 0x24,
	0xc2, 0xc5, 0x9a, 0x92, 0x9a, 0x97, 0x9f, 0x2b, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1,
	0x38, 0xf9, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x51, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x56, 0x62, 0x72, 0x76, 0x62, 0x8e, 0x4f,
	0x62, 0x52, 0xb1, 0x3e, 0xc4, 0x69, 0xba, 0x10, 0x3f, 0x54, 0xc0, 0x7d, 0x51, 0x52, 0x59, 0x90,
	0x5a, 0x9c, 0xc4, 0x06, 0x76, 0xbf, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x0c, 0x8f, 0x38,
	0xea, 0x00, 0x00, 0x00,
}

func (m *MintedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintMintedBlock(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Minted != 0 {
		i = encodeVarintMintedBlock(dAtA, i, uint64(m.Minted))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintMintedBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMintedBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovMintedBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MintedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovMintedBlock(uint64(m.Height))
	}
	if m.Minted != 0 {
		n += 1 + sovMintedBlock(uint64(m.Minted))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovMintedBlock(uint64(l))
	}
	return n
}

func sovMintedBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMintedBlock(x uint64) (n int) {
	return sovMintedBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MintedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMintedBlock
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
			return fmt.Errorf("proto: MintedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintedBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minted", wireType)
			}
			m.Minted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintedBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Minted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintedBlock
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
				return ErrInvalidLengthMintedBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMintedBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMintedBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMintedBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMintedBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMintedBlock
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
					return 0, ErrIntOverflowMintedBlock
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
					return 0, ErrIntOverflowMintedBlock
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
				return 0, ErrInvalidLengthMintedBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMintedBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMintedBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMintedBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMintedBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMintedBlock = fmt.Errorf("proto: unexpected end of group")
)
