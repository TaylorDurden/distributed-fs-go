package main

import (
	"log"

	"github.com/TaylorDurden/distributed-fs-go/p2p"
)

func main() {
	options := p2p.TCPTransportOptions{ListenAddr: ":3000", Decoder: p2p.GOBDecoder{}, HandShakeFunc: p2p.NOPHandShakeFunc}
	tcpTransport := p2p.NewTCPTransport(options)

	if err := tcpTransport.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
