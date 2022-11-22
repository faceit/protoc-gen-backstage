package main

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
	"google.golang.org/protobuf/proto"

	"github.com/faceit/protoc-gen-backstage/proto/gen/github.com/faceit/protoc-gen-backstage/efg/backstage"
)

func main() {
	if err := protokit.RunPlugin(&plugin{}); err != nil {
		log.Fatal(err)
	}
}

type plugin struct{}

type PluginOptions struct {
	OutputFile string
}

func (p *plugin) Generate(req *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	options, err := ParseOptions(req)
	if err != nil {
		return nil, err
	}

	descriptors := protokit.ParseCodeGenRequest(req)
	resp := new(plugin_go.CodeGeneratorResponse)

	for _, descriptor := range descriptors {
		content := "# Code generated by protoc-gen-backstage. DO NOT EDIT."
		if len(descriptor.Services) > 0 {
			for _, service := range descriptor.Services {
				options := service.ServiceDescriptorProto.Options
				owner := proto.GetExtension(options, backstage.E_Owner).(string)
				system := proto.GetExtension(options, backstage.E_System).(string)

				var lifecycle = "production"
				deprecated := options.GetDeprecated()
				if deprecated {
					lifecycle = "deprecated"
				}

				if owner == "" {
					content += `
# Owner is a required field your Service "` + service.FullName + `" must implement the custom option
# See github.com/faceit/protoc-gen-backstage for more info
`
				} else {
					content += `
---
apiVersion: backstage.io/v1alpha1
kind: API
metadata:
  name: ` + strings.Replace(service.FullName, ".", "-", -1) + `
spec:
  type: grpc
  lifecycle: ` + lifecycle +
						`
  owner: ` + owner
					if system != "" {
						content += `
  system: " + system
`
					}
					content += `
  definition:
    $text: /proto/` + descriptor.GetName() + `
`
				}
			}

			dir := strings.Split(descriptor.GetName(), ".")[0] + "/"

			resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
				Name:    proto.String(filepath.Join(dir, options.OutputFile)),
				Content: proto.String(content),
			})
		}
	}

	return resp, nil
}

func ParseOptions(req *plugin_go.CodeGeneratorRequest) (*PluginOptions, error) {
	options := &PluginOptions{OutputFile: "catalog-info.yaml"}

	params := req.GetParameter()

	if params != "" {
		options.OutputFile = params
	}

	return options, nil
}
