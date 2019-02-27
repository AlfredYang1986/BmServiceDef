package BmFactory

type BmFactoryInterface interface {
	GetModelByName(name string) interface{}
	GetResourceByName(name string) interface{}
	GetStorageByName(name string) interface{}
	GetDaemonByName(name string) interface{}
	GetFunctionByName(name string) interface{}
	GetMiddlewareByName(name string) interface{}
}
