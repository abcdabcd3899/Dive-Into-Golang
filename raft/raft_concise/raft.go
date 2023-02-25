/*
The code is a simple implementation of the Raft consensus algorithm in Go, but there are a few ways it can be improved to better resemble the actual Raft algorithm:

1. Add log replication: The current implementation only appends new entries to the local log, but in a real-world implementation of Raft, the leader should replicate the log to all the followers.

2. Implement leader election timeout: Currently, there is no mechanism to timeout a leader election if a leader is not elected. In a real-world implementation of Raft, each node should have a randomized election timeout and reset it upon receiving a message from a leader.

3. Implement heartbeat timeout: Currently, there is no mechanism for followers to detect a failed leader. In a real-world implementation of Raft, the leader should send periodic heartbeat messages to its followers, and followers should timeout the leader if no message is received.

4. Add a mechanism to stop the program: The current implementation will run indefinitely. It can be improved by adding a mechanism to stop the infinite loop in the main function.

5. Add mechanism for handling leader failure: The current implementation does not handle the case where the leader fails. A real-world implementation of raft should have a mechanism to detect the failure of a leader and start a new election.

6. Add mechanism for handling term changes: The current implementation does not handle the case where a node's term is out of sync with the leader's term. A real-world implementation of raft should have a mechanism for nodes to update their term when a leader with a higher term is elected.

7. Handle the case where a candidate receives vote from a node with a higher term, the candidate should step down and update its term.

8. Add mechanism for handling conflicting logs: The current implementation does not handle the case where a node's log is out of sync with the leader's log. A real-world implementation of Raft should have a mechanism for nodes to resolve conflicting logs.

9. Implement snapshotting mechanism: The current implementation does not handle the case where the log is getting too large. A real-world implementation of Raft should have a mechanism for taking snapshots of the state machine and discarding old logs.

10. Add mechanism for handling node failures: The current implementation does not handle the case where a node fails. A real-world implementation of Raft should have a mechanism for detecting and handling node failures.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	FOLLOWER = iota
	CANDIDATE
	LEADER
)

type StateMachine struct {
	// term  int
	// state int
	log []string
}

type RaftNode struct {
	id                   int
	stateMachine         *StateMachine
	currentTerm          int
	votedFor             int
	state                string
	leaderId             int
	peers                []*RaftNode
	electionTimeout      int
	heartbeatTimeout     int
	voteChan             chan bool
	appendChan           chan string
	stopChan             chan bool
	electionTimer        *time.Timer
	heartbeatChan        chan bool
	leaderFailureTimeout int
}

func (r *RaftNode) start() {
	r.electionTimer = time.NewTimer(5 * time.Second)
	for {
		select {
		case <-r.stopChan:
			return
		case <-r.electionTimer.C:
			r.state = "CANDIDATE"
			r.electionTimer.Reset(5 * time.Second)
		default:
			switch r.state {
			case "FOLLOWER":
				r.runFollower()
			case "CANDIDATE":
				r.runCandidate()
			case "LEADER":
				r.runLeader()
			}
		}
		if <-r.stopChan {
			break
		}
	}
}

func (r *RaftNode) runFollower() {
	heartbeatChan := make(chan bool)
	select {
	case <-time.After(time.Duration(r.electionTimeout) * time.Second):
		r.state = "CANDIDATE"
	case vote := <-r.voteChan:
		if vote {
			r.currentTerm++
			r.votedFor = r.id
		}
	case <-r.appendChan:
		r.stateMachine.log = append(r.stateMachine.log, "new entry")
	case <-heartbeatChan:
		r.electionTimer.Reset(time.Duration(r.heartbeatTimeout) * time.Second)
	case <-r.electionTimer.C:
		r.state = "CANDIDATE"
	}
}

func (r *RaftNode) runCandidate() {
	r.currentTerm++
	r.votedFor = r.id
	voteCount := 1
	for _, peer := range r.peers {
		go func(p *RaftNode) {
			p.voteChan <- true
		}(peer)
	}
	for {
		select {
		case vote := <-r.voteChan:
			if vote {
				voteCount++
				if voteCount > len(r.peers)/2 {
					r.state = "LEADER"
					r.leaderId = r.id
					return
				}
			}
		case <-time.After(time.Duration(r.electionTimeout) * time.Second):
			r.state = "FOLLOWER"
			return
		case <-r.electionTimer.C:
			r.state = "FOLLOWER"
			return
		}
	}
}

func (r *RaftNode) runLeader() {
	go func() {
		for {
			time.Sleep(time.Duration(r.heartbeatTimeout) * time.Second)
			for _, peer := range r.peers {
				go func(p *RaftNode) {
					p.heartbeatChan <- true
				}(peer)
			}
		}
	}()
	for _, peer := range r.peers {
		go func(p *RaftNode) {
			p.appendChan <- "new entry"
		}(peer)
	}
	r.stateMachine.log = append(r.stateMachine.log, "new entry")
	r.replicateLog()

	// Add a new goroutine to detect leader failure
	go func() {
		leaderFailureTimer := time.NewTimer(time.Duration(r.leaderFailureTimeout) * time.Second)
		for {
			select {
			case <-leaderFailureTimer.C:
				// Leader has failed to respond, start a new election
				r.state = "CANDIDATE"
				return
			case <-r.heartbeatChan:
				// Leader is still responding, reset the timer
				leaderFailureTimer.Reset(time.Duration(r.leaderFailureTimeout) * time.Second)
			}
		}
	}()

	for {
		select {
		case <-time.After(time.Duration(r.electionTimeout) * time.Second):
			for _, peer := range r.peers {
				go func(p *RaftNode) {
					p.appendChan <- "new entry"
				}(peer)
			}
		}
	}
}

func (r *RaftNode) replicateLog() {
	for _, peer := range r.peers {
		go func(p *RaftNode) {
			for _, entry := range r.stateMachine.log {
				p.appendChan <- entry
			}
		}(peer)
	}
}

/*
This code is a simple implementation of the Raft consensus algorithm in Go, it creates 3 raft nodes, and simulate the communication between these nodes.

It starts with a leader node, a follower node and a candidate node.
The leader node sends append entries to all the other peers, while the follower node waits for append entries and votes, and the candidate node waits for votes.

When you run this code, it will not produce any output and continue to run indefinitely.
It simulates the Raft algorithm, but it does not have a way to stop it. If you want the program to stop, you can add a mechanism to stop the infinite loop in the main function.

You should also note that this implementation is not production ready, as it is missing a lot of features that a real-world implementation of Raft should have such as log replication, leader election timeout and heartbeat timeout.
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	stopChan := make(chan bool)
	sm := &StateMachine{}
	node1 := &RaftNode{id: 1, stateMachine: sm, electionTimeout: 10, heartbeatTimeout: 5}
	node2 := &RaftNode{id: 2, stateMachine: sm, electionTimeout: 10, heartbeatTimeout: 5}
	node3 := &RaftNode{id: 3, stateMachine: sm, electionTimeout: 10, heartbeatTimeout: 5}
	node1.peers = []*RaftNode{node2, node3}
	node2.peers = []*RaftNode{node1, node3}
	node3.peers = []*RaftNode{node1, node2}
	node1.state = "LEADER"
	node2.state = "FOLLOWER"
	node3.state = "CANDIDATE"
	go func() {
		defer wg.Done()
		node1.start()
	}()
	go func() {
		defer wg.Done()
		node2.start()
	}()
	go func() {
		defer wg.Done()
		node3.start()
	}()
	go func() {
		time.Sleep(time.Second * 30)
		stopChan <- true
	}()
	for {
		select {
		case <-stopChan:
			fmt.Println("stop signal received, shutting down Raft cluster")
			wg.Wait() // added this line
			fmt.Println("Raft cluster is shut down")
			return
		}
	}
}
