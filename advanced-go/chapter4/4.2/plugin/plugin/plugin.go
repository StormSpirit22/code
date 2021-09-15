package plugin

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type NetrpcPlugin struct{ *generator.Generator }

func (p *NetrpcPlugin) Name() string                { return "netrpc" }
func (p *NetrpcPlugin) Init(g *generator.Generator) { p.Generator = g }

func (p *NetrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

func (p *NetrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *NetrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P("// TODO: import code")
}

func (p *NetrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	p.P("// TODO: service code, Name = " + svc.GetName())
}

