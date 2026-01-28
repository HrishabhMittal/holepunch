package main

import (
	"net"
	"time"
)
func punchHole(con *net.UDPConn,addr *net.UDPAddr) {
	buf := make([]byte, 128)
	for {
		con.SetReadDeadline(time.Now().Add(time.Second/2))
		_, _, err := con.ReadFromUDP(buf)
		if err != nil {
			return
		}
		con.WriteToUDP([]byte("hello"),addr)
	}
}

