package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	uuid "github.com/google/uuid"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		response := make(map[string]any)

		response["type"] = "generate_ok"
		// We need all of our nodes to be highly available regardless of cluster
		// presence, so since we can't establish strong consistency and consensus
		// without denying availability, UUIDs are a great way of guaranteeing
		// globally unique identifiers, as the chance of generating a duplicate,
		// even across multiple nodes are astronomically low.
		response["id"] = uuid.New().String()

		return n.Reply(msg, response)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
