package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	localPort := 8080
	lAddr, _ := net.ResolveUDPAddr("udp",fmt.Sprintf("0.0.0.0:%d",localPort))
	resp, err := runStunDiscovery(lAddr)
	if err != nil {
		fmt.Println("failed to fetch ip and port.")
		return
	}
	ip, port := decodeStunMessage(resp)
	if ip == "" && port == 0 {
		fmt.Println("failed to decode stun response")
		return
	}
	fmt.Printf("ip: %s\nport: %d\n",ip,port)
	var pip string
	var pport int
	fmt.Print("peer ip:port > ")
	fmt.Scanf("%s %d",&pip,&pport)
	fmt.Printf("ip: %s\nport: %d\n",pip,pport)
	con, _ := net.ListenUDP("udp",lAddr)
	pAddr, _ := net.ResolveUDPAddr("udp",fmt.Sprintf("%s:%d",pip,pport))

	punchHole(con,pAddr)

	for {
		var msg string
		fmt.Scanln(&msg)
		if strings.ToLower(msg)=="exit" {
			break
		}
		con.WriteToUDP([]byte(msg),pAddr)
		buf := make([]byte,128)
		for {	
			con.SetReadDeadline(time.Now().Add(time.Second/2))
			n, _, err := con.ReadFromUDP(buf)
			if err != nil {
				break
			}
			fmt.Println("-> "+string(buf[:n]))
		}
	}
}
