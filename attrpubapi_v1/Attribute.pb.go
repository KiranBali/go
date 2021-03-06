// Code generated by protoc-gen-go.
// source: Attribute.proto
// DO NOT EDIT!

/*
Package attrpubapi_v1 is a generated protocol buffer package.

It is generated from these files:
	Attribute.proto
	List.proto
	Signing.proto

It has these top-level messages:
	Attribute
	Anchor
	AttributeAndId
	AttributeAndIdList
	AttributeList
	AttributeSigning
*/
package attrpubapi_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ContentType indicates how to interpret the ‘Value’ field of an Attribute.
type ContentType int32

const (
	// UNDEFINED should not be seen, and is used as an error placeholder
	// value.
	ContentType_UNDEFINED ContentType = 0
	// STRING means the value is UTF-8 encoded text.
	ContentType_STRING ContentType = 1
	// JPEG indicates a standard .jpeg image.
	ContentType_JPEG ContentType = 2
	// Date as string in RFC3339 format (YYYY-MM-DD).
	ContentType_DATE ContentType = 3
	// PNG indicates a standard .png image.
	ContentType_PNG ContentType = 4
)

var ContentType_name = map[int32]string{
	0: "UNDEFINED",
	1: "STRING",
	2: "JPEG",
	3: "DATE",
	4: "PNG",
}
var ContentType_value = map[string]int32{
	"UNDEFINED": 0,
	"STRING":    1,
	"JPEG":      2,
	"DATE":      3,
	"PNG":       4,
}

func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}
func (ContentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Attribute struct {
	Name        string      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value       []byte      `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ContentType ContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,enum=attrpubapi_v1.ContentType" json:"content_type,omitempty"`
	Anchors     []*Anchor   `protobuf:"bytes,4,rep,name=anchors" json:"anchors,omitempty"`
}

func (m *Attribute) Reset()                    { *m = Attribute{} }
func (m *Attribute) String() string            { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()               {}
func (*Attribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Attribute) GetAnchors() []*Anchor {
	if m != nil {
		return m.Anchors
	}
	return nil
}

type Anchor struct {
	ArtifactLink      []byte   `protobuf:"bytes,1,opt,name=artifact_link,json=artifactLink,proto3" json:"artifact_link,omitempty"`
	OriginServerCerts [][]byte `protobuf:"bytes,2,rep,name=origin_server_certs,json=originServerCerts,proto3" json:"origin_server_certs,omitempty"`
	ArtifactSignature []byte   `protobuf:"bytes,3,opt,name=artifact_signature,json=artifactSignature,proto3" json:"artifact_signature,omitempty"`
	SubType           string   `protobuf:"bytes,4,opt,name=sub_type,json=subType" json:"sub_type,omitempty"`
	Signature         []byte   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	SignedTimeStamp   []byte   `protobuf:"bytes,6,opt,name=signed_time_stamp,json=signedTimeStamp,proto3" json:"signed_time_stamp,omitempty"`
}

func (m *Anchor) Reset()                    { *m = Anchor{} }
func (m *Anchor) String() string            { return proto.CompactTextString(m) }
func (*Anchor) ProtoMessage()               {}
func (*Anchor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Attribute)(nil), "attrpubapi_v1.Attribute")
	proto.RegisterType((*Anchor)(nil), "attrpubapi_v1.Anchor")
	proto.RegisterEnum("attrpubapi_v1.ContentType", ContentType_name, ContentType_value)
}

