// Code generated by github.com/Je33/gqlgenc, DO NOT EDIT.

package gen

import (
	"context"
	"net/http"

	"github.com/Je33/gqlgenc/clientv2"
	"github.com/Je33/gqlgenc/example/input-scalar/scalar"
)

type GithubGraphQLClient interface {
	GetNumber(ctx context.Context, number scalar.Number, interceptors ...clientv2.RequestInterceptor) (*GetNumber, error)
}

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli *http.Client, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) GithubGraphQLClient {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type GetNumber struct {
	EnumToNum int "json:\"enumToNum\" graphql:\"enumToNum\""
}

func (t *GetNumber) GetEnumToNum() int {
	if t == nil {
		t = &GetNumber{}
	}
	return t.EnumToNum
}

const GetNumberDocument = `query GetNumber ($number: Number!) {
	enumToNum(number: $number)
}
`

func (c *Client) GetNumber(ctx context.Context, number scalar.Number, interceptors ...clientv2.RequestInterceptor) (*GetNumber, error) {
	vars := map[string]any{
		"number": number,
	}

	var res GetNumber
	if err := c.Client.Post(ctx, "GetNumber", GetNumberDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

var DocumentOperationNames = map[string]string{
	GetNumberDocument: "GetNumber",
}
