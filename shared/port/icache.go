package port

// import (
// 	"manager/shared/connection/managercache"
// 	"manager/shared/types"
// )

// type ICache interface {
// 	Collection(name string) (*managercache.Collection, error)
// 	DefaultCollection() (*managercache.Collection, error)
// }

// type ICacheCollection interface {
// 	Delete(key string) error
// 	Get(key string) ([]byte, error)
// 	GetSubset(key string, offset, limit int, items ...*types.Pair) ([]byte, *types.CacheSubsetInfo, error)
// 	Set(key string, entry []byte) error
// }
