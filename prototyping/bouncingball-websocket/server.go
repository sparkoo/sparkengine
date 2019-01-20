package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
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
		//c.EnableWriteCompression(true)

		var balls = make([]*Ball, 1000)
		for bi := range balls {
			size := rand.Intn(15)
			balls[bi] = NewBall(FRAME_WIDTH/2, FRAME_HEIGHT/2,
				size, size,
				(rand.Float64()*10.0)-5.0, (rand.Float64()*10.0)-5.0)
			//fmt.Println(balls[bi])
		}

		for {
			// allocate new framebuffer for each frame to clean it
			// we can either choose to not clean it or to clean.
			// when clean, benchmark what is the most effective way!
			response := make([]byte, FRAME_BYTESIZE)

			w, err := c.NextWriter(websocket.BinaryMessage)
			if err != nil {
				log.Fatal(err)
			}
			for _, b := range balls {
				render(response, b)
			}
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

func render(res []byte, ball *Ball) {
	for _, p := range ball.getPixels() {
		x := ball.getXoffset() + p.x
		y := ball.getYoffset() + p.y
		i := (x + (FRAME_WIDTH * y)) * 4
		res[i] = p.color[0]
		res[i+1] = p.color[1]
		res[i+2] = p.color[2]
		res[i+3] = p.color[3]
	}

	xPot := int(ball.xpos + ball.xvel)
	if xPot < 0 || xPot+ball.xsize >= FRAME_WIDTH {
		ball.xvel *= -1
	}

	yPot := int(ball.ypos + ball.yvel)
	if yPot < 0 || yPot+ball.ysize >= FRAME_HEIGHT {
		ball.yvel *= -1
	}

	ball.xpos += ball.xvel
	ball.ypos += ball.yvel
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
