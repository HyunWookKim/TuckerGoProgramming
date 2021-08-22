package main

import (
	"goprojects/ex20/fedex"
	// "goprojects/ex20/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// sender := &koreaPost.PostSender{}
	sender := &fedex.FedexSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
