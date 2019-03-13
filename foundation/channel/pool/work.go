package pool

import "sync"

type Worker interface {
	Task()
}

type WorkPool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func NewWorkPool(maxGoroutines int) *WorkPool {
	p := &WorkPool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for work := range p.work {
				work.Task()
			}
			p.wg.Done()
		}()
	}
	return p
}

func (p *WorkPool) WorkPoolRun(w Worker) {
	p.work <-w
}

func (p *WorkPool) ShutDown() {
	close(p.work)
	p.wg.Wait()
}