package api

/************************************************
  generated by IDE. for [GSLBAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件のリセット
func (api *GSLBAPI) Reset() *GSLBAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *GSLBAPI) Offset(offset int) *GSLBAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *GSLBAPI) Limit(limit int) *GSLBAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *GSLBAPI) Include(key string) *GSLBAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *GSLBAPI) Exclude(key string) *GSLBAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルター
func (api *GSLBAPI) FilterBy(key string, value interface{}) *GSLBAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *GSLBAPI) FilterMultiBy(key string, value interface{}) *GSLBAPI {
	api.filterBy(key, value, true)
	return api
}

// WithNameLike 名称条件
func (api *GSLBAPI) WithNameLike(name string) *GSLBAPI {
	return api.FilterBy("Name", name)
}

// WithTag タグ条件
func (api *GSLBAPI) WithTag(tag string) *GSLBAPI {
	return api.FilterBy("Tags.Name", tag)
}

// WithTags タグ(複数)条件
func (api *GSLBAPI) WithTags(tags []string) *GSLBAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *GSLBAPI) WithSizeGib(size int) *GSLBAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *GSLBAPI) WithSharedScope() *GSLBAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *GSLBAPI) WithUserScope() *GSLBAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

// SortBy 指定キーでのソート
func (api *GSLBAPI) SortBy(key string, reverse bool) *GSLBAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名称でのソート
func (api *GSLBAPI) SortByName(reverse bool) *GSLBAPI {
	api.sortByName(reverse)
	return api
}

// func (api *GSLBAPI) SortBySize(reverse bool) *GSLBAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
   To support Setxxx interface for Find()
************************************************/

// SetEmpty 検索条件のリセット
func (api *GSLBAPI) SetEmpty() {
	api.reset()
}

// SetOffset オフセット
func (api *GSLBAPI) SetOffset(offset int) {
	api.offset(offset)
}

// SetLimit リミット
func (api *GSLBAPI) SetLimit(limit int) {
	api.limit(limit)
}

// SetInclude 取得する項目
func (api *GSLBAPI) SetInclude(key string) {
	api.include(key)
}

// SetExclude 除外する項目
func (api *GSLBAPI) SetExclude(key string) {
	api.exclude(key)
}

// SetFilterBy 指定キーでのフィルター
func (api *GSLBAPI) SetFilterBy(key string, value interface{}) {
	api.filterBy(key, value, false)
}

// SetFilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *GSLBAPI) SetFilterMultiBy(key string, value interface{}) {
	api.filterBy(key, value, true)
}

// SetNameLike 名称条件
func (api *GSLBAPI) SetNameLike(name string) {
	api.FilterBy("Name", name)
}

// SetTag タグ条件
func (api *GSLBAPI) SetTag(tag string) {
	api.FilterBy("Tags.Name", tag)
}

// SetTags タグ(複数)条件
func (api *GSLBAPI) SetTags(tags []string) {
	api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *GSLBAPI) SetSizeGib(size int)  {
// 	api.FilterBy("SizeMB", size*1024)
// }

// func (api *GSLBAPI) SetSharedScope() {
// 	api.FilterBy("Scope", "shared")
// }

// func (api *GSLBAPI) SetUserScope()  {
// 	api.FilterBy("Scope", "user")
// }

// SetSortBy 指定キーでのソート
func (api *GSLBAPI) SetSortBy(key string, reverse bool) {
	api.sortBy(key, reverse)
}

// SetSortByName 名称でのソート
func (api *GSLBAPI) SetSortByName(reverse bool) {
	api.sortByName(reverse)
}

// func (api *GSLBAPI) SetSortBySize(reverse bool)  {
// 	api.sortBy("SizeMB", reverse)
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// func (api *GSLBAPI) New() *sacloud.GSLB {
// 	return &sacloud.GSLB{}
// }

// func (api *GSLBAPI) Create(value *sacloud.GSLB) (*sacloud.GSLB, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.create(api.createRequest(value), res)
// 	})
// }

// func (api *GSLBAPI) Read(id string) (*sacloud.GSLB, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.read(id, nil, res)
// 	})
// }

// func (api *GSLBAPI) Update(id string, value *sacloud.GSLB) (*sacloud.GSLB, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.update(id, api.createRequest(value), res)
// 	})
// }

// func (api *GSLBAPI) Delete(id string) (*sacloud.GSLB, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.delete(id, nil, res)
// 	})
// }

/************************************************
  Inner functions
************************************************/

func (api *GSLBAPI) setStateValue(setFunc func(*sacloud.Request)) *GSLBAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

//func (api *GSLBAPI) request(f func(*sacloud.Response) error) (*sacloud.GSLB, error) {
//	res := &sacloud.Response{}
//	err := f(res)
//	if err != nil {
//		return nil, err
//	}
//	return res.GSLB, nil
//}
//
//func (api *GSLBAPI) createRequest(value *sacloud.GSLB) *sacloud.Request {
//	req := &sacloud.Request{}
//	req.GSLB = value
//	return req
//}
