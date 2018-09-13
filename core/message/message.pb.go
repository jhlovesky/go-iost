// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/message/message.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestBlock struct {
	BlockNumber          int64    `protobuf:"varint,1,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestBlock) Reset()         { *m = RequestBlock{} }
func (m *RequestBlock) String() string { return proto.CompactTextString(m) }
func (*RequestBlock) ProtoMessage()    {}
func (*RequestBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_91801a829f8d0ae6, []int{0}
}
func (m *RequestBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RequestBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestBlock.Merge(dst, src)
}
func (m *RequestBlock) XXX_Size() int {
	return m.Size()
}
func (m *RequestBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestBlock.DiscardUnknown(m)
}

var xxx_messageInfo_RequestBlock proto.InternalMessageInfo

func (m *RequestBlock) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *RequestBlock) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

type BlockHashQuery struct {
	ReqType              int32    `protobuf:"varint,1,opt,name=reqType,proto3" json:"reqType,omitempty"`
	Start                int64    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	Nums                 []int64  `protobuf:"varint,4,rep,packed,name=nums" json:"nums,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHashQuery) Reset()         { *m = BlockHashQuery{} }
func (m *BlockHashQuery) String() string { return proto.CompactTextString(m) }
func (*BlockHashQuery) ProtoMessage()    {}
func (*BlockHashQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_91801a829f8d0ae6, []int{1}
}
func (m *BlockHashQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHashQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHashQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BlockHashQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHashQuery.Merge(dst, src)
}
func (m *BlockHashQuery) XXX_Size() int {
	return m.Size()
}
func (m *BlockHashQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHashQuery.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHashQuery proto.InternalMessageInfo

func (m *BlockHashQuery) GetReqType() int32 {
	if m != nil {
		return m.ReqType
	}
	return 0
}

func (m *BlockHashQuery) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *BlockHashQuery) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *BlockHashQuery) GetNums() []int64 {
	if m != nil {
		return m.Nums
	}
	return nil
}

type BlockHash struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHash) Reset()         { *m = BlockHash{} }
func (m *BlockHash) String() string { return proto.CompactTextString(m) }
func (*BlockHash) ProtoMessage()    {}
func (*BlockHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_91801a829f8d0ae6, []int{2}
}
func (m *BlockHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BlockHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHash.Merge(dst, src)
}
func (m *BlockHash) XXX_Size() int {
	return m.Size()
}
func (m *BlockHash) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHash.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHash proto.InternalMessageInfo

