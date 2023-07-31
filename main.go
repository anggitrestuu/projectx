package main

import (
	"time"

	"github.com/anggitrestuu/projectx/network"
)

// server
// transport => tcp,udp
// Block
// TX
// Keypair

func main() {

	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {

		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello World"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transport: []network.Transport{
			trLocal,
		},
	}

	s := network.NewServer(opts)
	s.Start()

}
