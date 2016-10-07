package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/wuxc/gowebwx/webwx"
)

var upgrader = websocket.Upgrader{}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("connect:", err)
		return
	}
	log.Println("websocket connected.")
	defer ws.Close()
	for {
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Println("recieved:", mt, msg)
		err = ws.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	path := "res/index.html"
	if r.URL.Path != "/" {
		path = "res" + r.URL.Path
	}
	log.Println("serveStatic", path)
	http.ServeFile(w, r, path)
}

func main() {
	t := time.Now()
	var it int = int(t.UnixNano() / 1000)
	fmt.Println(^it + 1)
	fmt.Println(t, t.Unix())
	fmt.Println(strconv.FormatInt(time.Now().UnixNano(), 10))

	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/", serveStatic)

	w := webwx.New()
	w.Start()

	// handle Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Println("Ctrl+C received.")
			w.Stop()
			os.Exit(0)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
