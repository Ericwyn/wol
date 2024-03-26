package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var version = "wol (Wake-On-Line)\nuse -h to get help\n\nversion: 0.0.1, https://github.com/Ericwyn/wol"

var macAddr = flag.String("m", "", "MAC address")
var broadcastAddr = flag.String("b", "255.255.255.255", "Broadcast address")
var versionFlag = flag.Bool("v", false, "Show version")

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	if *macAddr == "" {
		fmt.Println(version)
		os.Exit(-1)
	}

	// 解析MAC地址
	hwAddr, err := net.ParseMAC(*macAddr)
	if err != nil {
		fmt.Printf("Invalid MAC address: %s\n", err)
		os.Exit(1)
	}

	// 构建魔术包
	magicPacket := []byte{}
	// 6字节的0xFF
	header := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	magicPacket = append(magicPacket, header...)
	// 16次重复MAC地址
	for i := 0; i < 16; i++ {
		magicPacket = append(magicPacket, hwAddr...)
	}

	var broadcastIp = buildIpAddr(*broadcastAddr)

	// 发送魔术包
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   broadcastIp,
		Port: 9,
	})
	if err != nil {
		fmt.Printf("Failed to dial UDP: %s\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 写入数据到连接
	bytesWritten, err := conn.Write(magicPacket)
	if err != nil {
		fmt.Printf("Failed to send magic packet: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Magic packet sent (%d bytes), mac: %s, broadcast: %s \n",
		bytesWritten, *macAddr, *broadcastAddr)
}

// buildIpAddr builds a net.IP from a string
func buildIpAddr(ip string) net.IP {
	return net.ParseIP(ip)
}