func (m *BlockHash) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type BlockHashResponse struct {
	BlockHashes          []*BlockHash `protobuf:"bytes,1,rep,name=blockHashes" json:"blockHashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BlockHashResponse) Reset()         { *m = BlockHashResponse{} }
func (m *BlockHashResponse) String() string { return proto.CompactTextString(m) }
func (*BlockHashResponse) ProtoMessage()    {}
func (*BlockHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_91801a829f8d0ae6, []int{3}
}
func (m *BlockHashResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHashResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BlockHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHashResponse.Merge(dst, src)
}
func (m *BlockHashResponse) XXX_Size() int {
	return m.Size()
}
func (m *BlockHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHashResponse proto.InternalMessageInfo

func (m *BlockHashResponse) GetBlockHashes() []*BlockHash {
	if m != nil {
		return m.BlockHashes
	}
	return nil
}

type SyncHeight struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncHeight) Reset()         { *m = SyncHeight{} }
func (m *SyncHeight) String() string { return proto.CompactTextString(m) }
func (*SyncHeight) ProtoMessage()    {}
func (*SyncHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_91801a829f8d0ae6, []int{4}
}
func (m *SyncHeight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncHeight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SyncHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncHeight.Merge(dst, src)
}
func (m *SyncHeight) XXX_Size() int {
	return m.Size()
}
func (m *SyncHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncHeight.DiscardUnknown(m)
}

var xxx_messageInfo_SyncHeight proto.InternalMessageInfo

func (m *SyncHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *SyncHeight) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestBlock)(nil), "message.RequestBlock")
	proto.RegisterType((*BlockHashQuery)(nil), "message.BlockHashQuery")
	proto.RegisterType((*BlockHash)(nil), "message.BlockHash")
	proto.RegisterType((*BlockHashResponse)(nil), "message.BlockHashResponse")
	proto.RegisterType((*SyncHeight)(nil), "message.SyncHeight")
}
func (m *RequestBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestBlock) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BlockNumber != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.BlockNumber))
	}
	if len(m.BlockHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.BlockHash)))
		i += copy(dAtA[i:], m.BlockHash)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BlockHashQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHashQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReqType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ReqType))
	}
	if m.Start != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Start))
	}
	if m.End != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.End))
	}
	if len(m.Nums) > 0 {
		dAtA2 := make([]byte, len(m.Nums)*10)
		var j1 int
		for _, num1 := range m.Nums {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintMessage(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BlockHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHash) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Height))
	}
	if len(m.Hash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BlockHashResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHashResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.BlockHashes) > 0 {
		for _, msg := range m.BlockHashes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMessage(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SyncHeight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncHeight) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Height))
	}
	if m.Time != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Time))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RequestBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockNumber != 0 {
		n += 1 + sovMessage(uint64(m.BlockNumber))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockHashQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReqType != 0 {
		n += 1 + sovMessage(uint64(m.ReqType))
	}
	if m.Start != 0 {
		n += 1 + sovMessage(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovMessage(uint64(m.End))
	}
	if len(m.Nums) > 0 {
		l = 0
		for _, e := range m.Nums {
			l += sovMessage(uint64(e))
		}
		n += 1 + sovMessage(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovMessage(uint64(m.Height))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockHashResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BlockHashes) > 0 {
		for _, e := range m.BlockHashes {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SyncHeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovMessage(uint64(m.Height))
	}
	if m.Time != 0 {
		n += 1 + sovMessage(uint64(m.Time))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *BlockHashQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockHashQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHashQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqType", wireType)
			}
			m.ReqType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqType |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Nums = append(m.Nums, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMessage
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Nums = append(m.Nums, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Nums", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *BlockHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *BlockHashResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockHashResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHashResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHashes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHashes = append(m.BlockHashes, &BlockHash{})
			if err := m.BlockHashes[len(m.BlockHashes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *SyncHeight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SyncHeight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncHeight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
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
				next, err := skipMessage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("core/message/message.proto", fileDescriptor_message_91801a829f8d0ae6) }

var fileDescriptor_message_91801a829f8d0ae6 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0x7f, 0xff, 0xee, 0x45, 0x9d, 0x56, 0xa8, 0x8c, 0x10, 0xb2, 0x10, 0x8a, 0xa2, 0xac,
	0xb2, 0x2a, 0x12, 0x20, 0xc1, 0x3a, 0xab, 0xb2, 0xa9, 0x84, 0xe1, 0x05, 0x92, 0x30, 0x34, 0x15,
	0xe4, 0x52, 0xdb, 0x59, 0xe4, 0x4d, 0x78, 0x24, 0x96, 0x3c, 0x02, 0x0a, 0x2f, 0x82, 0xe2, 0xe6,
	0xd2, 0x15, 0x2b, 0x9f, 0xef, 0xc8, 0x33, 0x73, 0xc6, 0x86, 0x8b, 0x38, 0x57, 0x74, 0x95, 0x92,
	0xd6, 0xe1, 0xb6, 0x3f, 0x57, 0x85, 0xca, 0x4d, 0x8e, 0xd3, 0x16, 0xbd, 0x0d, 0x2c, 0x24, 0xed,
	0x4b, 0xd2, 0x26, 0x78, 0xcf, 0xe3, 0x37, 0x74, 0x61, 0x1e, 0x35, 0x62, 0x53, 0xa6, 0x11, 0x29,
	0xc1, 0x5c, 0xe6, 0x73, 0x79, 0x6c, 0xe1, 0x25, 0xcc, 0x2c, 0xae, 0x43, 0x9d, 0x88, 0xff, 0x2e,
	0xf3, 0x17, 0x72, 0x30, 0xbc, 0x57, 0x38, 0x09, 0x3a, 0x78, 0x2c, 0x49, 0x55, 0x28, 0x60, 0xaa,
	0x68, 0xff, 0x5c, 0x15, 0x64, 0xbb, 0x8d, 0x65, 0x87, 0x78, 0x06, 0x63, 0x6d, 0x42, 0x65, 0x6c,
	0x17, 0x2e, 0x0f, 0x80, 0x4b, 0xe0, 0x94, 0xbd, 0x08, 0x6e, 0xbd, 0x46, 0x22, 0xc2, 0x28, 0x2b,
	0x53, 0x2d, 0x46, 0x2e, 0xf7, 0xb9, 0xb4, 0xda, 0xbb, 0x83, 0x59, 0x3f, 0x07, 0xcf, 0x61, 0x92,
	0xd0, 0x6e, 0x9b, 0x98, 0x36, 0x6f, 0x4b, 0x4d, 0x61, 0x32, 0xa4, 0xb4, 0xda, 0x7b, 0x80, 0xd3,
	0xbe, 0x50, 0x92, 0x2e, 0xf2, 0x4c, 0x13, 0xde, 0xb6, 0x5b, 0x37, 0x26, 0x69, 0xc1, 0x5c, 0xee,
	0xcf, 0xaf, 0x71, 0xd5, 0xbd, 0xd9, 0x50, 0x70, 0x7c, 0xcd, 0xbb, 0x07, 0x78, 0xaa, 0xb2, 0x78,
	0x7d, 0x18, 0xf6, 0x47, 0x08, 0xb3, 0x4b, 0xa9, 0x5d, 0xd2, 0xea, 0x60, 0xf9, 0x59, 0x3b, 0xec,
	0xab, 0x76, 0xd8, 0x77, 0xed, 0xb0, 0x8f, 0x1f, 0xe7, 0x5f, 0x34, 0xb1, 0xff, 0x72, 0xf3, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x61, 0xef, 0xaf, 0x75, 0xb5, 0x01, 0x00, 0x00,
}
