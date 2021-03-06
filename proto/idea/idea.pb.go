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
	AddRatingsRequest
	AddRatingsResponse
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
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim" bson:"_id" valid:"required~ID is required, uuidv4~ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"name" valid:"required~Name is required" conform:"trim"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" bson:"name" valid:"required~Name is required" conform:"trim" bson:"name" valid:"required~Name is required" conform:"trim"`
	// @inject_tag: bson:"slug" valid:"required~Slug is required" conform:"trim"
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty" bson:"slug" valid:"required~Slug is required" conform:"trim" bson:"slug" valid:"required~Slug is required" conform:"trim"`
	// @inject_tag: bson:"description" valid:"required~Description is required" conform:"trim"
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty" bson:"description" valid:"required~Description is required" conform:"trim" bson:"description" valid:"required~Description is required" conform:"trim"`
	// @inject_tag: bson:"hidden"
	Hidden bool `protobuf:"varint,5,opt,name=hidden" json:"hidden,omitempty" bson:"hidden" bson:"hidden"`
	// @inject_tag: bson:"app_types" valid:"required~App types is required"
	AppTypes []AppType `protobuf:"varint,6,rep,packed,name=app_types,json=appTypes,enum=go.micro.srv.ideas.AppType" json:"app_types,omitempty" bson:"app_types" valid:"required~App types is required" bson:"app_types" valid:"required~App types is required"`
	// @inject_tag: bson:"min_age" valid:"range(8|99)~Invalid min age range, must between 8 and 99"
	MinAge uint64 `protobuf:"varint,7,opt,name=min_age,json=minAge" json:"min_age,omitempty" bson:"min_age" valid:"range(8|99)~Invalid min age range, must between 8 and 99" bson:"min_age" valid:"range(8|99)~Invalid min age range, must between 8 and 99"`
	// @inject_tag: bson:"max_age" valid:"range(8|99)~Invalid max age range, must between 8 and 99"
	MaxAge uint64 `protobuf:"varint,8,opt,name=max_age,json=maxAge" json:"max_age,omitempty" bson:"max_age" valid:"range(8|99)~Invalid max age range, must between 8 and 99" bson:"max_age" valid:"range(8|99)~Invalid max age range, must between 8 and 99"`
	// @inject_tag: bson:"ratings"
	Ratings []*Rating `protobuf:"bytes,9,rep,name=ratings" json:"ratings,omitempty" bson:"ratings" bson:"ratings"`
	// @inject_tag: bson:"rating"
	Rating float32 `protobuf:"fixed32,10,opt,name=rating" json:"rating,omitempty" bson:"rating" bson:"rating"`
	// @inject_tag: bson:"stories"
	Stories []*Story `protobuf:"bytes,11,rep,name=stories" json:"stories,omitempty" bson:"stories" bson:"stories"`
	// @inject_tag: bson:"occupations"
	Occupations []*go_micro_srv_ideas.Occupation `protobuf:"bytes,12,rep,name=occupations" json:"occupations,omitempty" bson:"occupations" bson:"occupations"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at" bson:"updated_at"`
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

func (m *Idea) GetAppTypes() []AppType {
	if m != nil {
		return m.AppTypes
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

func (m *Idea) GetRatings() []*Rating {
	if m != nil {
		return m.Ratings
	}
	return nil
}

func (m *Idea) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
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
	Title string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty" bson:"title" valid:"required~Title is required" conform:"trim" bson:"title" valid:"required~Title is required" conform:"trim"`
	// @inject_tag: bson:"pirority" valid:"required~Pirority is required"
	Pirority Pirority `protobuf:"varint,2,opt,name=pirority,enum=go.micro.srv.ideas.Pirority" json:"pirority,omitempty" bson:"pirority" valid:"required~Pirority is required" bson:"pirority" valid:"required~Pirority is required"`
	// @inject_tag: bson:"role" valid:"required~Who is required" conform:"trim"
	Who string `protobuf:"bytes,3,opt,name=who" json:"who,omitempty" bson:"role" valid:"required~Who is required" conform:"trim" bson:"role" valid:"required~Who is required" conform:"trim"`
	// @inject_tag: bson:"action" valid:"required~What is required" conform:"trim"
	What string `protobuf:"bytes,4,opt,name=what" json:"what,omitempty" bson:"action" valid:"required~What is required" conform:"trim" bson:"action" valid:"required~What is required" conform:"trim"`
	// @inject_tag: bson:"goal" valid:"required~Goal is required" conform:"trim"
	Goal string `protobuf:"bytes,5,opt,name=goal" json:"goal,omitempty" bson:"goal" valid:"required~Goal is required" conform:"trim" bson:"goal" valid:"required~Goal is required" conform:"trim"`
	// @inject_tag: bson:"order"
	Order uint64 `protobuf:"varint,6,opt,name=order" json:"order,omitempty" bson:"order" bson:"order"`
	// @inject_tag: bson:"hidden"
	Hidden bool `protobuf:"varint,7,opt,name=hidden" json:"hidden,omitempty" bson:"hidden" bson:"hidden"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at" bson:"updated_at"`
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
	// @inject_tag: bson:"user_id" valid:"required~User ID is required, uuidv4~User ID must be a valid UUIDv4" conform:"trim"
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty" bson:"user_id" valid:"required~User ID is required, uuidv4~User ID must be a valid UUIDv4" conform:"trim" bson:"user_id" valid:"required~User ID is required, uuidv4~User ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"value" valid:"required~Rating value is required"
	Value float32 `protobuf:"fixed32,2,opt,name=value" json:"value,omitempty" bson:"value" valid:"required~Rating value is required" bson:"value" valid:"required~Rating value is required"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at" bson:"updated_at"`
}

func (m *Rating) Reset()                    { *m = Rating{} }
func (m *Rating) String() string            { return proto.CompactTextString(m) }
func (*Rating) ProtoMessage()               {}
func (*Rating) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Rating) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Rating) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
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

