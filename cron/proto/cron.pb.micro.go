// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cron.proto

package cron

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Cron service

func NewCronEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Cron service

type CronService interface {
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...client.CallOption) (*ScheduleResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Jobs(ctx context.Context, in *JobsRequest, opts ...client.CallOption) (*JobsResponse, error)
}

type cronService struct {
	c    client.Client
	name string
}

func NewCronService(name string, c client.Client) CronService {
	return &cronService{
		c:    c,
		name: name,
	}
}

func (c *cronService) Schedule(ctx context.Context, in *ScheduleRequest, opts ...client.CallOption) (*ScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "Cron.Schedule", in)
	out := new(ScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Cron.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronService) Jobs(ctx context.Context, in *JobsRequest, opts ...client.CallOption) (*JobsResponse, error) {
	req := c.c.NewRequest(c.name, "Cron.Jobs", in)
	out := new(JobsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cron service

type CronHandler interface {
	Schedule(context.Context, *ScheduleRequest, *ScheduleResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Jobs(context.Context, *JobsRequest, *JobsResponse) error
}

func RegisterCronHandler(s server.Server, hdlr CronHandler, opts ...server.HandlerOption) error {
	type cron interface {
		Schedule(ctx context.Context, in *ScheduleRequest, out *ScheduleResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Jobs(ctx context.Context, in *JobsRequest, out *JobsResponse) error
	}
	type Cron struct {
		cron
	}
	h := &cronHandler{hdlr}
	return s.Handle(s.NewHandler(&Cron{h}, opts...))
}

type cronHandler struct {
	CronHandler
}

func (h *cronHandler) Schedule(ctx context.Context, in *ScheduleRequest, out *ScheduleResponse) error {
	return h.CronHandler.Schedule(ctx, in, out)
}

func (h *cronHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.CronHandler.Delete(ctx, in, out)
}

func (h *cronHandler) Jobs(ctx context.Context, in *JobsRequest, out *JobsResponse) error {
	return h.CronHandler.Jobs(ctx, in, out)
}
