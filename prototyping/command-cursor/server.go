package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	FRAME_WIDTH    = 640
	FRAME_HEIGHT   = 480
	FRAME_BYTESIZE = FRAME_WIDTH * FRAME_HEIGHT * 4
)

var (
	addr     = flag.String("addr", ":8080", "http service address")
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: FRAME_BYTESIZE,
		CheckOrigin: func(r *http.Request) bool {
			fmt.Println(r)
			return true
		},
	}
)

var cursor = NewCursor(FRAME_WIDTH/2, FRAME_HEIGHT/2)

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
		//c.EnableWriteCompression(true)

		go read(c)

		for {
			// allocate new framebuffer for each frame to clean it
			// we can either choose to not clean it or to clean.
			// when clean, benchmark what is the most effective way!
			response := make([]byte, FRAME_BYTESIZE)

			w, err := c.NextWriter(websocket.BinaryMessage)
			if err != nil {
				log.Fatal(err)
			}

			render(response, cursor)

			if n, err := w.Write(response); n != FRAME_BYTESIZE || err != nil {
				log.Printf("written [%v], should write [%v]", n, FRAME_BYTESIZE)
				log.Fatal(err)
			}
			if err = w.Close(); err != nil {
				log.Fatal(err)
			}
			//fmt.Println(response)
			time.Sleep(15 * time.Millisecond)
		}
	}
}

func read(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		coords := strings.Split(string(message), ",")
		x, y := coords[0], coords[1]
		fx, err := strconv.ParseFloat(x, 64)
		if err == nil {
			cursor.xpos += fx
		} else {
			log.Fatal(err)
		}
		fy, err := strconv.ParseFloat(y, 64)
		if err == nil {
			cursor.ypos += fy
		} else {
			log.Fatal(err)
		}
	}
}

func render(res []byte, c *Cursor) {
	for _, p := range c.getPixels() {
		x := c.getXoffset() + p.x
		y := c.getYoffset() + p.y
		i := (x + (FRAME_WIDTH * y)) * 4
		res[i] = p.color[0]
		res[i+1] = p.color[1]
		res[i+2] = p.color[2]
		res[i+3] = p.color[3]
	}
}

func main() {
	log.Println("start")
	rand.Seed(time.Now().UnixNano())

	flag.Parse()

	http.HandleFunc("/ws", echo)
	http.HandleFunc("/webgl", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "webgl.html")
	})
	http.HandleFunc("/2d", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "2d.html")
	})
	http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "script.js")
	})
	log.Println("start listening on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
