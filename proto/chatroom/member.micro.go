// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: chatroom/member.proto

/*
Package chatroom is a generated protocol buffer package.

It is generated from these files:
	chatroom/member.proto

It has these top-level messages:
	Member
	MemberResponse
	SetDataRequest
	GetMemberRequest
	LoginResponse
	TokenRequest
*/
package chatroom

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "gitlab.srgow.com/warehouse/proto/common"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = common.Response{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for MemberService service

type MemberService interface {
	Login(ctx context.Context, in *Member, opts ...client.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*common.Response, error)
	SetData(ctx context.Context, in *SetDataRequest, opts ...client.CallOption) (*common.Response, error)
	GetMember(ctx context.Context, in *GetMemberRequest, opts ...client.CallOption) (*MemberResponse, error)
	Validate(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*MemberResponse, error)
}

type memberService struct {
	c    client.Client
	name string
}

func NewMemberService(name string, c client.Client) MemberService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "chatroom"
	}
	return &memberService{
		c:    c,
		name: name,
	}
}

func (c *memberService) Login(ctx context.Context, in *Member, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "MemberService.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) Logout(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "MemberService.Logout", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) SetData(ctx context.Context, in *SetDataRequest, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "MemberService.SetData", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) GetMember(ctx context.Context, in *GetMemberRequest, opts ...client.CallOption) (*MemberResponse, error) {
	req := c.c.NewRequest(c.name, "MemberService.GetMember", in)
	out := new(MemberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) Validate(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*MemberResponse, error) {
	req := c.c.NewRequest(c.name, "MemberService.Validate", in)
	out := new(MemberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MemberService service

type MemberServiceHandler interface {
	Login(context.Context, *Member, *LoginResponse) error
	Logout(context.Context, *TokenRequest, *common.Response) error
	SetData(context.Context, *SetDataRequest, *common.Response) error
	GetMember(context.Context, *GetMemberRequest, *MemberResponse) error
	Validate(context.Context, *TokenRequest, *MemberResponse) error
}

func RegisterMemberServiceHandler(s server.Server, hdlr MemberServiceHandler, opts ...server.HandlerOption) error {
	type memberService interface {
		Login(ctx context.Context, in *Member, out *LoginResponse) error
		Logout(ctx context.Context, in *TokenRequest, out *common.Response) error
		SetData(ctx context.Context, in *SetDataRequest, out *common.Response) error
		GetMember(ctx context.Context, in *GetMemberRequest, out *MemberResponse) error
		Validate(ctx context.Context, in *TokenRequest, out *MemberResponse) error
	}
	type MemberService struct {
		memberService
	}
	h := &memberServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MemberService{h}, opts...))
}

type memberServiceHandler struct {
	MemberServiceHandler
}

func (h *memberServiceHandler) Login(ctx context.Context, in *Member, out *LoginResponse) error {
	return h.MemberServiceHandler.Login(ctx, in, out)
}

func (h *memberServiceHandler) Logout(ctx context.Context, in *TokenRequest, out *common.Response) error {
	return h.MemberServiceHandler.Logout(ctx, in, out)
}

func (h *memberServiceHandler) SetData(ctx context.Context, in *SetDataRequest, out *common.Response) error {
	return h.MemberServiceHandler.SetData(ctx, in, out)
}

func (h *memberServiceHandler) GetMember(ctx context.Context, in *GetMemberRequest, out *MemberResponse) error {
	return h.MemberServiceHandler.GetMember(ctx, in, out)
}

func (h *memberServiceHandler) Validate(ctx context.Context, in *TokenRequest, out *MemberResponse) error {
	return h.MemberServiceHandler.Validate(ctx, in, out)
}
