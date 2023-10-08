package Countdown

import (
	"fmt"
	"io"
	"jazz-store/helpers"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(w io.Writer, sleeper helpers.Sleeper) {
	for i:=countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord)
}