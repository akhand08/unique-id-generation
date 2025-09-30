package main

import (
	"github.com/akhand08/unique-id-generation/protocols"
)

func main() {

	node := protocols.CreateNode()

	node.Run()

}
