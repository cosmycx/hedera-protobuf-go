// Code generated by protoc-gen-go. DO NOT EDIT.
// source: AdminDelete.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Delete a file or smart contract - can only be done with a Hedera admin multisig. When it is deleted, it immediately disappears from the system as seen by the user, but is still stored internally until the expiration time, at which time it is truly and permanently deleted. Until that time, it can be undeleted by the Hedera admin multisig. When a smart contract is deleted, the cryptocurrency account within it continues to exist, and is not affected by the expiration time here.
type AdminDeleteTransactionBody struct {
	// Types that are valid to be assigned to Id:
	//	*AdminDeleteTransactionBody_FileID
	//	*AdminDeleteTransactionBody_ContractID
	Id                   isAdminDeleteTransactionBody_Id `protobuf_oneof:"id"`
	ExpirationTime       *TimestampSeconds               `protobuf:"bytes,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *AdminDeleteTransactionBody) Reset()         { *m = AdminDeleteTransactionBody{} }
func (m *AdminDeleteTransactionBody) String() string { return proto.CompactTextString(m) }
func (*AdminDeleteTransactionBody) ProtoMessage()    {}
func (*AdminDeleteTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_de5caba83e5ed17e, []int{0}
}

func (m *AdminDeleteTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminDeleteTransactionBody.Unmarshal(m, b)
}
func (m *AdminDeleteTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminDeleteTransactionBody.Marshal(b, m, deterministic)
}
func (m *AdminDeleteTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminDeleteTransactionBody.Merge(m, src)
}
func (m *AdminDeleteTransactionBody) XXX_Size() int {
	return xxx_messageInfo_AdminDeleteTransactionBody.Size(m)
}
func (m *AdminDeleteTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminDeleteTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_AdminDeleteTransactionBody proto.InternalMessageInfo

type isAdminDeleteTransactionBody_Id interface {
	isAdminDeleteTransactionBody_Id()
}

type AdminDeleteTransactionBody_FileID struct {
	FileID *FileID `protobuf:"bytes,1,opt,name=fileID,proto3,oneof"`
}

type AdminDeleteTransactionBody_ContractID struct {
	ContractID *ContractID `protobuf:"bytes,2,opt,name=contractID,proto3,oneof"`
}

func (*AdminDeleteTransactionBody_FileID) isAdminDeleteTransactionBody_Id() {}

func (*AdminDeleteTransactionBody_ContractID) isAdminDeleteTransactionBody_Id() {}

func (m *AdminDeleteTransactionBody) GetId() isAdminDeleteTransactionBody_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AdminDeleteTransactionBody) GetFileID() *FileID {
	if x, ok := m.GetId().(*AdminDeleteTransactionBody_FileID); ok {
		return x.FileID
	}
	return nil
}

func (m *AdminDeleteTransactionBody) GetContractID() *ContractID {
	if x, ok := m.GetId().(*AdminDeleteTransactionBody_ContractID); ok {
		return x.ContractID
	}
	return nil
}

func (m *AdminDeleteTransactionBody) GetExpirationTime() *TimestampSeconds {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AdminDeleteTransactionBody) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AdminDeleteTransactionBody_OneofMarshaler, _AdminDeleteTransactionBody_OneofUnmarshaler, _AdminDeleteTransactionBody_OneofSizer, []interface{}{
		(*AdminDeleteTransactionBody_FileID)(nil),
		(*AdminDeleteTransactionBody_ContractID)(nil),
	}
}

func _AdminDeleteTransactionBody_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AdminDeleteTransactionBody)
	// id
	switch x := m.Id.(type) {
	case *AdminDeleteTransactionBody_FileID:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FileID); err != nil {
			return err
		}
	case *AdminDeleteTransactionBody_ContractID:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContractID); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AdminDeleteTransactionBody.Id has unexpected type %T", x)
	}
	return nil
}

func _AdminDeleteTransactionBody_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AdminDeleteTransactionBody)
	switch tag {
	case 1: // id.fileID
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FileID)
		err := b.DecodeMessage(msg)
		m.Id = &AdminDeleteTransactionBody_FileID{msg}
		return true, err
	case 2: // id.contractID
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContractID)
		err := b.DecodeMessage(msg)
		m.Id = &AdminDeleteTransactionBody_ContractID{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AdminDeleteTransactionBody_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AdminDeleteTransactionBody)
	// id
	switch x := m.Id.(type) {
	case *AdminDeleteTransactionBody_FileID:
		s := proto.Size(x.FileID)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdminDeleteTransactionBody_ContractID:
		s := proto.Size(x.ContractID)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AdminDeleteTransactionBody)(nil), "proto.AdminDeleteTransactionBody")
}

func init() { proto.RegisterFile("AdminDelete.proto", fileDescriptor_de5caba83e5ed17e) }

var fileDescriptor_de5caba83e5ed17e = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x9b, 0x02, 0x1d, 0x2e, 0xe2, 0xa7, 0x59, 0x88, 0x32, 0xa1, 0x0e, 0xc0, 0x94, 0x81,
	0x3e, 0x00, 0x22, 0x44, 0xa8, 0xdd, 0x50, 0xc8, 0x0b, 0x5c, 0xec, 0x0b, 0xbe, 0xa8, 0xfe, 0x91,
	0x6d, 0x21, 0xfa, 0x78, 0xbc, 0x19, 0xaa, 0x63, 0x55, 0x55, 0x27, 0xcb, 0xe7, 0x7c, 0x9f, 0x7d,
	0x60, 0xfe, 0x2c, 0x35, 0x9b, 0x8e, 0x36, 0x14, 0xa9, 0x71, 0xde, 0x46, 0x5b, 0x9e, 0xa5, 0xa3,
	0xbe, 0x6e, 0x31, 0xb0, 0x18, 0xb6, 0x8e, 0xc2, 0x58, 0xd4, 0x57, 0x03, 0x6b, 0x0a, 0x11, 0xb5,
	0x1b, 0x83, 0xc5, 0x5f, 0x01, 0xf5, 0x81, 0x3f, 0x78, 0x34, 0x01, 0x45, 0x64, 0x6b, 0x5a, 0x2b,
	0xb7, 0xe5, 0x3d, 0xcc, 0x3e, 0x79, 0x43, 0xeb, 0xae, 0x2a, 0x6e, 0x8b, 0x87, 0xf3, 0xc7, 0x8b,
	0x51, 0x6b, 0x5e, 0x53, 0xb8, 0x9a, 0xf4, 0xb9, 0x2e, 0x97, 0x00, 0xc2, 0x9a, 0xe8, 0x51, 0xc4,
	0x75, 0x57, 0x4d, 0x13, 0x3c, 0xcf, 0xf0, 0xcb, 0xbe, 0x58, 0x4d, 0xfa, 0x03, 0xac, 0x7c, 0x82,
	0x4b, 0xfa, 0x75, 0xec, 0x71, 0xf7, 0xdf, 0x6e, 0x59, 0x75, 0x92, 0xc4, 0x9b, 0x2c, 0xee, 0xc7,
	0xbe, 0x93, 0xb0, 0x46, 0x86, 0xfe, 0x08, 0x6f, 0x4f, 0x61, 0xca, 0xb2, 0xbd, 0x83, 0x85, 0xb0,
	0xba, 0x51, 0x24, 0xc9, 0xa3, 0xc2, 0xa0, 0xbe, 0x3c, 0x3a, 0xd5, 0xa0, 0xe3, 0xfc, 0xce, 0x37,
	0xfe, 0xe0, 0x5b, 0xf1, 0x31, 0x4b, 0xb7, 0xe5, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x86,
	0x79, 0x57, 0x31, 0x01, 0x00, 0x00,
}
