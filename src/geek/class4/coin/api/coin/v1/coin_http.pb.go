// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.3

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type CoinHTTPServer interface {
	AddCoin(context.Context, *AddCoinRequest) (*AddCoinReply, error)
	ReduceCoin(context.Context, *ReduceCoinRequest) (*ReduceCoinReply, error)
	ShowCoin(context.Context, *ShowCoinRequest) (*ShowCoinReply, error)
}

func RegisterCoinHTTPServer(s *http.Server, srv CoinHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/coin/add", _Coin_AddCoin0_HTTP_Handler(srv))
	r.POST("/v1/coin/reduce", _Coin_ReduceCoin0_HTTP_Handler(srv))
	r.GET("/v1/coin/show", _Coin_ShowCoin0_HTTP_Handler(srv))
}

func _Coin_AddCoin0_HTTP_Handler(srv CoinHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCoinRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.coin.v1.Coin/AddCoin")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddCoin(ctx, req.(*AddCoinRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddCoinReply)
		return ctx.Result(200, reply)
	}
}

func _Coin_ReduceCoin0_HTTP_Handler(srv CoinHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReduceCoinRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.coin.v1.Coin/ReduceCoin")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReduceCoin(ctx, req.(*ReduceCoinRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReduceCoinReply)
		return ctx.Result(200, reply)
	}
}

func _Coin_ShowCoin0_HTTP_Handler(srv CoinHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ShowCoinRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.coin.v1.Coin/ShowCoin")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ShowCoin(ctx, req.(*ShowCoinRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ShowCoinReply)
		return ctx.Result(200, reply)
	}
}

type CoinHTTPClient interface {
	AddCoin(ctx context.Context, req *AddCoinRequest, opts ...http.CallOption) (rsp *AddCoinReply, err error)
	ReduceCoin(ctx context.Context, req *ReduceCoinRequest, opts ...http.CallOption) (rsp *ReduceCoinReply, err error)
	ShowCoin(ctx context.Context, req *ShowCoinRequest, opts ...http.CallOption) (rsp *ShowCoinReply, err error)
}

type CoinHTTPClientImpl struct {
	cc *http.Client
}

func NewCoinHTTPClient(client *http.Client) CoinHTTPClient {
	return &CoinHTTPClientImpl{client}
}

func (c *CoinHTTPClientImpl) AddCoin(ctx context.Context, in *AddCoinRequest, opts ...http.CallOption) (*AddCoinReply, error) {
	var out AddCoinReply
	pattern := "/v1/coin/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.coin.v1.Coin/AddCoin"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoinHTTPClientImpl) ReduceCoin(ctx context.Context, in *ReduceCoinRequest, opts ...http.CallOption) (*ReduceCoinReply, error) {
	var out ReduceCoinReply
	pattern := "/v1/coin/reduce"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.coin.v1.Coin/ReduceCoin"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CoinHTTPClientImpl) ShowCoin(ctx context.Context, in *ShowCoinRequest, opts ...http.CallOption) (*ShowCoinReply, error) {
	var out ShowCoinReply
	pattern := "/v1/coin/show"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.coin.v1.Coin/ShowCoin"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
