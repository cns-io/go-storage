// Code generated by go generate via internal/cmd/service; DO NOT EDIT.
package qingstor

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"

	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/pkg/credential"
	"github.com/Xuanwo/storage/pkg/endpoint"
	"github.com/Xuanwo/storage/pkg/segment"
	"github.com/Xuanwo/storage/pkg/storageclass"
	"github.com/Xuanwo/storage/types"
	"github.com/Xuanwo/storage/types/metadata"
	ps "github.com/Xuanwo/storage/types/pairs"
)

var _ credential.Provider
var _ endpoint.Provider
var _ segment.Segment
var _ storage.Storager
var _ storageclass.Type

// Type is the type for qingstor
const Type = "qingstor"

type pairServiceCreate struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	Location string
}

func parseServicePairCreate(opts ...*types.Pair) (*pairServiceCreate, error) {
	result := &pairServiceCreate{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	v, ok = values[ps.Location]
	if !ok {
		return nil, types.NewErrPairRequired(ps.Location)
	}
	if ok {
		result.Location = v.(string)
	}
	return result, nil
}

type pairServiceDelete struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasLocation bool
	Location    string
}

func parseServicePairDelete(opts ...*types.Pair) (*pairServiceDelete, error) {
	result := &pairServiceDelete{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	v, ok = values[ps.Location]
	if ok {
		result.HasLocation = true
		result.Location = v.(string)
	}
	return result, nil
}

type pairServiceGet struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasLocation bool
	Location    string
}

func parseServicePairGet(opts ...*types.Pair) (*pairServiceGet, error) {
	result := &pairServiceGet{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	v, ok = values[ps.Location]
	if ok {
		result.HasLocation = true
		result.Location = v.(string)
	}
	return result, nil
}

type pairServiceList struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasLocation     bool
	Location        string
	HasStoragerFunc bool
	StoragerFunc    storage.StoragerFunc
}

func parseServicePairList(opts ...*types.Pair) (*pairServiceList, error) {
	result := &pairServiceList{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	v, ok = values[ps.Location]
	if ok {
		result.HasLocation = true
		result.Location = v.(string)
	}
	v, ok = values[ps.StoragerFunc]
	if ok {
		result.HasStoragerFunc = true
		result.StoragerFunc = v.(storage.StoragerFunc)
	}
	return result, nil
}

type pairServiceNew struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	Credential  *credential.Provider
	HasEndpoint bool
	Endpoint    endpoint.Provider
}

