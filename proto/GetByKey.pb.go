// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GetByKey.proto

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

// Get all accounts, claims, files, and smart contract instances whose associated keys include the given Key. The given Key must not be a contractID or a ThresholdKey. This is not yet implemented in the API, but will be in the future.
type GetByKeyQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Key                  *Key         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetByKeyQuery) Reset()         { *m = GetByKeyQuery{} }
func (m *GetByKeyQuery) String() string { return proto.CompactTextString(m) }
func (*GetByKeyQuery) ProtoMessage()    {}
func (*GetByKeyQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c7ecee7679d1dfd, []int{0}
}

func (m *GetByKeyQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByKeyQuery.Unmarshal(m, b)
}
func (m *GetByKeyQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByKeyQuery.Marshal(b, m, deterministic)
}
func (m *GetByKeyQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByKeyQuery.Merge(m, src)
}
func (m *GetByKeyQuery) XXX_Size() int {
	return xxx_messageInfo_GetByKeyQuery.Size(m)
}
func (m *GetByKeyQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByKeyQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GetByKeyQuery proto.InternalMessageInfo

func (m *GetByKeyQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetByKeyQuery) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

// the ID for a single entity (account, claim, file, or smart contract instance)
type EntityID struct {
	// Types that are valid to be assigned to Entity:
	//	*EntityID_AccountID
	//	*EntityID_Claim
	//	*EntityID_FileID
	//	*EntityID_ContractID
	Entity               isEntityID_Entity `protobuf_oneof:"entity"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EntityID) Reset()         { *m = EntityID{} }
func (m *EntityID) String() string { return proto.CompactTextString(m) }
func (*EntityID) ProtoMessage()    {}
func (*EntityID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c7ecee7679d1dfd, []int{1}
}

func (m *EntityID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityID.Unmarshal(m, b)
}
func (m *EntityID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityID.Marshal(b, m, deterministic)
}
func (m *EntityID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityID.Merge(m, src)
}
func (m *EntityID) XXX_Size() int {
	return xxx_messageInfo_EntityID.Size(m)
}
func (m *EntityID) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityID.DiscardUnknown(m)
}

var xxx_messageInfo_EntityID proto.InternalMessageInfo

type isEntityID_Entity interface {
	isEntityID_Entity()
}

type EntityID_AccountID struct {
	AccountID *AccountID `protobuf:"bytes,1,opt,name=accountID,proto3,oneof"`
}

type EntityID_Claim struct {
	Claim *Claim `protobuf:"bytes,2,opt,name=claim,proto3,oneof"`
}

type EntityID_FileID struct {
	FileID *FileID `protobuf:"bytes,3,opt,name=fileID,proto3,oneof"`
}

type EntityID_ContractID struct {
	ContractID *ContractID `protobuf:"bytes,4,opt,name=contractID,proto3,oneof"`
}

func (*EntityID_AccountID) isEntityID_Entity() {}

func (*EntityID_Claim) isEntityID_Entity() {}

func (*EntityID_FileID) isEntityID_Entity() {}

func (*EntityID_ContractID) isEntityID_Entity() {}

func (m *EntityID) GetEntity() isEntityID_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *EntityID) GetAccountID() *AccountID {
	if x, ok := m.GetEntity().(*EntityID_AccountID); ok {
		return x.AccountID
	}
	return nil
}

func (m *EntityID) GetClaim() *Claim {
	if x, ok := m.GetEntity().(*EntityID_Claim); ok {
		return x.Claim
	}
	return nil
}

func (m *EntityID) GetFileID() *FileID {
	if x, ok := m.GetEntity().(*EntityID_FileID); ok {
		return x.FileID
	}
	return nil
}

func (m *EntityID) GetContractID() *ContractID {
	if x, ok := m.GetEntity().(*EntityID_ContractID); ok {
		return x.ContractID
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EntityID) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EntityID_OneofMarshaler, _EntityID_OneofUnmarshaler, _EntityID_OneofSizer, []interface{}{
		(*EntityID_AccountID)(nil),
		(*EntityID_Claim)(nil),
		(*EntityID_FileID)(nil),
		(*EntityID_ContractID)(nil),
	}
}

func _EntityID_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EntityID)
	// entity
	switch x := m.Entity.(type) {
	case *EntityID_AccountID:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AccountID); err != nil {
			return err
		}
	case *EntityID_Claim:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Claim); err != nil {
			return err
		}
	case *EntityID_FileID:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FileID); err != nil {
			return err
		}
	case *EntityID_ContractID:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContractID); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EntityID.Entity has unexpected type %T", x)
	}
	return nil
}

func _EntityID_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EntityID)
	switch tag {
	case 1: // entity.accountID
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AccountID)
		err := b.DecodeMessage(msg)
		m.Entity = &EntityID_AccountID{msg}
		return true, err
	case 2: // entity.claim
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Claim)
		err := b.DecodeMessage(msg)
		m.Entity = &EntityID_Claim{msg}
		return true, err
	case 3: // entity.fileID
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FileID)
		err := b.DecodeMessage(msg)
		m.Entity = &EntityID_FileID{msg}
		return true, err
	case 4: // entity.contractID
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContractID)
		err := b.DecodeMessage(msg)
		m.Entity = &EntityID_ContractID{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EntityID_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EntityID)
	// entity
	switch x := m.Entity.(type) {
	case *EntityID_AccountID:
		s := proto.Size(x.AccountID)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EntityID_Claim:
		s := proto.Size(x.Claim)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EntityID_FileID:
		s := proto.Size(x.FileID)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EntityID_ContractID:
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

// Response when the client sends the node GetByKeyQuery
type GetByKeyResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Entities             []*EntityID     `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetByKeyResponse) Reset()         { *m = GetByKeyResponse{} }
func (m *GetByKeyResponse) String() string { return proto.CompactTextString(m) }
func (*GetByKeyResponse) ProtoMessage()    {}
func (*GetByKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c7ecee7679d1dfd, []int{2}
}

func (m *GetByKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByKeyResponse.Unmarshal(m, b)
}
func (m *GetByKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByKeyResponse.Marshal(b, m, deterministic)
}
func (m *GetByKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByKeyResponse.Merge(m, src)
}
func (m *GetByKeyResponse) XXX_Size() int {
	return xxx_messageInfo_GetByKeyResponse.Size(m)
}
func (m *GetByKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByKeyResponse proto.InternalMessageInfo

func (m *GetByKeyResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetByKeyResponse) GetEntities() []*EntityID {
	if m != nil {
		return m.Entities
	}
	return nil
}

func init() {
	proto.RegisterType((*GetByKeyQuery)(nil), "proto.GetByKeyQuery")
	proto.RegisterType((*EntityID)(nil), "proto.EntityID")
	proto.RegisterType((*GetByKeyResponse)(nil), "proto.GetByKeyResponse")
}

func init() { proto.RegisterFile("GetByKey.proto", fileDescriptor_0c7ecee7679d1dfd) }

var fileDescriptor_0c7ecee7679d1dfd = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4f, 0xea, 0x40,
	0x14, 0x85, 0x5b, 0x78, 0x34, 0xbc, 0xcb, 0xe3, 0x09, 0x13, 0x4d, 0x1a, 0xe2, 0x82, 0x34, 0x46,
	0x89, 0xc6, 0xc6, 0xc0, 0x2f, 0xa0, 0xa0, 0x96, 0xb0, 0xd1, 0x89, 0x1b, 0x97, 0xe3, 0xf4, 0x6a,
	0xab, 0xd0, 0x69, 0xa6, 0x83, 0xc9, 0xfc, 0x45, 0x7f, 0x95, 0x61, 0x3a, 0x45, 0x70, 0xd5, 0xf4,
	0xdc, 0xef, 0x9c, 0x7b, 0xe7, 0xc0, 0xff, 0x7b, 0x54, 0x91, 0x5e, 0xa2, 0x0e, 0x0b, 0x29, 0x94,
	0x20, 0x2d, 0xf3, 0x19, 0xf4, 0x22, 0x56, 0x66, 0xfc, 0x49, 0x17, 0x58, 0x56, 0x83, 0x41, 0xff,
	0x71, 0x83, 0x52, 0xc7, 0xc8, 0x12, 0x94, 0x56, 0x3a, 0xa6, 0x58, 0x16, 0x22, 0x2f, 0xf1, 0x50,
	0x9d, 0x49, 0x5d, 0x28, 0x31, 0x4d, 0x92, 0xd9, 0x8a, 0x65, 0xeb, 0x4a, 0x0d, 0x9e, 0xa1, 0x5b,
	0x6f, 0x32, 0x41, 0xe4, 0x12, 0xbc, 0xd4, 0xd8, 0x7c, 0x77, 0xe8, 0x8e, 0x3a, 0x63, 0x52, 0x81,
	0xe1, 0xde, 0x1a, 0x6a, 0x09, 0x72, 0x0a, 0xcd, 0x0f, 0xd4, 0x7e, 0xc3, 0x80, 0x60, 0xc1, 0x25,
	0x6a, 0xba, 0x95, 0x83, 0x2f, 0x17, 0xda, 0xb7, 0xb9, 0xca, 0x94, 0x5e, 0xcc, 0xc9, 0x0d, 0xfc,
	0x65, 0x9c, 0x8b, 0x4d, 0xae, 0x16, 0x73, 0x9b, 0xdc, 0xb3, 0x86, 0x69, 0xad, 0xc7, 0x0e, 0xfd,
	0x81, 0xc8, 0x19, 0xb4, 0xf8, 0xf6, 0x50, 0x1b, 0xff, 0xcf, 0xd2, 0xe6, 0xf8, 0xd8, 0xa1, 0xd5,
	0x90, 0x5c, 0x80, 0xf7, 0x9a, 0xad, 0x70, 0x31, 0xf7, 0x9b, 0x06, 0xeb, 0x5a, 0xec, 0xce, 0x88,
	0xb1, 0x43, 0xed, 0x98, 0x4c, 0x00, 0xb8, 0xc8, 0x95, 0x64, 0x7c, 0x7b, 0xc1, 0x1f, 0x03, 0xf7,
	0xeb, 0xcc, 0xdd, 0x20, 0x76, 0xe8, 0x1e, 0x16, 0xb5, 0xc1, 0x43, 0xf3, 0x82, 0x20, 0x87, 0x5e,
	0xdd, 0x53, 0xdd, 0x2e, 0xb9, 0xfe, 0x55, 0xd5, 0x89, 0x8d, 0x3b, 0xac, 0x7f, 0xd7, 0xd6, 0x15,
	0xb4, 0x4d, 0x58, 0x86, 0xa5, 0xdf, 0x18, 0x36, 0x47, 0x9d, 0xf1, 0x91, 0x35, 0xd4, 0x2d, 0xd1,
	0x1d, 0x10, 0x9d, 0x43, 0xc0, 0xc5, 0x3a, 0x4c, 0x31, 0x41, 0xc9, 0x52, 0x56, 0xa6, 0x6f, 0x92,
	0x15, 0x69, 0xc8, 0x8a, 0xcc, 0x7a, 0xde, 0xd9, 0x27, 0x7b, 0x70, 0x5f, 0x3c, 0xf3, 0x37, 0xf9,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x99, 0xa7, 0x7b, 0x90, 0x30, 0x02, 0x00, 0x00,
}
