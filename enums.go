package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
	StateDisconnected
)

var stateName = map[ServerState]string{
	StateIdle:         "Idle",
	StateConnected:    "Connected",
	StateError:        "Error",
	StateRetrying:     "Retrying",
	StateDisconnected: "Disconnected",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	case StateDisconnected:
		return StateDisconnected
	default:
		panic("Invalid state")
	}
}

func enums() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(StateConnected)
	fmt.Println(ns2)
}
