package nsqd

import (
	"bytes"
	"sync"
)

//真的是个buffer pool!
var bp sync.Pool

func init() {
	bp.New = func() interface{} {
		return &bytes.Buffer{}
	}
}

func bufferPoolGet() *bytes.Buffer {
	return bp.Get().(*bytes.Buffer)
}

func bufferPoolPut(b *bytes.Buffer) {
	bp.Put(b)
}
