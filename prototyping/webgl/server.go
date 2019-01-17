package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	timeout = 1 * time.Second
)

var (
	addr     = flag.String("addr", ":8080", "http service address")
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s, %v", message, mt)
		c.EnableWriteCompression(true)

		t1 := time.Now()
		for i := 0; i < 320 * 10 * 4; i++ {
			t2 := time.Now()
			diff := t2.Sub(t1).Nanoseconds() / 1000 / 1000
			t1 = t2
			log.Println(i, "time: ", diff, "ms")
			w, _ := c.NextWriter(websocket.BinaryMessage)
			rrr(w, mt)
			w.Close()
			time.Sleep(1 * time.Millisecond)
		}
	}

}

var lll = 0

func rrr(w io.WriteCloser, mt int) {
	size := 320 * 200 * 4
	response := make([]byte, size)
	for i := 0; i < lll*4; i += 4 {
		response[i] = 127
		response[i+1] = 127
		response[i+2] = 127
		response[i+3] = 255
	}
	lll += 4
	//err := conn.WriteMessage(websocket.BinaryMessage, response)
	//if err != nil {
	//	log.Println("write:", err)
	//}
	if _, err := w.Write(response); err != nil {
		log.Fatal(err)
	}

}

func main() {
	log.Println("start")

	flag.Parse()

	http.HandleFunc("/ws", echo)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/pixi.min.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pixi.min.js")
	})
	http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "script.js")
	})
	log.Println("start listening on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
