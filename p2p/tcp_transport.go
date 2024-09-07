package p2p

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection.
// TCPPeer代表了一个TCP连接上的远程节点
type TCPPeer struct {
	// conn is the underlying connection of the peer.
	// conn是对等体的基础连接
	conn net.Conn

	// if we dial and retrieve a conn => outbound == true
	// if we accept and retrieve a conn => outbound == false
	// 如果我们拨号并检索到一个连接 => outbound == true
	// 如果我们接受并检索到一个连接 => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listenser     net.Listener
	handshakeFunc HandshakeFunc

	mu    sync.Mutex
	peers map[net.Addr]Peer
}

func NewTCpTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		handshakeFunc: NOPHandshakeFunc,
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listenser, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listenser.Accept()
		if err != nil {
			continue
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.handshakeFunc(conn); err != nil {

	}

	buf := new(bytes.Buffer)

	fmt.Printf("new incoming connection %v\n", peer)
}
