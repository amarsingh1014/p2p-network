package main

import (
    "p2p-network/internal/peer"
)

func main() {
    // Create a new peer
    myPeer := peer.NewPeer("myPeer", "192.168.1.1:8080")

    // Initialize discovery
    discovery := peer.NewDiscovery()
    discovery.DiscoverPeers()

    // Add discovered peers to myPeer
    for _, discoveredPeer := range discovery.GetKnownPeers() {
        myPeer.AddPeer(discoveredPeer)
    }

    // Print known peers
    myPeer.PrintPeers()
}
