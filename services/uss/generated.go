// Code generated by go generate via internal/cmd/service; DO NOT EDIT.
package uss

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"

	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/pkg/credential"
	"github.com/Xuanwo/storage/pkg/endpoint"
	"github.com/Xuanwo/storage/pkg/httpclient"
	"github.com/Xuanwo/storage/pkg/segment"
	"github.com/Xuanwo/storage/pkg/storageclass"
	"github.com/Xuanwo/storage/services"
	"github.com/Xuanwo/storage/types"
	"github.com/Xuanwo/storage/types/metadata"
	ps "github.com/Xuanwo/storage/types/pairs"
)

var _ credential.Provider
var _ endpoint.Provider
var _ segment.Segment
var _ storage.Storager
var _ storageclass.Type
var _ services.ServiceError
var _ httpclient.Options

// Type is the type for uss
const Type = "uss"

var pairStorageDeleteMap = map[string]struct{}{
	// Meta-defined pairs
	"context": struct{}{},
}

type pairStorageDelete struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasContext bool
	Context    context.Context
}

func (s *Storage) parsePairDelete(opts ...*types.Pair) (*pairStorageDelete, error) {
	result := &pairStorageDelete{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageDeleteMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageListDirMap = map[string]struct{}{
	// Meta-defined pairs
	"context":   struct{}{},
	"dir_func":  struct{}{},
	"file_func": struct{}{},
}

type pairStorageListDir struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasContext  bool
	Context     context.Context
	HasDirFunc  bool
	DirFunc     types.ObjectFunc
	HasFileFunc bool
	FileFunc    types.ObjectFunc
}

func (s *Storage) parsePairListDir(opts ...*types.Pair) (*pairStorageListDir, error) {
	result := &pairStorageListDir{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageListDirMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}
	v, ok = values[ps.DirFunc]
	if ok {
		result.HasDirFunc = true
		result.DirFunc = v.(types.ObjectFunc)
	}
	v, ok = values[ps.FileFunc]
	if ok {
		result.HasFileFunc = true
		result.FileFunc = v.(types.ObjectFunc)
	}

	return result, nil
}

var pairStorageListPrefixMap = map[string]struct{}{
	// Meta-defined pairs
	"object_func": struct{}{},
	"context":     struct{}{},
}

type pairStorageListPrefix struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	ObjectFunc types.ObjectFunc
	HasContext bool
	Context    context.Context
}

func (s *Storage) parsePairListPrefix(opts ...*types.Pair) (*pairStorageListPrefix, error) {
	result := &pairStorageListPrefix{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageListPrefixMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs
	v, ok = values[ps.ObjectFunc]
	if !ok {
		return nil, services.NewPairRequiredError(ps.ObjectFunc)
	}
	if ok {
		result.ObjectFunc = v.(types.ObjectFunc)
	}
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageMetadataMap = map[string]struct{}{
	// Meta-defined pairs
	"context": struct{}{},
}

type pairStorageMetadata struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasContext bool
	Context    context.Context
}

func (s *Storage) parsePairMetadata(opts ...*types.Pair) (*pairStorageMetadata, error) {
	result := &pairStorageMetadata{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageMetadataMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageNewMap = map[string]struct{}{
	// Meta-defined pairs
	"credential":          struct{}{},
	"name":                struct{}{},
	"context":             struct{}{},
	"http_client_options": struct{}{},
	"work_dir":            struct{}{},
}

type pairStorageNew struct {
	// Pre-defined pairs

	// Meta-defined pairs
	Credential           *credential.Provider
	Name                 string
	HasContext           bool
	Context              context.Context
	HasHTTPClientOptions bool
	HTTPClientOptions    *httpclient.Options
	HasWorkDir           bool
	WorkDir              string
}

func parseStoragePairNew(opts ...*types.Pair) (*pairStorageNew, error) {
	result := &pairStorageNew{}

	values := make(map[string]interface{})
	for _, v := range opts {
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs

	// Parse meta-defined pairs
	v, ok = values[ps.Credential]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Credential)
	}
	if ok {
		result.Credential = v.(*credential.Provider)
	}
	v, ok = values[ps.Name]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Name)
	}
	if ok {
		result.Name = v.(string)
	}
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}
	v, ok = values[ps.HTTPClientOptions]
	if ok {
		result.HasHTTPClientOptions = true
		result.HTTPClientOptions = v.(*httpclient.Options)
	}
	v, ok = values[ps.WorkDir]
	if ok {
		result.HasWorkDir = true
		result.WorkDir = v.(string)
	}

	return result, nil
}

var pairStorageReadMap = map[string]struct{}{
	// Meta-defined pairs
	"context":            struct{}{},
	"read_callback_func": struct{}{},
}

type pairStorageRead struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasContext          bool
	Context             context.Context
	HasReadCallbackFunc bool
	ReadCallbackFunc    func([]byte)
}

func (s *Storage) parsePairRead(opts ...*types.Pair) (*pairStorageRead, error) {
	result := &pairStorageRead{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageReadMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}

	return result, nil
}

var pairStorageStatMap = map[string]struct{}{
	// Meta-defined pairs
	"context": struct{}{},
}

type pairStorageStat struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasContext bool
	Context    context.Context
}

func (s *Storage) parsePairStat(opts ...*types.Pair) (*pairStorageStat, error) {
	result := &pairStorageStat{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageStatMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageWriteMap = map[string]struct{}{
	// Meta-defined pairs
	"context":            struct{}{},
	"read_callback_func": struct{}{},
}

type pairStorageWrite struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasContext          bool
	Context             context.Context
	HasReadCallbackFunc bool
	ReadCallbackFunc    func([]byte)
}

func (s *Storage) parsePairWrite(opts ...*types.Pair) (*pairStorageWrite, error) {
	result := &pairStorageWrite{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageWriteMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Parse pre-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.Context = v.(context.Context)
	} else {
		result.Context = context.Background()
	}

	// Parse meta-defined pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}

	return result, nil
}

// DeleteWithContext adds context support for Delete.
func (s *Storage) DeleteWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/uss.storage.Delete")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Delete(path, pairs...)
}

// ListDirWithContext adds context support for ListDir.
func (s *Storage) ListDirWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/uss.storage.ListDir")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.ListDir(path, pairs...)
}

// ListPrefixWithContext adds context support for ListPrefix.
func (s *Storage) ListPrefixWithContext(ctx context.Context, prefix string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/uss.storage.ListPrefix")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.ListPrefix(prefix, pairs...)
}

// MetadataWithContext adds context support for Metadata.
func (s *Storage) MetadataWithContext(ctx context.Context, pairs ...*types.Pair) (m metadata.StorageMeta, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/uss.storage.Metadata")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Metadata(pairs...)
}

// ReadWithContext adds context support for Read.
func (s *Storage) ReadWithContext(ctx context.Context, path string, pairs ...*types.Pair) (r io.ReadCloser, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/uss.storage.Read")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Read(path, pairs...)
}

// StatWithContext adds context support for Stat.
func (s *Storage) StatWithContext(ctx context.Context, path string, pairs ...*types.Pair) (o *types.Object, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/uss.storage.Stat")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Stat(path, pairs...)
}

// WriteWithContext adds context support for Write.
func (s *Storage) WriteWithContext(ctx context.Context, path string, r io.Reader, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/uss.storage.Write")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Write(path, r, pairs...)
}
