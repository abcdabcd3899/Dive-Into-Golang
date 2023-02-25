package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Node struct {
	Name string
	Id   uint32
}

type ConsistentHash struct {
	Nodes       []Node
	NumReplicas int
}

func (c *ConsistentHash) AddNode(node Node) {
	c.Nodes = append(c.Nodes, node)
}

func (c *ConsistentHash) Build() {
	var nodes []Node
	for i := 0; i < c.NumReplicas; i++ {
		for _, node := range c.Nodes {
			nodes = append(nodes, Node{
				Name: node.Name,
				Id:   crc32.ChecksumIEEE([]byte(node.Name + string(rune(i)))),
			})
		}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Id < nodes[j].Id
	})
	c.Nodes = nodes
}

func (c *ConsistentHash) Get(key string) Node {
	if len(c.Nodes) == 0 {
		return Node{}
	}
	hash := crc32.ChecksumIEEE([]byte(key))
	i := sort.Search(len(c.Nodes), func(i int) bool {
		return c.Nodes[i].Id >= hash
	})
	if i == len(c.Nodes) {
		i = 0
	}
	return c.Nodes[i]
}

func (c *ConsistentHash) AffectedRange(node Node) (start uint32, end uint32) {
	var startIndex int
	var endIndex int

	for i, n := range c.Nodes {
		if n.Name == node.Name {
			startIndex = i - 1
			endIndex = i + 1
			break
		}
	}

	if startIndex == -1 {
		startIndex = len(c.Nodes) - 1
	}

	if endIndex == len(c.Nodes) {
		endIndex = 0
	}

	start = c.Nodes[startIndex].Id
	end = c.Nodes[endIndex].Id

	return
}

func (c *ConsistentHash) RemoveNode(node Node) {
	for i, n := range c.Nodes {
		if n.Name == node.Name {
			c.Nodes = append(c.Nodes[:i], c.Nodes[i+1:]...)
			break
		}
	}

	c.Build()
}

func main() {
	c := ConsistentHash{NumReplicas: 3}
	c.AddNode(Node{Name: "node1"})
	c.AddNode(Node{Name: "node2"})
	c.AddNode(Node{Name: "node3"})
	c.Build()

	// Test GetNode function
	node := c.Get("key1")
	if node.Name != "node1" {
		fmt.Println("GetNode failed, expected node1, got", node.Name)
	}

	node = c.Get("key2")
	if node.Name != "node2" {
		fmt.Println("GetNode failed, expected node2, got", node.Name)
	}

	node = c.Get("key3")
	if node.Name != "node3" {
		fmt.Println("GetNode failed, expected node3, got", node.Name)
	}

	// Test AffectedRange function
	start, end := c.AffectedRange(Node{Name: "node1"})
	if start != c.Nodes[2].Id || end != c.Nodes[1].Id {
		fmt.Printf("AffectedRange failed for node1, expected [%s, %s], got [%d, %d]\n",
			c.Nodes[2].Name, c.Nodes[1].Name, start, end)
	}

	start, end = c.AffectedRange(Node{Name: "node2"})
	if start != c.Nodes[0].Id || end != c.Nodes[2].Id {
		fmt.Printf("AffectedRange failed for node2, expected [%d, %d], got [%d, %d]\n",
			c.Nodes[0].Id, c.Nodes[2].Id, start, end)
	}

	start, end = c.AffectedRange(Node{Name: "node3"})
	if start != c.Nodes[1].Id || end != c.Nodes[0].Id {
		fmt.Printf("AffectedRange failed for node3, expected [%d, %d], got [%d, %d]\n",
			c.Nodes[1].Id, c.Nodes[0].Id, start, end)
	}

	// Test AddNode and RemoveNode functions
	c.AddNode(Node{Name: "node4"})
	node = c.Get("key4")
	if node.Name != "node4" {
		fmt.Println("GetNode failed after adding node4, expected node4, got", node.Name)
	}

	c.RemoveNode(Node{Name: "node4"})
	node = c.Get("key4")
	if node.Name == "node4" {
		fmt.Println("GetNode still returns node4 after removing node4")
	}
}
