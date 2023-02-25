package main

import (
	"hash/crc32"
	"sort"
	"sync"
)

type Node struct {
	Id string
	IP string
}

type ConsistentHash struct {
	nodes       map[uint32]Node
	NumReplicas int
	Ring        []uint32
	sync.RWMutex
}

func (c *ConsistentHash) nodeIdx(id string) int {
	for i, node := range c.Ring {
		if c.nodes[node].Id == id {
			return i
		}
	}
	return -1
}

func New(numReplicas int) *ConsistentHash {
	return &ConsistentHash{
		nodes:       make(map[uint32]Node),
		NumReplicas: numReplicas,
		Ring:        []uint32{},
	}
}

func (c *ConsistentHash) AddNode(node Node) {
	c.Lock()
	defer c.Unlock()

	for i := 0; i < c.NumReplicas; i++ {
		replica := []byte(node.IP + "-" + string(rune(i)))
		hash := crc32.ChecksumIEEE(replica)
		c.nodes[hash] = node
		c.Ring = append(c.Ring, hash)
	}

	sort.Slice(c.Ring, func(i, j int) bool {
		return c.Ring[i] < c.Ring[j]
	})
}

func (c *ConsistentHash) RemoveNode(node Node) {
	c.Lock()
	defer c.Unlock()

	nodeIdx := c.nodeIdx(node.Id)
	if nodeIdx == -1 {
		return
	}

	for i := 0; i < c.NumReplicas; i++ {
		replica := []byte(node.IP + "-" + string(rune(i)))
		hash := crc32.ChecksumIEEE(replica)

		for j, h := range c.Ring {
			if h == hash {
				c.Ring = append(c.Ring[:j], c.Ring[j+1:]...)
				delete(c.nodes, h)
				break
			}
		}
	}
}

func (c *ConsistentHash) GetNode(key string) Node {
	c.RLock()
	defer c.RUnlock()

	if len(c.Ring) == 0 {
		return Node{}
	}

	hash := crc32.ChecksumIEEE([]byte(key))
	i := sort.Search(len(c.Ring), func(i int) bool { return c.Ring[i] >= hash })

	if i == len(c.Ring) {
		i = 0
	}

	return c.nodes[c.Ring[i]]
}

func (c *ConsistentHash) AffectedRange(node Node) []Node {
	c.RLock()
	defer c.RUnlock()

	if len(c.Ring) == 0 {
		return []Node{}
	}

	affectedRange := []Node{}
	nodeIdx := c.nodeIdx(node.Id)
	if nodeIdx == -1 {
		return affectedRange
	}

	for i := nodeIdx - 1; i >= 0; i-- {
		affectedRange = append(affectedRange, c.nodes[c.Ring[i]])
	}

	for i := nodeIdx + 1; i < len(c.Ring); i++ {
		affectedRange = append(affectedRange, c.nodes[c.Ring[i]])
	}

	return affectedRange
}

func main() {
	c := New(3)

	node1 := Node{Id: "node1", IP: "192.168.1.1"}
	node2 := Node{Id: "node2", IP: "192.168.1.2"}
	node3 := Node{Id: "node3", IP: "192.168.1.3"}

	c.AddNode(node1)
	c.AddNode(node2)
	c.AddNode(node3)

	key := "mykey"
	node := c.GetNode(key)
	println("Node for key: ", key, " is: ", node.IP)

	affectedRange := c.AffectedRange(node)
	println("Affected range for node ", node.IP, " is: ")
	for _, n := range affectedRange {
		println(n.IP)
	}

	nodeToRemove := Node{Id: "node1", IP: "192.168.1.1"}
	c.RemoveNode(nodeToRemove)
}
