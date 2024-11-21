package peer

import (
	"fmt"
	"net"
)

type Peer struct {
	ID string
	Address string
	Conn net.Conn
	Peers []Peer
}

func NewPeer(id, address string) *Peer {
	return &Peer{
		ID: id,
		Address: address,
		Peers: []Peer{},
	}
}

func (p *Peer) AddPeer(newPeer Peer) {
	p.Peers = append(p.Peers, newPeer)
	fmt.Printf("Peer %s added peer %s\n", p.ID, newPeer.ID)
}

func (p *Peer) PrintPeers() {
	fmt.Println("Known Peers:")
	for _, peer := range p.Peers {
		fmt.Printf("Peer %s at %s\n", peer.ID, peer.Address)
	}
}

