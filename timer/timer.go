package timer

import "time"

const DOOR_OPEN_DURATION = 3 * time.Second

func initDoorTimer() *time.Timer {
	doorTimer := time.NewTimer(0 * time.Second)
	<-doorTimer.C
	return doorTimer
}

func DoorTimer(resetDoorTimer chan bool, doorTimeout chan bool) {
	doorTimer := initDoorTimer()
	for {
		select {
		case <-resetDoorTimer:
			doorTimer.Reset(DOOR_OPEN_DURATION)

		case <-doorTimer.C:
			doorTimeout <- true
		}
	}
}