func init() { proto.RegisterFile("Attribute.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x6e, 0x9b, 0x40,
	0x10, 0xc6, 0x8b, 0x21, 0x76, 0x3c, 0xc6, 0x0d, 0x9e, 0xfe, 0x11, 0xad, 0x7a, 0x40, 0xc9, 0x05,
	0x45, 0x2a, 0x55, 0xd3, 0x73, 0x0f, 0x4e, 0x4c, 0xad, 0x54, 0x15, 0xb2, 0x16, 0x7a, 0x46, 0x0b,
	0xdd, 0xa6, 0xab, 0x84, 0x05, 0x2d, 0x83, 0x25, 0x3f, 0x50, 0x1f, 0xb0, 0x6f, 0x50, 0xb1, 0x88,
	0xb8, 0xe9, 0x6d, 0xe6, 0xfb, 0x7d, 0x7c, 0x62, 0x66, 0x07, 0xce, 0xd6, 0x44, 0x5a, 0x16, 0x1d,
	0x89, 0xa8, 0xd1, 0x35, 0xd5, 0xb8, 0xe4, 0x44, 0xba, 0xe9, 0x0a, 0xde, 0xc8, 0x7c, 0xff, 0xf1,
	0xfc, 0xb7, 0x05, 0xf3, 0x47, 0x0b, 0x22, 0x38, 0x8a, 0x57, 0xc2, 0xb7, 0x02, 0x2b, 0x9c, 0x33,
	0x53, 0xe3, 0x4b, 0x38, 0xd9, 0xf3, 0x87, 0x4e, 0xf8, 0x93, 0xc0, 0x0a, 0x5d, 0x36, 0x34, 0xf8,
	0x19, 0xdc, 0xb2, 0x56, 0x24, 0x14, 0xe5, 0x74, 0x68, 0x84, 0x6f, 0x07, 0x56, 0xf8, 0xfc, 0xea,
	0x6d, 0xf4, 0x24, 0x3d, 0xba, 0x19, 0x2c, 0xd9, 0xa1, 0x11, 0x6c, 0x51, 0x1e, 0x1b, 0xfc, 0x00,
	0x33, 0xae, 0xca, 0x5f, 0xb5, 0x6e, 0x7d, 0x27, 0xb0, 0xc3, 0xc5, 0xd5, 0xab, 0xff, 0xbe, 0x5c,
	0x1b, 0xca, 0x46, 0xd7, 0xf9, 0x1f, 0x0b, 0xa6, 0x83, 0x86, 0x17, 0xb0, 0xe4, 0x9a, 0xe4, 0x4f,
	0x5e, 0x52, 0xfe, 0x20, 0xd5, 0xbd, 0xf9, 0x5b, 0x97, 0xb9, 0xa3, 0xf8, 0x4d, 0xaa, 0x7b, 0x8c,
	0xe0, 0x45, 0xad, 0xe5, 0x9d, 0x54, 0x79, 0x2b, 0xf4, 0x5e, 0xe8, 0xbc, 0x14, 0x9a, 0x5a, 0x7f,
	0x12, 0xd8, 0xa1, 0xcb, 0x56, 0x03, 0x4a, 0x0d, 0xb9, 0xe9, 0x01, 0xbe, 0x07, 0x7c, 0x0c, 0x6d,
	0xe5, 0x9d, 0xe2, 0xd4, 0xe9, 0x61, 0x2a, 0x97, 0xad, 0x46, 0x92, 0x8e, 0x00, 0xdf, 0xc0, 0x69,
	0xdb, 0x15, 0xc3, 0xe8, 0x8e, 0x59, 0xd6, 0xac, 0xed, 0x0a, 0x33, 0xda, 0x3b, 0x98, 0x1f, 0x03,
	0x4e, 0x4c, 0xc0, 0x51, 0xc0, 0x4b, 0x58, 0xf5, 0x8d, 0xf8, 0x91, 0x93, 0xac, 0x44, 0xde, 0x12,
	0xaf, 0x1a, 0x7f, 0x6a, 0x5c, 0x67, 0x03, 0xc8, 0x64, 0x25, 0xd2, 0x5e, 0xbe, 0x8c, 0x61, 0xf1,
	0xcf, 0x02, 0x71, 0x09, 0xf3, 0xef, 0xc9, 0x26, 0xfe, 0x72, 0x9b, 0xc4, 0x1b, 0xef, 0x19, 0x02,
	0x4c, 0xd3, 0x8c, 0xdd, 0x26, 0x5b, 0xcf, 0xc2, 0x53, 0x70, 0xbe, 0xee, 0xe2, 0xad, 0x37, 0xe9,
	0xab, 0xcd, 0x3a, 0x8b, 0x3d, 0x1b, 0x67, 0x60, 0xef, 0x92, 0xad, 0xe7, 0x5c, 0x5f, 0xc0, 0xeb,
	0xb2, 0xae, 0xa2, 0x43, 0x4d, 0xf2, 0xe9, 0x92, 0xaf, 0xcd, 0xcb, 0xef, 0xfa, 0xb3, 0x28, 0xa6,
	0xe6, 0x3a, 0x3e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x46, 0x81, 0xbe, 0xb5, 0x30, 0x02, 0x00,
	0x00,
}
