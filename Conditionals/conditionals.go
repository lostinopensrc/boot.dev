package Conditionals

import "fmt"

func IsConditionals (){
	messageLen := 10
	maxMessageLen := 20

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}