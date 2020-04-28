// Code generated by go generate via internal/cmd/service; DO NOT EDIT.
package qingstor

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
var _ httpclient.Options // Type is the type for qingstor
const Type = "qingstor"

var pairServiceCreateMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"location": struct{}{},
}

type pairServiceCreate struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	Location string
}

func (s *Service) parsePairCreate(opts ...*types.Pair) (*pairServiceCreate, error) {
	result := &pairServiceCreate{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairServiceCreateMap[v.Key]; !ok {
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
	v, ok = values[ps.Location]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Location)
	}
	if ok {
		result.Location = v.(string)
	}

	return result, nil
}

var pairServiceDeleteMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"location": struct{}{},
}

type pairServiceDelete struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasLocation bool
	Location    string
}

func (s *Service) parsePairDelete(opts ...*types.Pair) (*pairServiceDelete, error) {
	result := &pairServiceDelete{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairServiceDeleteMap[v.Key]; !ok {
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
	v, ok = values[ps.Location]
	if ok {
		result.HasLocation = true
		result.Location = v.(string)
	}

	return result, nil
}

var pairServiceGetMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"location": struct{}{},
}

type pairServiceGet struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasLocation bool
	Location    string
}

func (s *Service) parsePairGet(opts ...*types.Pair) (*pairServiceGet, error) {
	result := &pairServiceGet{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairServiceGetMap[v.Key]; !ok {
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
	v, ok = values[ps.Location]
	if ok {
		result.HasLocation = true
		result.Location = v.(string)
	}

	return result, nil
}

var pairServiceIsBucketNameValidMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairServiceIsBucketNameValid struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Service) parsePairIsBucketNameValid(opts ...*types.Pair) (*pairServiceIsBucketNameValid, error) {
	result := &pairServiceIsBucketNameValid{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairServiceIsBucketNameValidMap[v.Key]; !ok {
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

	return result, nil
}

var pairServiceListMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"location":      struct{}{},
	"storager_func": struct{}{},
}

type pairServiceList struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasLocation  bool
	Location     string
	StoragerFunc storage.StoragerFunc
}

func (s *Service) parsePairList(opts ...*types.Pair) (*pairServiceList, error) {
	result := &pairServiceList{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairServiceListMap[v.Key]; !ok {
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
	v, ok = values[ps.Location]
	if ok {
		result.HasLocation = true
		result.Location = v.(string)
	}
	v, ok = values[ps.StoragerFunc]
	if !ok {
		return nil, services.NewPairRequiredError(ps.StoragerFunc)
	}
	if ok {
		result.StoragerFunc = v.(storage.StoragerFunc)
	}

	return result, nil
}

var pairServiceNewMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"credential":          struct{}{},
	"endpoint":            struct{}{},
	"http_client_options": struct{}{},
}

type pairServiceNew struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	Credential           *credential.Provider
	HasEndpoint          bool
	Endpoint             endpoint.Provider
	HasHTTPClientOptions bool
	HTTPClientOptions    *httpclient.Options
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
		return nil, services.NewPairRequiredError(ps.Credential)
	}
	if ok {
		result.Credential = v.(*credential.Provider)
	}
	v, ok = values[ps.Endpoint]
	if ok {
		result.HasEndpoint = true
		result.Endpoint = v.(endpoint.Provider)
	}
	v, ok = values[ps.HTTPClientOptions]
	if ok {
		result.HasHTTPClientOptions = true
		result.HTTPClientOptions = v.(*httpclient.Options)
	}

	return result, nil
}

var pairServiceNewServicerMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairServiceNewServicer struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseServicePairNewServicer(opts ...*types.Pair) (*pairServiceNewServicer, error) {
	result := &pairServiceNewServicer{}

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

var pairServiceNewStoragerMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairServiceNewStorager struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func parseServicePairNewStorager(opts ...*types.Pair) (*pairServiceNewStorager, error) {
	result := &pairServiceNewStorager{}

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

var pairStorageAbortSegmentMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairStorageAbortSegment struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairAbortSegment(opts ...*types.Pair) (*pairStorageAbortSegment, error) {
	result := &pairStorageAbortSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageAbortSegmentMap[v.Key]; !ok {
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

	return result, nil
}

var pairStorageCompleteSegmentMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairStorageCompleteSegment struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairCompleteSegment(opts ...*types.Pair) (*pairStorageCompleteSegment, error) {
	result := &pairStorageCompleteSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageCompleteSegmentMap[v.Key]; !ok {
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

	return result, nil
}

var pairStorageCopyMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairStorageCopy struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairCopy(opts ...*types.Pair) (*pairStorageCopy, error) {
	result := &pairStorageCopy{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageCopyMap[v.Key]; !ok {
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

	return result, nil
}

var pairStorageDeleteMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairStorageDelete struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
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

	return result, nil
}

var pairStorageInitIndexSegmentMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairStorageInitIndexSegment struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairInitIndexSegment(opts ...*types.Pair) (*pairStorageInitIndexSegment, error) {
	result := &pairStorageInitIndexSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageInitIndexSegmentMap[v.Key]; !ok {
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

	return result, nil
}

var pairStorageListDirMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"dir_func":  struct{}{},
	"file_func": struct{}{},
}

type pairStorageListDir struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
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
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"object_func": struct{}{},
}

type pairStorageListPrefix struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	ObjectFunc types.ObjectFunc
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

	return result, nil
}

var pairStorageListPrefixSegmentsMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"segment_func": struct{}{},
}

type pairStorageListPrefixSegments struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	SegmentFunc segment.Func
}

func (s *Storage) parsePairListPrefixSegments(opts ...*types.Pair) (*pairStorageListPrefixSegments, error) {
	result := &pairStorageListPrefixSegments{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageListPrefixSegmentsMap[v.Key]; !ok {
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
	v, ok = values[ps.SegmentFunc]
	if !ok {
		return nil, services.NewPairRequiredError(ps.SegmentFunc)
	}
	if ok {
		result.SegmentFunc = v.(segment.Func)
	}

	return result, nil
}

var pairStorageMetadataMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairStorageMetadata struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
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

	return result, nil
}

var pairStorageMoveMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairStorageMove struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairMove(opts ...*types.Pair) (*pairStorageMove, error) {
	result := &pairStorageMove{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageMoveMap[v.Key]; !ok {
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

	return result, nil
}

var pairStorageNewMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"location": struct{}{},
	"name":     struct{}{},
	"work_dir": struct{}{},
}

type pairStorageNew struct {
	// Pre-defined pairs

	// Meta-defined pairs
	Location   string
	Name       string
	HasWorkDir bool
	WorkDir    string
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
	v, ok = values[ps.Location]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Location)
	}
	if ok {
		result.Location = v.(string)
	}
	v, ok = values[ps.Name]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Name)
	}
	if ok {
		result.Name = v.(string)
	}
	v, ok = values[ps.WorkDir]
	if ok {
		result.HasWorkDir = true
		result.WorkDir = v.(string)
	}

	return result, nil
}

var pairStorageReachMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"expire": struct{}{},
}

type pairStorageReach struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	Expire int
}

func (s *Storage) parsePairReach(opts ...*types.Pair) (*pairStorageReach, error) {
	result := &pairStorageReach{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageReachMap[v.Key]; !ok {
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
	v, ok = values[ps.Expire]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Expire)
	}
	if ok {
		result.Expire = v.(int)
	}

	return result, nil
}

var pairStorageReadMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"read_callback_func": struct{}{},
}

type pairStorageRead struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
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
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}

	return result, nil
}

var pairStorageStatMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairStorageStat struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
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

	return result, nil
}

var pairStorageStatisticalMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
}

type pairStorageStatistical struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
}

func (s *Storage) parsePairStatistical(opts ...*types.Pair) (*pairStorageStatistical, error) {
	result := &pairStorageStatistical{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageStatisticalMap[v.Key]; !ok {
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

	return result, nil
}

var pairStorageWriteMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"checksum":           struct{}{},
	"read_callback_func": struct{}{},
	"size":               struct{}{},
	"storage_class":      struct{}{},
}

type pairStorageWrite struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasChecksum         bool
	Checksum            string
	HasReadCallbackFunc bool
	ReadCallbackFunc    func([]byte)
	Size                int64
	HasStorageClass     bool
	StorageClass        storageclass.Type
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
	v, ok = values[ps.Checksum]
	if ok {
		result.HasChecksum = true
		result.Checksum = v.(string)
	}
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}
	v, ok = values[ps.Size]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Size)
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

var pairStorageWriteIndexSegmentMap = map[string]struct{}{
	// Pre-defined pairs
	"context": struct{}{},
	// Meta-defined pairs
	"read_callback_func": struct{}{},
}

type pairStorageWriteIndexSegment struct {
	// Pre-defined pairs
	Context context.Context

	// Meta-defined pairs
	HasReadCallbackFunc bool
	ReadCallbackFunc    func([]byte)
}

func (s *Storage) parsePairWriteIndexSegment(opts ...*types.Pair) (*pairStorageWriteIndexSegment, error) {
	result := &pairStorageWriteIndexSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageWriteIndexSegmentMap[v.Key]; !ok {
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
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}

	return result, nil
}

// CreateWithContext adds context support for Create.
func (s *Service) CreateWithContext(ctx context.Context, name string, pairs ...*types.Pair) (store storage.Storager, err error) {
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
func (s *Service) GetWithContext(ctx context.Context, name string, pairs ...*types.Pair) (store storage.Storager, err error) {
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
func (s *Storage) AbortSegmentWithContext(ctx context.Context, seg segment.Segment, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.AbortSegment")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.AbortSegment(seg, pairs...)
}

// CompleteSegmentWithContext adds context support for CompleteSegment.
func (s *Storage) CompleteSegmentWithContext(ctx context.Context, seg segment.Segment, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.CompleteSegment")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.CompleteSegment(seg, pairs...)
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

// InitIndexSegmentWithContext adds context support for InitIndexSegment.
func (s *Storage) InitIndexSegmentWithContext(ctx context.Context, path string, pairs ...*types.Pair) (seg segment.Segment, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.InitIndexSegment")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.InitIndexSegment(path, pairs...)
}

// ListDirWithContext adds context support for ListDir.
func (s *Storage) ListDirWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.ListDir")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.ListDir(path, pairs...)
}

// ListPrefixWithContext adds context support for ListPrefix.
func (s *Storage) ListPrefixWithContext(ctx context.Context, prefix string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.ListPrefix")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.ListPrefix(prefix, pairs...)
}

// ListPrefixSegmentsWithContext adds context support for ListPrefixSegments.
func (s *Storage) ListPrefixSegmentsWithContext(ctx context.Context, prefix string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.ListPrefixSegments")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.ListPrefixSegments(prefix, pairs...)
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

// WriteIndexSegmentWithContext adds context support for WriteIndexSegment.
func (s *Storage) WriteIndexSegmentWithContext(ctx context.Context, seg segment.Segment, r io.Reader, index int, size int64, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/qingstor.storage.WriteIndexSegment")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.WriteIndexSegment(seg, r, index, size, pairs...)
}
