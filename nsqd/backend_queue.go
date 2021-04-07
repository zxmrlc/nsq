package nsqd

//diskQueue,dummyBackendQueue实现了backendQueue接口。
//dummyBackendQueue channel topic 都实现了这么一套接口
// BackendQueue represents the behavior for the secondary message
// storage system
type BackendQueue interface {
	Put([]byte) error
	ReadChan() chan []byte // this is expected to be an *unbuffered* channel
	Close() error
	Delete() error
	Depth() int64
	Empty() error
}
