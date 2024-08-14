// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/shared/supplier.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
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

// Supplier is the type defining the actor in Pocket Network that provides RPC services.
type Supplier struct {
	// The address of the owner (i.e. staker, custodial) that owns the funds for staking.
	// By default, this address is the one that receives all the rewards unless owtherwise specified.
	// This property cannot be updated by the operator.
	OwnerAddress string `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	// The operator address of the supplier operator (i.e. the one managing the off-chain server).
	// The operator address can update the supplier's configurations excluding the owner address.
	// This property does not change over the supplier's lifespan, the supplier must be unstaked
	// and re-staked to effectively update this value.
	OperatorAddress string                   `protobuf:"bytes,2,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	Stake           *types.Coin              `protobuf:"bytes,3,opt,name=stake,proto3" json:"stake,omitempty"`
	Services        []*SupplierServiceConfig `protobuf:"bytes,4,rep,name=services,proto3" json:"services,omitempty"`
	// The session end height at which an actively unbonding supplier unbonds its stake.
	// If the supplier did not unstake, this value will be 0.
	UnstakeSessionEndHeight uint64 `protobuf:"varint,5,opt,name=unstake_session_end_height,json=unstakeSessionEndHeight,proto3" json:"unstake_session_end_height,omitempty"`
	// services_activation_heights_map is a map of serviceIds to the height at
	// which the staked supplier will become active for that service.
	// Activation heights are session start heights.
	ServicesActivationHeightsMap map[string]uint64 `protobuf:"bytes,6,rep,name=services_activation_heights_map,json=servicesActivationHeightsMap,proto3" json:"services_activation_heights_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *Supplier) Reset()         { *m = Supplier{} }
func (m *Supplier) String() string { return proto.CompactTextString(m) }
func (*Supplier) ProtoMessage()    {}
func (*Supplier) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a189b52ba503cf2, []int{0}
}
func (m *Supplier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Supplier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Supplier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Supplier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supplier.Merge(m, src)
}
func (m *Supplier) XXX_Size() int {
	return m.Size()
}
func (m *Supplier) XXX_DiscardUnknown() {
	xxx_messageInfo_Supplier.DiscardUnknown(m)
}

var xxx_messageInfo_Supplier proto.InternalMessageInfo

func (m *Supplier) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *Supplier) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *Supplier) GetStake() *types.Coin {
	if m != nil {
		return m.Stake
	}
	return nil
}

func (m *Supplier) GetServices() []*SupplierServiceConfig {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Supplier) GetUnstakeSessionEndHeight() uint64 {
	if m != nil {
		return m.UnstakeSessionEndHeight
	}
	return 0
}

func (m *Supplier) GetServicesActivationHeightsMap() map[string]uint64 {
	if m != nil {
		return m.ServicesActivationHeightsMap
	}
	return nil
}

func init() {
	proto.RegisterType((*Supplier)(nil), "poktroll.shared.Supplier")
	proto.RegisterMapType((map[string]uint64)(nil), "poktroll.shared.Supplier.ServicesActivationHeightsMapEntry")
}

func init() { proto.RegisterFile("poktroll/shared/supplier.proto", fileDescriptor_4a189b52ba503cf2) }

var fileDescriptor_4a189b52ba503cf2 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x8e, 0x9b, 0xa4, 0x2a, 0x2e, 0xa8, 0xd5, 0xaa, 0x12, 0xdb, 0x08, 0x96, 0xc0, 0x01, 0xe5,
	0x52, 0x5b, 0x2d, 0x17, 0x44, 0xc5, 0xa1, 0x89, 0x2a, 0xc1, 0x01, 0x21, 0x6d, 0x6e, 0x5c, 0x56,
	0xce, 0xee, 0x90, 0x58, 0x49, 0x6c, 0xcb, 0xe3, 0xa4, 0xe4, 0xca, 0x13, 0x70, 0xe2, 0x49, 0x78,
	0x08, 0x8e, 0x15, 0x27, 0x8e, 0x28, 0x79, 0x11, 0xb4, 0x6b, 0x6f, 0x90, 0x40, 0x34, 0xb7, 0x9d,
	0xfd, 0x7e, 0x66, 0x3e, 0xcf, 0xd0, 0xc4, 0xe8, 0xa9, 0xb3, 0x7a, 0x36, 0xe3, 0x38, 0x11, 0x16,
	0x0a, 0x8e, 0x0b, 0x63, 0x66, 0x12, 0x2c, 0x33, 0x56, 0x3b, 0x1d, 0x1d, 0xd5, 0x38, 0xf3, 0x78,
	0xe7, 0x34, 0xd7, 0x38, 0xd7, 0x98, 0x55, 0x30, 0xf7, 0x85, 0xe7, 0x76, 0x12, 0x5f, 0xf1, 0x91,
	0x40, 0xe0, 0xcb, 0xf3, 0x11, 0x38, 0x71, 0xce, 0x73, 0x2d, 0x55, 0xc0, 0x1f, 0xff, 0xd3, 0x0b,
	0xec, 0x52, 0xe6, 0xe0, 0xe1, 0x67, 0x5f, 0x5b, 0xf4, 0x60, 0x18, 0xba, 0x47, 0xaf, 0xe9, 0x03,
	0x7d, 0xa3, 0xc0, 0x66, 0xa2, 0x28, 0x2c, 0x20, 0xc6, 0xa4, 0x4b, 0x7a, 0xf7, 0xfa, 0xf1, 0x8f,
	0x6f, 0x67, 0x27, 0xa1, 0xe9, 0x95, 0x47, 0x86, 0xce, 0x4a, 0x35, 0x4e, 0xef, 0x57, 0xf4, 0xf0,
	0x2f, 0x1a, 0xd0, 0x63, 0x6d, 0xc0, 0x0a, 0xa7, 0xff, 0x38, 0xec, 0xed, 0x70, 0x38, 0xaa, 0x15,
	0xb5, 0x09, 0xa7, 0x6d, 0x74, 0x62, 0x0a, 0x71, 0xb3, 0x4b, 0x7a, 0x87, 0x17, 0xa7, 0x2c, 0xc8,
	0xca, 0x7c, 0x2c, 0xe4, 0x63, 0x03, 0x2d, 0x55, 0xea, 0x79, 0x51, 0x9f, 0x1e, 0x84, 0x48, 0x18,
	0xb7, 0xba, 0xcd, 0xde, 0xe1, 0xc5, 0x73, 0xf6, 0xd7, 0xfb, 0xb1, 0x3a, 0xe1, 0xd0, 0x13, 0x07,
	0x5a, 0x7d, 0x94, 0xe3, 0x74, 0xab, 0x8b, 0x2e, 0x69, 0x67, 0xa1, 0x2a, 0xbb, 0x0c, 0x01, 0x51,
	0x6a, 0x95, 0x81, 0x2a, 0xb2, 0x09, 0xc8, 0xf1, 0xc4, 0xc5, 0xed, 0x2e, 0xe9, 0xb5, 0xd2, 0x87,
	0x81, 0x31, 0xf4, 0x84, 0x6b, 0x55, 0xbc, 0xa9, 0xe0, 0xe8, 0x33, 0xa1, 0x4f, 0x6a, 0xa7, 0x4c,
	0xe4, 0x4e, 0x2e, 0x85, 0x2b, 0x1d, 0xbc, 0x1a, 0xb3, 0xb9, 0x30, 0xf1, 0x7e, 0x35, 0xd8, 0xe5,
	0x7f, 0x07, 0x63, 0x61, 0x32, 0xbc, 0xda, 0xea, 0xbd, 0x3b, 0xbe, 0x13, 0xe6, 0x5a, 0x39, 0xbb,
	0x4a, 0x1f, 0xe1, 0x1d, 0x94, 0xce, 0x7b, 0xfa, 0x74, 0xa7, 0x45, 0x74, 0x4c, 0x9b, 0x53, 0x58,
	0xf9, 0xad, 0xa6, 0xe5, 0x67, 0x74, 0x42, 0xdb, 0x4b, 0x31, 0x5b, 0x40, 0xb5, 0xa7, 0x56, 0xea,
	0x8b, 0x57, 0x7b, 0x2f, 0x49, 0xff, 0xed, 0xf7, 0x75, 0x42, 0x6e, 0xd7, 0x09, 0xf9, 0xb5, 0x4e,
	0xc8, 0x97, 0x4d, 0xd2, 0xb8, 0xdd, 0x24, 0x8d, 0x9f, 0x9b, 0xa4, 0xf1, 0x81, 0x8f, 0xa5, 0x9b,
	0x2c, 0x46, 0x2c, 0xd7, 0x73, 0x5e, 0xe6, 0x39, 0x53, 0xe0, 0x6e, 0xb4, 0x9d, 0xf2, 0xed, 0xa5,
	0x7d, 0xaa, 0x6f, 0xcd, 0xad, 0x0c, 0xe0, 0x68, 0xbf, 0x3a, 0xb5, 0x17, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0xa8, 0xdf, 0xfb, 0xf7, 0x02, 0x00, 0x00,
}

func (m *Supplier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Supplier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Supplier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServicesActivationHeightsMap) > 0 {
		for k := range m.ServicesActivationHeightsMap {
			v := m.ServicesActivationHeightsMap[k]
			baseI := i
			i = encodeVarintSupplier(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintSupplier(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintSupplier(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.UnstakeSessionEndHeight != 0 {
		i = encodeVarintSupplier(dAtA, i, uint64(m.UnstakeSessionEndHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Services[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSupplier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Stake != nil {
		{
			size, err := m.Stake.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSupplier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintSupplier(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintSupplier(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSupplier(dAtA []byte, offset int, v uint64) int {
	offset -= sovSupplier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Supplier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovSupplier(uint64(l))
	}
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovSupplier(uint64(l))
	}
	if m.Stake != nil {
		l = m.Stake.Size()
		n += 1 + l + sovSupplier(uint64(l))
	}
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovSupplier(uint64(l))
		}
	}
	if m.UnstakeSessionEndHeight != 0 {
		n += 1 + sovSupplier(uint64(m.UnstakeSessionEndHeight))
	}
	if len(m.ServicesActivationHeightsMap) > 0 {
		for k, v := range m.ServicesActivationHeightsMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovSupplier(uint64(len(k))) + 1 + sovSupplier(uint64(v))
			n += mapEntrySize + 1 + sovSupplier(uint64(mapEntrySize))
		}
	}
	return n
}

func sovSupplier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSupplier(x uint64) (n int) {
	return sovSupplier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Supplier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupplier
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
			return fmt.Errorf("proto: Supplier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Supplier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stake == nil {
				m.Stake = &types.Coin{}
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &SupplierServiceConfig{})
			if err := m.Services[len(m.Services)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakeSessionEndHeight", wireType)
			}
			m.UnstakeSessionEndHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnstakeSessionEndHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServicesActivationHeightsMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ServicesActivationHeightsMap == nil {
				m.ServicesActivationHeightsMap = make(map[string]uint64)
			}
			var mapkey string
			var mapvalue uint64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSupplier
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
							return ErrIntOverflowSupplier
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
						return ErrInvalidLengthSupplier
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthSupplier
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSupplier
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipSupplier(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthSupplier
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ServicesActivationHeightsMap[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSupplier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSupplier
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
func skipSupplier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSupplier
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
					return 0, ErrIntOverflowSupplier
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
					return 0, ErrIntOverflowSupplier
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
				return 0, ErrInvalidLengthSupplier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSupplier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSupplier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSupplier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSupplier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSupplier = fmt.Errorf("proto: unexpected end of group")
)
