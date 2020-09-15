package timeutils

import "time"

func ImmediatelyTick(execution func(), duration time.Duration) {

	t := time.NewTicker(duration)

	defer t.Stop()
	execution()

	go func() {
		for {
			<-t.C

			execution()

		}
	}()

}
