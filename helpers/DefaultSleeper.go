package helpers

import "time"

type Sleeper interface {
	Sleep()
}
type ConfigurableSleeper struct{
	duration time.Duration
	sleep func(time.Duration)
}

func NewConfigurableSleeper(duration time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{
		duration: duration,
		sleep:    sleep,
		}
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type DefaultSleeper struct{
	duration time.Duration
	sleep func(time.Duration)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}