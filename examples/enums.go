//enum - An enum is a type that has a fixed number of possible values, each with a distinct name

package main

import "fmt"

type ServerState interface

//iota keyword generates successivce contsnat values automatically (0,1,2,...)
const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)


//hard for many values
//use  stringer tool can be used in conjunction with go:generate to automate the process.
var StateName = map[ServerState]string{
	StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

func (ss ServerState) string() string {
	return StateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Printl(ns)

	ns1 := transition(ns)
	fmt.Println(ns1)
}

func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unwknown state: %s", s))
    }
    return StateConnected
}