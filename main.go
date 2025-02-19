package main

import "github.com/fmelihh/decentralized-poker-game-go/p2p"

func main() {
	cfg := p2p.ServerConfig{
		ListenAddr: ":3000",
	}
	s := p2p.NewServer(cfg)
	s.Start()
}
