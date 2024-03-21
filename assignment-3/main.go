package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
    "math/rand"
    "time"
	"io/ioutil"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type Status struct {
    Water int `json:"water"`
    Wind  int `json:"wind"`
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/ws", wsHandler)
    go updateStatus()

    fmt.Println("Server is running on :8080")
    http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error upgrading to WebSocket:", err)
        return
    }
    defer conn.Close()

    // Loop forever, sending current status every second
    for {
        status := Status{
            Water: rand.Intn(100) + 1,
            Wind:  rand.Intn(100) + 1,
        }

        jsonStatus, err := json.Marshal(status)
        if err != nil {
            fmt.Println("Error marshalling JSON:", err)
            return
        }

        if err := conn.WriteMessage(websocket.TextMessage, jsonStatus); err != nil {
            fmt.Println("Error writing JSON to WebSocket:", err)
            return
        }

        time.Sleep(1 * time.Second)
    }
}

func updateStatus() {
    // Loop forever, updating status every 15 seconds
    for {
        status := Status{
            Water: rand.Intn(100) + 1,
            Wind:  rand.Intn(100) + 1,
        }

        jsonStatus, err := json.Marshal(status)
        if err != nil {
            fmt.Println("Error marshalling JSON:", err)
            return
        }

        if err := ioutil.WriteFile("status.json", jsonStatus, 0644); err != nil {
            fmt.Println("Error writing JSON to file:", err)
            return
        }

        time.Sleep(15 * time.Second)
    }
}