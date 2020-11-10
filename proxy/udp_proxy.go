package proxy

import (
	"container/list"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	reuse "github.com/libp2p/go-reuseport"
)

var vpnTable *list.List

type udpProxy struct {
	backend     *Backend
	srcConn     *net.UDPConn
	dstConn     *net.UDPConn
	connections chan *net.UDPConn
	group       *sync.WaitGroup
	started     bool
	t0          time.Time
}

type proxyPair struct {
	remoteAddr  string //remote EM address with IP and Port
	backendAddr string // container address with IP only
}

// InitVpnTable
func InitVpnTable() {
	vpnTable = list.New()
}

// AddVpnEntry
func AddVpnEntry(pair proxyPair) {
	vpnTable.PushBack(pair)
}

// CheckVpnEntry
func CheckVpnEntry(pair proxyPair) int {
	var next *list.Element
	var result int
	result = 0
	for e := vpnTable.Front(); e != nil; e = next {
		next = e.Next()
		vPair := e.Value.(proxyPair)
		if pair.remoteAddr == vPair.remoteAddr {
			if pair.backendAddr == vPair.backendAddr {
				result = 1
				break
			} else {
				result = 2
				break
			}
		}
	}
	return result
}

// RemoveVpnEntry
func RemoveVpnEntry(pair proxyPair) {
	var next *list.Element
	for e := vpnTable.Front(); e != nil; e = next {
		next = e.Next()
		vPair := e.Value.(proxyPair)
		if pair == vPair {
			vpnTable.Remove(e)
			break
		}
	}
}

// RemoveAllVpnEntry
func RemoveAllVpnEntry() {
	var next *list.Element
	for e := vpnTable.Front(); e != nil; e = next {
		next = e.Next()
		vpnTable.Remove(e)
	}
}

func newUDPPRoxy(backend *Backend) (*udpProxy, error) {
	if backend.ConnectionBuffer == 0 {
		backend.ConnectionBuffer = 100
	}
	return &udpProxy{
		backend:     backend,
		connections: make(chan *net.UDPConn, backend.ConnectionBuffer),
		group:       &sync.WaitGroup{},
	}, nil
}

func (p *udpProxy) Close() error {
	//
	logger.Infof("Closing proxy")

	if p.connections != nil {
		close(p.connections)
		p.connections = nil
	}
	if p.dstConn != nil {
		p.dstConn.Close()
		p.dstConn = nil
	}
	if p.srcConn != nil {
		p.srcConn.Close()
		p.srcConn = nil
	}
	p.started = false
	//p.group.Wait()

	logger.Infof("Closing proxy done!!")
	return nil //err
}

func (p *udpProxy) Backend() *Backend {
	return p.backend
}

func (p *udpProxy) Start() (err error) {
	if p.started {
		return fmt.Errorf("proxy has already been started")
	}
	p.started = true
	p.srcConn = nil
	p.dstConn = nil

	for i := 0; i < p.backend.MaxConcurrent; i++ {
		logger.Infof("starting worker %d", i)

		bindAddress := p.backend.BindIP.String() + ":" + strconv.Itoa(p.backend.BindPort)
		conn, err := reuse.ListenPacket("udp", bindAddress)
		if err != nil {
			logger.Errorf(" Listen error type %s", err)
			return err
		}
		p.srcConn = conn.(*net.UDPConn)

		p.group.Add(1)

		worker := newWorker(p)
		p.connections <- p.srcConn
		//conn := p.srcConn
		go worker.work()
	}

	logger.Infof("start proxy")
	return nil
}
