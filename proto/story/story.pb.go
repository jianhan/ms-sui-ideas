// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/story/story.proto

/*
Package story is a generated protocol buffer package.

It is generated from these files:
	proto/story/story.proto

It has these top-level messages:
	DeleteStoriesRequest
	DeleteStoriesResponse
	Stories
	Story
*/
package story

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

// Pirority defines a set of pirorities
type Pirority int32

const (
	Pirority_UNSET        Pirority = 0
	Pirority_NICE_TO_HAVE Pirority = 1
	Pirority_SHOULD_HAVE  Pirority = 2
	Pirority_MUST_HAVE    Pirority = 3
)

var Pirority_name = map[int32]string{
	0: "UNSET",
	1: "NICE_TO_HAVE",
	2: "SHOULD_HAVE",
	3: "MUST_HAVE",
}
var Pirority_value = map[string]int32{
	"UNSET":        0,
	"NICE_TO_HAVE": 1,
	"SHOULD_HAVE":  2,
	"MUST_HAVE":    3,
}

func (x Pirority) String() string {
	return proto.EnumName(Pirority_name, int32(x))
}
func (Pirority) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DeleteStoriesRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *DeleteStoriesRequest) Reset()                    { *m = DeleteStoriesRequest{} }
func (m *DeleteStoriesRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteStoriesRequest) ProtoMessage()               {}
func (*DeleteStoriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteStoriesRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type DeleteStoriesResponse struct {
}

func (m *DeleteStoriesResponse) Reset()                    { *m = DeleteStoriesResponse{} }
func (m *DeleteStoriesResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteStoriesResponse) ProtoMessage()               {}
func (*DeleteStoriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Stories struct {
	Stories []*Story `protobuf:"bytes,1,rep,name=stories" json:"stories,omitempty"`
}

func (m *Stories) Reset()                    { *m = Stories{} }
func (m *Stories) String() string            { return proto.CompactTextString(m) }
func (*Stories) ProtoMessage()               {}
func (*Stories) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Stories) GetStories() []*Story {
	if m != nil {
		return m.Stories
	}
	return nil
}

type Story struct {
	// @inject_tag: bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"title" valid:"required~Title is required" conform:"trim"
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty" bson:"title" valid:"required~Title is required" conform:"trim"`
	// @inject_tag: bson:"pirority" valid:"required~Pirority is required"
	Pirority Pirority `protobuf:"varint,3,opt,name=pirority,enum=go.micro.srv.ideas.Pirority" json:"pirority,omitempty" bson:"pirority" valid:"required~Pirority is required"`
	// @inject_tag: bson:"role" valid:"required~Who is required" conform:"trim"
	Who string `protobuf:"bytes,4,opt,name=who" json:"who,omitempty" bson:"role" valid:"required~Who is required" conform:"trim"`
	// @inject_tag: bson:"action" valid:"required~What is required" conform:"trim"
	What string `protobuf:"bytes,5,opt,name=what" json:"what,omitempty" bson:"action" valid:"required~What is required" conform:"trim"`
	// @inject_tag: bson:"goal" valid:"required~Goal is required" conform:"trim"
	Goal string `protobuf:"bytes,6,opt,name=goal" json:"goal,omitempty" bson:"goal" valid:"required~Goal is required" conform:"trim"`
	// @inject_tag: bson:"order"
	Order uint64 `protobuf:"varint,7,opt,name=order" json:"order,omitempty" bson:"order"`
	// @inject_tag: bson:"idea_id" valid:"required~Idea ID is required, uuidv4~Idea id must be a valid UUIDv4" conform:"trim"
	IdeaId string `protobuf:"bytes,8,opt,name=idea_id,json=ideaId" json:"idea_id,omitempty" bson:"idea_id" valid:"required~Idea ID is required, uuidv4~Idea id must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at"`
}

func (m *Story) Reset()                    { *m = Story{} }
func (m *Story) String() string            { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()               {}
func (*Story) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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

func (m *Story) GetPirority() Pirority {
	if m != nil {
		return m.Pirority
	}
	return Pirority_UNSET
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

func (m *Story) GetIdeaId() string {
	if m != nil {
		return m.IdeaId
	}
	return ""
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
	proto.RegisterType((*DeleteStoriesRequest)(nil), "go.micro.srv.ideas.DeleteStoriesRequest")
	proto.RegisterType((*DeleteStoriesResponse)(nil), "go.micro.srv.ideas.DeleteStoriesResponse")
	proto.RegisterType((*Stories)(nil), "go.micro.srv.ideas.stories")
	proto.RegisterType((*Story)(nil), "go.micro.srv.ideas.Story")
	proto.RegisterEnum("go.micro.srv.ideas.Pirority", Pirority_name, Pirority_value)
}

func init() { proto.RegisterFile("proto/story/story.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x8f, 0xd2, 0x40,
	0x14, 0xc7, 0x69, 0xa1, 0x94, 0x3e, 0x64, 0x25, 0x2f, 0x6b, 0x18, 0xd1, 0xc4, 0xa6, 0xa7, 0xea,
	0xa1, 0x24, 0xec, 0x45, 0x2f, 0x26, 0x28, 0x44, 0x88, 0xca, 0x9a, 0x02, 0x1e, 0xbc, 0x90, 0xee,
	0x76, 0x96, 0x9d, 0x04, 0x9c, 0x3a, 0x33, 0x48, 0xf8, 0x06, 0x7e, 0x32, 0x3f, 0x97, 0x99, 0x99,
	0x76, 0x13, 0x95, 0xa8, 0x07, 0x2f, 0xe4, 0xbd, 0x3f, 0xbf, 0xf7, 0xe6, 0xcd, 0xfc, 0x5f, 0xa1,
	0x57, 0x08, 0xae, 0xf8, 0x40, 0x2a, 0x2e, 0x8e, 0xf6, 0x37, 0x31, 0x0a, 0xe2, 0x86, 0x27, 0x3b,
	0x76, 0x2d, 0x78, 0x22, 0xc5, 0xd7, 0x84, 0xe5, 0x34, 0x93, 0xfd, 0x27, 0x1b, 0xce, 0x37, 0x5b,
	0x3a, 0x30, 0xc4, 0xd5, 0xfe, 0x66, 0xa0, 0xd8, 0x8e, 0x4a, 0x95, 0xed, 0x0a, 0x5b, 0x14, 0xc5,
	0x70, 0x3e, 0xa6, 0x5b, 0xaa, 0xe8, 0x42, 0x71, 0xc1, 0xa8, 0x4c, 0xe9, 0x97, 0x3d, 0x95, 0x0a,
	0xbb, 0x50, 0x67, 0xb9, 0x24, 0x4e, 0x58, 0x8f, 0x83, 0x54, 0x87, 0x51, 0x0f, 0x1e, 0xfc, 0x42,
	0xca, 0x82, 0x7f, 0x96, 0x34, 0x7a, 0x09, 0xbe, 0xb4, 0x12, 0x5e, 0xdc, 0x85, 0xa6, 0xb2, 0x3d,
	0x7c, 0x98, 0xfc, 0x3e, 0x54, 0xa2, 0x1b, 0x1c, 0xd3, 0x8a, 0x8c, 0xbe, 0xbb, 0xe0, 0x19, 0x09,
	0xcf, 0xc0, 0x9d, 0x8d, 0x89, 0x13, 0x3a, 0x71, 0x90, 0xba, 0xb3, 0x31, 0x9e, 0x83, 0xa7, 0x98,
	0xda, 0x52, 0xe2, 0x1a, 0xc9, 0x26, 0xf8, 0x1c, 0x5a, 0x05, 0x13, 0x5c, 0x30, 0x75, 0x24, 0xf5,
	0xd0, 0x89, 0xcf, 0x86, 0x8f, 0x4f, 0x9d, 0xf2, 0xa1, 0x64, 0xd2, 0x3b, 0x5a, 0x5f, 0xea, 0x70,
	0xcb, 0x49, 0xc3, 0x74, 0xd3, 0x21, 0x22, 0x34, 0x0e, 0xb7, 0x99, 0x22, 0x9e, 0x91, 0x4c, 0xac,
	0xb5, 0x0d, 0xcf, 0xb6, 0xa4, 0x69, 0x35, 0x1d, 0xeb, 0x49, 0xb8, 0xc8, 0xa9, 0x20, 0x7e, 0xe8,
	0xc4, 0x8d, 0xd4, 0x26, 0xd8, 0x03, 0x5f, 0x9f, 0xb5, 0x66, 0x39, 0x69, 0x19, 0xb8, 0xa9, 0xd3,
	0x59, 0x8e, 0x2f, 0x00, 0xae, 0x05, 0xcd, 0x14, 0xcd, 0xd7, 0x99, 0x22, 0x41, 0xe8, 0xc4, 0xed,
	0x61, 0x3f, 0xb1, 0x5e, 0x24, 0x95, 0x17, 0xc9, 0xb2, 0xf2, 0x22, 0x0d, 0x4a, 0x7a, 0xa4, 0x74,
	0xe9, 0xbe, 0xc8, 0xab, 0x52, 0xf8, 0x7b, 0x69, 0x49, 0x8f, 0xd4, 0xb3, 0x37, 0xd0, 0xaa, 0x2e,
	0x8d, 0x01, 0x78, 0xab, 0xf9, 0x62, 0xb2, 0xec, 0xd6, 0xb0, 0x0b, 0xf7, 0xe6, 0xb3, 0xd7, 0x93,
	0xf5, 0xf2, 0x72, 0x3d, 0x1d, 0x7d, 0x9c, 0x74, 0x1d, 0xbc, 0x0f, 0xed, 0xc5, 0xf4, 0x72, 0xf5,
	0x6e, 0x6c, 0x05, 0x17, 0x3b, 0x10, 0xbc, 0x5f, 0x2d, 0x96, 0x36, 0xad, 0x0f, 0xbf, 0xb9, 0xe0,
	0x97, 0x2e, 0xe3, 0x14, 0x60, 0x4e, 0x0f, 0x55, 0xf6, 0xe8, 0xd4, 0x4b, 0x97, 0x46, 0xf6, 0xff,
	0xf4, 0x67, 0x54, 0xc3, 0xb7, 0xd0, 0x59, 0x99, 0x59, 0xff, 0x47, 0xb3, 0x1b, 0xe8, 0xfc, 0xb4,
	0x8d, 0x18, 0x9f, 0xe2, 0x4f, 0xad, 0x76, 0xff, 0xe9, 0x3f, 0x90, 0xe5, 0x6a, 0xd7, 0x5e, 0xf9,
	0x9f, 0x3c, 0xf3, 0x8d, 0x5d, 0x35, 0xcd, 0xdb, 0x5f, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc6,
	0x5c, 0x4d, 0x64, 0x7f, 0x03, 0x00, 0x00,
}
