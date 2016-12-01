// Code generated by protoc-gen-go.
// source: common/common.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto
	common/configuration.proto

It has these top-level messages:
	Header
	ChainHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigurationEnvelope
	SignedConfigurationItem
	ConfigurationItem
	ConfigurationSignature
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_UNKNOWN               Status = 0
	Status_SUCCESS               Status = 200
	Status_BAD_REQUEST           Status = 400
	Status_FORBIDDEN             Status = 403
	Status_NOT_FOUND             Status = 404
	Status_INTERNAL_SERVER_ERROR Status = 500
	Status_SERVICE_UNAVAILABLE   Status = 503
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	200: "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	500: "INTERNAL_SERVER_ERROR",
	503: "SERVICE_UNAVAILABLE",
}
var Status_value = map[string]int32{
	"UNKNOWN":               0,
	"SUCCESS":               200,
	"BAD_REQUEST":           400,
	"FORBIDDEN":             403,
	"NOT_FOUND":             404,
	"INTERNAL_SERVER_ERROR": 500,
	"SERVICE_UNAVAILABLE":   503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HeaderType int32

const (
	HeaderType_MESSAGE                   HeaderType = 0
	HeaderType_CONFIGURATION_TRANSACTION HeaderType = 1
	HeaderType_CONFIGURATION_ITEM        HeaderType = 2
	HeaderType_ENDORSER_TRANSACTION      HeaderType = 3
)

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIGURATION_TRANSACTION",
	2: "CONFIGURATION_ITEM",
	3: "ENDORSER_TRANSACTION",
}
var HeaderType_value = map[string]int32{
	"MESSAGE":                   0,
	"CONFIGURATION_TRANSACTION": 1,
	"CONFIGURATION_ITEM":        2,
	"ENDORSER_TRANSACTION":      3,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}
