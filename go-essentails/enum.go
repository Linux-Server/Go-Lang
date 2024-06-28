package main

import "fmt"

const (
	pending = iota //0
	processing
	success
	failed
	refund // 4
)

type ServerState int

var state_status = map[ServerState]string{
	pending:    "Payment Pending",
	processing: "Payment Processing",
	success:    "payment Success",
	failed:     "Payment Failed",
	refund:     "Paymnet Refunded",
}

func my_enum() {

	curr_state := transition(processing)
	_ = curr_state
	fmt.Println("The curr is : ", curr_state)
	data := state_status[curr_state]
	fmt.Println("The data is :", data)

}

// func (ss ServerState) String() string {
// 	fmt.Println("calling..")
// 	return state_status[ss]
// }

func transition(ss ServerState) ServerState {
	switch ss {
	case pending:
		return processing
	case processing:
		return success
	case success, failed:
		return pending
	default:
		return refund
	}

}
