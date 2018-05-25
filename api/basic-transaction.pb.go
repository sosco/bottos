// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bottos-project/bottos/api/basic-transaction.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// BasicTransaction definition for serialization
type BasicTransaction struct {
	Version     uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       []byte `protobuf:"bytes,8,opt,name=param,proto3" json:"param"`
	SigAlg      uint32 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
}

func (m *BasicTransaction) Reset()                    { *m = BasicTransaction{} }
func (m *BasicTransaction) String() string            { return proto.CompactTextString(m) }
func (*BasicTransaction) ProtoMessage()               {}
func (*BasicTransaction) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *BasicTransaction) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BasicTransaction) GetCursorNum() uint32 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *BasicTransaction) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *BasicTransaction) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *BasicTransaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *BasicTransaction) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *BasicTransaction) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *BasicTransaction) GetParam() []byte {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *BasicTransaction) GetSigAlg() uint32 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func init() {
	proto.RegisterType((*BasicTransaction)(nil), "api.BasicTransaction")
}

func init() {
	proto.RegisterFile("github.com/bottos-project/bottos/api/basic-transaction.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3d, 0x4b, 0xf4, 0x40,
	0x14, 0x85, 0xc9, 0x7e, 0x24, 0x9b, 0x79, 0xf7, 0x05, 0x19, 0x44, 0x07, 0x41, 0x88, 0x22, 0x92,
	0x66, 0x37, 0x85, 0xad, 0x8d, 0x5b, 0x8b, 0x45, 0xb0, 0xb2, 0x59, 0x6e, 0x66, 0xc7, 0xec, 0x95,
	0xcc, 0xdc, 0x30, 0x73, 0xe3, 0x3f, 0xf0, 0x7f, 0x4b, 0x3e, 0x5c, 0x5b, 0xcb, 0xe7, 0x3c, 0xf7,
	0xc0, 0xe5, 0x88, 0xc7, 0x1a, 0xf9, 0xd8, 0x55, 0x5b, 0x4d, 0xb6, 0xa8, 0x88, 0x99, 0xc2, 0xa6,
	0xf5, 0xf4, 0x61, 0x34, 0x4f, 0x58, 0x40, 0x8b, 0x45, 0x05, 0x01, 0xf5, 0x86, 0x3d, 0xb8, 0x00,
	0x9a, 0x91, 0xdc, 0xb6, 0xf5, 0xc4, 0x24, 0xe7, 0xd0, 0xe2, 0xed, 0xd7, 0x4c, 0x9c, 0xed, 0xfa,
	0x83, 0xd7, 0x5f, 0x2f, 0x95, 0x48, 0x3e, 0x8d, 0x0f, 0x48, 0x4e, 0x45, 0x59, 0x94, 0xff, 0x2f,
	0x7f, 0x50, 0x5e, 0x0b, 0xa1, 0x3b, 0x1f, 0xc8, 0xef, 0x5d, 0x67, 0xd5, 0x6c, 0x90, 0xe9, 0x98,
	0xbc, 0x74, 0x56, 0xde, 0x88, 0xf5, 0xa4, 0x1b, 0xa8, 0x4c, 0xa3, 0xe6, 0xc3, 0xc1, 0xbf, 0x31,
	0x7b, 0xee, 0x23, 0x79, 0x25, 0x56, 0x0d, 0xbe, 0x1b, 0x46, 0x6b, 0xd4, 0x22, 0x8b, 0xf2, 0x45,
	0x79, 0x62, 0x79, 0x21, 0xe2, 0x60, 0xdc, 0xc1, 0x78, 0xb5, 0xcc, 0xa2, 0x3c, 0x2d, 0x27, 0xea,
	0x3b, 0x9a, 0x1c, 0x7b, 0xd0, 0xac, 0xe2, 0xc1, 0x9c, 0xb8, 0xef, 0x58, 0xc3, 0x47, 0x3a, 0xa8,
	0x64, 0xec, 0x8c, 0x24, 0xcf, 0xc5, 0xb2, 0x05, 0x0f, 0x56, 0xad, 0xb2, 0x28, 0x5f, 0x97, 0x23,
	0xc8, 0x4b, 0x91, 0x04, 0xac, 0xf7, 0xd0, 0xd4, 0x2a, 0x1d, 0x7e, 0x8b, 0x03, 0xd6, 0x4f, 0x4d,
	0xbd, 0xbb, 0x7f, 0xbb, 0xfb, 0xcb, 0x98, 0x55, 0x3c, 0x6c, 0xf7, 0xf0, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x10, 0x9f, 0x46, 0xa6, 0x7b, 0x01, 0x00, 0x00,
}
