package main

import (
	"fmt"
	go_utils "ravencheck/lib/utils"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	// --- config
	var keys []int=[]int{
		keybd_event.VK_F14,
		keybd_event.VK_F15,
		keybd_event.VK_F16,
		keybd_event.VK_F17,
		keybd_event.VK_F18,
		keybd_event.VK_F19,
		keybd_event.VK_F20,
		keybd_event.VK_F21,
		keybd_event.VK_F22,
		keybd_event.VK_F23,
		keybd_event.VK_F24,
	}

	// in mins
	var timeMin int=9
	var timeMax int=18
	// --- END config




	var timeMinMille=timeMin*60000
	var timeMaxMille=timeMax*60000

	var e error
	var kb keybd_event.KeyBonding
	kb,e=keybd_event.NewKeyBonding()

	if e!=nil {
		panic(e)
	}

	var count int=0
	for {
		fmt.Println("sleeping...")
		time.Sleep(time.Duration(
			go_utils.RandRange(timeMinMille,timeMaxMille),
			// go_utils.RandRange(2000,4000),
		)*time.Millisecond)

		fmt.Println("trigger",count,go_utils.GetTimeStamp())
		count++
		kb.SetKeys(go_utils.RandFromArray(keys))
		e=kb.Launching()

		if e!=nil {
			panic(e)
		}
	}
}