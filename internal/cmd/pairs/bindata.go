// Code generated by go-bindata. DO NOT EDIT.
// sources:
// pair.tmpl (818B)

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

var _pairTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x50\x4d\x8f\xd3\x30\x14\xbc\xfb\x57\x8c\xa2\x0a\x6d\xd1\x12\xdf\x8b\xf6\x80\x28\x07\x84\x44\xf7\x50\x01\xd7\x57\xe7\x91\x5a\x71\x6d\xcb\x76\xb2\x44\xc1\xff\x1d\x39\xd9\x0d\x1f\x5a\xb1\xcd\x29\x9e\x79\x6f\x66\xde\x48\x89\xf7\xae\x61\xb4\x6c\x39\x50\xe2\x06\xa7\x11\xad\x5b\xdf\xd0\x36\x71\xb0\x64\xa4\xba\x34\xd2\x93\x0e\xf1\x2d\xf6\x07\x7c\x3e\x1c\xf1\x61\xff\xf1\x58\x0b\x4f\xaa\xa3\x96\x31\x73\x42\xe8\x8b\x77\x21\xe1\x46\x00\x40\xa5\x9c\x4d\xfc\x23\x55\x62\x79\xb6\x3a\x9d\xfb\x53\xad\xdc\x45\x7e\xeb\xc9\x3e\x38\x19\x93\x0b\xd4\x72\xf5\x02\x2f\x7d\xd7\xca\xc8\xed\x85\x6d\xba\x6a\x96\x6d\xe3\x9d\xbe\x72\x58\x05\x6e\xd8\x26\x4d\xe6\xba\x1c\xcb\xbf\x32\x14\xe3\x8b\x0b\x69\xf4\x1c\x2b\xb1\x15\x42\x4a\xbc\x33\x06\x34\x90\x36\x74\x32\x8f\x8d\xd5\x42\x39\x1b\x4b\x61\xd3\xf4\x06\x81\x6c\xcb\xd8\x74\xb7\xd8\x0c\xd8\xdd\xa1\xde\x53\x22\xe4\x3c\xbb\x48\x89\x69\xc2\xa6\xc3\x4f\x24\x77\x4f\x51\x91\x41\xce\x78\xd0\xc6\xcc\xc4\x50\xef\x39\xaa\xa0\x7d\xd2\xce\x3e\x2d\x3d\xb7\x71\x87\x6a\x81\x73\xae\x66\x5b\xb6\x4d\x99\xdf\x8a\xff\x87\x28\xe4\xc6\x77\x05\xfb\x57\xb3\x5c\xf7\x55\xa7\x73\xd1\xf5\xdd\x1a\x8b\xbc\x37\x23\x9e\xcc\x30\x90\xe9\x19\xc9\xe1\x30\x67\x8c\x65\xeb\x78\xd6\x71\xae\x02\x3a\xa2\x8f\xdc\x14\xfe\xd9\x73\xbe\xf7\x56\xfd\x6d\x72\x33\x3c\x4e\x1e\x47\xcf\xc8\x79\x8b\xd7\x73\xdf\xf5\x7d\xd1\x9b\xe6\x02\x02\xa7\x3e\x58\xbc\xfa\x4d\x2c\x78\xf9\x3e\xf1\xb8\xc3\xaa\x76\xbb\xe2\x5f\x4a\xce\x1d\x86\x05\xc9\x22\xff\xd1\xd2\xaf\x00\x00\x00\xff\xff\xeb\x74\xcb\xbd\x32\x03\x00\x00")

func pairTmplBytes() ([]byte, error) {
	return bindataRead(
		_pairTmpl,
		"pair.tmpl",
	)
}

func pairTmpl() (*asset, error) {
	bytes, err := pairTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pair.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7, 0x44, 0x8e, 0x3e, 0x6a, 0x49, 0x14, 0x4d, 0xf3, 0x91, 0x8b, 0x4a, 0xf3, 0xbc, 0xde, 0x59, 0x77, 0x81, 0x99, 0x63, 0xa6, 0x94, 0x1f, 0x6c, 0x1c, 0xff, 0xdb, 0x33, 0x62, 0x11, 0xdd, 0xe2}}
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
	"pair.tmpl": pairTmpl,
}

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
	"pair.tmpl": &bintree{pairTmpl, map[string]*bintree{}},
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
