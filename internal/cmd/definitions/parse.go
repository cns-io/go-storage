package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Xuanwo/templateutils"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

const (
	pairPath      = "pairs.hcl"
	infoPath      = "infos.hcl"
	operationPath = "operations.hcl"
)

func parse() (data *Data) {
	// Parse pairs
	pairSpec := &PairSpec{}
	content, err := ioutil.ReadFile(pairPath)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
	err = parseHCL(content, pairPath, pairSpec)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	// Parse metadata
	metaSpec := &InfoSpec{}
	content, err = ioutil.ReadFile(infoPath)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
	err = parseHCL(content, infoPath, metaSpec)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	// Parse operations
	operationsSpec := &OperationsSpec{}
	content, err = ioutil.ReadFile(operationPath)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
	err = parseHCL(content, operationPath, operationsSpec)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	// Parse service
	serviceSpecs := make([]*ServiceSpec, 0)
	files, err := ioutil.ReadDir("services")
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
	for _, v := range files {
		if !strings.HasSuffix(v.Name(), ".hcl") {
			continue
		}

		filePath := "services/" + v.Name()
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("parse: %v", err)
		}

		srv := &ServiceSpec{}
		err = parseHCL(content, filePath, srv)
		if err != nil {
			log.Fatalf("parse: %v", err)
		}

		serviceSpecs = append(serviceSpecs, srv)
	}

	data = FormatData(pairSpec, metaSpec, operationsSpec, serviceSpecs)
	return data
}

func injectReadCallbackFunc(ops []*Op) {
	for _, op := range ops {
		fn := op.Func
		if fn == nil {
			continue
		}
		if strings.Contains(fn.Params, "io.Reader") ||
			strings.Contains(fn.Returns, "io.ReadCloser") {
			op.Generated = append(op.Generated, "read_callback_func")
		}
	}
}

func injectContext(ops []*Op) {
	for _, op := range ops {
		op.Generated = append(op.Generated, "context")
	}
}

func injectHTTPClientOptions(srv *Service) {
	// We don't need to inject http client into fs
	if srv.Name == "fs" {
		return
	}

	fn := func(ops []*Op) {
		for _, op := range ops {
			if op.Name != "new" {
				continue
			}
			op.Generated = append(op.Generated, "http_client_options")
			break
		}
	}

	// If service exist, inject into service level; Or inject into storage level.
	if len(srv.Service) > 0 {
		fn(srv.Service)
	} else {
		fn(srv.Storage)
	}
}

func parseHCL(src []byte, filename string, in interface{}) (err error) {
	var diags hcl.Diagnostics
	defer func() {
		if err != nil {
			err = fmt.Errorf("parse hcl: %w", err)
		}
	}()

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return diags
	}

	diags = gohcl.DecodeBody(file.Body, nil, in)
	if diags.HasErrors() {
		return diags
	}

	return nil
}

func parseFunc(service, name string) map[string]*templateutils.Method {
	data := make(map[string]*templateutils.Method)
	filename := fmt.Sprintf("../services/%s/%s.go", service, name)

	content, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		return data
	}
	if err != nil {
		log.Fatalf("read file failed: %v", err)
	}

	source := &templateutils.Source{}
	err = source.ParseContent(filename, content)
	if err != nil {
		log.Fatalf("parse content: %v", err)
	}

	for _, v := range source.Methods {
		v := v
		data[v.Name] = v
	}
	return data
}
