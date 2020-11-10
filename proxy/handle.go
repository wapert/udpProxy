package proxy

import (
	"fmt"
	"net"
	"strconv"
	"time"

	reuse "github.com/libp2p/go-reuseport"
)

func (w *worker) handleConn(rawConn *net.UDPConn) error {

	if rawConn == nil {
		return nil
	}
	conn := rawConn
	start := time.Now()
	var buffer [10]byte
	var remoteAddr *net.UDPAddr
	var vPair proxyPair

	backendAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", w.p.backend.IP, w.p.backend.Port))
	if err != nil {
		logger.Errorf(" Resolve error type %s", err)
	} else {
		backendAddr = &net.UDPAddr{
			IP:   w.p.backend.IP,
			Port: w.p.backend.Port}
	}
	logger.Infof("%s waiting for remote packet ...", backendAddr.String())

	// check remote source address & target backend address pair
	for {
		_, clientAddr, err := conn.ReadFromUDP(buffer[0:])
		if checkreport(1, err) {
			if w.p.srcConn != nil {
				time.Sleep(50 * time.Millisecond)
				continue
			} else {
				return nil
			}
		} else {
			vPair.remoteAddr = clientAddr.String()
			vPair.backendAddr = w.p.backend.IP.String()
			res := CheckVpnEntry(vPair)
			if res < 2 {
				if res < 1 {
					AddVpnEntry(vPair)
				}
				logger.Infof("%s Received remote packet from %s!!!", backendAddr.String(), clientAddr.String())
				remoteAddr = clientAddr
				break

			} else {
				logger.Infof("%s Received CONFLICT remote packet from %s!!!", backendAddr.String(), clientAddr.String())
				udpRequestTimer.UpdateSince(start)
				time.Sleep(50 * time.Millisecond)
				continue
			}
		}
	}

	w.p.t0 = time.Now()

	udpLiveConnections.Inc(1)
	w.liveConns.Inc(1)
	w.totalConns.Inc(1)

	/* check source IP of UDP stream */

	remoteIP := remoteAddr.String()
	fmt.Printf("stream from remoteIP %s\n", remoteIP)
	if w.p.backend.RemoteIP != nil {
		if remoteIP != w.p.backend.RemoteIP.String() {
			fmt.Printf("%s not listed in remoteIP \n", remoteIP)
			return nil
		}
	}
	w.group.Add(3)

	logger.Infof("Dial to backend %s ", backendAddr.String())
	dst, err := reuse.Dial("udp", "", backendAddr.String())
	if err != nil {
		logger.Errorf(" Dial backend error type %s", err)
		return err
	}
	dst.Write(buffer[0:])

	logger.Infof("Dial to remote %s ", remoteAddr.String())
	src, err := reuse.Dial("udp", "0.0.0.0:"+strconv.Itoa(w.p.backend.Port), remoteAddr.String())
	if err != nil {
		logger.Errorf(" Dial remote error type %s", err)
		return err
	}

	//close listen socket and create dial socket
	conn.Close()
	w.p.connections <- nil
	w.p.srcConn = src.(*net.UDPConn)
	w.p.dstConn = dst.(*net.UDPConn)
	//w.p.connections <- w.p.srcConn

	c2 := &udpConn{
		readCon:    src.(*net.UDPConn),
		closeCon:   src.(*net.UDPConn),
		remoteAddr: nil, //remoteAddr,
	}

	d2 := &udpConn{
		readCon:    dst.(*net.UDPConn),
		closeCon:   dst.(*net.UDPConn),
		remoteAddr: nil,
	}

	defer func() {
		udpLiveConnections.Dec(1)
		w.liveConns.Dec(1)
		logger.Infof("Handler Done!!!")
	}()

	go transfer(c2, d2, w.group, w.p)
	go transfer(d2, c2, w.group, w.p)
	go checkTimeOut(w.group, w.p)

	w.group.Wait()

	logger.Infof("Ending udp transfer")

	RemoveVpnEntry(vPair)
	udpRequestTimer.UpdateSince(start)

	return nil
}
