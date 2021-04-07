package nsqd

/*see:https://nsq.io/overview/design.html#bounded-memory-footprint
Note, a topic/channel whose name ends in the string #ephemeral will not be buffered to disk and
will instead drop messages after passing the mem-queue-size.
This enables consumers which do not need message guarantees to subscribe to a channel.
*/
//dummyBackendQueue提供给 #ephemeral 后缀的topic/channel
//最大的特点是不会缓冲到磁盘
type dummyBackendQueue struct {
	readChan chan []byte
}

func newDummyBackendQueue() BackendQueue {
	return &dummyBackendQueue{readChan: make(chan []byte)}
}

func (d *dummyBackendQueue) Put([]byte) error {
	return nil
}

func (d *dummyBackendQueue) ReadChan() chan []byte {
	return d.readChan
}

func (d *dummyBackendQueue) Close() error {
	return nil
}

func (d *dummyBackendQueue) Delete() error {
	return nil
}

func (d *dummyBackendQueue) Depth() int64 {
	return int64(0)
}

func (d *dummyBackendQueue) Empty() error {
	return nil
}
