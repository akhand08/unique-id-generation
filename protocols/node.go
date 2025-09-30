package protocols

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

type Node struct {
	msg_counter int64
}

func CreateNode() *Node {
	return &Node{}
}

func (n *Node) Run() {

	reader := bufio.NewReader(os.Stdin)
	decoder := json.NewDecoder(reader)
	var wg sync.WaitGroup

	fmt.Fprintln(os.Stderr, "Server is started to running")

	for decoder.More() {

		var request Request
		err := decoder.Decode(&request)
		wg.Add(1)
		go request.RequestHandler(n, &wg)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding JSON message: %v\n", err)
			continue
		}

	}

	wg.Wait()
}
