package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	numbers := make([]uint64, 0)
	neighbors := make([]string, 0)

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		var response = make(map[string]any)

		response["type"] = "broadcast_ok"
		var number = uint64(body["message"].(float64))
		numbers = append(numbers, number)

		if err := replicate(n, neighbors, number); err != nil {
			return err
		}

		return n.Reply(msg, response)
	})

	n.Handle("replicate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		var number = uint64(body["message"].(float64))
		numbers = append(numbers, number)

		return nil
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		var err error
		if err, neighbors = ParseNeighbors(n.ID(), msg.Body); err != nil {
			return err
		}

		var response = make(map[string]any)
		response["type"] = "topology_ok"
		return n.Reply(msg, response)
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		var response = make(map[string]any)
		response["type"] = "read_ok"
		response["messages"] = numbers

		return n.Reply(msg, response)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

func replicate(node *maelstrom.Node, neighbors []string, number uint64) error {
	for _, neighbor := range neighbors {
		var payload = make(map[string]any)
		payload["type"] = "replicate"
		payload["message"] = number

		if err := node.Send(neighbor, payload); err != nil {
			return err
		}
	}

	return nil
}

func ParseNeighbors(self string, body_content []byte) (error, []string) {
	neighbors := make([]string, 0)

	// TODO: fix JSON topology parsing
	var body map[string]interface{}
	if err := json.Unmarshal(body_content, &body); err != nil {
		return err, nil
	}

	bs, _ := json.Marshal(body["topology"])
	var topology map[string][]string
	if err := json.Unmarshal(bs, &topology); err != nil {
		return err, nil
	}

	for node, node_peers := range topology {
		if node == self {
			continue;
		}

		for _, peer := range node_peers {
			if peer == self {
				continue
			}

			neighbors = append(neighbors, node)
		}
	}

	return nil, neighbors
}
