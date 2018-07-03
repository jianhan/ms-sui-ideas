// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/idea/idea.proto

/*
Package idea is a generated protocol buffer package.

It is generated from these files:
	proto/idea/idea.proto

It has these top-level messages:
	Idea
	Story
	Rating
	IdeaFilter
	ListIdeasResponse
	ShowIdeasRequest
	UpsertIdeasRequest
	CreateIdeasResponse
	UpdateIdeasResponse
	DeleteIdeasResponse
*/
package idea

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import go_micro_srv_ideas "github.com/jianhan/ms-sui-ideas/proto/occupation"
import _ "github.com/golang/protobuf/ptypes/empty"

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
func (Pirority) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Idea struct {
	// @inject_tag: bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"name" valid:"required~Name is required" conform:"trim"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" bson:"name" valid:"required~Name is required" conform:"trim"`
	// @inject_tag: bson:"slug" valid:"required~Slug is required" conform:"trim"
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty" bson:"slug" valid:"required~Slug is required" conform:"trim"`
	// @inject_tag: bson:"description" valid:"Description is required" conform:"trim"
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty" bson:"description" valid:"Description is required" conform:"trim"`
	// @inject_tag: bson:"hidden"
	Hidden bool `protobuf:"varint,5,opt,name=hidden" json:"hidden,omitempty" bson:"hidden"`
	// @inject_tag: bson:"app_type" valid:"App type is required"
	Type []AppType `protobuf:"varint,6,rep,packed,name=type,enum=go.micro.srv.ideas.AppType" json:"type,omitempty" bson:"app_type" valid:"App type is required"`
	// @inject_tag: bson:"min_age" valid:"range(8|99)~Invalid min age range"
	MinAge uint64 `protobuf:"varint,7,opt,name=min_age,json=minAge" json:"min_age,omitempty" bson:"min_age" valid:"range(8|99)~Invalid min age range"`
	// @inject_tag: bson:"max_age" valid:"range(8|99)~Invalid max age range"
	MaxAge uint64 `protobuf:"varint,8,opt,name=max_age,json=maxAge" json:"max_age,omitempty" bson:"max_age" valid:"range(8|99)~Invalid max age range"`
	// @inject_tag: bson:"rating"
	Rating []*Rating `protobuf:"bytes,9,rep,name=rating" json:"rating,omitempty" bson:"rating"`
	// @inject_tag: bson:"stories"
	Stories []*Story `protobuf:"bytes,10,rep,name=stories" json:"stories,omitempty" bson:"stories"`
	// @inject_tag: bson:"occupations"
	Occupations []*go_micro_srv_ideas.Occupation `protobuf:"bytes,11,rep,name=occupations" json:"occupations,omitempty" bson:"occupations"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at"`
}

func (m *Idea) Reset()                    { *m = Idea{} }
func (m *Idea) String() string            { return proto.CompactTextString(m) }
func (*Idea) ProtoMessage()               {}
func (*Idea) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *Idea) GetRating() []*Rating {
	if m != nil {
		return m.Rating
	}
	return nil
}

func (m *Idea) GetStories() []*Story {
	if m != nil {
		return m.Stories
	}
	return nil
}