type AddRatingsRequest struct {
	IdeaId  string    `protobuf:"bytes,1,opt,name=idea_id,json=ideaId" json:"idea_id,omitempty"`
	Ratings []*Rating `protobuf:"bytes,2,rep,name=ratings" json:"ratings,omitempty"`
}

func (m *AddRatingsRequest) Reset()                    { *m = AddRatingsRequest{} }
func (m *AddRatingsRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRatingsRequest) ProtoMessage()               {}
func (*AddRatingsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AddRatingsRequest) GetIdeaId() string {
	if m != nil {
		return m.IdeaId
	}
	return ""
}

func (m *AddRatingsRequest) GetRatings() []*Rating {
	if m != nil {
		return m.Ratings
	}
	return nil
}

type AddRatingsResponse struct {
	Idea *Idea `protobuf:"bytes,1,opt,name=idea" json:"idea,omitempty"`
}

func (m *AddRatingsResponse) Reset()                    { *m = AddRatingsResponse{} }
func (m *AddRatingsResponse) String() string            { return proto.CompactTextString(m) }
func (*AddRatingsResponse) ProtoMessage()               {}
func (*AddRatingsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AddRatingsResponse) GetIdea() *Idea {
	if m != nil {
		return m.Idea
	}
	return nil
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
	proto.RegisterType((*AddRatingsRequest)(nil), "go.micro.srv.ideas.AddRatingsRequest")
	proto.RegisterType((*AddRatingsResponse)(nil), "go.micro.srv.ideas.AddRatingsResponse")
	proto.RegisterEnum("go.micro.srv.ideas.AppType", AppType_name, AppType_value)
	proto.RegisterEnum("go.micro.srv.ideas.Pirority", Pirority_name, Pirority_value)
}

