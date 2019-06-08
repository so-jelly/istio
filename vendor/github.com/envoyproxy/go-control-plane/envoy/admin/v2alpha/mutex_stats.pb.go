// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v2alpha/mutex_stats.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	io "io"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Proto representation of the statistics collected upon absl::Mutex contention, if Envoy is run
// under :option:`--enable-mutex-tracing`. For more information, see the `absl::Mutex`
// [docs](https://abseil.io/about/design/mutex#extra-features).
//
// *NB*: The wait cycles below are measured by `absl::base_internal::CycleClock`, and may not
// correspond to core clock frequency. For more information, see the `CycleClock`
// [docs](https://github.com/abseil/abseil-cpp/blob/master/absl/base/internal/cycleclock.h).
type MutexStats struct {
	// The number of individual mutex contentions which have occurred since startup.
	NumContentions uint64 `protobuf:"varint,1,opt,name=num_contentions,json=numContentions,proto3" json:"num_contentions,omitempty"`
	// The length of the current contention wait cycle.
	CurrentWaitCycles uint64 `protobuf:"varint,2,opt,name=current_wait_cycles,json=currentWaitCycles,proto3" json:"current_wait_cycles,omitempty"`
	// The lifetime total of all contention wait cycles.
	LifetimeWaitCycles   uint64   `protobuf:"varint,3,opt,name=lifetime_wait_cycles,json=lifetimeWaitCycles,proto3" json:"lifetime_wait_cycles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutexStats) Reset()         { *m = MutexStats{} }
func (m *MutexStats) String() string { return proto.CompactTextString(m) }
func (*MutexStats) ProtoMessage()    {}
func (*MutexStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c1145f4b9ed4752, []int{0}
}
func (m *MutexStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MutexStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MutexStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MutexStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutexStats.Merge(m, src)
}
func (m *MutexStats) XXX_Size() int {
	return m.Size()
}
func (m *MutexStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MutexStats.DiscardUnknown(m)
}

var xxx_messageInfo_MutexStats proto.InternalMessageInfo

func (m *MutexStats) GetNumContentions() uint64 {
	if m != nil {
		return m.NumContentions
	}
	return 0
}

func (m *MutexStats) GetCurrentWaitCycles() uint64 {
	if m != nil {
		return m.CurrentWaitCycles
	}
	return 0
}

func (m *MutexStats) GetLifetimeWaitCycles() uint64 {
	if m != nil {
		return m.LifetimeWaitCycles
	}
	return 0
}

func init() {
	proto.RegisterType((*MutexStats)(nil), "envoy.admin.v2alpha.MutexStats")
}

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/mutex_stats.proto", fileDescriptor_2c1145f4b9ed4752)
}

var fileDescriptor_2c1145f4b9ed4752 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0x86, 0x89, 0x8a, 0x43, 0x06, 0x8b, 0xa9, 0x43, 0xa7, 0xa2, 0x82, 0xe8, 0x94, 0x8a, 0xbe,
	0x41, 0x3b, 0x0b, 0x45, 0x07, 0xc7, 0x10, 0x63, 0xc4, 0x40, 0x73, 0x52, 0x92, 0x93, 0xda, 0x3e,
	0x89, 0xaf, 0xe4, 0xe8, 0x23, 0x48, 0x9f, 0x44, 0x1a, 0x7a, 0x29, 0x17, 0xee, 0x7a, 0xfe, 0xef,
	0x83, 0xf3, 0xd1, 0x1b, 0x0d, 0x83, 0x9b, 0x2a, 0xf9, 0x6e, 0x0d, 0x54, 0xc3, 0x83, 0xec, 0xfa,
	0x4f, 0x59, 0xd9, 0x88, 0x7a, 0x14, 0x01, 0x25, 0x06, 0xde, 0x7b, 0x87, 0x8e, 0xe5, 0x09, 0xe3,
	0x09, 0xe3, 0x2b, 0x76, 0xfd, 0x4d, 0x28, 0x7d, 0x5a, 0xd0, 0x97, 0x85, 0x64, 0xb7, 0x34, 0x83,
	0x68, 0x85, 0x72, 0x80, 0x1a, 0xd0, 0x38, 0x08, 0x05, 0xb9, 0x24, 0x77, 0x27, 0xcf, 0x67, 0x10,
	0x6d, 0xb3, 0x5d, 0x19, 0xa7, 0xb9, 0x8a, 0xde, 0x6b, 0x40, 0xf1, 0x25, 0x0d, 0x0a, 0x35, 0xa9,
	0x4e, 0x87, 0xe2, 0x28, 0xc1, 0xe7, 0xeb, 0xf4, 0x2a, 0x0d, 0x36, 0x69, 0x60, 0xf7, 0xf4, 0xa2,
	0x33, 0x1f, 0x1a, 0x8d, 0xd5, 0x7b, 0xc2, 0x71, 0x12, 0xd8, 0x6e, 0xdb, 0x8c, 0xba, 0xfe, 0x99,
	0x4b, 0xf2, 0x3b, 0x97, 0xe4, 0x6f, 0x2e, 0x09, 0xbd, 0x32, 0x8e, 0xa7, 0xff, 0x7b, 0xef, 0xc6,
	0x89, 0x1f, 0x48, 0xa9, 0xb3, 0xad, 0xa3, 0x5d, 0x82, 0x5b, 0xf2, 0x76, 0x9a, 0xca, 0x1f, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x60, 0xe7, 0x4c, 0x18, 0x22, 0x01, 0x00, 0x00,
}

func (m *MutexStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MutexStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NumContentions != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMutexStats(dAtA, i, uint64(m.NumContentions))
	}
	if m.CurrentWaitCycles != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMutexStats(dAtA, i, uint64(m.CurrentWaitCycles))
	}
	if m.LifetimeWaitCycles != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMutexStats(dAtA, i, uint64(m.LifetimeWaitCycles))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMutexStats(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MutexStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumContentions != 0 {
		n += 1 + sovMutexStats(uint64(m.NumContentions))
	}
	if m.CurrentWaitCycles != 0 {
		n += 1 + sovMutexStats(uint64(m.CurrentWaitCycles))
	}
	if m.LifetimeWaitCycles != 0 {
		n += 1 + sovMutexStats(uint64(m.LifetimeWaitCycles))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMutexStats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMutexStats(x uint64) (n int) {
	return sovMutexStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MutexStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMutexStats
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
			return fmt.Errorf("proto: MutexStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MutexStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumContentions", wireType)
			}
			m.NumContentions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutexStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumContentions |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentWaitCycles", wireType)
			}
			m.CurrentWaitCycles = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutexStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentWaitCycles |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LifetimeWaitCycles", wireType)
			}
			m.LifetimeWaitCycles = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutexStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LifetimeWaitCycles |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMutexStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMutexStats
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMutexStats
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
func skipMutexStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMutexStats
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
					return 0, ErrIntOverflowMutexStats
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
					return 0, ErrIntOverflowMutexStats
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
				return 0, ErrInvalidLengthMutexStats
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMutexStats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMutexStats
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
				next, err := skipMutexStats(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMutexStats
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
	ErrInvalidLengthMutexStats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMutexStats   = fmt.Errorf("proto: integer overflow")
)