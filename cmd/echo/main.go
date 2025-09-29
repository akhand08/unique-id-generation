package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/akhand08/unique-id-generation/protocols"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	decoder := json.NewDecoder(reader)

	fmt.Fprintln(os.Stderr, "Server is started to running")

	for decoder.More() {

		var request protocols.Request
		err := decoder.Decode(&request)

		if request.Body.Type == "init" {
			initBody := protocols.InitResponseBody{Type: "init_ok", InReplyTo: request.Body.MessageID}
			initResponse := protocols.InitResponse{Source: request.Destination, Destination: request.Source, Body: initBody}

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
