package main

import (
	"sync"
	"testing"
	"time"
)

func TestRaftNode_start(t *testing.T) {
	// Create a new RaftNode
	rn := &RaftNode{
		id:                   1,
		stateMachine:         &StateMachine{},
		currentTerm:          0,
		votedFor:             0,
		state:                "FOLLOWER",
		leaderId:             0,
		peers:                []*RaftNode{},
		electionTimeout:      5,
		heartbeatTimeout:     1,
		voteChan:             make(chan bool),
		appendChan:           make(chan string),
		stopChan:             make(chan bool),
		electionTimer:        time.NewTimer(5 * time.Second),
		heartbeatChan:        make(chan bool),
		leaderFailureTimeout: 10,
	}

	// Start the RaftNode
	rn.start()

	// Check that the state is set to "CANDIDATE" after the election timeout
	if rn.state != "CANDIDATE" {
		t.Errorf("Expected state to be CANDIDATE, got %s", rn.state)
	}

	// Send a vote to the RaftNode
	rn.voteChan <- true

	// Check that the currentTerm is incremented
	if rn.currentTerm != 1 {
		t.Errorf("Expected currentTerm to be 1, got %d", rn.currentTerm)
	}

	// Check that the votedFor value is set to the RaftNode's id
	if rn.votedFor != rn.id {
		t.Errorf("Expected votedFor to be %d, got %d", rn.id, rn.votedFor)
	}

	// Send an append message to the RaftNode
	rn.appendChan <- "new entry"

	// Check that the log is appended with the new entry
	if len(rn.stateMachine.log) != 1 || rn.stateMachine.log[0] != "new entry" {
		t.Errorf("Expected log to be [new entry], got %v", rn.stateMachine.log)
	}

	// Send a heartbeat message to the RaftNode
	rn.heartbeatChan <- true

	// Check that the election timer is reset
	if !rn.electionTimer.Stop() {
		t.Error("Expected election timer to be stopped")
	}
}

func TestRunCandidate(t *testing.T) {
	// create a RaftNode with 5 peers
	peers := make([]*RaftNode, 5)
	for i := 0; i < 5; i++ {
		peers[i] = &RaftNode{id: i}
	}
	// create a RaftNode with id 0 as the candidate
	candidate := &RaftNode{
		id:              0,
		currentTerm:     1,
		state:           "CANDIDATE",
		electionTimeout: 5,
		peers:           peers,
		voteChan:        make(chan bool, 5),
	}

	// Start the runCandidate method
	go candidate.runCandidate()

	// Wait for the method to complete
	time.Sleep(time.Duration(candidate.electionTimeout) * time.Second)

	// Check if the state of the candidate is "LEADER"
	if candidate.state != "LEADER" {
		t.Errorf("Expected state to be 'LEADER', got '%s'", candidate.state)
	}

	// Check if the leaderId is set to the candidate's id
	if candidate.leaderId != candidate.id {
		t.Errorf("Expected leaderId to be '%d', got '%d'", candidate.id, candidate.leaderId)
	}
}

func TestRunLeader(t *testing.T) {
	// Create a mock state machine
	mockStateMachine := &StateMachine{}
	// Create a leader node and set its state to "LEADER"
	leader := &RaftNode{
		id:                   1,
		stateMachine:         mockStateMachine,
		currentTerm:          1,
		votedFor:             1,
		state:                "LEADER",
		leaderId:             1,
		peers:                []*RaftNode{},
		electionTimeout:      5,
		heartbeatTimeout:     1,
		voteChan:             make(chan bool),
		appendChan:           make(chan string),
		stopChan:             make(chan bool),
		electionTimer:        time.NewTimer(5 * time.Second),
		heartbeatChan:        make(chan bool),
		leaderFailureTimeout: 10,
	}

	// Add some peers to the leader node
	peer1 := &RaftNode{id: 2}
	peer2 := &RaftNode{id: 3}
	leader.peers = append(leader.peers, peer1, peer2)

	// Start the leader node
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		leader.start()
		wg.Done()
	}()

	// Wait for a short period of time to allow the leader to send out heartbeats
	time.Sleep(2 * time.Second)

	// Check that the leader has sent out heartbeats to its peers
	if len(peer1.heartbeatChan) != 1 || len(peer2.heartbeatChan) != 1 {
		t.Error("Leader did not send out heartbeats to its peers")
	}

	// Check that the leader has replicated its log to its peers
	if len(peer1.appendChan) != 1 || len(peer2.appendChan) != 1 {
		t.Error("Leader did not replicate its log to its peers")
	}

	// Check that the leader has added an entry to its own log
	if len(mockStateMachine.log) != 1 {
		t.Error("Leader did not add an entry to its own log")
	}

	// Stop the leader node
	leader.stopChan <- true
	wg.Wait()

}
