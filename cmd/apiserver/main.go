


package main

import (
	// Make sure dep tools picks up these dependencies
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "github.com/go-openapi/loads"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/cmd/server"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // Enable cloud provider auth

	"github.com/nightfury1204/test-apiserver/pkg/apis"
	"github.com/nightfury1204/test-apiserver/pkg/openapi"
)

func main() {
	version := "v0"
	server.StartApiServer("/registry/nahid.try", apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions, "Api", version)
}
