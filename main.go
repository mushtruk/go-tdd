package main

import (
	"jazz-store/Countdown"
	"jazz-store/helpers"
	"os"
	"time"
)

func main() {
	sleeper := helpers.NewConfigurableSleeper(1 * time.Second, time.Sleep)
	Countdown.Countdown(os.Stdout, sleeper)
}
