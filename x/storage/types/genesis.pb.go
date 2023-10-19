// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine_chain/storage/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the storage module's genesis state.
type GenesisState struct {
	Params              Params               `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FileList            []UnifiedFile        `protobuf:"bytes,2,rep,name=file_list,json=fileList,proto3" json:"file_list"`
	ProvidersList       []Providers          `protobuf:"bytes,3,rep,name=providers_list,json=providersList,proto3" json:"providers_list"`
	PaymentInfoList     []StoragePaymentInfo `protobuf:"bytes,4,rep,name=payment_info_list,json=paymentInfoList,proto3" json:"payment_info_list"`
	CollateralList      []Collateral         `protobuf:"bytes,5,rep,name=collateral_list,json=collateralList,proto3" json:"collateral_list"`
	ActiveProvidersList []ActiveProviders    `protobuf:"bytes,6,rep,name=active_providers_list,json=activeProvidersList,proto3" json:"active_providers_list"`
	ReportForms         []ReportForm         `protobuf:"bytes,7,rep,name=report_forms,json=reportForms,proto3" json:"report_forms"`
	AttestForms         []AttestationForm    `protobuf:"bytes,8,rep,name=attest_forms,json=attestForms,proto3" json:"attest_forms"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52da9fdb04c8b35, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFileList() []UnifiedFile {
	if m != nil {
		return m.FileList
	}
	return nil
}

func (m *GenesisState) GetProvidersList() []Providers {
	if m != nil {
		return m.ProvidersList
	}
	return nil
}

func (m *GenesisState) GetPaymentInfoList() []StoragePaymentInfo {
	if m != nil {
		return m.PaymentInfoList
	}
	return nil
}

func (m *GenesisState) GetCollateralList() []Collateral {
	if m != nil {
		return m.CollateralList
	}
	return nil
}

func (m *GenesisState) GetActiveProvidersList() []ActiveProviders {
	if m != nil {
		return m.ActiveProvidersList
	}
	return nil
}

func (m *GenesisState) GetReportForms() []ReportForm {
	if m != nil {
		return m.ReportForms
	}
	return nil
}

func (m *GenesisState) GetAttestForms() []AttestationForm {
	if m != nil {
		return m.AttestForms
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "canine_chain.storage.GenesisState")
}

func init() {
	proto.RegisterFile("canine_chain/storage/genesis.proto", fileDescriptor_a52da9fdb04c8b35)
}

var fileDescriptor_a52da9fdb04c8b35 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x86, 0xe3, 0xaf, 0x6d, 0xbe, 0xb2, 0x09, 0xad, 0x30, 0x45, 0x8a, 0x22, 0xe4, 0xa6, 0x15,
	0x88, 0x5c, 0xb0, 0xa5, 0x70, 0xe3, 0x46, 0x41, 0x45, 0x95, 0x22, 0xa8, 0x5a, 0x71, 0xe9, 0xc5,
	0x9a, 0xac, 0xd7, 0xee, 0x82, 0xbd, 0x6b, 0xed, 0x0e, 0x15, 0xf9, 0x05, 0x5c, 0xf9, 0x59, 0x3d,
	0xf6, 0xc8, 0x09, 0xa1, 0xe4, 0x8f, 0xa0, 0xee, 0x6e, 0xdc, 0x50, 0x2d, 0x3e, 0x25, 0x1a, 0x3d,
	0xef, 0xb3, 0xe3, 0x57, 0x43, 0x0e, 0x29, 0x08, 0x2e, 0x58, 0x4a, 0x2f, 0x81, 0x8b, 0x44, 0xa3,
	0x54, 0x50, 0xb0, 0xa4, 0x60, 0x82, 0x69, 0xae, 0xe3, 0x5a, 0x49, 0x94, 0xe1, 0xde, 0x3a, 0x13,
	0x3b, 0x66, 0xb8, 0x57, 0xc8, 0x42, 0x1a, 0x20, 0xb9, 0xfd, 0x67, 0xd9, 0xe1, 0x81, 0xd7, 0x57,
	0x83, 0x82, 0xca, 0xe9, 0x86, 0xcf, 0xbc, 0x08, 0x95, 0x02, 0x15, 0x50, 0x5c, 0x51, 0x2f, 0xbc,
	0x14, 0x50, 0xe4, 0x57, 0x2c, 0xcd, 0x18, 0x94, 0xed, 0xba, 0x5a, 0xc9, 0x2b, 0x9e, 0x31, 0xa5,
	0x5b, 0xf7, 0xd2, 0xa8, 0x60, 0xbe, 0x42, 0xfc, 0x55, 0xe4, 0x3c, 0x4b, 0x29, 0xcf, 0x5a, 0xb7,
	0xaa, 0x61, 0x5e, 0x31, 0x81, 0x29, 0x17, 0xb9, 0xeb, 0xe1, 0xf0, 0xfb, 0x16, 0xe9, 0xbf, 0xb7,
	0x2d, 0x9e, 0x23, 0x20, 0x0b, 0x5f, 0x93, 0xae, 0x6d, 0x61, 0x10, 0x8c, 0x82, 0x71, 0x6f, 0xf2,
	0x34, 0xf6, 0xb5, 0x1a, 0x9f, 0x1a, 0xe6, 0x68, 0xf3, 0xfa, 0xd7, 0x7e, 0xe7, 0xcc, 0x25, 0xc2,
	0x77, 0xe4, 0x41, 0xce, 0x4b, 0x96, 0x96, 0x5c, 0xe3, 0xe0, 0xbf, 0xd1, 0xc6, 0xb8, 0x37, 0x39,
	0xf0, 0xc7, 0x3f, 0x09, 0x9e, 0x73, 0x96, 0x1d, 0xf3, 0x92, 0x39, 0xc7, 0xf6, 0x6d, 0x72, 0xca,
	0x35, 0x86, 0x53, 0xb2, 0xd3, 0xb4, 0x62, 0x55, 0x1b, 0x46, 0xb5, 0xff, 0x8f, 0x4d, 0x56, 0xac,
	0x13, 0x3d, 0x6c, 0xc2, 0xc6, 0x76, 0x41, 0x1e, 0xad, 0x7f, 0xb6, 0x15, 0x6e, 0x1a, 0xe1, 0xd8,
	0x2f, 0x3c, 0xb7, 0xbf, 0xa7, 0x36, 0x75, 0x22, 0x72, 0xe9, 0xcc, 0xbb, 0xf5, 0xdd, 0xc8, 0xb8,
	0x3f, 0x92, 0x5d, 0x2a, 0xcb, 0x12, 0x90, 0x29, 0x28, 0xad, 0x79, 0xcb, 0x98, 0x47, 0x7e, 0xf3,
	0xdb, 0x06, 0x76, 0xc6, 0x9d, 0xbb, 0xb8, 0x11, 0xa6, 0xe4, 0x89, 0xbb, 0x9c, 0x7b, 0x0d, 0x74,
	0x8d, 0xf6, 0xb9, 0x5f, 0xfb, 0xc6, 0x44, 0xee, 0xf7, 0xf0, 0x18, 0xfe, 0x1e, 0x9b, 0x07, 0x4e,
	0x48, 0x5f, 0xb1, 0x5a, 0x2a, 0x4c, 0x73, 0xa9, 0x2a, 0x3d, 0xf8, 0xbf, 0x6d, 0xdd, 0x33, 0x43,
	0x1e, 0x4b, 0x55, 0x39, 0x65, 0x4f, 0x35, 0x13, 0x1d, 0x7e, 0x20, 0x7d, 0x40, 0x64, 0x7a, 0xa5,
	0xda, 0x6e, 0x5d, 0xd1, 0x90, 0x80, 0x5c, 0x8a, 0x75, 0x9f, 0x15, 0x18, 0xdf, 0xd1, 0xf4, 0x7a,
	0x11, 0x05, 0x37, 0x8b, 0x28, 0xf8, 0xbd, 0x88, 0x82, 0x1f, 0xcb, 0xa8, 0x73, 0xb3, 0x8c, 0x3a,
	0x3f, 0x97, 0x51, 0xe7, 0x62, 0x52, 0x70, 0xbc, 0xfc, 0x3a, 0x8b, 0xa9, 0xac, 0x92, 0xcf, 0x40,
	0xbf, 0x40, 0x39, 0x85, 0x99, 0x4e, 0xec, 0x43, 0x2f, 0xed, 0x89, 0x7f, 0x6b, 0x8e, 0x1c, 0xe7,
	0x35, 0xd3, 0xb3, 0xae, 0x39, 0xef, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x72, 0x7e, 0x95,
	0xad, 0x38, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AttestForms) > 0 {
		for iNdEx := len(m.AttestForms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AttestForms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.ReportForms) > 0 {
		for iNdEx := len(m.ReportForms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReportForms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ActiveProvidersList) > 0 {
		for iNdEx := len(m.ActiveProvidersList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActiveProvidersList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.CollateralList) > 0 {
		for iNdEx := len(m.CollateralList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollateralList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PaymentInfoList) > 0 {
		for iNdEx := len(m.PaymentInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PaymentInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ProvidersList) > 0 {
		for iNdEx := len(m.ProvidersList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProvidersList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FileList) > 0 {
		for iNdEx := len(m.FileList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FileList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FileList) > 0 {
		for _, e := range m.FileList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProvidersList) > 0 {
		for _, e := range m.ProvidersList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PaymentInfoList) > 0 {
		for _, e := range m.PaymentInfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CollateralList) > 0 {
		for _, e := range m.CollateralList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ActiveProvidersList) > 0 {
		for _, e := range m.ActiveProvidersList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ReportForms) > 0 {
		for _, e := range m.ReportForms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AttestForms) > 0 {
		for _, e := range m.AttestForms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileList = append(m.FileList, UnifiedFile{})
			if err := m.FileList[len(m.FileList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvidersList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProvidersList = append(m.ProvidersList, Providers{})
			if err := m.ProvidersList[len(m.ProvidersList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentInfoList = append(m.PaymentInfoList, StoragePaymentInfo{})
			if err := m.PaymentInfoList[len(m.PaymentInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralList = append(m.CollateralList, Collateral{})
			if err := m.CollateralList[len(m.CollateralList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveProvidersList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActiveProvidersList = append(m.ActiveProvidersList, ActiveProviders{})
			if err := m.ActiveProvidersList[len(m.ActiveProvidersList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportForms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReportForms = append(m.ReportForms, ReportForm{})
			if err := m.ReportForms[len(m.ReportForms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestForms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttestForms = append(m.AttestForms, AttestationForm{})
			if err := m.AttestForms[len(m.AttestForms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
