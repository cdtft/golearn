package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	mutex            sync.Mutex
	resourcesChannel chan io.Closer
	factory          func() (io.Closer, error)
	closed           bool
}

var ErrPoolClosed = errors.New("pool has been closed")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("poll size is too small")
	}
	return &Pool{
		factory:   fn,
		resourcesChannel: make(chan io.Closer, size),
	}, nil
}

//获取资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case resource, ok := <-p.resourcesChannel:
		if !ok {
			return nil, ErrPoolClosed
		}
		return resource, nil
	default:
		log.Println("Acquire:", "New resources")
		return p.factory()
	}
}

//释放资源, 保证线程安全
func (p *Pool) Release(resource io.Closer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	//池是否关闭
	if p.closed {
		resource.Close()
		return
	}

	select {
	case p.resourcesChannel <- resource:
		log.Println("release", "In Queue")
	default:
		//缓冲队列已经满了
		resource.Close()
	}
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return
	}
	//关闭通道
	p.closed = true
	close(p.resourcesChannel)
	for r := range p.resourcesChannel {
		r.Close()
	}
}
