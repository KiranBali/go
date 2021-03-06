// Code generated by protoc-gen-go.
// source: common-public-api/compubapi_v1/EncryptedData.proto
// DO NOT EDIT!

/*
Package compubapi_v1 is a generated protocol buffer package.

It is generated from these files:
	common-public-api/compubapi_v1/EncryptedData.proto

It has these top-level messages:
	EncryptedData
*/
package compubapi_v1

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

type EncryptedData struct {
	// the iv will be used in conjunction with the secret key
	// received via other channel in order to decrypt the cipher_text
	Iv []byte `protobuf:"bytes,1,opt,name=iv,proto3" json:"iv,omitempty"`
	// block of bytes to be decrypted
	CipherText []byte `protobuf:"bytes,2,opt,name=cipher_text,json=cipherText,proto3" json:"cipher_text,omitempty"`
}

func (m *EncryptedData) Reset()                    { *m = EncryptedData{} }
func (m *EncryptedData) String() string            { return proto.CompactTextString(m) }
func (*EncryptedData) ProtoMessage()               {}
func (*EncryptedData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*EncryptedData)(nil), "compubapi_v1.EncryptedData")
}

func init() { proto.RegisterFile("common-public-api/compubapi_v1/EncryptedData.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2d, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xd6, 0x4d, 0x2c, 0xc8, 0xd4, 0x4f, 0xce,
	0xcf, 0x2d, 0x28, 0x4d, 0x4a, 0x2c, 0xc8, 0x8c, 0x2f, 0x33, 0xd4, 0x77, 0xcd, 0x4b, 0x2e, 0xaa,
	0x2c, 0x28, 0x49, 0x4d, 0x71, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x56, 0xa1, 0xe4, 0xc0, 0xc5, 0x8b, 0xa2, 0x48, 0x88, 0x8f, 0x8b, 0x29, 0xb3, 0x4c, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x27, 0x88, 0x29, 0xb3, 0x4c, 0x48, 0x9e, 0x8b, 0x3b, 0x39, 0xb3, 0x20,
	0x23, 0xb5, 0x28, 0xbe, 0x24, 0xb5, 0xa2, 0x44, 0x82, 0x09, 0x2c, 0xc1, 0x05, 0x11, 0x0a, 0x49,
	0xad, 0x28, 0x71, 0xd2, 0xe6, 0x12, 0x4d, 0xce, 0xcf, 0xd5, 0xab, 0xcc, 0x2f, 0xc9, 0xd4, 0x43,
	0x36, 0xda, 0x49, 0x08, 0xc5, 0xe0, 0x00, 0x90, 0xe5, 0x49, 0x6c, 0x60, 0x37, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0x08, 0xe9, 0x4b, 0xb9, 0x00, 0x00, 0x00,
}
