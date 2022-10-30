// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine-chain/dsig/user_uploads.proto

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

type UserUploads struct {
	Fid       string `protobuf:"bytes,1,opt,name=fid,proto3" json:"fid,omitempty"`
	Cid       string `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (m *UserUploads) Reset()         { *m = UserUploads{} }
func (m *UserUploads) String() string { return proto.CompactTextString(m) }
func (*UserUploads) ProtoMessage()    {}
func (*UserUploads) Descriptor() ([]byte, []int) {
	return fileDescriptor_4294eb12daef75fa, []int{0}
}
func (m *UserUploads) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserUploads) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserUploads.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserUploads) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUploads.Merge(m, src)
}
func (m *UserUploads) XXX_Size() int {
	return m.Size()
}
func (m *UserUploads) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUploads.DiscardUnknown(m)
}

var xxx_messageInfo_UserUploads proto.InternalMessageInfo

func (m *UserUploads) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

func (m *UserUploads) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *UserUploads) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*UserUploads)(nil), "jackaldao.canine.dsig.UserUploads")
}

func init() {
	proto.RegisterFile("canine-chain/dsig/user_uploads.proto", fileDescriptor_4294eb12daef75fa)
}

var fileDescriptor_4294eb12daef75fa = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x4e, 0xcc, 0xcb,
	0xcc, 0x4b, 0xd5, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x29, 0xce, 0x4c, 0xd7, 0x2f, 0x2d,
	0x4e, 0x2d, 0x8a, 0x2f, 0x2d, 0xc8, 0xc9, 0x4f, 0x4c, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0xcd, 0x4a, 0x4c, 0xce, 0x4e, 0xcc, 0x49, 0x49, 0xcc, 0xd7, 0x83, 0xa8, 0xd7, 0x03,
	0xa9, 0x54, 0xf2, 0xe7, 0xe2, 0x0e, 0x2d, 0x4e, 0x2d, 0x0a, 0x85, 0xa8, 0x15, 0x12, 0xe0, 0x62,
	0x4e, 0xcb, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x41, 0x22, 0xc9, 0x99,
	0x29, 0x12, 0x4c, 0x10, 0x91, 0xe4, 0xcc, 0x14, 0x21, 0x19, 0x2e, 0xce, 0xe4, 0xa2, 0xd4, 0xc4,
	0x92, 0xd4, 0x14, 0xc7, 0x12, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x84, 0x80, 0x93, 0xf3,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c,
	0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x69, 0xa6, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x1c, 0xa3, 0x9b, 0x92, 0x98, 0xaf, 0x0f, 0x71, 0x8d,
	0x7e, 0x05, 0xc4, 0xe5, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x37, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x91, 0x06, 0xe8, 0xdb, 0x00, 0x00, 0x00,
}

func (m *UserUploads) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserUploads) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserUploads) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintUserUploads(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintUserUploads(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Fid) > 0 {
		i -= len(m.Fid)
		copy(dAtA[i:], m.Fid)
		i = encodeVarintUserUploads(dAtA, i, uint64(len(m.Fid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUserUploads(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserUploads(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserUploads) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Fid)
	if l > 0 {
		n += 1 + l + sovUserUploads(uint64(l))
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovUserUploads(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovUserUploads(uint64(m.CreatedAt))
	}
	return n
}

func sovUserUploads(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserUploads(x uint64) (n int) {
	return sovUserUploads(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserUploads) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserUploads
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
			return fmt.Errorf("proto: UserUploads: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserUploads: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserUploads
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
				return ErrInvalidLengthUserUploads
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserUploads
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserUploads
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
				return ErrInvalidLengthUserUploads
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserUploads
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserUploads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUserUploads(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserUploads
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
func skipUserUploads(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserUploads
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
					return 0, ErrIntOverflowUserUploads
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
					return 0, ErrIntOverflowUserUploads
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
				return 0, ErrInvalidLengthUserUploads
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserUploads
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserUploads
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserUploads        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserUploads          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserUploads = fmt.Errorf("proto: unexpected end of group")
)
