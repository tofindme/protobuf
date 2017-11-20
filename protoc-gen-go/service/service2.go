package service

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	. "github.com/golang/protobuf/protoc-gen-go/descriptor"
	. "github.com/golang/protobuf/protoc-gen-go/generator"
	"strconv"
)

const generatedCodeVersion = 5

// ServicePlugin2 produce the Service interface.
type ServicePlugin2 struct {
	*Generator
	ServiceName map[string]uint32
	ServiceId   map[uint32]uint32
}

// Name returns the name of the plugin.
func (p *ServicePlugin2) Name() string {
	return "ServiceId"
}

// Init is called once after data structures are built but before
// code generation begins.
func (p *ServicePlugin2) Init(g *Generator) {
	p.Generator = g
}

// Generate produces the code generated by the plugin for this file.
func (p *ServicePlugin2) GenerateImports(file *FileDescriptor) {
	//
}

// Generate generates the Service interface.
func (p *ServicePlugin2) Generate(file *FileDescriptor) {
	p.P("package ", p.PackageName, ";", "\n\n")
	p.P("enum ", "CMD {")
	for _, svc := range file.Service {
		for _, m := range svc.Method {
			method := CamelCase(*m.Name)
			if p.ServiceName[method] != 0 {
				panic(fmt.Sprintf("service name: %v already exist", method))
			}

			opt := p.getServiceMethodOption(m)
			if opt == nil {
				continue
			}
			id := *opt.Id
			if p.ServiceId[id] != 0 {
				panic(fmt.Sprintf("service id: %d for name: %v already exist", id, method))
			}
			p.ServiceName[method] = id
			p.In()
			p.P(method, "\t\t= ", strconv.Itoa(int(id)), " ;")
			p.Out()
		}
	}
	p.P("}")
}

func (p *ServicePlugin2) getServiceMethodOption(m *MethodDescriptorProto) *MethodOption {
	if m.Options != nil && proto.HasExtension(m.Options, E_MethodOption) {
		if ext, _ := proto.GetExtension(m.Options, E_MethodOption); ext != nil {
			if x, _ := ext.(*MethodOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func NewService() *ServicePlugin2 {
	return &ServicePlugin2{
		ServiceName: make(map[string]uint32),
		ServiceId:   make(map[uint32]uint32),
	}
}

func init() {
	RegisterPlugin(NewService())
}