package main

import (
	"fmt"
	"net"
	"time"
)
func punchHole(con *net.UDPConn,addr *net.UDPAddr) (bool) {
	buf := make([]byte, 128)
	// for {
		con.WriteToUDP([]byte("hello"),addr)
		fmt.Println("send hello to addr")
		con.SetReadDeadline(time.Now().Add(time.Second/2))
		n, _, err := con.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ERROR:")
			fmt.Println(err)
			return false
		}
		fmt.Printf("read %d bytes\n",n)
		return true
	// }
}

