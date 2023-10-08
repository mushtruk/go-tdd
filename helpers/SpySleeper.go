package helpers

import "time"

type SpyTime struct {
	DurationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.DurationSlept = duration
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, Sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, Write)
	return
}

const Write = "write"
const Sleep = "sleep"