func parseServicePairNew(opts ...*types.Pair) (*pairServiceNew, error) {
	result := &pairServiceNew{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	v, ok = values[ps.Credential]
	if !ok {
		return nil, types.NewErrPairRequired(ps.Credential)
	}
	if ok {
		result.Credential = v.(*credential.Provider)
	}
	v, ok = values[ps.Endpoint]
	if ok {
		result.HasEndpoint = true
		result.Endpoint = v.(endpoint.Provider)
	}
	return result, nil
}

type pairStorageAbortSegment struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairAbortSegment(opts ...*types.Pair) (*pairStorageAbortSegment, error) {
	result := &pairStorageAbortSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

type pairStorageCompleteSegment struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairCompleteSegment(opts ...*types.Pair) (*pairStorageCompleteSegment, error) {
	result := &pairStorageCompleteSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

type pairStorageCopy struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairCopy(opts ...*types.Pair) (*pairStorageCopy, error) {
	result := &pairStorageCopy{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

type pairStorageDelete struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairDelete(opts ...*types.Pair) (*pairStorageDelete, error) {
	result := &pairStorageDelete{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

type pairStorageInitSegment struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	PartSize int64
}

func parseStoragePairInitSegment(opts ...*types.Pair) (*pairStorageInitSegment, error) {
	result := &pairStorageInitSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	v, ok = values[ps.PartSize]
	if !ok {
		return nil, types.NewErrPairRequired(ps.PartSize)
	}
	if ok {
		result.PartSize = v.(int64)
	}
	return result, nil
}

type pairStorageList struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasDirFunc  bool
	DirFunc     types.ObjectFunc
	HasFileFunc bool
	FileFunc    types.ObjectFunc
}

func parseStoragePairList(opts ...*types.Pair) (*pairStorageList, error) {
	result := &pairStorageList{}

	values := make(map[string]interface{})
	for _, v := range opts {
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

type pairStorageListSegments struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasSegmentFunc bool
	SegmentFunc    segment.Func
}

func parseStoragePairListSegments(opts ...*types.Pair) (*pairStorageListSegments, error) {
	result := &pairStorageListSegments{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	v, ok = values[ps.SegmentFunc]
	if ok {
		result.HasSegmentFunc = true
		result.SegmentFunc = v.(segment.Func)
	}
	return result, nil
}

type pairStorageMetadata struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairMetadata(opts ...*types.Pair) (*pairStorageMetadata, error) {
	result := &pairStorageMetadata{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

type pairStorageMove struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairMove(opts ...*types.Pair) (*pairStorageMove, error) {
	result := &pairStorageMove{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

type pairStorageNewStorager struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairNewStorager(opts ...*types.Pair) (*pairStorageNewStorager, error) {
	result := &pairStorageNewStorager{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

type pairStorageReach struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	Expire int
}

func parseStoragePairReach(opts ...*types.Pair) (*pairStorageReach, error) {
	result := &pairStorageReach{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	v, ok = values[ps.Expire]
	if !ok {
		return nil, types.NewErrPairRequired(ps.Expire)
	}
	if ok {
		result.Expire = v.(int)
	}
	return result, nil
}

type pairStorageRead struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairRead(opts ...*types.Pair) (*pairStorageRead, error) {
	result := &pairStorageRead{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

type pairStorageStat struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairStat(opts ...*types.Pair) (*pairStorageStat, error) {
	result := &pairStorageStat{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

type pairStorageStatistical struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairStatistical(opts ...*types.Pair) (*pairStorageStatistical, error) {
	result := &pairStorageStatistical{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

type pairStorageWrite struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasChecksum     bool
	Checksum        string
	Size            int64
	HasStorageClass bool
	StorageClass    storageclass.Type
}

func parseStoragePairWrite(opts ...*types.Pair) (*pairStorageWrite, error) {
	result := &pairStorageWrite{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	v, ok = values[ps.Checksum]
	if ok {
		result.HasChecksum = true
		result.Checksum = v.(string)
	}
	v, ok = values[ps.Size]
	if !ok {
		return nil, types.NewErrPairRequired(ps.Size)
	}
	if ok {
		result.Size = v.(int64)
	}
	v, ok = values[ps.StorageClass]
	if ok {
		result.HasStorageClass = true
		result.StorageClass = v.(storageclass.Type)
	}
	return result, nil
}

type pairStorageWriteSegment struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseStoragePairWriteSegment(opts ...*types.Pair) (*pairStorageWriteSegment, error) {
	result := &pairStorageWriteSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
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
	return result, nil
}

// CreateWithContext adds context support for Create.
func (s *Service) CreateWithContext(ctx context.Context, name string, pairs ...*types.Pair) (storage.Storager, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.service.Create")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Create(name, pairs...)
}

// DeleteWithContext adds context support for Delete.
func (s *Service) DeleteWithContext(ctx context.Context, name string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.service.Delete")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Delete(name, pairs...)
}

// GetWithContext adds context support for Get.
func (s *Service) GetWithContext(ctx context.Context, name string, pairs ...*types.Pair) (storage.Storager, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.service.Get")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Get(name, pairs...)
}

// ListWithContext adds context support for List.
func (s *Service) ListWithContext(ctx context.Context, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.service.List")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.List(pairs...)
}

// AbortSegmentWithContext adds context support for AbortSegment.
func (s *Storage) AbortSegmentWithContext(ctx context.Context, id string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.AbortSegment")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.AbortSegment(id, pairs...)
}

// CompleteSegmentWithContext adds context support for CompleteSegment.
func (s *Storage) CompleteSegmentWithContext(ctx context.Context, id string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.CompleteSegment")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.CompleteSegment(id, pairs...)
}

// CopyWithContext adds context support for Copy.
func (s *Storage) CopyWithContext(ctx context.Context, src, dst string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.Copy")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Copy(src, dst, pairs...)
}

// DeleteWithContext adds context support for Delete.
func (s *Storage) DeleteWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.Delete")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Delete(path, pairs...)
}

// InitSegmentWithContext adds context support for InitSegment.
func (s *Storage) InitSegmentWithContext(ctx context.Context, path string, pairs ...*types.Pair) (id string, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.InitSegment")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.InitSegment(path, pairs...)
}

// ListWithContext adds context support for List.
func (s *Storage) ListWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.List")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.List(path, pairs...)
}

// ListSegmentsWithContext adds context support for ListSegments.
func (s *Storage) ListSegmentsWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.ListSegments")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.ListSegments(path, pairs...)
}

// MetadataWithContext adds context support for Metadata.
func (s *Storage) MetadataWithContext(ctx context.Context, pairs ...*types.Pair) (m metadata.StorageMeta, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.Metadata")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Metadata(pairs...)
}

// MoveWithContext adds context support for Move.
func (s *Storage) MoveWithContext(ctx context.Context, src, dst string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.Move")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Move(src, dst, pairs...)
}

// ReachWithContext adds context support for Reach.
func (s *Storage) ReachWithContext(ctx context.Context, path string, pairs ...*types.Pair) (url string, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.Reach")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Reach(path, pairs...)
}

// ReadWithContext adds context support for Read.
func (s *Storage) ReadWithContext(ctx context.Context, path string, pairs ...*types.Pair) (r io.ReadCloser, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.Read")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Read(path, pairs...)
}

// StatWithContext adds context support for Stat.
func (s *Storage) StatWithContext(ctx context.Context, path string, pairs ...*types.Pair) (o *types.Object, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.Stat")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Stat(path, pairs...)
}

// StatisticalWithContext adds context support for Statistical.
func (s *Storage) StatisticalWithContext(ctx context.Context, pairs ...*types.Pair) (m metadata.StorageStatistic, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.Statistical")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Statistical(pairs...)
}

// WriteWithContext adds context support for Write.
func (s *Storage) WriteWithContext(ctx context.Context, path string, r io.Reader, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.Write")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Write(path, r, pairs...)
}

// WriteSegmentWithContext adds context support for WriteSegment.
func (s *Storage) WriteSegmentWithContext(ctx context.Context, id string, offset, size int64, r io.Reader, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.WriteSegment")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.WriteSegment(id, offset, size, r, pairs...)
}
