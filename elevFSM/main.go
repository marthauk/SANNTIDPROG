package main

import (
	. "./elevDrivers"
	"fmt"
	"elevFSM"
)

func main() {
	Elev_init()
	fmt.Println("Press STOP button to stop elevator and exit program.\n")
	Elev_set_motor_direction(1)
	tall := 10
	Stupidshit(tall)
	fmt.Printf("tall:%d",tall)

	for {
		if Elev_get_floor_sensor_signal() == N_FLOORS-1 {
			Elev_set_motor_direction(DIRN_DOWN)
		} else if Elev_get_floor_sensor_signal() == 0 {
			Elev_set_motor_direction(DIRN_UP)
		}

		if Elev_get_stop_signal() == 1 {
			Elev_set_motor_direction(DIRN_STOP)
		}
	}
}
