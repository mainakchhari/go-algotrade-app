package datasource

import "sync"

type IStream[T IData] interface {
	IHasDataSink[T]
	Process(wg *sync.WaitGroup)
}

type IData interface {
	Get() interface{}
}

type IHasDataSource[T IData] interface {
	GetDataSource() chan T
}

type IHasDataSink[T IData] interface {
	GetDataSink() chan T
}
