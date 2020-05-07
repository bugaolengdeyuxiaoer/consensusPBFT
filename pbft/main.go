package main

import (
	"github.com/bugaolengdeyuxiaoer/consensusPBFT/pbft/network"
	"os"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID)

	server.Start()
}