func (m *Idea) GetOccupations() []*go_micro_srv_ideas.Occupation {
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

type Story struct {
	// @inject_tag: bson:"title" valid:"required~Title is required" conform:"trim"
	Title string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty" bson:"title" valid:"required~Title is required" conform:"trim"`
	// @inject_tag: bson:"pirority" valid:"required~Pirority is required"
	Pirority Pirority `protobuf:"varint,2,opt,name=pirority,enum=go.micro.srv.ideas.Pirority" json:"pirority,omitempty" bson:"pirority" valid:"required~Pirority is required"`
	// @inject_tag: bson:"role" valid:"required~Who is required" conform:"trim"
	Who string `protobuf:"bytes,3,opt,name=who" json:"who,omitempty" bson:"role" valid:"required~Who is required" conform:"trim"`
	// @inject_tag: bson:"action" valid:"required~What is required" conform:"trim"
	What string `protobuf:"bytes,4,opt,name=what" json:"what,omitempty" bson:"action" valid:"required~What is required" conform:"trim"`
	// @inject_tag: bson:"goal" valid:"required~Goal is required" conform:"trim"
	Goal string `protobuf:"bytes,5,opt,name=goal" json:"goal,omitempty" bson:"goal" valid:"required~Goal is required" conform:"trim"`
	// @inject_tag: bson:"order"
	Order uint64 `protobuf:"varint,6,opt,name=order" json:"order,omitempty" bson:"order"`
	// @inject_tag: bson:"hidden"
	Hidden bool `protobuf:"varint,7,opt,name=hidden" json:"hidden,omitempty" bson:"hidden"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at"`
}

func (m *Story) Reset()                    { *m = Story{} }
func (m *Story) String() string            { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()               {}
func (*Story) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

func (m *Story) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
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

type Rating struct {
	// @inject_tag: bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"user_id" valid:"required~User ID is required, uuidv4~User ID must be a valid UUIDv4" conform:"trim"
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty" bson:"user_id" valid:"required~User ID is required, uuidv4~User ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"idea_id" valid:"required~Idea ID is required, uuidv4~Idea id must be a valid UUIDv4" conform:"trim"
	IdeaId string `protobuf:"bytes,3,opt,name=idea_id,json=ideaId" json:"idea_id,omitempty" bson:"idea_id" valid:"required~Idea ID is required, uuidv4~Idea id must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at"`
}

func (m *Rating) Reset()                    { *m = Rating{} }
func (m *Rating) String() string            { return proto.CompactTextString(m) }
func (*Rating) ProtoMessage()               {}
func (*Rating) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Rating) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Rating) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Rating) GetIdeaId() string {
	if m != nil {
		return m.IdeaId
	}
	return ""
}

func (m *Rating) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Rating) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type IdeaFilter struct {
	Ids         []string                         `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	Names       []string                         `protobuf:"bytes,2,rep,name=names" json:"names,omitempty"`
	Hidden      bool                             `protobuf:"varint,3,opt,name=hidden" json:"hidden,omitempty"`
	AppTypes    []AppType                        `protobuf:"varint,4,rep,packed,name=app_types,json=appTypes,enum=go.micro.srv.ideas.AppType" json:"app_types,omitempty"`
	MinAge      uint64                           `protobuf:"varint,5,opt,name=min_age,json=minAge" json:"min_age,omitempty"`
	MaxAge      uint64                           `protobuf:"varint,6,opt,name=max_age,json=maxAge" json:"max_age,omitempty"`
	MinRating   float32                          `protobuf:"fixed32,7,opt,name=min_rating,json=minRating" json:"min_rating,omitempty"`
	MaxRating   float32                          `protobuf:"fixed32,8,opt,name=max_rating,json=maxRating" json:"max_rating,omitempty"`
	Occupations []*go_micro_srv_ideas.Occupation `protobuf:"bytes,9,rep,name=occupations" json:"occupations,omitempty"`
	CurrentPage uint64                           `protobuf:"varint,10,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	PerPage     uint64                           `protobuf:"varint,11,opt,name=per_page,json=perPage" json:"per_page,omitempty"`
	Query       string                           `protobuf:"bytes,12,opt,name=query" json:"query,omitempty"`
	Sorts       []string                         `protobuf:"bytes,13,rep,name=sorts" json:"sorts,omitempty"`
}

func (m *IdeaFilter) Reset()                    { *m = IdeaFilter{} }
func (m *IdeaFilter) String() string            { return proto.CompactTextString(m) }
func (*IdeaFilter) ProtoMessage()               {}
func (*IdeaFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IdeaFilter) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *IdeaFilter) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *IdeaFilter) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *IdeaFilter) GetAppTypes() []AppType {
	if m != nil {
		return m.AppTypes
	}
	return nil
}

func (m *IdeaFilter) GetMinAge() uint64 {
	if m != nil {
		return m.MinAge
	}
	return 0
}

func (m *IdeaFilter) GetMaxAge() uint64 {
	if m != nil {
		return m.MaxAge
	}
	return 0
}

func (m *IdeaFilter) GetMinRating() float32 {
	if m != nil {
		return m.MinRating
	}
	return 0
}

func (m *IdeaFilter) GetMaxRating() float32 {
	if m != nil {
		return m.MaxRating
	}
	return 0
}

func (m *IdeaFilter) GetOccupations() []*go_micro_srv_ideas.Occupation {
	if m != nil {
		return m.Occupations
	}
	return nil
}

func (m *IdeaFilter) GetCurrentPage() uint64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *IdeaFilter) GetPerPage() uint64 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *IdeaFilter) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *IdeaFilter) GetSorts() []string {
	if m != nil {
		return m.Sorts
	}
	return nil
}

