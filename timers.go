package main

import "time"

func Timers() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
}
