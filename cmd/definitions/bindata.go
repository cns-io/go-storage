// Code generated by go-bindata. DO NOT EDIT.
// sources:
// cmd/definitions/tmpl/function.tmpl (567B)
// cmd/definitions/tmpl/info.tmpl (1.702kB)
// cmd/definitions/tmpl/object.tmpl (1.907kB)
// cmd/definitions/tmpl/operation.tmpl (1.698kB)
// cmd/definitions/tmpl/pair.tmpl (502B)
// cmd/definitions/tmpl/service.tmpl (11.378kB)

// +build tools

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cmdDefinitionsTmplFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\x41\x4b\xc4\x30\x10\x85\xef\xfd\x15\x8f\x25\x87\xae\xec\xe6\x07\x08\x9e\x8a\x82\xb0\xc8\xa2\xde\x25\x64\xa7\x6b\xb0\x99\x94\x66\x5a\x17\x62\xfe\xbb\xa4\x55\x17\x11\x0f\xde\x3c\x25\xcc\xcc\x7b\xf3\x3e\x26\xa5\x2d\x06\xc3\x47\x82\x7a\xda\x40\x4d\xb8\xbc\x82\xd2\x37\x23\xdb\x88\x9c\xab\xd2\x76\x2d\x38\x08\xd4\xa4\x6f\x7d\xdf\x91\x27\x16\x3a\x7c\x36\x55\xcb\x2f\xb3\x66\xd2\x77\xc6\x13\xde\x20\xa1\x31\x9e\xba\x65\xa0\x88\xd5\xa4\x77\xc1\x9a\xb9\xd2\x8e\x6c\x51\x47\x5c\xa4\x04\x75\x56\xec\x4d\x5c\x06\xd6\x48\xa9\x58\xe6\x5c\xa7\xa4\x26\xbd\x37\x83\xf1\x51\x3f\x0e\xce\xef\x4c\x14\xfd\x20\x83\xe3\xe3\x35\x1f\xe2\xab\x93\xe7\x26\x78\x6f\x72\x46\xe8\x05\xbd\x71\xc3\x2f\xa6\xa5\x5c\x62\x7e\xdf\xb4\x2c\xb8\xa7\x38\x76\x12\x3f\x8c\xe7\x00\x15\x00\xf4\x86\x9d\xad\x57\x05\xdc\x9d\xa9\x57\xeb\x6a\xa6\xa2\x2e\xd2\x1f\x71\xac\x9c\x60\x03\x0b\x9d\x44\x37\xcb\xbb\xc1\x3f\x66\xdc\x82\xf8\xeb\xca\x3f\xbe\xef\x01\x00\x00\xff\xff\x41\xc8\xae\xbf\x37\x02\x00\x00")

func cmdDefinitionsTmplFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplFunctionTmpl,
		"cmd/definitions/tmpl/function.tmpl",
	)
}

func cmdDefinitionsTmplFunctionTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/function.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc1, 0xfc, 0xb, 0xf4, 0xfb, 0xf, 0x94, 0xca, 0x64, 0x7a, 0xb8, 0x48, 0xf9, 0x69, 0xf1, 0x88, 0x46, 0x59, 0xd3, 0xbc, 0x0, 0x4e, 0xe7, 0x8, 0x1e, 0xdf, 0x27, 0x59, 0x81, 0x75, 0x52, 0xa2}}
	return a, nil
}

var _cmdDefinitionsTmplInfoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\x5f\x6f\xd3\x30\x14\xc5\xdf\xfd\x29\x0e\x55\x85\x1a\xd4\x35\x9b\x84\x78\x80\xe5\x69\x1b\x68\x42\xdb\x90\x36\xf1\x82\x10\x72\x92\x9b\xca\x34\xb6\x23\xdb\x89\x56\x32\x7f\x77\xe4\x24\x74\x4d\xf7\x07\x8a\x78\xe1\xcd\x37\xf6\xbd\xe7\x77\x8f\xaf\x13\xc7\x38\xd1\x39\x61\x49\x8a\x0c\x77\x94\x23\x5d\x63\xa9\x37\x31\x32\x99\xc7\x39\x15\x42\x09\x27\xb4\xb2\xef\x70\x7a\x85\xcb\xab\x1b\x9c\x9d\x9e\xdf\x2c\x58\xc5\xb3\x15\x5f\x12\xdc\xba\x22\xcb\x98\x90\x95\x36\x0e\x33\x06\x00\x93\x42\xba\x09\x8b\x18\x6b\xdb\x03\x18\xae\x96\x84\xe9\x6a\x8e\xa9\x50\x85\xb6\x78\x9b\x60\x71\x1e\x56\x17\xbc\x82\xf7\xac\x6d\x31\xb5\x64\x1a\x91\xd1\x25\x97\x14\xf6\xa7\x2b\xdc\xc1\xe9\x13\x2e\xa9\x0c\x47\x58\x1c\xe3\xbd\xa0\x32\x87\x50\x39\xdd\x42\x28\xb4\xed\x76\x92\xf7\x48\x85\x63\x99\x56\x36\x40\xec\xe8\x36\x5d\xcd\x5e\xdd\xfb\x0e\x71\x37\xfd\x3c\xd4\x0d\x24\xcd\xa2\x83\x08\xf2\x9f\xb8\xcd\x78\xd0\x47\x82\xa3\xe3\xe3\xb0\xbb\xea\x81\x0f\x40\x2a\x0f\xcb\x88\xb1\x60\x00\x76\x7b\x18\xa7\x5b\x67\xea\xcc\xa1\x1d\x94\x37\x6c\xdf\x9e\x62\x0b\x18\x37\xeb\xaa\xaf\xe5\xfd\xd6\x97\xfb\x33\x1b\x86\x2e\x8e\xe3\x60\x00\x6a\x4b\x39\xb8\x05\x0f\x91\xe4\x15\x0a\x6d\xa0\xd3\xef\x94\x39\x34\xbc\xac\x69\x8e\x43\x48\xe2\xca\x42\x69\x07\x4b\x6e\x8e\xa3\xe1\x83\x25\xd7\x95\xea\xea\x08\xe5\xde\xbc\xee\x42\x09\xc9\xab\x2f\xd6\x19\xa1\x96\x5f\x85\x72\x64\x0a\x9e\x51\xeb\xd9\xa0\xfc\xbc\xd7\x61\x57\x14\x81\xfe\xec\xb6\x9b\x10\xef\x59\x51\xab\x0c\x33\x89\x57\xcf\xba\x16\xe1\x03\xb9\xbe\xf1\x53\x61\xab\x92\xaf\x07\x37\x66\xd1\xd8\x8f\xc1\x57\x43\xae\x36\x0a\x72\xf1\xc0\xbe\x40\xfa\xa7\x9a\xd7\x4f\x68\x36\x63\xcd\xe8\x37\x85\x06\xa6\x47\x60\x90\xa0\x19\xf1\xb2\x61\xa0\x4a\xdb\xb1\xfe\x03\x77\x66\x23\xd4\x39\x52\xad\xcb\x68\x20\x12\x05\xe4\x22\xdc\xf0\xcb\x3d\x9f\xc0\x8b\x04\x87\x43\x8d\xe7\xdd\x9e\xc3\x99\x9a\xba\x83\x7e\xbb\xd1\x2d\xa8\x3b\xfc\x20\xa3\x3f\x87\x79\xec\x12\x0a\x5e\x5a\xda\xe7\x96\x2e\x6a\xeb\xf6\x9b\x8e\xbf\xee\x3b\x19\xf7\x5d\x71\x25\xb2\x59\x21\xdd\xe2\xba\x32\x42\xb9\x62\x36\x79\x8c\xf5\x23\xa5\x3c\xbd\x7f\xb9\xbf\xee\x5e\x6c\x9e\xdd\x24\x8a\x1e\x5a\xf4\xdf\xcc\x6e\x6f\xe5\x5d\xb2\x9f\x97\x8f\x8e\x7d\xff\x0f\xdb\xf9\xa5\xdd\x2f\x7f\x06\x00\x00\xff\xff\x26\xf5\xd0\xb3\xa6\x06\x00\x00")

func cmdDefinitionsTmplInfoTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplInfoTmpl,
		"cmd/definitions/tmpl/info.tmpl",
	)
}

func cmdDefinitionsTmplInfoTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplInfoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/info.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa6, 0x9d, 0xdc, 0x29, 0x17, 0xc8, 0x8a, 0x14, 0x25, 0x13, 0xa5, 0x85, 0x13, 0x6, 0xf, 0xb3, 0x4a, 0x21, 0xd1, 0x20, 0x4e, 0x6f, 0x92, 0x60, 0xe2, 0x6e, 0x29, 0x96, 0x4a, 0xee, 0x99, 0x66}}
	return a, nil
}

var _cmdDefinitionsTmplObjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x51\x6f\xdb\x36\x10\x7e\x36\x7f\xc5\x37\x23\xd8\xa4\xc1\x91\x92\x6d\xd8\x43\x16\x3f\x0c\x71\xb6\x06\x68\xec\x02\x76\xfb\x4e\x53\x27\x9b\x8d\x44\x0a\x24\xe5\xda\x75\xf4\xdf\x0b\x4a\xb2\x23\x27\x4e\xe1\xa0\x2f\x7d\x23\x79\x77\xdf\xdd\xf7\xdd\x91\x8c\x63\xdc\xe8\x84\xb0\x20\x45\x86\x3b\x4a\x30\xdf\x60\xa1\xf7\x7b\x88\x3c\x89\x13\x4a\xa5\x92\x4e\x6a\x65\xff\xc1\x68\x82\xf1\x64\x86\xdb\xd1\xdd\x2c\x62\x05\x17\x0f\x7c\x41\x70\x9b\x82\x2c\x63\x32\x2f\xb4\x71\x08\x18\x00\xf4\xd3\xdc\xf5\x9b\x95\x93\x39\xb5\x4b\xbb\x51\xa2\xcf\x42\xc6\xe2\x18\xff\x49\xca\x12\x48\x95\xd0\x1a\x52\x41\xcf\x3f\x93\x70\x98\x4b\xc7\x84\x56\xd6\xe3\x6c\xb7\xe7\x30\x5c\x2d\x08\x67\x0f\x03\x9c\xad\x70\x35\x44\x34\xa9\xfd\xee\xc9\x71\x54\x55\x8d\xda\x44\xde\x79\xa0\xed\x16\x67\xab\x68\xcc\x73\xc2\x23\x9c\xfe\xc0\xad\xe0\x19\xaa\x0a\xa5\x54\xee\xef\xbf\x30\xc4\xe5\xf5\xb5\x77\x7a\xf0\xc1\x1e\x9f\x54\xe2\x97\x4d\x49\x0d\x36\xa4\x85\x5b\x12\x6c\xce\xb3\x8c\xac\x43\xa9\xa4\xf3\x25\x2e\xf4\xb9\x75\xda\xf0\x05\x45\x2c\x8e\x7d\xc0\x78\x32\xbb\x9d\x5e\xf9\x15\x70\xde\x86\xff\x66\x91\x7a\x6a\x16\xd3\x77\x93\x8f\xef\x47\x50\xda\x61\x4e\x10\x4b\x4f\x25\x81\x2e\x9d\x95\x09\xc1\x92\x59\x49\x41\x36\x3a\x0c\xc7\xcd\xbf\x63\x2f\xb1\x8f\xd0\x85\xa4\xe4\x99\x59\x5a\x08\xad\x44\x69\x0c\x29\x07\xcb\x53\x8a\x98\x6f\xc0\xce\x6e\x9d\x29\x85\xc3\xf6\x54\xf5\xbc\x9b\x4c\xbd\x6c\x23\xb2\xc2\xc8\xc2\x77\xfa\xc9\xf8\xaa\x61\x2f\xdd\x93\xdf\x6c\x53\x50\xad\x7d\x55\x75\x4e\x9e\x29\xcd\x7a\x71\x0c\x91\x49\x5f\x7d\x2b\xf4\x6e\xa7\xf0\x65\x29\xc5\xb2\xc3\x94\x67\x72\x45\x11\xeb\xb5\x1e\xd3\x46\x7d\xc3\xea\xa4\x71\xec\xa7\x05\xa5\xa5\x04\xdc\x82\xfb\x5d\xce\x0b\xa4\xda\xec\xa6\x69\xc5\xb3\x92\x06\xb8\x40\x4e\x5c\xd9\xba\x13\x96\xdc\x00\x97\xed\x81\x25\x57\x43\xd5\x38\xf5\x88\xb0\x5e\xa2\x15\xd5\x9b\x3f\xff\x60\xbd\xdc\x5b\xfd\xd4\x46\xf7\xa5\xa3\x35\xab\x18\x3b\x45\xd8\x27\x51\x6f\xd7\xf5\xa5\xa8\x2a\x96\x96\x4a\x20\xd0\xf8\xbd\x71\x0d\xf1\x3f\xb9\x56\x5f\x69\x8b\x8c\x6f\x5a\xe9\x82\xf0\x50\x3c\x6c\xeb\x12\x0d\xb9\xd2\x28\xe8\xe8\x85\xd6\xbe\xa8\x17\xe0\xd3\x57\xc0\x57\x87\xe0\xe1\x2e\xa2\xcd\x72\x04\x1e\x43\xac\x0e\x2a\x60\x6d\x3f\x33\x5b\x67\x7f\x0b\xb1\xe0\x20\xf9\x00\x73\xad\xb3\x70\x9f\xda\x3a\xee\x82\xb0\x69\xae\x4c\xa1\x23\xdf\x96\x5f\x4f\xba\xdf\xbf\x0c\x71\xd1\xe2\x7c\x5f\xac\x01\x9c\x29\xa9\x76\xac\x58\x97\x56\xa7\xb2\x47\x7c\x25\xa3\x3f\xf9\xd9\xa9\x23\x52\x9e\x59\x3a\xaa\xf2\x7d\x69\xdd\xdb\xda\xf8\x43\x2c\x87\x87\x2c\x0b\xae\xa4\x08\xd2\xdc\x45\xd3\xc2\x48\xe5\xd2\xa0\xdf\xce\x7d\x07\xa1\xaa\xfc\x45\x6a\x27\xbf\x1f\x86\x2d\xf5\x9f\x63\xa4\x1a\xf2\x8f\xc3\x53\xd8\x1f\x9d\xc1\xe6\x4d\xe9\x3e\x2f\x2f\xca\x16\x99\x56\x14\xac\x3b\x27\xdb\xfd\x13\x76\xca\xef\x72\xb4\xfc\xf5\x31\xd1\x9e\x3d\x8c\x1d\x86\x75\x80\xff\xda\xaa\x6f\x01\x00\x00\xff\xff\x94\x03\xb8\x29\x73\x07\x00\x00")

func cmdDefinitionsTmplObjectTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplObjectTmpl,
		"cmd/definitions/tmpl/object.tmpl",
	)
}

func cmdDefinitionsTmplObjectTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplObjectTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/object.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x0, 0x64, 0x33, 0x31, 0x2b, 0x9e, 0xa5, 0x1c, 0x46, 0x19, 0x57, 0xa8, 0x78, 0xb3, 0xc9, 0x92, 0xd7, 0x11, 0xee, 0xde, 0xf1, 0xa5, 0xf2, 0x26, 0x22, 0x58, 0x32, 0x8d, 0xfb, 0xb5, 0x58, 0xdc}}
	return a, nil
}

var _cmdDefinitionsTmplOperationTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x4d\x8f\xd3\x30\x10\xbd\xfb\x57\x8c\xac\x1e\x12\x69\x49\xef\x48\x7b\x82\x45\x5a\x09\x75\x2b\x10\xe2\x88\x5c\x67\xda\xb5\x48\x6c\x33\x9e\xec\x87\x82\xff\x3b\xb2\x93\x76\xd3\xaa\x10\x15\x10\xe2\x16\x7b\x66\xde\x9b\xf7\xc6\x13\xaf\xf4\x57\xb5\x43\xe0\x67\x8f\x41\x08\xd3\x7a\x47\x0c\x85\x00\x00\x90\xda\x59\xc6\x27\x96\xc3\xc9\x38\x29\x4a\x21\xfa\xfe\x15\x90\xb2\x3b\x84\xc5\x97\x2b\x58\x18\x78\x7d\x0d\xd5\xad\x65\xa4\xad\xd2\x18\x20\x46\xd1\xf7\xb0\x30\xd5\x5b\x0c\x9a\x8c\x67\xe3\x6c\xba\x4c\x0c\x30\x46\x4c\xf0\x8d\x7a\x5e\xa9\x16\x21\x46\x30\xfb\x62\xe8\x33\x53\x62\x30\x5b\x70\x04\x05\x7e\x4b\xf9\x39\x51\x06\xa4\x07\xa3\x91\x64\x79\x72\xcf\x8e\xd4\x2e\xdd\xc7\x98\xeb\x3f\x32\x19\xbb\x2b\x4a\x08\xf9\xe3\x80\x89\xb6\x4e\x8d\x8c\xe7\xa9\x08\xe7\x93\x8a\x85\xa9\xee\x7c\x16\x90\x32\x96\xcb\xdc\xad\xf3\x03\xcd\x77\x60\xb7\x56\x41\xab\x26\xb5\x3c\x46\x4e\x24\x8e\xc0\xe7\x6b\x8a\x31\xf2\xce\x51\xab\x78\xad\x48\xb5\x89\xab\x84\xe3\xc0\x07\x0c\x5d\xc3\xe1\xb3\xe1\xfb\xf5\x30\x9b\x23\x95\x32\x95\x4c\x5d\xb2\x8e\x73\xf5\x7b\x37\xd0\xcc\xf6\x9e\x90\xdf\x0c\x83\xfd\x1d\x19\x93\xf2\x42\xf3\x13\x8c\x6f\xa4\x1a\xef\xae\xfe\xba\xca\xfd\xd8\x8e\x4f\xf9\xd8\x76\x81\x6f\xda\x0d\xd6\x9f\xac\x69\x7d\x83\x2d\x5a\xc6\xfa\xdc\x1b\x2b\x4a\x11\x85\x58\x2e\x61\x36\x33\x83\xc2\x06\x01\x13\x70\x8d\x35\xb0\x83\x7b\xf5\x80\xb0\x75\xf4\xa8\xa8\x06\xed\x5a\xaf\xd8\x6c\x1a\x84\x03\x96\x4a\xde\x85\x6a\x78\xe4\xf3\x1c\x81\xa9\xd3\x0c\x7d\x14\x62\xdb\x59\x0d\x45\x98\x2f\x2a\x2f\x92\x7b\x29\xf6\xc9\xce\x8c\x9b\x48\xc8\x1d\x59\x90\xb3\x00\x32\xd9\xfb\xeb\x9d\xba\xa0\x99\x7f\xb2\x44\xc7\x3f\x9b\xb3\x6b\x84\x44\x70\x0d\x2b\x7c\xbc\xf3\x48\x79\xc4\x2b\xc7\xb7\x2f\xdd\xdf\x10\x39\x2a\xe4\xb4\xdb\x18\x65\x79\xfa\xc3\x79\x71\x52\x44\xf1\x33\xbe\x0b\xec\xf9\x0f\x76\x73\x34\xef\x4f\x0c\x9a\x38\xb2\x37\x6a\xfa\x75\x70\xef\x47\x00\x00\x00\xff\xff\xbd\xe8\x37\x32\xa2\x06\x00\x00")

func cmdDefinitionsTmplOperationTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplOperationTmpl,
		"cmd/definitions/tmpl/operation.tmpl",
	)
}

func cmdDefinitionsTmplOperationTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplOperationTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/operation.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc2, 0xaf, 0x9a, 0xb8, 0xf7, 0x9f, 0x1b, 0x53, 0xaa, 0x36, 0x95, 0x6d, 0xe6, 0x56, 0x7d, 0xb9, 0x8e, 0xd7, 0x4, 0x74, 0xce, 0x5b, 0x9, 0x4f, 0xb6, 0x92, 0xc, 0x1e, 0x49, 0x5, 0x4c, 0x4c}}
	return a, nil
}

var _cmdDefinitionsTmplPairTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x41\x6b\xe3\x30\x10\x85\xef\xfa\x15\x0f\xe3\x43\x02\x89\x75\xd9\x53\x96\x3d\x6d\xf6\xb0\x14\x92\x1c\x42\x7b\x2c\x8a\x3c\x51\x44\x6c\x49\xc8\x63\xb7\xc6\xf5\x7f\x2f\xb2\x93\x40\xe9\xa5\x3a\xcd\x7c\x33\xef\xf1\x46\x52\xe2\xaf\x2f\x09\x86\x1c\x45\xc5\x54\xe2\xd4\xc3\xf8\x47\x0f\x5d\x97\xb2\xa4\xb3\x75\x96\xad\x77\xcd\x6f\x6c\xf7\xd8\xed\x8f\xf8\xb7\xfd\x7f\x2c\x44\x50\xfa\xaa\x0c\x21\x28\x1b\x1b\x21\x6c\x1d\x7c\x64\x2c\x04\x00\x64\xda\x3b\xa6\x77\xce\xc4\xdc\x1a\xcb\x97\xf6\x54\x68\x5f\xcb\x13\xf5\xde\x95\x0d\xfb\xa8\x0c\x49\xe3\xd7\xf7\xb2\xfb\x25\xc3\xd5\xc8\x0b\x73\xd0\x95\x25\xc7\xd9\xa4\x2d\x7e\xac\xe6\x3e\x50\x93\x09\xb1\x14\x62\x18\xd6\x88\xca\x19\x42\xfe\xba\x42\xde\x61\xf3\x07\xc5\x21\x05\xc5\x38\x4e\xd3\x3c\x38\x55\x53\xe2\x79\x57\xec\x52\xf9\x01\xf6\x07\xd5\x68\x55\xa5\x1d\x29\xf1\x62\xf9\x32\x0c\xf7\xcd\x71\xc4\x9b\xad\x2a\xa8\x10\xaa\x1e\x89\xdf\x74\xe3\x88\x4e\x55\x2d\x81\x3d\xf6\x61\xfa\xa9\x42\x48\x29\xe6\x95\x2d\x35\x3a\xda\x09\x27\xdb\x73\xeb\xf4\x37\xe3\x45\x77\xf3\x3b\xf6\x21\xf5\x4b\xa4\xa8\x18\xa6\xfb\x23\x71\x1b\xdd\x44\x66\x90\xde\x13\xf5\x1b\x64\x5f\x42\x64\xab\xc7\xf4\x39\xc5\xd9\xa0\x9b\xc9\x28\xe6\x8b\xc9\x95\x29\xc1\x67\x00\x00\x00\xff\xff\x81\x9d\x53\x8d\xf6\x01\x00\x00")

func cmdDefinitionsTmplPairTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplPairTmpl,
		"cmd/definitions/tmpl/pair.tmpl",
	)
}

func cmdDefinitionsTmplPairTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplPairTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/pair.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1b, 0x1a, 0xbf, 0xf0, 0xe9, 0x47, 0x35, 0xe6, 0xa7, 0xc, 0xe6, 0xba, 0xc0, 0x5e, 0x39, 0xda, 0x2d, 0xb9, 0xd0, 0xee, 0xbc, 0x1, 0xc0, 0xd1, 0x5e, 0xa2, 0xbf, 0x4d, 0xcd, 0x4, 0xe9, 0xc0}}
	return a, nil
}

var _cmdDefinitionsTmplServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5b\x6f\xdc\xb8\x15\x7e\x9f\x5f\x71\x76\x60\x2c\x66\x82\x09\xd5\x02\x7d\x9a\xc2\x0f\x5b\xe7\x52\x63\xd3\xc4\x88\xb3\xdd\x87\xdd\x85\x41\x4b\x47\x33\xac\x29\x52\x4b\x52\xe3\x18\x13\xfd\xf7\x82\x17\xdd\xa5\xb9\xa4\x29\x36\x2d\xe2\x17\x8f\xc8\x73\x0e\x0f\xbf\x73\x25\xa5\x28\x82\x2b\x99\x20\x6c\x50\xa0\xa2\x06\x13\xb8\x7f\x82\x8d\xac\x9f\x61\xc7\x28\xc4\x59\x12\x25\x98\x32\xc1\x0c\x93\x42\xff\x15\x5e\xbc\x83\xb7\xef\x3e\xc0\xcb\x17\xd7\x1f\xc8\x2c\xa7\xf1\x03\xdd\x20\xec\xf7\x40\xde\xd2\x0c\xa1\x2c\x67\x33\x96\xe5\x52\x19\x58\xcc\x00\x00\xe6\xb1\x14\x06\x3f\x9a\xb9\x7f\x62\x72\x3e\xf3\xbf\x36\xcc\x6c\x8b\x7b\x12\xcb\x2c\xba\xc7\x27\x29\x12\x6d\xa4\xa2\x1b\x8c\x36\xf2\x79\xf5\x73\xf7\x97\x28\x7f\xd8\x44\x28\x92\x5c\x32\x51\xc9\x38\x83\x33\x56\x98\xa0\x30\x8c\xf2\xf3\x79\xb7\xc6\xe4\x31\x67\x78\xee\xba\x1a\xd5\x8e\xc5\xa8\x3d\x17\x39\x99\xcf\x3c\xe5\x96\x69\x39\x9b\xed\xa8\x82\x3b\x68\x74\x27\x37\x4a\xee\x58\x82\x2a\xcc\x54\x78\x90\x7f\x52\x5e\x60\x18\xbc\xf5\x92\x2a\x9a\x4a\x0b\x72\xeb\x7f\xbc\x54\x4a\x56\x73\xcd\xce\xc8\xbb\xdc\x99\x75\x36\x8b\x22\xf8\xf0\x94\x23\x30\x0d\x66\x8b\x60\x95\x81\x54\xaa\x8e\x65\x63\x29\xb4\xf1\x64\x97\x30\x6f\xcd\xcc\x1d\xff\xbb\xfb\x7f\x61\x6c\xfe\x81\x86\x26\xd4\x50\xb0\x5b\x43\x5d\x29\x02\x59\x35\x6e\xa5\x4a\x47\x4a\x66\x6e\x99\x01\x9f\x2a\x62\x03\xfb\xd9\x7e\xff\x1c\x14\x15\x1b\x84\x8b\xbb\x15\x5c\xec\x60\x7d\x09\xe4\x5a\xa4\x52\x5b\x65\x2c\xb8\x96\x82\xa5\x20\xa4\x81\x8b\x1d\x79\xcd\xe5\x3d\xe5\xed\xb9\x8b\x5c\x58\x05\xd7\x97\x76\xda\xe9\xfa\x09\x8c\xbc\xa1\x3a\xee\xd2\xb1\xd4\x12\xbc\x60\x3a\xe7\xf4\xa9\xda\x2d\x84\xbf\x96\xa0\xcb\x09\x32\x4b\x82\x22\x69\x1e\x1d\x1d\xea\x58\x31\x87\x6f\x7b\xc2\x4b\x2a\xcb\x40\xe5\xd0\x1c\x8a\x69\xfd\x2c\x1d\xb8\xaf\xd1\xf4\x70\x7a\x64\x9c\xc3\x06\x4d\x1f\xbf\x54\xc9\x2c\x8c\x91\x59\x14\x59\xe6\xe7\xf0\x61\xcb\x34\xa4\x85\x88\x9d\x3a\x7a\x2b\x0b\x9e\x38\xdc\xee\x11\x62\xca\xb9\x0f\xfe\xca\x56\x2c\xcb\x39\x66\x28\x0c\x2a\x52\xf1\x23\x28\x34\x85\x12\x4c\x6c\xfa\x2b\x32\x0d\x0a\x69\x02\x52\xf0\x27\xa0\x22\xe9\xc9\xcf\x64\xc2\x52\x86\x09\x99\x59\x05\x86\x3b\x59\x48\x78\xe6\x47\x96\x7d\xc9\x7b\x87\x8b\xcc\x56\x20\x1f\xac\x19\x25\x79\x8d\x26\x78\x74\xcd\xbe\x74\x44\x2c\xb5\x34\xfb\xda\x6c\x5e\x5b\x90\x19\x59\x74\x85\x7a\x72\x8f\x78\x20\xea\x12\xec\x2b\xcc\xf5\x38\xe6\x7a\x88\x39\x13\x46\x9e\x86\xb9\xc3\xa8\x01\x5d\x8a\x18\x57\x90\x73\xa4\x1a\x21\xa3\x0f\x08\xba\x50\x08\x94\x73\x70\x82\xb7\x54\xc3\x3d\xa2\x80\x47\xc5\x8c\x41\x01\xf7\x98\x4a\x85\x56\x87\x00\xe7\x40\xc9\x06\xce\x15\xd4\x9e\x50\x6f\xbe\x82\x94\xdc\x0e\x81\x94\xd9\xd2\xee\x7c\x34\xee\x6e\x28\x53\xba\xf2\xcc\xb1\x98\x3b\x25\xde\x9a\x58\xbb\x92\x22\xe5\x2c\x36\x76\x34\x8a\xe0\x05\xe6\x0a\x63\x5b\x83\xd6\xf0\x93\x46\xc8\xed\x6a\xe4\x67\x66\xb6\x9d\x90\x61\x42\x1b\xa4\x89\x83\xb8\x15\x21\x51\x04\x03\x52\x67\x29\x9a\xe7\xfc\x29\x04\x5a\x88\x57\xd8\xd9\x9c\x09\xd6\x5c\x3e\xf5\x05\x61\x23\x11\xeb\xe0\xed\x0b\x5e\xec\xba\x81\xbb\x04\x8b\xcc\xd0\xef\xec\x68\x33\xf8\x23\x3e\xad\x03\xdf\xab\x82\xf3\xa0\xcb\xaa\x9e\x77\x89\x7c\x0d\xbb\x55\xf0\xcd\x4e\x02\x68\xfd\x1c\xb7\x8d\x15\xa7\x73\x1a\x63\x27\x31\x1a\xcc\x72\x6e\xcb\xf8\x9c\xd9\x40\x4e\xed\xfc\xdc\xf9\xd8\x2d\xb7\x41\x5e\x61\x72\xb1\x23\xd7\x35\xc1\xb8\x80\x14\xa9\x29\xd4\x24\xfb\xab\x42\xc4\x13\x9c\xd6\x90\x77\x02\x1f\xdb\x9c\x8b\x31\xef\x58\xba\x41\x7c\x9c\x16\xf3\x19\xab\x57\xc1\x77\x94\xb5\x0f\xb1\xeb\x78\x7a\xc8\x35\x3e\x6e\x41\x67\x22\xc1\x8f\x40\xe0\x4f\xf5\xb8\xa3\xd5\xed\xb9\x3f\xdb\x39\x5b\x72\x17\xb5\x6e\x7d\xdb\x55\x5c\x41\xfd\xbb\xb6\xb3\x76\x62\x07\x2e\xe1\x7b\xef\x87\xdd\xf1\xfd\xa0\x74\x2c\x27\x77\xd3\x98\x71\x6a\x2f\xc3\x70\xbd\x48\xc5\x70\x4f\xbe\x6e\x07\x75\xca\xf2\x55\x90\xdb\x94\x6e\xab\xd2\x1b\x29\x35\xbe\xcb\x6d\x2b\xc9\xa4\xf8\x81\x73\xb8\x97\x92\x4f\x23\x61\x17\x0a\x30\x74\x59\x27\x21\xe9\xc8\xab\x36\xec\xc2\x89\x29\x53\x50\x7e\xd2\xe2\xa9\xe8\xaf\x5e\x65\xa9\x54\x90\x5b\x96\x15\xdc\xb5\xc6\x61\xae\x2f\xd9\xea\x96\x8a\x53\x95\x3b\xa0\xab\xcd\x17\x9f\xa5\x26\x47\xe1\x54\x08\x72\xba\x7a\xba\x2c\x34\xa5\xe2\x60\xb9\x9e\x96\x23\x16\xaa\x56\x81\x33\x96\x39\xcb\x7c\xbd\xc7\xe9\xc8\x6c\x12\xcb\xb1\xb8\x4c\xc5\xd0\x7f\xc3\xc4\x43\xb5\xad\x91\x5a\x65\xab\x4a\x1e\xf6\xe5\xbc\xdc\x6f\xf0\xc1\xd5\x21\xdf\x24\xe7\x54\x69\x4c\x82\xd7\xfb\x98\xe8\x71\x58\x86\xb2\xec\xc6\x85\xcb\x65\xf0\xcb\x6f\x16\x34\xef\x02\x51\x04\xef\xf1\xf7\x82\x29\x4c\xfc\xec\x98\x09\xec\x44\xa5\x6e\x4d\x1d\x10\xfb\x3b\xd5\x6e\x51\xca\xd4\x18\xd0\xd0\x86\xfa\x10\x59\x33\x3d\xd1\x96\x06\x6d\x7d\xe9\xa4\xfc\x34\x6d\x6b\xea\x2f\xad\xed\x71\x7d\x2b\x33\x2a\x8d\x37\x13\xb6\x74\x8d\x82\xa3\xf0\x95\x5c\xbb\x22\xe1\x1a\xba\x67\x13\xf6\xf7\x9d\xc1\x21\xa9\x0b\x99\x9b\xca\xc6\x4b\x58\x4c\xc8\x59\x01\xda\x53\x59\xd5\x93\x29\xd4\x05\x37\x16\xb5\x09\xfa\xa6\x9f\x70\xb8\xaf\xc1\xae\x52\xb5\x0c\xee\x9f\x3d\x5b\xdd\xad\xc0\x85\xab\xb7\x86\x53\xa4\x61\xd4\x8f\xcc\xc4\x5b\xd8\x91\x1f\xf1\xa9\x35\x3c\xee\x81\x67\x7a\xa1\xfd\x8b\x6d\x1f\x5b\x5b\xa5\xd5\xeb\xac\x6b\x92\xd0\xac\xfb\xdd\x92\x23\xce\xb0\xef\xb0\xb9\x15\xa4\x30\x4c\x14\xd8\x99\x28\x3b\x4f\xa7\x89\xbe\x04\xa3\x7a\x62\x02\xe3\x61\xae\x9d\x3f\x76\x93\x45\xdf\xf9\x96\x1d\xd0\x5a\x11\x33\x1d\x35\x67\x46\xce\x37\x84\x0f\x20\x5c\xb6\x03\xe1\x0c\xb7\x65\x29\x7c\x77\x2e\x54\xa1\xc9\x9f\x8a\xd3\x72\xd5\xdc\xbf\xd8\x14\x50\x2d\xe9\x2e\x61\xf6\xf6\x34\xa0\xd7\xbf\xfc\xa6\x8d\x62\x62\xb3\x0f\x17\x0c\x1d\x6b\x3e\x2f\x4b\x28\xeb\x1d\xc1\x48\xdb\x10\x34\xf0\x8a\xaf\x40\x30\x7e\xa4\x58\xfa\xae\x0f\x9a\x0b\x8d\x43\x9d\x5f\x4d\xe6\x0e\xa1\x83\xf2\x19\x1c\xfa\x05\xa6\xb4\xe0\xa6\x46\xc0\x9f\x11\x99\xb6\xeb\xda\x89\x50\xea\x6c\x52\xd2\x39\xc6\x2c\x65\x31\x50\xd7\x90\x3b\x09\xae\x60\x8e\xcb\xe8\x14\xcc\xd1\x5e\xc4\x29\xd6\xb9\xa5\x99\xe8\x32\x42\xa1\x1d\x71\x9a\x71\x57\x99\x12\xdf\x3f\x8e\x74\x8e\x14\xb9\x80\x8b\xdd\xc1\x5b\x9c\x81\x3d\xbe\xee\xc6\xa5\xe6\x38\xd2\xba\xfc\x11\x4d\xca\x69\x6d\xca\xd7\xd9\x90\x1c\xd1\xac\xd7\x5e\xff\x9f\x74\x4a\x0b\x0d\xcf\xea\xc9\xe5\xff\x64\xdf\x14\x45\xc0\xf4\x4f\x42\x17\x79\x2e\x95\xc1\xc4\x01\xa0\x30\x96\x2a\xd1\xf0\xb8\x45\xb3\x45\x05\x71\xa1\x14\x0a\x9f\xf6\x6c\xa8\x15\x0d\x3d\xa9\x25\x0d\xc5\xac\x2f\x21\xa5\x5c\xe3\xec\x58\x97\xf6\xad\x17\xfb\x8c\x4e\x61\x54\xcf\x91\xf6\xe1\x5b\x1b\xf6\xc7\x80\xdb\x4b\x79\x5f\x0d\xb6\x51\x04\xd7\x29\x14\x1a\x15\xa0\xa0\xf7\x1c\x7d\xe9\xdc\x05\x75\xdd\x16\xc2\x85\xda\x0a\x1e\x11\x62\x6a\xfb\x41\x3d\x42\xe5\x32\x24\x33\xa4\xaf\xbc\x26\xd5\x85\x1c\xe9\x5d\x02\x7d\xfa\x34\x31\xd9\xe4\xb4\x13\xc9\xce\x83\xe5\xb3\xfd\xe9\x4b\xfb\xd4\x89\x36\x1a\xe6\xd2\x9e\x6e\x23\xae\x18\x7a\xd3\xf5\x79\x82\xca\x26\x39\xdb\xf3\xc2\x90\x7e\x7f\x38\x26\x5a\xfc\x13\x8e\xc5\xa5\xd4\xb6\xec\x84\x8b\xc5\x8e\x6f\xb9\xca\xcb\x36\x42\x2a\x5f\x7a\x5b\x8b\xbb\x23\x05\x99\x8d\xbb\xd5\xf0\x0a\xb6\xeb\x32\xc3\x7b\xd6\x50\xea\x8f\xee\xa6\x31\xfa\x59\x07\xa1\xbe\xe2\x7b\x3b\xb8\xde\x75\xfa\xf1\x28\x82\xab\x2d\xc6\x0f\xa0\x3a\x37\x12\xe4\xdb\xb9\x6e\x78\xa1\xdf\xbc\xdf\xf8\xa2\x67\xbb\x93\x8f\x44\xfe\x65\xde\x1b\x19\x77\x13\x78\xf7\xb4\xc4\xed\xf4\x5d\xa5\xeb\xe4\xb9\xa9\x46\x83\xeb\xc1\xfb\xf7\xe1\x2b\x9d\xe3\x62\x46\x2f\x94\x8f\x21\xf9\xb9\xa7\xb2\x53\x0e\x65\xed\xf8\xf2\x97\xe5\xbd\x17\x8e\xfe\xc5\x71\xf7\xb5\xb1\x0b\xfd\x58\xa1\xdd\x3c\x85\xf0\x31\x0d\xdc\x3f\x55\x79\x8c\xb4\x5b\xec\x8b\x5c\x94\xe5\xb2\xb5\xd0\xc2\x5d\x41\x93\x1b\xaa\x68\xa6\xc9\xad\x73\x4c\x4b\x11\xc6\xdf\x3b\x2f\x6b\x4f\xf8\x50\x88\xcd\x47\xbb\x97\xb0\x1a\xf9\x1b\x8d\x1f\x36\x4a\x16\x22\x09\xef\xda\x83\x93\x6a\xd2\xac\xf4\x33\x33\xdb\x2b\x4f\xbf\x88\xcd\xc7\x15\x74\x56\xbe\xa2\x9c\xa3\xb2\x59\xbe\x0f\x45\x8b\x6f\x02\x95\x43\xfb\xeb\xad\x5a\x6b\x1c\xc6\x7a\x5a\x9c\xbc\xff\x04\x53\x54\xce\x06\x8b\x65\xb7\xfd\xbe\xc8\xa9\xd9\x56\x76\x0e\x62\x6f\xa8\xd9\xfa\x0d\x8e\x84\x07\x15\x09\x2c\xf0\xf7\xc0\x38\x9f\x2f\xc3\x93\x80\x79\x78\xfd\x3e\x5f\xb6\xd9\xa0\x3e\xab\x99\x2d\x5c\xc2\x7c\xf5\xeb\xfc\xd7\xf9\xbc\x2f\xb8\x57\xd3\x50\xd9\x82\xa5\x49\x2a\x55\x46\x8d\xcb\x46\x8b\xb9\xdf\xa2\xf5\xc5\xb2\x9c\xbb\x23\x53\x23\xb8\x2c\x21\x7c\x06\xb1\x58\x36\x21\x3f\xf5\xd2\xcb\x57\xac\xaa\xa7\x19\x94\x28\x7b\xd4\x11\xd2\x84\x6a\x96\x34\x15\x8b\xc6\x06\x38\x7b\x40\x30\xd6\xa5\x1b\xbe\x40\xdf\x7c\x65\x12\x8e\x47\x36\x63\x0f\x1b\x9a\x4e\xfd\xfa\xfe\xfb\x83\x24\xa3\x25\xcc\xa3\xf3\x16\x1f\x6b\xb2\xb7\xd2\x5c\x37\x8b\xd7\x78\x35\xc1\x5b\x96\xf3\x65\xaf\x28\x1c\xc8\xd9\xd5\x27\x11\xe8\x4f\x0b\xfe\x83\x0b\x99\x20\xcc\xbb\x79\x39\xb3\x63\xd5\x99\xa2\xa1\x1a\xa6\x67\x0b\x84\x24\x76\x92\x5c\xbb\xca\xe5\x38\xcb\x72\xb1\x1c\x6c\xab\xae\x47\x8d\xc0\x6b\xb1\xa3\x9c\x85\xa2\xf4\xf2\x63\x8e\xb1\xfb\xae\xc2\x4e\xb5\x64\xad\xe0\x87\xd8\x62\xb7\x06\xbf\x52\xbf\xa6\x1f\xd8\xaf\xbf\x06\xba\x04\x9a\xe7\x28\x12\x77\x4a\xd7\x2b\xd0\x24\xa4\x24\x77\x85\xd7\x4a\x0d\x84\x10\x0f\xe6\x8e\x2a\x7b\xae\x9e\x2a\xaf\x5e\xb8\xcc\xcd\xaa\x76\xe8\x83\x77\x05\x6e\xdd\xfa\xe3\x1f\xcb\xf2\xdd\xa5\x2d\x98\x83\x7a\xde\xee\x2e\x7a\xa9\xab\xc9\xd5\x57\x34\x43\x6b\x80\x91\xec\xf5\x41\xb1\xec\x0d\xd5\x26\xa4\xb1\x97\x22\xb1\x07\xf4\xed\x95\xcc\x32\x5a\x96\x56\xe3\xe5\x81\x22\xdd\x2f\x7f\x87\x2a\x75\x7b\x6e\xac\xd0\x54\xd9\xe1\x48\xb1\x09\x41\x7b\xb4\xe0\x78\xba\x8a\xfc\xdc\xc2\x63\xd9\xfe\x0b\xc5\xe7\x3f\xf1\xb0\x93\xbd\xac\x4a\x6b\xbe\xa7\x76\xd7\x49\xf0\xb8\x65\x1c\x61\x4b\x45\xc2\x99\xd8\x80\xb3\x9b\xdd\x60\xf8\x4c\xa8\x62\x73\x0e\x7a\x77\xb2\x7b\xf6\x7b\x4b\xa7\xf7\x84\xeb\x9d\xed\x75\xde\xb1\xdb\x9e\xe7\x0c\xc2\x04\x33\x75\xae\x38\xf1\x93\x21\xfb\x57\x27\x93\xf7\xb8\x61\xda\xa0\x9a\xba\x55\x57\x0b\x7b\x6c\x5b\xd9\xa4\x3a\x49\xb2\x1c\x5e\x31\xfe\x3b\x00\x00\xff\xff\x3f\x23\x19\xe1\x72\x2c\x00\x00")

func cmdDefinitionsTmplServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplServiceTmpl,
		"cmd/definitions/tmpl/service.tmpl",
	)
}

func cmdDefinitionsTmplServiceTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/service.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x93, 0x3d, 0x71, 0x6a, 0xc8, 0x48, 0x4a, 0x66, 0x7c, 0x29, 0xc7, 0xb4, 0x70, 0x65, 0x64, 0x54, 0x29, 0x73, 0xf4, 0xb3, 0xf5, 0x9d, 0x26, 0x1b, 0xca, 0xd5, 0x26, 0xb4, 0x61, 0x7d, 0x91, 0xcd}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"cmd/definitions/tmpl/function.tmpl":  cmdDefinitionsTmplFunctionTmpl,
	"cmd/definitions/tmpl/info.tmpl":      cmdDefinitionsTmplInfoTmpl,
	"cmd/definitions/tmpl/object.tmpl":    cmdDefinitionsTmplObjectTmpl,
	"cmd/definitions/tmpl/operation.tmpl": cmdDefinitionsTmplOperationTmpl,
	"cmd/definitions/tmpl/pair.tmpl":      cmdDefinitionsTmplPairTmpl,
	"cmd/definitions/tmpl/service.tmpl":   cmdDefinitionsTmplServiceTmpl,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"cmd": {nil, map[string]*bintree{
		"definitions": {nil, map[string]*bintree{
			"tmpl": {nil, map[string]*bintree{
				"function.tmpl": {cmdDefinitionsTmplFunctionTmpl, map[string]*bintree{}},
				"info.tmpl": {cmdDefinitionsTmplInfoTmpl, map[string]*bintree{}},
				"object.tmpl": {cmdDefinitionsTmplObjectTmpl, map[string]*bintree{}},
				"operation.tmpl": {cmdDefinitionsTmplOperationTmpl, map[string]*bintree{}},
				"pair.tmpl": {cmdDefinitionsTmplPairTmpl, map[string]*bintree{}},
				"service.tmpl": {cmdDefinitionsTmplServiceTmpl, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
