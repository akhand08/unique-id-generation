package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type RequestBody struct {
	Type      string `json:"type"`
	MessageID int    `json:"msg_id"`
	InReplyTo int    `json:"in_reply_to"`
}

type RequestNode struct {
	Source      string `json:"src"`
	Destination string `json:"dest"`
	Body        RequestBody
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	decoder := json.NewDecoder(reader)

	fmt.Fprintln(os.Stderr, "Server is started to running")

	for decoder.More() {

		var request RequestNode
		err := decoder.Decode(&request)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding JSON message: %v\n", err)
			continue
		}

		fmt.Println("Response: ")
		fmt.Println("-------------------------------------------------------")
		fmt.Println(request)
		fmt.Println("-------------------------------------------------------")

	}

}
