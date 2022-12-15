package client

import (
	"fmt"
	"github.com/rancher/apiserver/pkg/types"
	"net/url"
)

type APIWarnings struct {
	ar *types.APIRequest
}

func (am APIWarnings) HandleWarningHeader(code int, agent string, message string) {
	if code != 299 || len(message) == 0 {
		return
	}

	fmt.Println("jiandao =====")
	fmt.Println(message)
	fmt.Println(url.QueryEscape(message))

	am.ar.Response.Header().Add("X-API-Warnings", url.QueryEscape(message))
}
