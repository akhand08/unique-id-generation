package protocols

import (
	"encoding/json"
	"os"
	"sync"
	"sync/atomic"

	"github.com/akhand08/unique-id-generation/internals/echo"
)

type RequestBody struct {
	Type      string   `json:"type"`
	MessageID int      `json:"msg_id"`
	Echo      string   `json:"echo"`
	NodeID    string   `json:"node_id"`
	NodeIDs   []string `json:"node_ids"`
}

type InitResponseBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
}
type InitResponse struct {
	Source      string           `json:"src"`
	Destination string           `json:"dest"`
	Body        InitResponseBody `json:"body"`
}

type Request struct {
	Source      string      `json:"src"`
	Destination string      `json:"dest"`
	Body        RequestBody `json:"body"`
}

func (req *Request) RequestHandler(node *Node, wg *sync.WaitGroup) {

	defer wg.Done()

	if req.Body.Type == "init" {

		req.InitHanlder()

	} else if req.Body.Type == "echo" {

		req.EchoRequestPrcocessor(node)

	}

}

func (req *Request) InitHanlder() {

	initBody := InitResponseBody{Type: "init_ok", InReplyTo: req.Body.MessageID}
	initResponse := InitResponse{Source: req.Destination, Destination: req.Source, Body: initBody}

	json.NewEncoder(os.Stdout).Encode(initResponse)

}

func (req *Request) EchoRequestPrcocessor(node *Node) {

	atomic.AddInt64(&node.msg_counter, 1)

	responseBody := echo.ResponseBody{Type: "echo_ok", MessageID: int(node.msg_counter), InReplyTo: req.Body.MessageID, Echo: req.Body.Echo}
	response := echo.Response{Source: req.Destination, Destination: req.Source, Body: responseBody}

	json.NewEncoder(os.Stdout).Encode(response)

}
