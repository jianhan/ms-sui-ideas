// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/tag/tag.proto

/*
Package idea is a generated protocol buffer package.

It is generated from these files:
	proto/tag/tag.proto

It has these top-level messages:
	Tag
	TagFilter
	ListTagsResponse
	UpsertTagsRequest
	UpsertTagsResponse
	DeleteTagsResponse
*/
package idea

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = google_protobuf.Empty{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for IdeaService service

type IdeaServiceClient interface {
	ListTags(ctx context.Context, in *TagFilter, opts ...client.CallOption) (*ListTagsResponse, error)
	CreateTags(ctx context.Context, in *UpsertTagsRequest, opts ...client.CallOption) (*UpsertTagsResponse, error)
	UpdateTags(ctx context.Context, in *UpsertTagsRequest, opts ...client.CallOption) (*UpsertTagsResponse, error)
	DeleteTags(ctx context.Context, in *TagFilter, opts ...client.CallOption) (*DeleteTagsResponse, error)
	HideTags(ctx context.Context, in *TagFilter, opts ...client.CallOption) (*google_protobuf.Empty, error)
	ShowTags(ctx context.Context, in *TagFilter, opts ...client.CallOption) (*google_protobuf.Empty, error)
}

type ideaServiceClient struct {
	c           client.Client
	serviceName string
}

func NewIdeaServiceClient(serviceName string, c client.Client) IdeaServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.ideas"
	}
	return &ideaServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *ideaServiceClient) ListTags(ctx context.Context, in *TagFilter, opts ...client.CallOption) (*ListTagsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "IdeaService.ListTags", in)
	out := new(ListTagsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ideaServiceClient) CreateTags(ctx context.Context, in *UpsertTagsRequest, opts ...client.CallOption) (*UpsertTagsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "IdeaService.CreateTags", in)
	out := new(UpsertTagsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ideaServiceClient) UpdateTags(ctx context.Context, in *UpsertTagsRequest, opts ...client.CallOption) (*UpsertTagsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "IdeaService.UpdateTags", in)
	out := new(UpsertTagsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ideaServiceClient) DeleteTags(ctx context.Context, in *TagFilter, opts ...client.CallOption) (*DeleteTagsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "IdeaService.DeleteTags", in)
	out := new(DeleteTagsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ideaServiceClient) HideTags(ctx context.Context, in *TagFilter, opts ...client.CallOption) (*google_protobuf.Empty, error) {
	req := c.c.NewRequest(c.serviceName, "IdeaService.HideTags", in)
	out := new(google_protobuf.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ideaServiceClient) ShowTags(ctx context.Context, in *TagFilter, opts ...client.CallOption) (*google_protobuf.Empty, error) {
	req := c.c.NewRequest(c.serviceName, "IdeaService.ShowTags", in)
	out := new(google_protobuf.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IdeaService service

type IdeaServiceHandler interface {
	ListTags(context.Context, *TagFilter, *ListTagsResponse) error
	CreateTags(context.Context, *UpsertTagsRequest, *UpsertTagsResponse) error
	UpdateTags(context.Context, *UpsertTagsRequest, *UpsertTagsResponse) error
	DeleteTags(context.Context, *TagFilter, *DeleteTagsResponse) error
	HideTags(context.Context, *TagFilter, *google_protobuf.Empty) error
	ShowTags(context.Context, *TagFilter, *google_protobuf.Empty) error
}

func RegisterIdeaServiceHandler(s server.Server, hdlr IdeaServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&IdeaService{hdlr}, opts...))
}

type IdeaService struct {
	IdeaServiceHandler
}

func (h *IdeaService) ListTags(ctx context.Context, in *TagFilter, out *ListTagsResponse) error {
	return h.IdeaServiceHandler.ListTags(ctx, in, out)
}

func (h *IdeaService) CreateTags(ctx context.Context, in *UpsertTagsRequest, out *UpsertTagsResponse) error {
	return h.IdeaServiceHandler.CreateTags(ctx, in, out)
}

func (h *IdeaService) UpdateTags(ctx context.Context, in *UpsertTagsRequest, out *UpsertTagsResponse) error {
	return h.IdeaServiceHandler.UpdateTags(ctx, in, out)
}

func (h *IdeaService) DeleteTags(ctx context.Context, in *TagFilter, out *DeleteTagsResponse) error {
	return h.IdeaServiceHandler.DeleteTags(ctx, in, out)
}

func (h *IdeaService) HideTags(ctx context.Context, in *TagFilter, out *google_protobuf.Empty) error {
	return h.IdeaServiceHandler.HideTags(ctx, in, out)
}

func (h *IdeaService) ShowTags(ctx context.Context, in *TagFilter, out *google_protobuf.Empty) error {
	return h.IdeaServiceHandler.ShowTags(ctx, in, out)
}