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
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x9b, 0xf4, 0x6f, 0xa6, 0xdb, 0xa5, 0x1a, 0x2d, 0xaa, 0x29, 0x48, 0x44, 0x39, 0x05,
	0x0e, 0xa9, 0xd4, 0xbd, 0xc0, 0x05, 0xa9, 0xd0, 0x8a, 0x56, 0x40, 0x17, 0x25, 0x2d, 0x07, 0x2e,
	0x55, 0xb6, 0xf1, 0x76, 0x2d, 0xb5, 0x38, 0xd8, 0xee, 0x56, 0x7d, 0x08, 0x5e, 0x8b, 0xe7, 0x42,
	0xb6, 0x93, 0x95, 0x80, 0x0a, 0x38, 0x70, 0x89, 0x66, 0xbe, 0xfc, 0x66, 0x3c, 0xce, 0x37, 0x81,
	0x5e, 0x2e, 0xb8, 0xe2, 0x03, 0xa9, 0xb8, 0x38, 0xda, 0x67, 0x64, 0x14, 0xc4, 0x0d, 0x8f, 0x76,
	0x6c, 0x2d, 0x78, 0x24, 0xc5, 0x5d, 0xc4, 0x32, 0x9a, 0xca, 0xfe, 0xd3, 0x0d, 0xe7, 0x9b, 0x2d,
	0x1d, 0x18, 0xe2, 0x7a, 0x7f, 0x33, 0x50, 0x6c, 0x47, 0xa5, 0x4a, 0x77, 0xb9, 0x2d, 0x0a, 0x42,
	0xb8, 0x18, 0xd3, 0x2d, 0x55, 0x34, 0x51, 0x5c, 0x30, 0x2a, 0x63, 0xfa, 0x75, 0x4f, 0xa5, 0xc2,
	0x2e, 0x54, 0x59, 0x26, 0x89, 0xe3, 0x57, 0x43, 0x2f, 0xd6, 0x61, 0xd0, 0x83, 0x87, 0xbf, 0x90,
	0x32, 0xe7, 0x5f, 0x24, 0x0d, 0x5e, 0x41, 0x53, 0x5a, 0x09, 0x2f, 0xef, 0x43, 0x53, 0xd9, 0x1e,
	0x3e, 0x8a, 0x7e, 0x1f, 0x2a, 0xd2, 0x0d, 0x8e, 0x71, 0x49, 0x06, 0xdf, 0x5d, 0xa8, 0x1b, 0x09,
	0xcf, 0xc1, 0x9d, 0x8d, 0x89, 0xe3, 0x3b, 0xa1, 0x17, 0xbb, 0xb3, 0x31, 0x5e, 0x40, 0x5d, 0x31,
	0xb5, 0xa5, 0xc4, 0x35, 0x92, 0x4d, 0xf0, 0x05, 0xb4, 0x72, 0x26, 0xb8, 0x60, 0xea, 0x48, 0xaa,
	0xbe, 0x13, 0x9e, 0x0f, 0x9f, 0x9c, 0x3a, 0xe5, 0x63, 0xc1, 0xc4, 0xf7, 0xb4, 0xbe, 0xd4, 0xe1,
	0x96, 0x93, 0x9a, 0xe9, 0xa6, 0x43, 0x44, 0xa8, 0x1d, 0x6e, 0x53, 0x45, 0xea, 0x46, 0x32, 0xb1,
	0xd6, 0x36, 0x3c, 0xdd, 0x92, 0x86, 0xd5, 0x74, 0xac, 0x27, 0xe1, 0x22, 0xa3, 0x82, 0x34, 0x7d,
	0x27, 0xac, 0xc5, 0x36, 0xc1, 0x1e, 0x34, 0xf5, 0x59, 0x2b, 0x96, 0x91, 0x96, 0x81, 0x1b, 0x3a,
	0x9d, 0x65, 0xf8, 0x12, 0x60, 0x2d, 0x68, 0xaa, 0x68, 0xb6, 0x4a, 0x15, 0xf1, 0x7c, 0x27, 0x6c,
	0x0f, 0xfb, 0x91, 0xf5, 0x22, 0x2a, 0xbd, 0x88, 0x16, 0xa5, 0x17, 0xb1, 0x57, 0xd0, 0x23, 0xa5,
	0x4b, 0xf7, 0x79, 0x56, 0x96, 0xc2, 0xdf, 0x4b, 0x0b, 0x7a, 0xa4, 0x9e, 0xbf, 0x85, 0x56, 0x79,
	0x69, 0xf4, 0xa0, 0xbe, 0x9c, 0x27, 0x93, 0x45, 0xb7, 0x82, 0x5d, 0x38, 0x9b, 0xcf, 0xde, 0x4c,
	0x56, 0x8b, 0xab, 0xd5, 0x74, 0xf4, 0x69, 0xd2, 0x75, 0xf0, 0x01, 0xb4, 0x93, 0xe9, 0xd5, 0xf2,
	0xfd, 0xd8, 0x0a, 0x2e, 0x76, 0xc0, 0xfb, 0xb0, 0x4c, 0x16, 0x36, 0xad, 0x0e, 0xbf, 0xb9, 0x70,
	0x66, 0x1c, 0x49, 0xa8, 0xb8, 0x63, 0x6b, 0x8a, 0x53, 0x80, 0x39, 0x3d, 0x14, 0xc6, 0xe3, 0xe3,
	0x53, 0x9f, 0xbb, 0x70, 0xb3, 0xff, 0xa7, 0x97, 0x41, 0x05, 0xdf, 0x41, 0x67, 0x69, 0x06, 0xfe,
	0x1f, 0xcd, 0x6e, 0xa0, 0xf3, 0xd3, 0x4a, 0x62, 0x78, 0x8a, 0x3f, 0xb5, 0xdf, 0xfd, 0x67, 0xff,
	0x40, 0x16, 0xfb, 0x5d, 0x79, 0xdd, 0xfc, 0x5c, 0x37, 0x3f, 0xda, 0x75, 0xc3, 0x18, 0x70, 0xf9,
	0x23, 0x00, 0x00, 0xff, 0xff, 0x06, 0x1c, 0xd4, 0x58, 0x84, 0x03, 0x00, 0x00,
}