func init() { proto.RegisterFile("proto/idea/idea.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1069 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x5b, 0x6f, 0x1b, 0xc5,
	0x17, 0x8f, 0xbd, 0xbe, 0xed, 0xd9, 0x24, 0x7f, 0x77, 0xda, 0x7f, 0xbb, 0x75, 0xa1, 0x18, 0x0b,
	0x8a, 0x55, 0x51, 0x5b, 0x4a, 0x11, 0x2a, 0x6f, 0x38, 0xb1, 0x21, 0x16, 0x69, 0x1d, 0xad, 0x1d,
	0x40, 0x48, 0xc8, 0x4c, 0xbc, 0x53, 0x7b, 0x90, 0xf7, 0xd2, 0x99, 0xd9, 0x24, 0x16, 0x1f, 0x06,
	0x89, 0x77, 0x1e, 0xf9, 0x5a, 0x7c, 0x06, 0x34, 0x17, 0x3b, 0x9b, 0x78, 0x43, 0x2e, 0x7d, 0xe0,
	0xc5, 0x9a, 0x73, 0x3f, 0x73, 0xce, 0xef, 0x37, 0x6b, 0xf8, 0x7f, 0xcc, 0x22, 0x11, 0xb5, 0xa9,
	0x4f, 0xb0, 0xfa, 0x69, 0x29, 0x19, 0xa1, 0x69, 0xd4, 0x0a, 0xe8, 0x84, 0x45, 0x2d, 0xce, 0x4e,
	0x5a, 0xd2, 0xc0, 0x6b, 0x1f, 0x4d, 0xa3, 0x68, 0x3a, 0x27, 0x6d, 0xe5, 0x71, 0x9c, 0xbc, 0x6d,
	0x0b, 0x1a, 0x10, 0x2e, 0x70, 0x10, 0xeb, 0xa0, 0x5a, 0x67, 0x4a, 0xc5, 0x2c, 0x39, 0x6e, 0x4d,
	0xa2, 0xa0, 0xfd, 0x2b, 0xc5, 0xe1, 0x0c, 0x87, 0xed, 0x80, 0xbf, 0xe0, 0x09, 0x7d, 0xa1, 0x32,
	0xe8, 0xc8, 0x76, 0x34, 0x99, 0x24, 0x31, 0x16, 0x34, 0x0a, 0x53, 0x47, 0x93, 0xe2, 0xc9, 0xe5,
	0x1a, 0x24, 0x88, 0xc5, 0x42, 0x1b, 0x1b, 0xbf, 0x17, 0xa0, 0xd0, 0xf7, 0x09, 0x46, 0xdb, 0x90,
	0xef, 0x77, 0xdd, 0x5c, 0x3d, 0xd7, 0xb4, 0xbd, 0x7c, 0xbf, 0x8b, 0x10, 0x14, 0x42, 0x1c, 0x10,
	0x37, 0xaf, 0x34, 0xea, 0x2c, 0x75, 0x7c, 0x9e, 0x4c, 0x5d, 0x4b, 0xeb, 0xe4, 0x19, 0xd5, 0xc1,
	0xf1, 0x09, 0x9f, 0x30, 0x1a, 0xcb, 0x92, 0x6e, 0x41, 0x99, 0xd2, 0x2a, 0xf4, 0x10, 0x4a, 0x33,
	0xea, 0xfb, 0x24, 0x74, 0x8b, 0xf5, 0x5c, 0xb3, 0xe2, 0x19, 0x09, 0xbd, 0x02, 0x1b, 0xc7, 0xf1,
	0x58, 0x2c, 0x62, 0xc2, 0xdd, 0x52, 0xdd, 0x6a, 0x6e, 0xef, 0x3c, 0x69, 0xad, 0xcf, 0xa8, 0xd5,
	0x89, 0xe3, 0xd1, 0x22, 0x26, 0x5e, 0x05, 0xeb, 0x03, 0x47, 0x8f, 0xa0, 0x1c, 0xd0, 0x70, 0x8c,
	0xa7, 0xc4, 0x2d, 0xd7, 0x73, 0xcd, 0x82, 0x57, 0x0a, 0x68, 0xd8, 0x99, 0x12, 0x65, 0xc0, 0x67,
	0xca, 0x50, 0x31, 0x06, 0x7c, 0x26, 0x0d, 0x5f, 0x40, 0x99, 0x61, 0x41, 0xc3, 0x29, 0x77, 0xed,
	0xba, 0xd5, 0x74, 0x76, 0x6a, 0x59, 0x95, 0x3c, 0xe5, 0xe2, 0x2d, 0x5d, 0x65, 0xe7, 0xfa, 0xe8,
	0x42, 0x3d, 0xd7, 0xcc, 0x7b, 0x46, 0x42, 0x2f, 0xa1, 0xcc, 0x45, 0xc4, 0x28, 0xe1, 0xae, 0xa3,
	0xb2, 0x3d, 0xce, 0xca, 0x36, 0x14, 0x11, 0x5b, 0x78, 0x4b, 0x4f, 0xf4, 0x35, 0x38, 0xe7, 0xab,
	0xe1, 0xee, 0xa6, 0x0a, 0x7c, 0x9a, 0x15, 0x38, 0x58, 0xb9, 0x79, 0xe9, 0x10, 0xf4, 0x15, 0xc0,
	0x84, 0x11, 0x2c, 0x88, 0x3f, 0xc6, 0xc2, 0xdd, 0xaa, 0xe7, 0xcc, 0x3d, 0xe4, 0x76, 0x5b, 0xcb,
	0xed, 0xb6, 0x46, 0x4b, 0x04, 0x79, 0xb6, 0xf1, 0xee, 0x08, 0x19, 0x9a, 0xc4, 0xfe, 0x32, 0x74,
	0xfb, 0xfa, 0x50, 0xe3, 0xdd, 0x11, 0x8d, 0xbf, 0xf2, 0x50, 0x54, 0x57, 0x41, 0x0f, 0xa0, 0x28,
	0xa8, 0x98, 0x13, 0x83, 0x12, 0x2d, 0xa0, 0x57, 0x50, 0x89, 0x29, 0x8b, 0x18, 0x15, 0x0b, 0x05,
	0x96, 0xed, 0x9d, 0x0f, 0xb2, 0x2e, 0x75, 0x68, 0x7c, 0xbc, 0x95, 0x37, 0xaa, 0x82, 0x75, 0x3a,
	0x8b, 0x0c, 0x9a, 0xe4, 0x51, 0x02, 0xec, 0x74, 0x86, 0x85, 0x41, 0x91, 0x3a, 0x4b, 0xdd, 0x34,
	0xc2, 0x73, 0x05, 0x1e, 0xdb, 0x53, 0x67, 0xd9, 0x49, 0xc4, 0x7c, 0xc2, 0xdc, 0x92, 0xda, 0xb2,
	0x16, 0x52, 0x40, 0x2b, 0x5f, 0x00, 0xda, 0xc5, 0xb9, 0x55, 0xee, 0x3e, 0x37, 0xfb, 0x36, 0x73,
	0xfb, 0x33, 0x07, 0x25, 0x0d, 0x28, 0x09, 0xcb, 0x84, 0x13, 0x36, 0xa6, 0xbe, 0x19, 0x5d, 0x49,
	0x8a, 0x7d, 0x5f, 0xde, 0xe3, 0x04, 0xcf, 0x13, 0xcd, 0xb2, 0xbc, 0xa7, 0x85, 0x4b, 0xfd, 0x5a,
	0x77, 0xef, 0xb7, 0x70, 0x9b, 0x7e, 0xff, 0xb0, 0x00, 0xe4, 0x4b, 0xf0, 0x0d, 0x9d, 0x0b, 0xc2,
	0xe4, 0x72, 0xa8, 0xcf, 0xdd, 0x5c, 0xdd, 0x92, 0xcb, 0xa1, 0x3e, 0x97, 0xcd, 0xca, 0x57, 0x80,
	0xbb, 0x79, 0xa5, 0xd3, 0x42, 0x6a, 0xe8, 0xd6, 0xd5, 0xec, 0x2e, 0xdc, 0x91, 0xdd, 0xc5, 0xab,
	0xd8, 0x5d, 0xba, 0xc0, 0xee, 0x0f, 0x01, 0x64, 0x84, 0xe1, 0x6a, 0x59, 0xcd, 0xd2, 0x0e, 0x68,
	0x68, 0xc6, 0x2f, 0xcd, 0xf8, 0x6c, 0x69, 0xae, 0x18, 0x33, 0x3e, 0x33, 0xe6, 0x4b, 0xc4, 0xb4,
	0x6f, 0x4f, 0xcc, 0x8f, 0x61, 0x73, 0x92, 0x30, 0x46, 0x42, 0x31, 0x8e, 0x65, 0x77, 0xa0, 0xba,
	0x73, 0x8c, 0xee, 0x10, 0x4f, 0x09, 0x7a, 0x0c, 0x95, 0x98, 0x30, 0x6d, 0x76, 0x94, 0xb9, 0x1c,
	0x13, 0xa6, 0x4c, 0x0f, 0xa0, 0xf8, 0x2e, 0x21, 0x6c, 0xe1, 0x6e, 0x6a, 0x5a, 0x29, 0x41, 0x6a,
	0x79, 0xc4, 0x04, 0x77, 0xb7, 0xf4, 0xb4, 0x95, 0xd0, 0xd8, 0x83, 0x7b, 0x07, 0x94, 0x0b, 0xb9,
	0x27, 0xee, 0x11, 0x1e, 0x47, 0x21, 0x27, 0xa8, 0x05, 0x45, 0xd5, 0x9f, 0x5a, 0x96, 0xb3, 0xe3,
	0x66, 0xb5, 0x2e, 0x23, 0x3c, 0xed, 0xd6, 0xf8, 0x04, 0xaa, 0xc3, 0x59, 0x74, 0x6a, 0x92, 0xbc,
	0x4b, 0x08, 0x17, 0xeb, 0xeb, 0x6e, 0x74, 0x01, 0x1d, 0xc5, 0x9c, 0x30, 0x71, 0xc1, 0xef, 0xb6,
	0xb5, 0x7e, 0x83, 0xfb, 0x7b, 0x0a, 0x9d, 0x17, 0x5b, 0x76, 0xe5, 0x2a, 0xc5, 0x64, 0x46, 0x34,
	0x23, 0x2c, 0x6f, 0x29, 0xa2, 0x1a, 0x54, 0x82, 0xc8, 0xa7, 0x6f, 0x29, 0xf1, 0x15, 0x2b, 0x2c,
	0x6f, 0x25, 0x9f, 0x17, 0xb7, 0x6e, 0x5c, 0xfc, 0x48, 0xe1, 0xfb, 0xbf, 0x28, 0xde, 0x86, 0xfb,
	0x5d, 0x32, 0x27, 0x19, 0xc5, 0x7d, 0xa5, 0x5e, 0x15, 0x37, 0x62, 0xe3, 0x18, 0xee, 0x75, 0x7c,
	0x5f, 0x83, 0x72, 0x35, 0xef, 0x47, 0x50, 0x96, 0xe9, 0x52, 0x4f, 0x87, 0x14, 0xfb, 0x7e, 0xfa,
	0x8b, 0x96, 0xbf, 0xf1, 0x17, 0xad, 0xb1, 0x0b, 0x28, 0x5d, 0xc3, 0xf4, 0xf4, 0x39, 0x14, 0xa4,
	0xbb, 0xaa, 0xf0, 0x6f, 0x37, 0x53, 0x5e, 0xcf, 0xbf, 0x84, 0xb2, 0x21, 0x2d, 0x2a, 0x83, 0xf5,
	0x43, 0x6f, 0xb7, 0xba, 0x81, 0x00, 0x4a, 0xaf, 0x07, 0xbb, 0xfd, 0x83, 0x5e, 0x35, 0x87, 0x1c,
	0x28, 0x77, 0x7b, 0xc3, 0xef, 0x46, 0x83, 0xc3, 0x6a, 0x1e, 0xd9, 0x50, 0x1c, 0x8c, 0xf6, 0x7b,
	0x5e, 0xd5, 0x7a, 0xfe, 0x2d, 0x54, 0x96, 0x1f, 0x01, 0xa9, 0x3e, 0x7a, 0x33, 0xec, 0x8d, 0xaa,
	0x1b, 0xa8, 0x0a, 0x9b, 0x6f, 0xfa, 0x7b, 0xbd, 0xf1, 0x68, 0x30, 0xde, 0xef, 0x7c, 0x2f, 0x13,
	0xfc, 0x0f, 0x9c, 0xe1, 0xfe, 0xe0, 0xe8, 0xa0, 0xab, 0x15, 0x79, 0xb4, 0x05, 0xf6, 0xeb, 0xa3,
	0xe1, 0x48, 0x8b, 0xd6, 0xce, 0xdf, 0x05, 0x70, 0x64, 0x3f, 0x43, 0xc2, 0x4e, 0xe8, 0x84, 0xa0,
	0x11, 0xd8, 0x2b, 0x52, 0xa0, 0xa7, 0x57, 0x75, 0xaf, 0xdf, 0xb5, 0xda, 0xa7, 0x59, 0xf6, 0x35,
	0x4e, 0x35, 0x36, 0xd0, 0x2f, 0xe0, 0xa4, 0x90, 0x8b, 0x9e, 0x65, 0xc5, 0xad, 0x13, 0xa4, 0xf6,
	0x59, 0x96, 0x5f, 0x06, 0x05, 0x74, 0x85, 0x14, 0x3c, 0xdf, 0xaf, 0x42, 0x06, 0xce, 0x1b, 0x1b,
	0xe8, 0x47, 0x70, 0x52, 0x18, 0xbc, 0x76, 0x36, 0x99, 0x99, 0x33, 0x40, 0xdc, 0xd8, 0x40, 0x3d,
	0xb0, 0xf7, 0xa9, 0x7f, 0xc3, 0xbc, 0x0f, 0xd7, 0xbe, 0x40, 0x3d, 0xf9, 0x17, 0x54, 0xa7, 0x59,
	0x3d, 0x45, 0xef, 0x91, 0xe6, 0x67, 0x80, 0x73, 0x58, 0xa3, 0xcc, 0x15, 0xaf, 0x51, 0xab, 0xf6,
	0xec, 0x3a, 0xb7, 0xe5, 0x65, 0x77, 0x4b, 0x3f, 0x29, 0xe4, 0x1f, 0x97, 0x54, 0xe1, 0x97, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x33, 0x79, 0xd5, 0xbe, 0xe1, 0x0b, 0x00, 0x00,
}
