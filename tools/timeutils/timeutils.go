package timeutils

import "time"

func ImmediatelyTick(execution func(), duration time.Duration) {

	t := time.NewTicker(duration)

	execution()

	go func() {
		for {
			<-t.C

			execution()

		}

		t.Stop()
	}()

}
func DelayTick(execution func(), duration time.Duration) {

	t := time.NewTicker(duration)


	go func() {
		for {
			<-t.C

			execution()

		}

		t.Stop()
	}()

}
