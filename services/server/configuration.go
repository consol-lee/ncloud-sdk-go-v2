/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-19T10:11:03Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

import (
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"os"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

func NewConfiguration(apiKey *ncloud.APIKey) *ncloud.Configuration {
	cfg := &ncloud.Configuration{
		BasePath:      "https://ncloud.apigw.ntruss.com/server/v2",
		DefaultHeader: make(map[string]string),
		UserAgent:     "server/1.0.0/go",
		APIKey:        apiKey,
	}
	if os.Getenv("NCLOUD_API_GW") != "" {
		cfg.BasePath = os.Getenv("NCLOUD_API_GW") + "/server/v2"
	}
	return cfg
}
