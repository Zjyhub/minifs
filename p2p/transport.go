package p2p

// peer is an interface that represents the remote node
// peer是一个代表远程节点的接口
type Peer interface {
}

// Transport is anything that handles the communication
// between nodes in the network.
// This can be of the form of a TCP connection, a UDP connection,websockets,...
// Transport是处理网络中节点之间通信的任何东西。
// 这可以是TCP连接、UDP连接、websockets等形式。
type Transport interface {
}
