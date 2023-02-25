package main

import (
	"testing"
)

func TestConsistentHash(t *testing.T) {
	c := New(3)

	node1 := Node{Id: "node1", IP: "192.168.1.1"}
	node2 := Node{Id: "node2", IP: "192.168.1.2"}
	node3 := Node{Id: "node3", IP: "192.168.1.3"}

	c.AddNode(node1)
	c.AddNode(node2)
	c.AddNode(node3)

	// Test GetNode
	key := "myKey"
	node := c.GetNode(key)
	if node.IP != node3.IP {
		t.Errorf("Expected node IP to be %s, but got %s", node1.IP, node.IP)
	}

	// Test AffectedRange
	affectedRange := c.AffectedRange(node)
	if len(affectedRange) != 2 {
		t.Errorf("Expected affected range length to be 2, but got %d", len(affectedRange))
	}
	if affectedRange[0].IP != node3.IP {
		t.Errorf("Expected affected range node IP to be %s, but got %s", node3.IP, affectedRange[0].IP)
	}
	if affectedRange[1].IP != node2.IP {
		t.Errorf("Expected affected range node IP to be %s, but got %s", node2.IP, affectedRange[1].IP)
	}

	// Test RemoveNode
	c.RemoveNode(node1)
	key = "myKey"
	node = c.GetNode(key)
	if node.IP != node2.IP {
		t.Errorf("Expected node IP to be %s, but got %s", node2.IP, node.IP)
	}

	// Test AffectedRange after RemoveNode
	affectedRange = c.AffectedRange(node)
	if len(affectedRange) != 1 {
		t.Errorf("Expected affected range length to be 1, but got %d", len(affectedRange))
	}
	if affectedRange[0].IP != node3.IP {
		t.Errorf("Expected affected range node IP to be %s, but got %s", node3.IP, affectedRange[0].IP)
	}
}
