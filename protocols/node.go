package protocols

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Node struct {
}

func CreateNode() *Node {
	return &Node{}
}

func (n *Node) Run() {

	reader := bufio.NewReader(os.Stdin)
	decoder := json.NewDecoder(reader)

	fmt.Fprintln(os.Stderr, "Server is started to running")

	for decoder.More() {

		var request Request
		err := decoder.Decode(&request)

		if request.Body.Type == "init" {
			initBody := InitResponseBody{Type: "init_ok", InReplyTo: request.Body.MessageID}
			initResponse := InitResponse{Source: request.Destination, Destination: request.Source, Body: initBody}

			json.NewEncoder(os.Stdout).Encode(initResponse)
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding JSON message: %v\n", err)
			continue
		}

	}
}
