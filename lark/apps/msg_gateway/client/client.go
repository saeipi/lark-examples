// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:7301", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	headers := map[string][]string{}
	token := "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODEyOTY3NTMsImlhdCI6MTY4MDY5MTk1MywiaXNzIjoibGFyay5jb20iLCJwbGF0Zm9ybSI6MSwic2Vzc2lvbl9pZCI6IjUxM2JjMjFiZTNiZjE2MGM5YzE1YWRkY2U2YWQ0ZWNjIiwidWlkIjoiMTY0MzU2NzI4OTg2MzI0NTgyNCJ9.cwMB4KRHHDDCEz0UlKSyTlHVp1g0Jrjkbqc2g1XPkuk"
	headers["Cookie"] = []string{token}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println(t)
			msg := "请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我我请我bbb"
			err := c.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
