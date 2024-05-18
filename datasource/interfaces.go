package datasource

import "sync"

type IStream[T IData] interface {
	GetDataChan() chan T
	Process(wg *sync.WaitGroup)
}

type IData interface {
	Get() interface{}
}
