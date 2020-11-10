package proxy

import (
	"io"
	"net"
	"sync"
	"syscall"
	"time"
)

func checkreport(level int, err error) bool {
	if err == nil {
		return false
	}
	logger.Infof("Error %s", err)
	return true
}

func checkTimeOut(group *sync.WaitGroup, p *udpProxy) {
	for {
		time.Sleep(500 * time.Millisecond)
		if p.backend.Port == 500 || p.backend.BindPort == 500 {
			if time.Since(p.t0).Seconds() > 30 {
				if p.srcConn != nil {
					p.srcConn.Close()
					p.srcConn = nil
				}
				if p.dstConn != nil {
					p.dstConn.Close()
					p.dstConn = nil
				}
				logger.Infof("500 Port timeout!!!")
				break
			}
		} else {
			break
		}
	}
	group.Done()
	return
}

// transfer bytes from one udp connection to another
func transfer(from, to *udpConn, group *sync.WaitGroup, p *udpProxy) {

	for {
		_, err := io.Copy(to, from)
		if err != nil {
			if err, ok := err.(*net.OpError); ok && err.Err == syscall.EPIPE {
				logger.Errorf("Unexpected network error type %s", err)
			}
			break
		}
	}
	group.Done()
	return
}

// udpConn is used to handle normal communication
type udpConn struct {
	readCon    *net.UDPConn
	closeCon   *net.UDPConn
	remoteAddr net.Addr
}

func (t *udpConn) Read(b []byte) (int, error) {
	//logger.Infof("starting udp read")
	return t.readCon.Read(b)
}

func (t *udpConn) Write(b []byte) (int, error) {
	return t.readCon.Write(b)
}

func (t *udpConn) LocalAddr() net.Addr {
	return t.readCon.LocalAddr()
}

func (t *udpConn) RemoteAddr() net.Addr {
	return t.readCon.RemoteAddr()
}

func (t *udpConn) SetDeadline(tm time.Time) error {
	return t.readCon.SetDeadline(tm)
}

func (t *udpConn) SetReadDeadline(tm time.Time) error {
	return t.readCon.SetReadDeadline(tm)
}

func (t *udpConn) SetWriteDeadline(tm time.Time) error {
	return t.readCon.SetWriteDeadline(tm)
}

func (t *udpConn) CloseRead() error {
	logger.Infof("starting udp closeread")
	return t.closeCon.Close()
}

func (t *udpConn) CloseWrite() error {
	logger.Infof("starting udp closewrite")
	return t.closeCon.Close()
}

func (t *udpConn) Close() error {
	logger.Infof("starting udp close")
	return t.closeCon.Close()
}