func (HeaderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Header struct {
	ChainHeader     *ChainHeader     `protobuf:"bytes,1,opt,name=chainHeader" json:"chainHeader,omitempty"`
	SignatureHeader *SignatureHeader `protobuf:"bytes,2,opt,name=signatureHeader" json:"signatureHeader,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Header) GetChainHeader() *ChainHeader {
	if m != nil {
		return m.ChainHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() *SignatureHeader {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChainHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// Identifier of the chain this message is bound for
	ChainID []byte `protobuf:"bytes,4,opt,name=chainID,proto3" json:"chainID,omitempty"`
	// An unique identifier that is used end-to-end.
	//  -  set by higher layers such as end user or SDK
	//  -  passed to the endorser (which will check for uniqueness)
	//  -  as the header is passed along unchanged, it will be
	//     be retrieved by the committer (uniqueness check here as well)
	//  -  to be stored in the ledger
	TxID string `protobuf:"bytes,5,opt,name=txID" json:"txID,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch uint64 `protobuf:"varint,6,opt,name=epoch" json:"epoch,omitempty"`
	// Extension that may be attached based on the header type
	Extension []byte `protobuf:"bytes,7,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *ChainHeader) Reset()                    { *m = ChainHeader{} }
func (m *ChainHeader) String() string            { return proto.CompactTextString(m) }
func (*ChainHeader) ProtoMessage()               {}
func (*ChainHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChainHeader) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type SignatureHeader struct {
	// Creator of the message, specified as a certificate chain
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *SignatureHeader) Reset()                    { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string            { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()               {}
func (*SignatureHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Payload is the message contents (and header to allow for signing)
type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Payload) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

// Envelope wraps a Payload with a signature so that the message may be authenticated
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// This is finalized block structure to be shared among the orderer and peer
// Note that the BlockHeader chains to the previous BlockHeader, and the BlockData hash is embedded
// in the BlockHeader.  This makes it natural and obvious that the Data is included in the hash, but
// the Metadata is not.
type Block struct {
	Header   *BlockHeader   `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"`
	Data     *BlockData     `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Metadata *BlockMetadata `protobuf:"bytes,3,opt,name=Metadata" json:"Metadata,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetMetadata() *BlockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type BlockHeader struct {
	Number       uint64 `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	PreviousHash []byte `protobuf:"bytes,2,opt,name=PreviousHash,proto3" json:"PreviousHash,omitempty"`
	DataHash     []byte `protobuf:"bytes,3,opt,name=DataHash,proto3" json:"DataHash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type BlockData struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *BlockData) Reset()                    { *m = BlockData{} }
func (m *BlockData) String() string            { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()               {}
func (*BlockData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type BlockMetadata struct {
	Metadata [][]byte `protobuf:"bytes,1,rep,name=Metadata,proto3" json:"Metadata,omitempty"`
}

func (m *BlockMetadata) Reset()                    { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string            { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()               {}
func (*BlockMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*Header)(nil), "common.Header")
	proto.RegisterType((*ChainHeader)(nil), "common.ChainHeader")
	proto.RegisterType((*SignatureHeader)(nil), "common.SignatureHeader")
	proto.RegisterType((*Payload)(nil), "common.Payload")
	proto.RegisterType((*Envelope)(nil), "common.Envelope")
	proto.RegisterType((*Block)(nil), "common.Block")
	proto.RegisterType((*BlockHeader)(nil), "common.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "common.BlockData")
	proto.RegisterType((*BlockMetadata)(nil), "common.BlockMetadata")
	proto.RegisterEnum("common.Status", Status_name, Status_value)
	proto.RegisterEnum("common.HeaderType", HeaderType_name, HeaderType_value)
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xad, 0xeb, 0x3c, 0x9a, 0xeb, 0xd0, 0x9a, 0xe9, 0x03, 0x13, 0x81, 0x1a, 0x59, 0x02, 0x45,
	0xad, 0x48, 0x44, 0x11, 0x12, 0x5b, 0x27, 0x9e, 0xb6, 0x16, 0xad, 0x5d, 0xc6, 0x4e, 0x91, 0xd8,
	0x44, 0x4e, 0x32, 0x4d, 0x02, 0x89, 0x27, 0x72, 0x9c, 0xaa, 0xdd, 0xb2, 0x45, 0x42, 0x48, 0xf0,
	0x51, 0xfc, 0x01, 0x3f, 0x82, 0xc4, 0x16, 0xcd, 0x8c, 0x9d, 0x47, 0x57, 0x99, 0x73, 0xcf, 0x7d,
	0x9c, 0x73, 0xc7, 0x19, 0xd8, 0xed, 0xb1, 0xc9, 0x84, 0x45, 0x0d, 0xf9, 0x53, 0x9f, 0xc6, 0x2c,
	0x61, 0xa8, 0x20, 0x51, 0xe5, 0x70, 0xc0, 0xd8, 0x60, 0x4c, 0x1b, 0x22, 0xda, 0x9d, 0xdf, 0x34,
	0x92, 0xd1, 0x84, 0xce, 0x92, 0x70, 0x32, 0x95, 0x89, 0xe6, 0x57, 0x05, 0x0a, 0xe7, 0x34, 0xec,
	0xd3, 0x18, 0xbd, 0x05, 0xad, 0x37, 0x0c, 0x47, 0x91, 0x84, 0x86, 0x52, 0x55, 0x6a, 0xda, 0xc9,
	0x6e, 0x3d, 0xed, 0xdb, 0x5a, 0x52, 0x64, 0x35, 0x0f, 0x59, 0xb0, 0x33, 0x1b, 0x0d, 0xa2, 0x30,
	0x99, 0xc7, 0x34, 0x2d, 0xdd, 0x14, 0xa5, 0x4f, 0xb2, 0x52, 0x7f, 0x9d, 0x26, 0x0f, 0xf3, 0xcd,
	0x3f, 0x0a, 0x68, 0x2b, 0xfd, 0x11, 0x82, 0x5c, 0x72, 0x3f, 0xa5, 0x42, 0x42, 0x9e, 0x88, 0x33,
	0x32, 0xa0, 0x78, 0x4b, 0xe3, 0xd9, 0x88, 0x45, 0xa2, 0x7d, 0x9e, 0x64, 0x10, 0xbd, 0x83, 0xd2,
	0xc2, 0x95, 0xa1, 0x8a, 0xd1, 0x95, 0xba, 0xf4, 0x5d, 0xcf, 0x7c, 0xd7, 0x83, 0x2c, 0x83, 0x2c,
	0x93, 0x79, 0x4f, 0xe1, 0xc4, 0xb1, 0x8d, 0x5c, 0x55, 0xa9, 0x95, 0x49, 0x06, 0x85, 0x82, 0x3b,
	0xc7, 0x36, 0xf2, 0x55, 0xa5, 0x56, 0x22, 0xe2, 0x8c, 0xf6, 0x20, 0x4f, 0xa7, 0xac, 0x37, 0x34,
	0x0a, 0x55, 0xa5, 0x96, 0x23, 0x12, 0xa0, 0x67, 0x50, 0xa2, 0x77, 0x09, 0x8d, 0x84, 0xb2, 0xa2,
	0xe8, 0xb2, 0x0c, 0x98, 0x16, 0xec, 0x3c, 0x70, 0x2f, 0x86, 0xc6, 0x34, 0x4c, 0x98, 0x5c, 0x31,
	0x1f, 0x2a, 0x21, 0x1f, 0x10, 0xb1, 0xa8, 0x47, 0x85, 0xc1, 0x32, 0x91, 0xc0, 0xc4, 0x50, 0xbc,
	0x0a, 0xef, 0xc7, 0x2c, 0xec, 0xa3, 0x97, 0x50, 0x18, 0xae, 0x5e, 0xce, 0x76, 0xb6, 0xe1, 0x74,
	0xb1, 0x29, 0xcb, 0xd5, 0xf7, 0xc3, 0x24, 0x4c, 0xfb, 0x88, 0xb3, 0xd9, 0x84, 0x2d, 0x1c, 0xdd,
	0xd2, 0x31, 0x93, 0xbb, 0x9c, 0xca, 0x96, 0x99, 0x84, 0x14, 0x72, 0x37, 0x8b, 0xcb, 0x49, 0xcb,
	0x97, 0x01, 0xf3, 0xbb, 0x02, 0xf9, 0xe6, 0x98, 0xf5, 0xbe, 0xa0, 0xe3, 0xec, 0xab, 0x79, 0xf8,
	0x99, 0x08, 0x3a, 0x93, 0x93, 0x3a, 0x7e, 0x01, 0x39, 0x3b, 0x93, 0xa3, 0x9d, 0x3c, 0x5e, 0x4b,
	0xe5, 0x04, 0x11, 0x34, 0x7a, 0x0d, 0x5b, 0x97, 0x34, 0x09, 0x85, 0x72, 0x79, 0x8d, 0xfb, 0x6b,
	0xa9, 0x19, 0x49, 0x16, 0x69, 0x26, 0x05, 0x6d, 0x65, 0x20, 0x3a, 0x80, 0x82, 0x3b, 0x9f, 0x74,
	0x53, 0x55, 0x39, 0x92, 0x22, 0x64, 0x42, 0xf9, 0x2a, 0xa6, 0xb7, 0x23, 0x36, 0x9f, 0x9d, 0x87,
	0xb3, 0x61, 0x6a, 0x6c, 0x2d, 0x86, 0x2a, 0xb0, 0xc5, 0x55, 0x08, 0x5e, 0x15, 0xfc, 0x02, 0x9b,
	0x87, 0x50, 0x5a, 0x88, 0xe5, 0xcb, 0x15, 0x6e, 0x94, 0xaa, 0xca, 0x97, 0xcb, 0xcf, 0xe6, 0x31,
	0x3c, 0x5a, 0x93, 0xc8, 0xbb, 0x2d, 0xbc, 0xc8, 0xc4, 0x05, 0x3e, 0xfa, 0xa6, 0x40, 0xc1, 0x4f,
	0xc2, 0x64, 0x3e, 0x43, 0x1a, 0x14, 0xdb, 0xee, 0x7b, 0xd7, 0xfb, 0xe8, 0xea, 0x1b, 0xa8, 0x0c,
	0x45, 0xbf, 0xdd, 0x6a, 0x61, 0xdf, 0xd7, 0x7f, 0x2b, 0x48, 0x07, 0xad, 0x69, 0xd9, 0x1d, 0x82,
	0x3f, 0xb4, 0xb1, 0x1f, 0xe8, 0x3f, 0x54, 0xb4, 0x0d, 0xa5, 0x53, 0x8f, 0x34, 0x1d, 0xdb, 0xc6,
	0xae, 0xfe, 0x53, 0x60, 0xd7, 0x0b, 0x3a, 0xa7, 0x5e, 0xdb, 0xb5, 0xf5, 0x5f, 0x2a, 0xaa, 0xc0,
	0xbe, 0xe3, 0x06, 0x98, 0xb8, 0xd6, 0x45, 0xc7, 0xc7, 0xe4, 0x1a, 0x93, 0x0e, 0x26, 0xc4, 0x23,
	0xfa, 0x5f, 0x15, 0x19, 0xb0, 0xcb, 0x43, 0x4e, 0x0b, 0x77, 0xda, 0xae, 0x75, 0x6d, 0x39, 0x17,
	0x56, 0xf3, 0x02, 0xeb, 0xff, 0xd4, 0xa3, 0xcf, 0x00, 0x72, 0x7b, 0x01, 0xff, 0x97, 0x69, 0x50,
	0xbc, 0xc4, 0xbe, 0x6f, 0x9d, 0x61, 0x7d, 0x03, 0x3d, 0x87, 0xa7, 0x2d, 0xcf, 0x3d, 0x75, 0xce,
	0xda, 0xc4, 0x0a, 0x1c, 0xcf, 0xed, 0x04, 0xc4, 0x72, 0x7d, 0xab, 0xc5, 0xcf, 0xba, 0x82, 0x0e,
	0x00, 0xad, 0xd3, 0x4e, 0x80, 0x2f, 0xf5, 0x4d, 0x64, 0xc0, 0x1e, 0x76, 0x6d, 0x8f, 0xf8, 0x98,
	0xac, 0x55, 0xa8, 0xcd, 0x57, 0x9f, 0x8e, 0x07, 0xa3, 0x64, 0x38, 0xef, 0xf2, 0x7b, 0x6d, 0x0c,
	0xef, 0xa7, 0x34, 0x1e, 0xd3, 0xfe, 0x80, 0xc6, 0x8d, 0x9b, 0xb0, 0x1b, 0x8f, 0x7a, 0xf2, 0x99,
	0x9a, 0xa5, 0x4f, 0x59, 0xb7, 0x20, 0xe0, 0x9b, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x46,
	0x37, 0x79, 0xe2, 0x04, 0x00, 0x00,
}