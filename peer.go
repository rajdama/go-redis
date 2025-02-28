package main

import (
	"fmt"
	"log/slog"
	"net"
)

type Peer struct {
	conn  net.Conn
	msgch chan []byte
}

func NewPeer(conn net.Conn, msgch chan []byte) *Peer {
	return &Peer{
		conn:  conn,
		msgch: msgch,
	}
}

func (p *Peer) readLoop() error {
	buf := make([]byte, 1024)

	for {

		n, err := p.conn.Read(buf)
		if err != nil {
			slog.Error("peer read error", "err")
			return err
		}
		fmt.Println(string(buf[:n]))
		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])
		p.msgch <- msgBuf

	}
}
