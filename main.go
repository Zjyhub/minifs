package main

import (
	"log"
	"minifs/p2p"
)

func main() {
	tr := p2p.NewTCpTransport(":4000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
