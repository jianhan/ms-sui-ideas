// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/story/story.proto

/*
Package ideas is a generated protocol buffer package.

It is generated from these files:
	proto/story/story.proto

It has these top-level messages:
	Story
*/
package ideas

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import go_micro_srv_ideas "github.com/jianhan/ms-sui-ideas/proto/idea"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Story struct {
	// @inject_tag: bson:"_id" valid:"uuidv4~ID must be a valid UUIDv4" conform:"trim"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"uuidv4~ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"title" valid:"required~Title is required" conform:"trim"
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty" bson:"title" valid:"required~Title is required" conform:"trim"`
	// @inject_tag: bson:"pirority" valid:"required~Pirority is required"
	Pirority go_micro_srv_ideas.Pirority `protobuf:"varint,3,opt,name=pirority,enum=go.micro.srv.ideas.Pirority" json:"pirority,omitempty" bson:"pirority" valid:"required~Pirority is required"`
	// @inject_tag: bson:"role" valid:"required~Who is required" conform:"trim"
	Who string `protobuf:"bytes,4,opt,name=who" json:"who,omitempty" bson:"role" valid:"required~Who is required" conform:"trim"`
	// @inject_tag: bson:"action" valid:"required~What is required" conform:"trim"
	What string `protobuf:"bytes,5,opt,name=what" json:"what,omitempty" bson:"action" valid:"required~What is required" conform:"trim"`
	// @inject_tag: bson:"goal" valid:"required~Goal is required" conform:"trim"
	Goal string `protobuf:"bytes,6,opt,name=goal" json:"goal,omitempty" bson:"goal" valid:"required~Goal is required" conform:"trim"`
	// @inject_tag: bson:"order"
	Order uint64 `protobuf:"varint,7,opt,name=order" json:"order,omitempty" bson:"order"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at"`
}

func (m *Story) Reset()                    { *m = Story{} }
func (m *Story) String() string            { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()               {}
func (*Story) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Story) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Story) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Story) GetPirority() go_micro_srv_ideas.Pirority {
	if m != nil {
		return m.Pirority
	}
	return go_micro_srv_ideas.Pirority_UNSET
}

func (m *Story) GetWho() string {
	if m != nil {
		return m.Who
	}
	return ""
}

func (m *Story) GetWhat() string {
	if m != nil {
		return m.What
	}
	return ""
}

func (m *Story) GetGoal() string {
	if m != nil {
		return m.Goal
	}
	return ""
}

func (m *Story) GetOrder() uint64 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Story) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Story) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Story)(nil), "go.micro.srv.ideas.Story")
}

func init() { proto.RegisterFile("proto/story/story.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x3f, 0x6b, 0xf3, 0x30,
	0x10, 0xc6, 0xb1, 0x13, 0xe7, 0x8f, 0x5e, 0x08, 0x2f, 0xa2, 0x50, 0x11, 0x0a, 0x35, 0x9d, 0xbc,
	0x44, 0x82, 0x94, 0x42, 0x3b, 0xa6, 0x64, 0xc9, 0x56, 0xdc, 0x4e, 0x5d, 0x8a, 0x12, 0xab, 0xb6,
	0x8a, 0x9d, 0x33, 0xd2, 0xb9, 0x21, 0x1f, 0xaa, 0xdf, 0xb1, 0x48, 0xb2, 0xbb, 0x74, 0xe8, 0x72,
	0x3c, 0xf7, 0xf0, 0xfc, 0x4e, 0xa7, 0x23, 0x97, 0xad, 0x01, 0x04, 0x61, 0x11, 0xcc, 0x39, 0x54,
	0xee, 0x1d, 0x4a, 0x4b, 0xe0, 0x8d, 0x3e, 0x18, 0xe0, 0xd6, 0x7c, 0x72, 0x5d, 0x28, 0x69, 0x97,
	0xd7, 0x25, 0x40, 0x59, 0x2b, 0xe1, 0x13, 0xfb, 0xee, 0x5d, 0xa0, 0x6e, 0x94, 0x45, 0xd9, 0xb4,
	0x01, 0x5a, 0xde, 0x95, 0x1a, 0xab, 0x6e, 0xcf, 0x0f, 0xd0, 0x88, 0x0f, 0x2d, 0x8f, 0x95, 0x3c,
	0x8a, 0xc6, 0xae, 0x6c, 0xa7, 0x57, 0x7e, 0x42, 0x20, 0x85, 0xd3, 0xbe, 0x04, 0xec, 0xe6, 0x2b,
	0x26, 0xc9, 0xb3, 0x7b, 0x9b, 0x2e, 0x48, 0xbc, 0xdb, 0xb2, 0x28, 0x8d, 0xb2, 0x79, 0x1e, 0xef,
	0xb6, 0xf4, 0x82, 0x24, 0xa8, 0xb1, 0x56, 0x2c, 0xf6, 0x56, 0x68, 0xe8, 0x3d, 0x99, 0xb5, 0xda,
	0x80, 0xd1, 0x78, 0x66, 0xa3, 0x34, 0xca, 0x16, 0xeb, 0x2b, 0xfe, 0x7b, 0x5d, 0xfe, 0xd4, 0x67,
	0xf2, 0x9f, 0x34, 0xfd, 0x4f, 0x46, 0xa7, 0x0a, 0xd8, 0xd8, 0x4f, 0x73, 0x92, 0x52, 0x32, 0x3e,
	0x55, 0x12, 0x59, 0xe2, 0x2d, 0xaf, 0x9d, 0x57, 0x82, 0xac, 0xd9, 0x24, 0x78, 0x4e, 0xbb, 0x4d,
	0xc0, 0x14, 0xca, 0xb0, 0x69, 0x1a, 0x65, 0xe3, 0x3c, 0x34, 0xf4, 0x81, 0x90, 0x83, 0x51, 0x12,
	0x55, 0xf1, 0x26, 0x91, 0xcd, 0xd2, 0x28, 0xfb, 0xb7, 0x5e, 0xf2, 0x70, 0x26, 0x3e, 0x9c, 0x89,
	0xbf, 0x0c, 0x67, 0xca, 0xe7, 0x7d, 0x7a, 0x83, 0x0e, 0xed, 0xda, 0x62, 0x40, 0xe7, 0x7f, 0xa3,
	0x7d, 0x7a, 0x83, 0x8f, 0xd3, 0xd7, 0xc4, 0xff, 0x70, 0x3f, 0xf1, 0xb9, 0xdb, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x46, 0x63, 0x88, 0xf3, 0xc6, 0x01, 0x00, 0x00,
}
