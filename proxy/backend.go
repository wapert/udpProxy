package proxy

import "net"

type Backend struct {
	Name             string `json:"name,omitempty"`
	Proto            string `json:"proto,omitempty"`
	BindIP           net.IP `json:"bind_ip,omitempty"`
	BindPort         int    `json:"bind_port,omitempty"`
	IP               net.IP `json:"backend_ip,omitempty"`
	Port             int    `json:"backend_port,omitempty"`
	RemoteIP         net.IP `json:"remote_ip,omitempty"`
	MaxConcurrent    int    `json:"max_concurrent,omitempty"`
	ConnectionBuffer int    `json:"connection_buffer,omitempty"`
}
