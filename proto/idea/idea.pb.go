// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/idea/idea.proto

/*
Package idea is a generated protocol buffer package.

It is generated from these files:
	proto/idea/idea.proto

It has these top-level messages:
	ShowIdeasRequest
	ShowIdeasResponse
	CreateIdeasRequest
	CreateIdeasResponse
	UpdateIdeasRequest
	UpdateIdeasResponse
	HideIdeasResponse
	DeleteIdeasRequest
	DeleteIdeasResponse
	Ideas
	HideIdeasRequest
	Idea
*/
package idea

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import go_micro_srv_ideas "github.com/jianhan/ms-sui-ideas/proto/rating"
import go_micro_srv_ideas1 "github.com/jianhan/ms-sui-ideas/proto/story"
import go_micro_srv_ideas2 "github.com/jianhan/ms-sui-ideas/proto/occupation"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AppType defines application type
type AppType int32

const (
	AppType_WEB     AppType = 0
	AppType_MOBILE  AppType = 1
	AppType_DESKTOP AppType = 2
	AppType_OTHER   AppType = 3
)

var AppType_name = map[int32]string{
	0: "WEB",
	1: "MOBILE",
	2: "DESKTOP",
	3: "OTHER",
}
var AppType_value = map[string]int32{
	"WEB":     0,
	"MOBILE":  1,
	"DESKTOP": 2,
	"OTHER":   3,
}

