package main

import (
	"fmt"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	kb,e:=keybd_event.NewKeyBonding()

	if e!=nil {
		panic(e)
	}

	kb.SetKeys(
		keybd_event.VK_F20,
	)

	fmt.Println("preparing")
	time.Sleep(5*time.Second)
	e=kb.Launching()

	if e!=nil {
		panic(e)
	}
}