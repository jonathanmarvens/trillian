// Code generated by protoc-gen-go.
// source: keyspb.proto
// DO NOT EDIT!

/*
Package keyspb is a generated protocol buffer package.

It is generated from these files:
	keyspb.proto

It has these top-level messages:
	PEMKeyFile
	PublicKey
*/
package keyspb

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

// PEMKeyFile identifies a private key stored in a PEM-encoded file.
type PEMKeyFile struct {
	// File path of the private key.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// Password for decrypting the private key.
	// If empty, indicates that the private key is not encrypted.
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *PEMKeyFile) Reset()                    { *m = PEMKeyFile{} }
func (m *PEMKeyFile) String() string            { return proto.CompactTextString(m) }
func (*PEMKeyFile) ProtoMessage()               {}
func (*PEMKeyFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PEMKeyFile) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PEMKeyFile) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// PublicKey is a public key, used for verifying signatures.
type PublicKey struct {
	// The key in DER-encoded PKIX format.
	Der []byte `protobuf:"bytes,1,opt,name=der,proto3" json:"der,omitempty"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PublicKey) GetDer() []byte {
	if m != nil {
		return m.Der
	}
	return nil
}

func init() {
	proto.RegisterType((*PEMKeyFile)(nil), "keyspb.PEMKeyFile")
	proto.RegisterType((*PublicKey)(nil), "keyspb.PublicKey")
}

func init() { proto.RegisterFile("keyspb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4e, 0xad, 0x2c,
	0x2e, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x6c, 0xb8, 0xb8,
	0x02, 0x5c, 0x7d, 0xbd, 0x53, 0x2b, 0xdd, 0x32, 0x73, 0x52, 0x85, 0x84, 0xb8, 0x58, 0x0a, 0x12,
	0x4b, 0x32, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x29, 0x2e, 0x8e, 0x82,
	0xc4, 0xe2, 0xe2, 0xf2, 0xfc, 0xa2, 0x14, 0x09, 0x26, 0xb0, 0x38, 0x9c, 0xaf, 0x24, 0xcb, 0xc5,
	0x19, 0x50, 0x9a, 0x94, 0x93, 0x99, 0xec, 0x9d, 0x5a, 0x29, 0x24, 0xc0, 0xc5, 0x9c, 0x92, 0x5a,
	0x04, 0xd6, 0xcb, 0x13, 0x04, 0x62, 0x26, 0xb1, 0x81, 0xed, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xf6, 0xa6, 0x63, 0x12, 0x7b, 0x00, 0x00, 0x00,
}