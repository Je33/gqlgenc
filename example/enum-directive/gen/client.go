// Code generated by github.com/Je33/gqlgenc, DO NOT EDIT.

package gen

import (
	"net/http"

	"github.com/Je33/gqlgenc/clientv2"
)

type GithubGraphQLClient interface {
}

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli *http.Client, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) GithubGraphQLClient {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

var DocumentOperationNames = map[string]string{}