func (x AppType) String() string {
	return proto.EnumName(AppType_name, int32(x))
}
func (AppType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ShowIdeasRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *ShowIdeasRequest) Reset()                    { *m = ShowIdeasRequest{} }
func (m *ShowIdeasRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowIdeasRequest) ProtoMessage()               {}
func (*ShowIdeasRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ShowIdeasRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ShowIdeasResponse struct {
	Matched  int64   `protobuf:"varint,1,opt,name=matched" json:"matched,omitempty"`
	Modified int64   `protobuf:"varint,2,opt,name=modified" json:"modified,omitempty"`
	Ideas    []*Idea `protobuf:"bytes,3,rep,name=ideas" json:"ideas,omitempty"`
}

func (m *ShowIdeasResponse) Reset()                    { *m = ShowIdeasResponse{} }
func (m *ShowIdeasResponse) String() string            { return proto.CompactTextString(m) }
func (*ShowIdeasResponse) ProtoMessage()               {}
func (*ShowIdeasResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ShowIdeasResponse) GetMatched() int64 {
	if m != nil {
		return m.Matched
	}
	return 0
}

func (m *ShowIdeasResponse) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

func (m *ShowIdeasResponse) GetIdeas() []*Idea {
	if m != nil {
		return m.Ideas
	}
	return nil
}

type CreateIdeasRequest struct {
	Ideas []*Idea `protobuf:"bytes,1,rep,name=ideas" json:"ideas,omitempty"`
}

func (m *CreateIdeasRequest) Reset()                    { *m = CreateIdeasRequest{} }
func (m *CreateIdeasRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateIdeasRequest) ProtoMessage()               {}
func (*CreateIdeasRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateIdeasRequest) GetIdeas() []*Idea {
	if m != nil {
		return m.Ideas
	}
	return nil
}

type CreateIdeasResponse struct {
	Matched  int64   `protobuf:"varint,1,opt,name=matched" json:"matched,omitempty"`
	Modified int64   `protobuf:"varint,2,opt,name=modified" json:"modified,omitempty"`
	Ideas    []*Idea `protobuf:"bytes,3,rep,name=ideas" json:"ideas,omitempty"`
}

func (m *CreateIdeasResponse) Reset()                    { *m = CreateIdeasResponse{} }
func (m *CreateIdeasResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateIdeasResponse) ProtoMessage()               {}
func (*CreateIdeasResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateIdeasResponse) GetMatched() int64 {
	if m != nil {
		return m.Matched
	}
	return 0
}

func (m *CreateIdeasResponse) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

func (m *CreateIdeasResponse) GetIdeas() []*Idea {
	if m != nil {
		return m.Ideas
	}
	return nil
}

type UpdateIdeasRequest struct {
	Ideas []*Idea `protobuf:"bytes,1,rep,name=ideas" json:"ideas,omitempty"`
}

func (m *UpdateIdeasRequest) Reset()                    { *m = UpdateIdeasRequest{} }
func (m *UpdateIdeasRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateIdeasRequest) ProtoMessage()               {}
func (*UpdateIdeasRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateIdeasRequest) GetIdeas() []*Idea {
	if m != nil {
		return m.Ideas
	}
	return nil
}

type UpdateIdeasResponse struct {
	Matched  int64   `protobuf:"varint,1,opt,name=matched" json:"matched,omitempty"`
	Modified int64   `protobuf:"varint,2,opt,name=modified" json:"modified,omitempty"`
	Ideas    []*Idea `protobuf:"bytes,3,rep,name=ideas" json:"ideas,omitempty"`
}

func (m *UpdateIdeasResponse) Reset()                    { *m = UpdateIdeasResponse{} }
func (m *UpdateIdeasResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateIdeasResponse) ProtoMessage()               {}
func (*UpdateIdeasResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateIdeasResponse) GetMatched() int64 {
	if m != nil {
		return m.Matched
	}
	return 0
}

func (m *UpdateIdeasResponse) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

func (m *UpdateIdeasResponse) GetIdeas() []*Idea {
	if m != nil {
		return m.Ideas
	}
	return nil
}

type HideIdeasResponse struct {
	Matched  int64   `protobuf:"varint,1,opt,name=matched" json:"matched,omitempty"`
	Modified int64   `protobuf:"varint,2,opt,name=modified" json:"modified,omitempty"`
	Ideas    []*Idea `protobuf:"bytes,3,rep,name=ideas" json:"ideas,omitempty"`
}

func (m *HideIdeasResponse) Reset()                    { *m = HideIdeasResponse{} }
func (m *HideIdeasResponse) String() string            { return proto.CompactTextString(m) }
func (*HideIdeasResponse) ProtoMessage()               {}
func (*HideIdeasResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *HideIdeasResponse) GetMatched() int64 {
	if m != nil {
		return m.Matched
	}
	return 0
}

func (m *HideIdeasResponse) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

func (m *HideIdeasResponse) GetIdeas() []*Idea {
	if m != nil {
		return m.Ideas
	}
	return nil
}

type DeleteIdeasRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *DeleteIdeasRequest) Reset()                    { *m = DeleteIdeasRequest{} }
func (m *DeleteIdeasRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteIdeasRequest) ProtoMessage()               {}
func (*DeleteIdeasRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteIdeasRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type DeleteIdeasResponse struct {
	Matched  int64 `protobuf:"varint,1,opt,name=matched" json:"matched,omitempty"`
	Modified int64 `protobuf:"varint,2,opt,name=modified" json:"modified,omitempty"`
}

func (m *DeleteIdeasResponse) Reset()                    { *m = DeleteIdeasResponse{} }
func (m *DeleteIdeasResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteIdeasResponse) ProtoMessage()               {}
func (*DeleteIdeasResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeleteIdeasResponse) GetMatched() int64 {
	if m != nil {
		return m.Matched
	}
	return 0
}

func (m *DeleteIdeasResponse) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

type Ideas struct {
	Ideas []*Idea `protobuf:"bytes,1,rep,name=ideas" json:"ideas,omitempty"`
}

func (m *Ideas) Reset()                    { *m = Ideas{} }
func (m *Ideas) String() string            { return proto.CompactTextString(m) }
func (*Ideas) ProtoMessage()               {}
func (*Ideas) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Ideas) GetIdeas() []*Idea {
	if m != nil {
		return m.Ideas
	}
	return nil
}

type HideIdeasRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *HideIdeasRequest) Reset()                    { *m = HideIdeasRequest{} }
func (m *HideIdeasRequest) String() string            { return proto.CompactTextString(m) }
func (*HideIdeasRequest) ProtoMessage()               {}
func (*HideIdeasRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *HideIdeasRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Idea struct {
	// @inject_tag: bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim" bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"name" valid:"required~Name is required" conform:"trim"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" bson:"name" valid:"required~Name is required" conform:"trim" bson:"name" valid:"required~Name is required" conform:"trim"`
	// @inject_tag: bson:"slug" valid:"required~Slug is required" conform:"trim"
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty" bson:"slug" valid:"required~Slug is required" conform:"trim" bson:"slug" valid:"required~Slug is required" conform:"trim"`
	// @inject_tag: bson:"description" valid:"Description is required" conform:"trim"
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty" bson:"description" valid:"Description is required" conform:"trim" bson:"description" valid:"Description is required" conform:"trim"`
	// @inject_tag: bson:"hidden"
	Hidden bool `protobuf:"varint,5,opt,name=hidden" json:"hidden,omitempty" bson:"hidden" bson:"hidden"`
	// @inject_tag: bson:"app_type" valid:"App type is required"
	Type []AppType `protobuf:"varint,6,rep,packed,name=type,enum=go.micro.srv.ideas.AppType" json:"type,omitempty" bson:"app_type" valid:"App type is required" bson:"app_type" valid:"App type is required"`
	// @inject_tag: bson:"min_age"
	MinAge uint64 `protobuf:"varint,7,opt,name=min_age,json=minAge" json:"min_age,omitempty" bson:"min_age" bson:"min_age"`
	// @inject_tag: bson:"max_age"
	MaxAge uint64 `protobuf:"varint,8,opt,name=max_age,json=maxAge" json:"max_age,omitempty" bson:"max_age" bson:"max_age"`
	// @inject_tag: bson:"rating"
	Rating []*go_micro_srv_ideas.Rating `protobuf:"bytes,9,rep,name=rating" json:"rating,omitempty" bson:"rating" bson:"rating"`
	// @inject_tag: bson:"stories"
	Stories []*go_micro_srv_ideas1.Story `protobuf:"bytes,10,rep,name=stories" json:"stories,omitempty" bson:"stories" bson:"stories"`
	// @inject_tag: bson:"occupations"
	Occupations []*go_micro_srv_ideas2.Occupation `protobuf:"bytes,11,rep,name=occupations" json:"occupations,omitempty" bson:"occupations" bson:"occupations"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at" bson:"updated_at"`
}

func (m *Idea) Reset()                    { *m = Idea{} }
func (m *Idea) String() string            { return proto.CompactTextString(m) }
func (*Idea) ProtoMessage()               {}
func (*Idea) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Idea) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Idea) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Idea) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Idea) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Idea) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *Idea) GetType() []AppType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Idea) GetMinAge() uint64 {
	if m != nil {
		return m.MinAge
	}
	return 0
}

func (m *Idea) GetMaxAge() uint64 {
	if m != nil {
		return m.MaxAge
	}
	return 0
}

func (m *Idea) GetRating() []*go_micro_srv_ideas.Rating {
	if m != nil {
		return m.Rating
	}
	return nil
}

func (m *Idea) GetStories() []*go_micro_srv_ideas1.Story {
	if m != nil {
		return m.Stories
	}
	return nil
}

func (m *Idea) GetOccupations() []*go_micro_srv_ideas2.Occupation {
	if m != nil {
		return m.Occupations
	}
	return nil
}

func (m *Idea) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Idea) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowIdeasRequest)(nil), "go.micro.srv.ideas.ShowIdeasRequest")
	proto.RegisterType((*ShowIdeasResponse)(nil), "go.micro.srv.ideas.ShowIdeasResponse")
	proto.RegisterType((*CreateIdeasRequest)(nil), "go.micro.srv.ideas.CreateIdeasRequest")
	proto.RegisterType((*CreateIdeasResponse)(nil), "go.micro.srv.ideas.CreateIdeasResponse")
	proto.RegisterType((*UpdateIdeasRequest)(nil), "go.micro.srv.ideas.UpdateIdeasRequest")
	proto.RegisterType((*UpdateIdeasResponse)(nil), "go.micro.srv.ideas.UpdateIdeasResponse")
	proto.RegisterType((*HideIdeasResponse)(nil), "go.micro.srv.ideas.HideIdeasResponse")
	proto.RegisterType((*DeleteIdeasRequest)(nil), "go.micro.srv.ideas.DeleteIdeasRequest")
	proto.RegisterType((*DeleteIdeasResponse)(nil), "go.micro.srv.ideas.DeleteIdeasResponse")
	proto.RegisterType((*Ideas)(nil), "go.micro.srv.ideas.ideas")
	proto.RegisterType((*HideIdeasRequest)(nil), "go.micro.srv.ideas.HideIdeasRequest")
	proto.RegisterType((*Idea)(nil), "go.micro.srv.ideas.Idea")
	proto.RegisterEnum("go.micro.srv.ideas.AppType", AppType_name, AppType_value)
}

