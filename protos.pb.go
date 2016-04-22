// Code generated by protoc-gen-go.
// source: protos.proto
// DO NOT EDIT!

/*
Package gam is a generated protocol buffer package.

It is generated from these files:
	protos.proto

It has these top-level messages:
	PID
	Restarting
	Stopping
	Stopped
	PoisonPill
	Started
*/
package gam

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type PID struct {
	Node string `protobuf:"bytes,1,opt,name=Node,json=node" json:"Node,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=Host,json=host" json:"Host,omitempty"`
	Id   uint64 `protobuf:"varint,3,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *PID) Reset()                    { *m = PID{} }
func (m *PID) String() string            { return proto.CompactTextString(m) }
func (*PID) ProtoMessage()               {}
func (*PID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// user messages
type Restarting struct {
}

func (m *Restarting) Reset()                    { *m = Restarting{} }
func (m *Restarting) String() string            { return proto.CompactTextString(m) }
func (*Restarting) ProtoMessage()               {}
func (*Restarting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Stopping struct {
}

func (m *Stopping) Reset()                    { *m = Stopping{} }
func (m *Stopping) String() string            { return proto.CompactTextString(m) }
func (*Stopping) ProtoMessage()               {}
func (*Stopping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Stopped struct {
}

func (m *Stopped) Reset()                    { *m = Stopped{} }
func (m *Stopped) String() string            { return proto.CompactTextString(m) }
func (*Stopped) ProtoMessage()               {}
func (*Stopped) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PoisonPill struct {
}

func (m *PoisonPill) Reset()                    { *m = PoisonPill{} }
func (m *PoisonPill) String() string            { return proto.CompactTextString(m) }
func (*PoisonPill) ProtoMessage()               {}
func (*PoisonPill) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Started struct {
}

func (m *Started) Reset()                    { *m = Started{} }
func (m *Started) String() string            { return proto.CompactTextString(m) }
func (*Started) ProtoMessage()               {}
func (*Started) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*PID)(nil), "gam.PID")
	proto.RegisterType((*Restarting)(nil), "gam.Restarting")
	proto.RegisterType((*Stopping)(nil), "gam.Stopping")
	proto.RegisterType((*Stopped)(nil), "gam.Stopped")
	proto.RegisterType((*PoisonPill)(nil), "gam.PoisonPill")
	proto.RegisterType((*Started)(nil), "gam.Started")
}

var fileDescriptor0 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x03, 0x53, 0x42, 0xcc, 0xe9, 0x89, 0xb9, 0x4a, 0xb6, 0x5c, 0xcc, 0x01, 0x9e,
	0x2e, 0x42, 0x42, 0x5c, 0x2c, 0x7e, 0xf9, 0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x2c, 0x79, 0x40, 0x36, 0x48, 0xcc, 0x23, 0xbf, 0xb8, 0x44, 0x82, 0x09, 0x22, 0x96, 0x01, 0x64,
	0x0b, 0xf1, 0x71, 0x31, 0x79, 0xa6, 0x48, 0x30, 0x03, 0x45, 0x58, 0x82, 0x98, 0x32, 0x53, 0x94,
	0x78, 0xb8, 0xb8, 0x82, 0x52, 0x8b, 0x4b, 0x12, 0x8b, 0x4a, 0x32, 0xf3, 0xd2, 0x95, 0xb8, 0xb8,
	0x38, 0x82, 0x4b, 0xf2, 0x0b, 0x0a, 0x40, 0x6c, 0x4e, 0x2e, 0x76, 0x30, 0x3b, 0x15, 0xac, 0x28,
	0x20, 0x3f, 0xb3, 0x38, 0x3f, 0x2f, 0x20, 0x33, 0x27, 0x07, 0x22, 0x01, 0xd4, 0x90, 0x9a, 0x92,
	0xc4, 0x06, 0x76, 0x88, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x5b, 0x37, 0x4e, 0x98, 0x00,
	0x00, 0x00,
}
