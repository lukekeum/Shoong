package main

import (
	"fmt"

	transcodepb "shoong.com/proto/transcode"
)

func main() {
	msg := &transcodepb.SendStreamDataRequest{
		Uuid:      "1234",
		StreamKey: "abcd",
	}

	fmt.Println("SendDataRequest:", msg)
}
