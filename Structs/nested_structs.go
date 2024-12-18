// Complete the canSendMessage function. 
// It should return true only if the sender and recipient fields each contain a name and a number. 
// If any of the default zero values are present, return false instead.

package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	// my solution:
	if mToSend.sender.name != "" && mToSend.sender.number != 0 && 
	   mToSend.recipient.name != "" && mToSend.recipient.number != 0 {
		return true
	}
	return false
}