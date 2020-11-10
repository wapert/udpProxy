package proxy

import (
	"fmt"
	"sync"

	"github.com/rcrowley/go-metrics"
)

type worker struct {
	sync.Mutex
	p      *udpProxy
	group  *sync.WaitGroup
	closed bool

	liveConns  metrics.Counter
	totalConns metrics.Counter
}

func newWorker(p *udpProxy) *worker {
	w := &worker{
		p:     p,
		group: &sync.WaitGroup{},
	}

	w.registerCounters()

	return w
}

func (w *worker) registerCounters() {
	w.liveConns = metrics.NewCounter()
	w.totalConns = metrics.NewCounter()

	metrics.Register(fmt.Sprintf("%s_udp_live_connections", w.p.backend.Name), w.liveConns)
	metrics.Register(fmt.Sprintf("%s_udp_total_connections", w.p.backend.Name), w.totalConns)
}

func (w *worker) unregisterCounters() {
	metrics.Unregister(fmt.Sprintf("%s_udp_live_connections", w.p.backend.Name))
	metrics.Unregister(fmt.Sprintf("%s_udp_total_connections", w.p.backend.Name))
}

// work processes connections from the channel until it is closed
func (w *worker) work() {
	for conn := range w.p.connections {
		if err := w.handleConn(conn); err != nil {
			logger.WithField("error", err).Errorf("handle connection")
		}
	}
	w.closed = true

	w.unregisterCounters()

	// signal to the proxy that we are done processing all open connections
	w.p.group.Done()
}