type ListIdeasResponse struct {
	Ideas []*Idea `protobuf:"bytes,1,rep,name=ideas" json:"ideas,omitempty"`
}

func (m *ListIdeasResponse) Reset()                    { *m = ListIdeasResponse{} }
func (m *ListIdeasResponse) String() string            { return proto.CompactTextString(m) }
func (*ListIdeasResponse) ProtoMessage()               {}
func (*ListIdeasResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListIdeasResponse) GetIdeas() []*Idea {
	if m != nil {
		return m.Ideas
	}
	return nil
}

type ShowIdeasRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *ShowIdeasRequest) Reset()                    { *m = ShowIdeasRequest{} }
func (m *ShowIdeasRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowIdeasRequest) ProtoMessage()               {}
func (*ShowIdeasRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ShowIdeasRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type UpsertIdeasRequest struct {
	Ideas []*Idea `protobuf:"bytes,1,rep,name=ideas" json:"ideas,omitempty"`
}

func (m *UpsertIdeasRequest) Reset()                    { *m = UpsertIdeasRequest{} }
func (m *UpsertIdeasRequest) String() string            { return proto.CompactTextString(m) }
func (*UpsertIdeasRequest) ProtoMessage()               {}
func (*UpsertIdeasRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpsertIdeasRequest) GetIdeas() []*Idea {
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
func (*CreateIdeasResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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

type UpdateIdeasResponse struct {
	Matched  int64   `protobuf:"varint,1,opt,name=matched" json:"matched,omitempty"`
	Modified int64   `protobuf:"varint,2,opt,name=modified" json:"modified,omitempty"`
	Ideas    []*Idea `protobuf:"bytes,3,rep,name=ideas" json:"ideas,omitempty"`
}

func (m *UpdateIdeasResponse) Reset()                    { *m = UpdateIdeasResponse{} }
func (m *UpdateIdeasResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateIdeasResponse) ProtoMessage()               {}
func (*UpdateIdeasResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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

type DeleteIdeasResponse struct {
	Deleted int64 `protobuf:"varint,1,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *DeleteIdeasResponse) Reset()                    { *m = DeleteIdeasResponse{} }
func (m *DeleteIdeasResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteIdeasResponse) ProtoMessage()               {}
func (*DeleteIdeasResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteIdeasResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func init() {
	proto.RegisterType((*Idea)(nil), "go.micro.srv.ideas.Idea")
	proto.RegisterType((*Story)(nil), "go.micro.srv.ideas.Story")
	proto.RegisterType((*Rating)(nil), "go.micro.srv.ideas.Rating")
	proto.RegisterType((*IdeaFilter)(nil), "go.micro.srv.ideas.IdeaFilter")
	proto.RegisterType((*ListIdeasResponse)(nil), "go.micro.srv.ideas.ListIdeasResponse")
	proto.RegisterType((*ShowIdeasRequest)(nil), "go.micro.srv.ideas.ShowIdeasRequest")
	proto.RegisterType((*UpsertIdeasRequest)(nil), "go.micro.srv.ideas.UpsertIdeasRequest")
	proto.RegisterType((*CreateIdeasResponse)(nil), "go.micro.srv.ideas.CreateIdeasResponse")
	proto.RegisterType((*UpdateIdeasResponse)(nil), "go.micro.srv.ideas.UpdateIdeasResponse")
	proto.RegisterType((*DeleteIdeasResponse)(nil), "go.micro.srv.ideas.DeleteIdeasResponse")
	proto.RegisterEnum("go.micro.srv.ideas.AppType", AppType_name, AppType_value)
	proto.RegisterEnum("go.micro.srv.ideas.Pirority", Pirority_name, Pirority_value)
}

func init() { proto.RegisterFile("proto/idea/idea.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1007 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdd, 0x6e, 0x1a, 0x47,
	0x14, 0x36, 0x2c, 0x2c, 0xec, 0xd9, 0xd8, 0xa5, 0x93, 0x34, 0xd9, 0x90, 0x36, 0xa5, 0xa8, 0x3f,
	0x56, 0xa4, 0x80, 0x44, 0xa4, 0x2a, 0xbd, 0x2b, 0x36, 0xb4, 0x46, 0x75, 0x82, 0xb5, 0x40, 0x5b,
	0xf5, 0x86, 0x8e, 0xd9, 0xc9, 0x32, 0x15, 0xfb, 0x93, 0x99, 0xd9, 0xd8, 0xa8, 0x0f, 0xd1, 0x77,
	0xe8, 0x3b, 0xf4, 0xbe, 0x0f, 0xd6, 0x8b, 0x6a, 0x7e, 0x20, 0x60, 0x36, 0x72, 0xec, 0x5e, 0xf4,
	0x06, 0xcd, 0x39, 0xdf, 0x39, 0x73, 0xce, 0x9c, 0xf3, 0x7d, 0x2b, 0xe0, 0xa3, 0x94, 0x25, 0x22,
	0x69, 0xd3, 0x80, 0x60, 0xf5, 0xd3, 0x52, 0x36, 0x42, 0x61, 0xd2, 0x8a, 0xe8, 0x8c, 0x25, 0x2d,
	0xce, 0xde, 0xb4, 0x24, 0xc0, 0xeb, 0x9f, 0x86, 0x49, 0x12, 0x2e, 0x48, 0x5b, 0x45, 0x9c, 0x67,
	0xaf, 0xda, 0x82, 0x46, 0x84, 0x0b, 0x1c, 0xa5, 0x3a, 0xa9, 0xde, 0x0d, 0xa9, 0x98, 0x67, 0xe7,
	0xad, 0x59, 0x12, 0xb5, 0x7f, 0xa3, 0x38, 0x9e, 0xe3, 0xb8, 0x1d, 0xf1, 0xa7, 0x3c, 0xa3, 0x4f,
	0xd5, 0x0d, 0x3a, 0xb3, 0x9d, 0xcc, 0x66, 0x59, 0x8a, 0x05, 0x4d, 0xe2, 0x8d, 0xa3, 0xb9, 0xe2,
	0xd1, 0xd5, 0x1a, 0x24, 0x4a, 0xc5, 0x52, 0x83, 0xcd, 0x7f, 0x2c, 0x28, 0x0d, 0x02, 0x82, 0xd1,
	0x01, 0x14, 0x07, 0x3d, 0xaf, 0xd0, 0x28, 0x1c, 0x3a, 0x7e, 0x71, 0xd0, 0x43, 0x08, 0x4a, 0x31,
	0x8e, 0x88, 0x57, 0x54, 0x1e, 0x75, 0x96, 0x3e, 0xbe, 0xc8, 0x42, 0xcf, 0xd2, 0x3e, 0x79, 0x46,
	0x0d, 0x70, 0x03, 0xc2, 0x67, 0x8c, 0xa6, 0xb2, 0xa4, 0x57, 0x52, 0xd0, 0xa6, 0x0b, 0xdd, 0x07,
	0x7b, 0x4e, 0x83, 0x80, 0xc4, 0x5e, 0xb9, 0x51, 0x38, 0xac, 0xfa, 0xc6, 0x42, 0x6d, 0x28, 0x89,
	0x65, 0x4a, 0x3c, 0xbb, 0x61, 0x1d, 0x1e, 0x74, 0x1e, 0xb5, 0x76, 0xc7, 0xd3, 0xea, 0xa6, 0xe9,
	0x78, 0x99, 0x12, 0x5f, 0x05, 0xa2, 0x07, 0x50, 0x89, 0x68, 0x3c, 0xc5, 0x21, 0xf1, 0x2a, 0x8d,
	0xc2, 0x61, 0xc9, 0xb7, 0x23, 0x1a, 0x77, 0x43, 0x0d, 0xe0, 0x4b, 0x05, 0x54, 0x0d, 0x80, 0x2f,
	0x25, 0xd0, 0x01, 0x9b, 0x61, 0x41, 0xe3, 0xd0, 0x73, 0x1a, 0xd6, 0xa1, 0xdb, 0xa9, 0xe7, 0x15,
	0xf1, 0x55, 0x84, 0x6f, 0x22, 0xd1, 0x33, 0xa8, 0x70, 0x91, 0x30, 0x4a, 0xb8, 0x07, 0x2a, 0xe9,
	0x61, 0x5e, 0xd2, 0x48, 0x24, 0x6c, 0xe9, 0xaf, 0x22, 0xd1, 0xb7, 0xe0, 0xbe, 0x9d, 0x3b, 0xf7,
	0x5c, 0x95, 0xf8, 0x38, 0x2f, 0x71, 0xb8, 0x0e, 0xf3, 0x37, 0x53, 0xd0, 0x37, 0x00, 0x33, 0x46,
	0xb0, 0x20, 0xc1, 0x14, 0x0b, 0xef, 0x4e, 0xa3, 0x60, 0xda, 0x95, 0xab, 0x6b, 0xad, 0x56, 0xd7,
	0x1a, 0xaf, 0xe8, 0xe1, 0x3b, 0x26, 0xba, 0x2b, 0x64, 0x6a, 0x96, 0x06, 0xab, 0xd4, 0xfd, 0xeb,
	0x53, 0x4d, 0x74, 0x57, 0x34, 0xff, 0x2a, 0x42, 0x59, 0x3d, 0x05, 0xdd, 0x83, 0xb2, 0xa0, 0x62,
	0x41, 0x0c, 0x05, 0xb4, 0x81, 0x9e, 0x43, 0x35, 0xa5, 0x2c, 0x61, 0x54, 0x2c, 0x15, 0x13, 0x0e,
	0x3a, 0x1f, 0xe7, 0x3d, 0xea, 0xcc, 0xc4, 0xf8, 0xeb, 0x68, 0x54, 0x03, 0xeb, 0x62, 0x9e, 0x18,
	0xaa, 0xc8, 0xa3, 0x64, 0xcf, 0xc5, 0x1c, 0x0b, 0x43, 0x11, 0x75, 0x96, 0xbe, 0x30, 0xc1, 0x0b,
	0xc5, 0x0c, 0xc7, 0x57, 0x67, 0xd9, 0x49, 0xc2, 0x02, 0xc2, 0x3c, 0x5b, 0xed, 0x52, 0x1b, 0x1b,
	0x2c, 0xaa, 0x6c, 0xb1, 0x68, 0x7b, 0x6e, 0xd5, 0xdb, 0xcf, 0xcd, 0xb9, 0xc9, 0xdc, 0xfe, 0x2e,
	0x80, 0xad, 0x79, 0xb3, 0x23, 0x9c, 0x07, 0x50, 0xc9, 0x38, 0x61, 0x53, 0x1a, 0x18, 0xed, 0xd8,
	0xd2, 0x1c, 0x04, 0x12, 0x90, 0xd3, 0x92, 0x80, 0x9e, 0x8a, 0x2d, 0xcd, 0x41, 0x70, 0xe5, 0x09,
	0xa5, 0xdb, 0x3f, 0xa1, 0x7c, 0x93, 0x27, 0xfc, 0x69, 0x01, 0x48, 0xe5, 0x7f, 0x47, 0x17, 0x82,
	0x30, 0xb9, 0x2f, 0x1a, 0x70, 0xaf, 0xd0, 0xb0, 0xe4, 0xbe, 0x68, 0xc0, 0xe5, 0x1e, 0xa4, 0xea,
	0xb9, 0x57, 0x54, 0x3e, 0x6d, 0x6c, 0xec, 0xc1, 0xda, 0xda, 0xc3, 0x73, 0x70, 0x70, 0x9a, 0x4e,
	0xa5, 0x50, 0xb9, 0x57, 0xba, 0x5e, 0xd2, 0x55, 0xac, 0x0f, 0x7c, 0x53, 0xd6, 0xe5, 0x77, 0xc9,
	0xda, 0xde, 0x92, 0xf5, 0x27, 0x00, 0x32, 0xc3, 0x48, 0x5b, 0xf2, 0xa1, 0xe8, 0x3b, 0x11, 0x8d,
	0xcd, 0x46, 0x24, 0x8c, 0x2f, 0x57, 0x70, 0xd5, 0xc0, 0xf8, 0xd2, 0xc0, 0x57, 0xb4, 0xea, 0xdc,
	0x5c, 0xab, 0x9f, 0xc1, 0x9d, 0x59, 0xc6, 0x18, 0x89, 0xc5, 0x34, 0x95, 0xdd, 0x81, 0xea, 0xce,
	0x35, 0xbe, 0x33, 0x1c, 0x12, 0xf4, 0x10, 0xaa, 0x29, 0x61, 0x1a, 0x76, 0x15, 0x5c, 0x49, 0x09,
	0x53, 0xd0, 0x3d, 0x28, 0xbf, 0xce, 0x08, 0x5b, 0x2a, 0x91, 0x3b, 0xbe, 0x36, 0xa4, 0x97, 0x27,
	0x4c, 0x70, 0x6f, 0x5f, 0x4f, 0x5b, 0x19, 0xcd, 0x63, 0xf8, 0xf0, 0x94, 0x72, 0x21, 0xf7, 0xc4,
	0x7d, 0xc2, 0xd3, 0x24, 0xe6, 0x04, 0xb5, 0xa0, 0xac, 0xfa, 0x53, 0xcb, 0x72, 0x3b, 0x5e, 0x5e,
	0xeb, 0x32, 0xc3, 0xd7, 0x61, 0xcd, 0xcf, 0xa1, 0x36, 0x9a, 0x27, 0x17, 0xe6, 0x92, 0xd7, 0x19,
	0xe1, 0x62, 0x77, 0xdd, 0xcd, 0x1e, 0xa0, 0x49, 0xca, 0x09, 0x13, 0x5b, 0x71, 0x37, 0xad, 0xf5,
	0x3b, 0xdc, 0x3d, 0x56, 0xec, 0xdc, 0x6e, 0xd9, 0x93, 0xab, 0x14, 0xb3, 0x39, 0x09, 0x94, 0x52,
	0x2c, 0x7f, 0x65, 0xa2, 0x3a, 0x54, 0xa3, 0x24, 0xa0, 0xaf, 0x28, 0xd1, 0x7a, 0xb1, 0xfc, 0xb5,
	0xfd, 0xb6, 0xb8, 0xf5, 0xde, 0xc5, 0x27, 0x8a, 0xdf, 0xff, 0x47, 0xf1, 0x36, 0xdc, 0xed, 0x91,
	0x05, 0xc9, 0x29, 0x1e, 0x28, 0xf7, 0xba, 0xb8, 0x31, 0x9f, 0x7c, 0x0d, 0x15, 0x23, 0x06, 0x54,
	0x01, 0xeb, 0xa7, 0xfe, 0x51, 0x6d, 0x0f, 0x01, 0xd8, 0x2f, 0x86, 0x47, 0x83, 0xd3, 0x7e, 0xad,
	0x80, 0x5c, 0xa8, 0xf4, 0xfa, 0xa3, 0x1f, 0xc6, 0xc3, 0xb3, 0x5a, 0x11, 0x39, 0x50, 0x1e, 0x8e,
	0x4f, 0xfa, 0x7e, 0xcd, 0x7a, 0xf2, 0x3d, 0x54, 0x57, 0xdf, 0x5b, 0xe9, 0x9e, 0xbc, 0x1c, 0xf5,
	0xc7, 0xb5, 0x3d, 0x54, 0x83, 0x3b, 0x2f, 0x07, 0xc7, 0xfd, 0xe9, 0x78, 0x38, 0x3d, 0xe9, 0xfe,
	0x28, 0x2f, 0xf8, 0x00, 0xdc, 0xd1, 0xc9, 0x70, 0x72, 0xda, 0xd3, 0x8e, 0x22, 0xda, 0x07, 0xe7,
	0xc5, 0x64, 0x34, 0xd6, 0xa6, 0xd5, 0xf9, 0xa3, 0x04, 0xae, 0x6c, 0x76, 0x44, 0xd8, 0x1b, 0x3a,
	0x23, 0x68, 0x0c, 0xce, 0x9a, 0x6c, 0xe8, 0xf1, 0xbb, 0xde, 0xab, 0xbf, 0x17, 0xf5, 0x2f, 0xf2,
	0xf0, 0x1d, 0xae, 0x36, 0xf7, 0xd0, 0xaf, 0xe0, 0x6e, 0x30, 0x02, 0x7d, 0x99, 0x97, 0xb7, 0x4b,
	0xbc, 0xfa, 0x57, 0x79, 0x71, 0x39, 0xd4, 0xd2, 0x15, 0x36, 0xd6, 0xfe, 0xdf, 0x2a, 0xe4, 0xf0,
	0xa7, 0xb9, 0x87, 0x7e, 0x06, 0x77, 0x63, 0xb7, 0xd7, 0xce, 0x26, 0xf7, 0xe6, 0x1c, 0x72, 0x34,
	0xf7, 0x50, 0x1f, 0x9c, 0x13, 0x1a, 0xbc, 0xe7, 0xbd, 0xf7, 0x77, 0xbe, 0xec, 0x7d, 0xf9, 0x57,
	0x4e, 0x5f, 0xb3, 0x96, 0xf8, 0xed, 0xaf, 0x39, 0xb2, 0x7f, 0x29, 0xc9, 0xe8, 0x73, 0x5b, 0x21,
	0xcf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x27, 0x4e, 0x8b, 0x43, 0xca, 0x0a, 0x00, 0x00,
}
