package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"net"
)

const (
	STUN_SERVER = "stun.l.google.com:19302"
	MAGIC_COOKIE = 0x2112A442
)

func decodeStunMessage(data []byte) (string, int) {
	if len(data) < 20 {
		return "", 0
	}
	pointer := 20
	for pointer < len(data) {
		attrType := binary.BigEndian.Uint16(data[pointer:pointer+2])
		attrLen := int(binary.BigEndian.Uint16(data[pointer+2:pointer+4]))
	
		if attrType == 0x0020 {
			attrData := data[pointer+4:pointer+4+attrLen]			
			rawPort := binary.BigEndian.Uint16(attrData[2:4])
			finalPort := rawPort^(MAGIC_COOKIE>>16)

			xIP := attrData[4:8]
			cookieBytes := make([]byte,4)
			binary.BigEndian.PutUint32(cookieBytes,uint32(MAGIC_COOKIE))
			finalIP := make([]byte,4)
			for i := range 4 {
				finalIP[i] = xIP[i]^cookieBytes[i]
			}
			return net.IP(finalIP).String(), int(finalPort)
		}
		pointer+=4+((attrLen+3)& ^3)
	}
	return "", 0
}
func runStunDiscovery(localAddr *net.UDPAddr) ([]byte,error) {
	con, err := net.ListenUDP("udp",localAddr)
	if err!=nil {
		return nil,err
	}
	defer con.Close()

	req := make([]byte,12)
	rand.Read(req)
	packet := new(bytes.Buffer)
	binary.Write(packet,binary.BigEndian,uint16(0x0001))
	binary.Write(packet,binary.BigEndian,uint16(0))
	binary.Write(packet,binary.BigEndian,uint32(MAGIC_COOKIE))
	packet.Write(req)

	serverAddr, err := net.ResolveUDPAddr("udp",STUN_SERVER)
	if err != nil {
		return nil, err
	}
	_, err = con.WriteToUDP(packet.Bytes(),serverAddr)
	if err != nil {
		return nil, err
	}
	buf := make([]byte,1024)
	n, _, err := con.ReadFromUDP(buf)
	return buf[:n], err
}
