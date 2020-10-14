package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func formatGlobal(data *Data) {
	// Generate pairs
	hf := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(data.pairSpec, hf.Body())

	formatBody(hf.Body())

	content := hclwrite.Format(hf.Bytes())
	err := ioutil.WriteFile(pairPath, content, 0644)
	if err != nil {
		log.Fatalf("formatGlobal: %v", err)
	}

	// Generate metadata
	hf = hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(data.infoSpec, hf.Body())

	formatBody(hf.Body())

	content = hclwrite.Format(hf.Bytes())
	err = ioutil.WriteFile(infoPath, content, 0644)
	if err != nil {
		log.Fatalf("formatGlobal: %v", err)
	}

	// Generate operations
	hf = hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(data.operationsSpec, hf.Body())

	formatBody(hf.Body())

	content = hclwrite.Format(hf.Bytes())
	err = ioutil.WriteFile(operationPath, content, 0644)
	if err != nil {
		log.Fatalf("formatGlobal: %v", err)
	}
}

func formatService(data *Data) {
	filePath := fmt.Sprintf(os.Args[1])

	hf := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(data.serviceSpec, hf.Body())

	formatBody(hf.Body())

	content := hclwrite.Format(hf.Bytes())
	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		log.Fatalf("formatGlobal: %v", err)
	}
}

func formatBody(body *hclwrite.Body) {
	for k, v := range body.Attributes() {
		if isAttrEmpty(v) {
			body.RemoveAttribute(k)
			continue
		}
	}
	for _, v := range body.Blocks() {
		if isBlockEmpty(v) {
			body.RemoveBlock(v)
			continue
		}

		formatBody(v.Body())
	}
}

func isBlockEmpty(block *hclwrite.Block) bool {
	if len(block.Body().Blocks()) > 0 {
		return false
	}

	attrs := block.Body().Attributes()
	for _, v := range attrs {
		if !isAttrEmpty(v) {
			return false
		}
	}
	return true
}

func isAttrEmpty(attr *hclwrite.Attribute) bool {
	tokens := attr.Expr().BuildTokens(make([]*hclwrite.Token, 0))

	// xxx = null
	if len(tokens) == 1 && tokens[0].Type == hclsyntax.TokenIdent {
		s := string(tokens[0].Bytes)
		switch s {
		case "true":
			return false
		case "null", "false":
			return true
		default:
			log.Fatalf("not handled token: %s", s)
		}
	}
	// xxx = ""
	if len(tokens) == 2 &&
		tokens[0].Type == hclsyntax.TokenOQuote &&
		tokens[1].Type == hclsyntax.TokenCQuote {
		return true
	}
	// xxx = []
	if len(tokens) == 2 &&
		tokens[0].Type == hclsyntax.TokenOBrack &&
		tokens[1].Type == hclsyntax.TokenCBrack {
		return true
	}

	return false
}