func init() { proto.RegisterFile("proto/idea/idea.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xef, 0x6f, 0x12, 0x31,
	0x18, 0xe6, 0x38, 0x06, 0xe3, 0x45, 0x17, 0xd6, 0x45, 0xad, 0x98, 0xe8, 0x85, 0xcc, 0x49, 0x4c,
	0x76, 0x24, 0x2c, 0x71, 0xd9, 0x37, 0x99, 0x90, 0x8c, 0x4c, 0x83, 0x29, 0x18, 0x93, 0x7d, 0x99,
	0xb7, 0xbb, 0xee, 0xa8, 0xd9, 0xfd, 0xf0, 0x5a, 0xe6, 0x88, 0x7f, 0xb2, 0xff, 0x82, 0x1f, 0x4c,
	0x7b, 0x07, 0x1e, 0xa3, 0x66, 0x33, 0x9a, 0xf9, 0xe5, 0x68, 0xfb, 0x3e, 0xcf, 0xfb, 0x8b, 0xbe,
	0x4f, 0xe1, 0x41, 0x9c, 0x44, 0x22, 0x6a, 0x33, 0x8f, 0x3a, 0xea, 0x63, 0xab, 0x3d, 0x42, 0x7e,
	0x64, 0x07, 0xcc, 0x4d, 0x22, 0x9b, 0x27, 0x97, 0xb6, 0x34, 0xf0, 0xc6, 0x33, 0x3f, 0x8a, 0xfc,
	0x0b, 0xda, 0x56, 0x88, 0xb3, 0xe9, 0x79, 0x5b, 0xb0, 0x80, 0x72, 0xe1, 0x04, 0x71, 0x4a, 0x6a,
	0x1c, 0xf8, 0x4c, 0x4c, 0xa6, 0x67, 0xb6, 0x1b, 0x05, 0xed, 0xcf, 0xcc, 0x09, 0x27, 0x4e, 0xd8,
	0x0e, 0xf8, 0x2e, 0x9f, 0xb2, 0x5d, 0xe5, 0x21, 0x65, 0xb6, 0x13, 0x47, 0xb0, 0xd0, 0xcf, 0x7e,
	0x32, 0xea, 0xfe, 0xed, 0xa8, 0x5c, 0x44, 0xc9, 0x2c, 0xfd, 0x66, 0xc4, 0xee, 0xed, 0x88, 0x91,
	0xeb, 0x4e, 0x63, 0x47, 0xb0, 0x28, 0xcc, 0x2d, 0x53, 0x17, 0xcd, 0x6d, 0xa8, 0x8f, 0x26, 0xd1,
	0xd7, 0x81, 0x84, 0x13, 0xfa, 0x65, 0x4a, 0xb9, 0x40, 0x75, 0x30, 0x99, 0xc7, 0xb1, 0x61, 0x99,
	0xad, 0x2a, 0x91, 0xcb, 0xe6, 0x0c, 0x36, 0x73, 0x28, 0x1e, 0x47, 0x21, 0xa7, 0x08, 0x43, 0x25,
	0x70, 0x84, 0x3b, 0xa1, 0x1e, 0x36, 0x2c, 0xa3, 0x65, 0x92, 0xf9, 0x16, 0x35, 0x60, 0x3d, 0x88,
	0x3c, 0x76, 0xce, 0xa8, 0x87, 0x8b, 0xca, 0xb4, 0xd8, 0x23, 0x1b, 0xd6, 0x54, 0x6e, 0xd8, 0xb4,
	0xcc, 0x56, 0xad, 0x83, 0xed, 0xd5, 0x66, 0xdb, 0x32, 0x0e, 0x49, 0x61, 0xcd, 0x1e, 0xa0, 0x37,
	0x09, 0x75, 0x04, 0x5d, 0x4a, 0x71, 0xe1, 0xc5, 0xb8, 0x9d, 0x97, 0x6f, 0xb0, 0xb5, 0xe4, 0xe5,
	0xae, 0x4b, 0xf8, 0x10, 0x7b, 0xff, 0xa0, 0x84, 0x25, 0x2f, 0x77, 0x5a, 0xc2, 0x0c, 0x36, 0x8f,
	0x98, 0xf7, 0x5f, 0x42, 0xef, 0x00, 0xea, 0xd1, 0x0b, 0x7a, 0xad, 0x7b, 0xab, 0x77, 0xf4, 0x18,
	0xb6, 0x96, 0x70, 0x7f, 0x93, 0x64, 0x73, 0x3f, 0x4b, 0xf2, 0x8f, 0xff, 0xa5, 0x6d, 0xa8, 0xe7,
	0x1a, 0xf5, 0xbb, 0x5c, 0x7f, 0x98, 0x50, 0x92, 0x10, 0xb4, 0x01, 0xc5, 0x41, 0x4f, 0x25, 0x56,
	0x25, 0xc5, 0x41, 0x0f, 0x21, 0x28, 0x85, 0x4e, 0x40, 0x55, 0x3e, 0x55, 0xa2, 0xd6, 0xf2, 0x8c,
	0x5f, 0x4c, 0x7d, 0x6c, 0xa6, 0x67, 0x72, 0x8d, 0x2c, 0xa8, 0x79, 0x94, 0xbb, 0x09, 0x8b, 0xe5,
	0x2c, 0xe3, 0x92, 0x32, 0xe5, 0x8f, 0xd0, 0x43, 0x28, 0x4f, 0x98, 0xe7, 0xd1, 0x10, 0xaf, 0x59,
	0x46, 0x6b, 0x9d, 0x64, 0x3b, 0xd4, 0x86, 0x92, 0x98, 0xc5, 0x14, 0x97, 0x2d, 0xb3, 0xb5, 0xd1,
	0x79, 0xa2, 0xab, 0xa7, 0x1b, 0xc7, 0xe3, 0x59, 0x4c, 0x89, 0x02, 0xa2, 0x47, 0x50, 0x09, 0x58,
	0x78, 0xea, 0xf8, 0x14, 0x57, 0x2c, 0xa3, 0x55, 0x22, 0xe5, 0x80, 0x85, 0x5d, 0x3f, 0x35, 0x38,
	0x57, 0xca, 0xb0, 0x9e, 0x19, 0x9c, 0x2b, 0x69, 0xe8, 0x40, 0x39, 0xd5, 0x37, 0x5c, 0x55, 0x4d,
	0x6b, 0xe8, 0x82, 0x10, 0x85, 0x20, 0x19, 0x12, 0xed, 0x41, 0x45, 0x2a, 0x1b, 0xa3, 0x1c, 0x83,
	0x22, 0x3d, 0xd6, 0x91, 0x46, 0x52, 0xfc, 0xc8, 0x1c, 0x89, 0x5e, 0x43, 0xed, 0x97, 0xa0, 0x71,
	0x5c, 0x53, 0xc4, 0xa7, 0x3a, 0xe2, 0x70, 0x01, 0x23, 0x79, 0x0a, 0x3a, 0x00, 0x70, 0x95, 0x2e,
	0x78, 0xa7, 0x8e, 0xc0, 0xf7, 0x2c, 0x23, 0x4b, 0x57, 0x6a, 0xbd, 0x3d, 0xd7, 0x7a, 0x7b, 0x3c,
	0xd7, 0x7a, 0x52, 0xcd, 0xd0, 0x5d, 0x21, 0xa9, 0x53, 0x35, 0x8f, 0x8a, 0x7a, 0xff, 0x66, 0x6a,
	0x86, 0xee, 0x8a, 0x97, 0xaf, 0xa0, 0x92, 0xf5, 0x18, 0x55, 0xc0, 0xfc, 0xd8, 0x3f, 0xac, 0x17,
	0x10, 0x40, 0xf9, 0xdd, 0xf0, 0x70, 0xf0, 0xb6, 0x5f, 0x37, 0x50, 0x0d, 0x2a, 0xbd, 0xfe, 0xe8,
	0x78, 0x3c, 0x7c, 0x5f, 0x2f, 0xa2, 0x2a, 0xac, 0x0d, 0xc7, 0x47, 0x7d, 0x52, 0x37, 0x3b, 0xdf,
	0x4d, 0xa8, 0xc9, 0x6b, 0x33, 0xa2, 0xc9, 0x25, 0x73, 0x29, 0xfa, 0x04, 0xb5, 0x9c, 0xaa, 0xa1,
	0x1d, 0x5d, 0xe5, 0xab, 0xe2, 0xd9, 0x78, 0x71, 0x23, 0x2e, 0x9d, 0x9d, 0x66, 0x41, 0x46, 0xc8,
	0x89, 0x8e, 0x3e, 0xc2, 0xaa, 0xb6, 0xe9, 0x23, 0x68, 0xd4, 0x2b, 0x8d, 0x90, 0x1b, 0x5b, 0x7d,
	0x84, 0xd5, 0xf9, 0xd7, 0x47, 0xd0, 0xcc, 0x7f, 0xb3, 0x80, 0x4e, 0xa0, 0xba, 0x18, 0x49, 0xb4,
	0xad, 0xe3, 0x5d, 0x9f, 0xd8, 0xc6, 0xf3, 0x1b, 0x50, 0x79, 0xdf, 0x8b, 0x87, 0x51, 0xef, 0xfb,
	0xfa, 0xeb, 0xaa, 0xf7, 0xbd, 0xf2, 0xba, 0x36, 0x0b, 0x87, 0xe5, 0x93, 0x92, 0x34, 0x9e, 0x95,
	0xd5, 0x65, 0xda, 0xfb, 0x19, 0x00, 0x00, 0xff, 0xff, 0x37, 0xdd, 0x12, 0xe3, 0xae, 0x08, 0x00,
	0x00,
}
