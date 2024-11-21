package peer

import (
	"fmt"
)

type Discovery struct {
	knownPeers map[string]Peer
}

func NewDiscovery() *Discovery {
	return &Discovery{
		knownPeers: make(map[string]Peer),
	}
}

func (d *Discovery) DiscoverPeers() {
	peers := []Peer{
		*NewPeer("1", "localhost:8080"),
		*NewPeer("2", "localhost:8081"),
	}

	for _, peer := range peers {
		d.knownPeers[peer.ID] = peer
		fmt.Printf("Discovered peer %s at %s\n", peer.ID, peer.Address)
	}
}

func (d *Discovery) GetKnownPeers() []Peer {
	peers := []Peer{}
	for _, peer := range d.knownPeers {
		peers = append(peers, peer)
	}
	return peers
}