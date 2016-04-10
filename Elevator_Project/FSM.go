package main

import (
	. "./elevDrivers"
	//."./elevOrders"
	"fmt"
	"time"
)

const (
	IDLE = 0
	DRIVING = 1
	DOOR_TIMER = 2

	NONE = 0
	FLOOR_ARRIVAL = 1
	NEW_ORDER = 2
	TIMER_TIMEOUT = 3
)

type Elevator struct {
	STATE 			int
	EVENT 			int
	CURRENT_FLOOR 	int
	DIRECTION 		int
}



func main(){
	/* INITIALIZATION */
	Elev_init()
	Elev_set_motor_direction(DIRN_DOWN)
	/* DRIVES THE ELEVATOR TO THE FIRST FLOOR DOWNWARDS */
	for{
		if (Elev_get_floor_sensor_signal() != -1){
			Elev_set_motor_direction(DIRN_STOP)
			break
		}
	}
	/* SETS INITIAL STATE VARIABLES */
	Elevator := 				Elevator{IDLE, NONE, Elev_get_floor_sensor_signal(), DIRN_STOP}
	Button_Press_Chan :=		make(chan Button, 10)
	//Floor_Arrival_Chan :=		make(chan int)
	//Door_Close_Chan :=		make(chan int)
	//Objective_Chan :=			make(chan Button)
	//Elevator_Stop_Chan:=		make(chan int)
	
	/* STARTS ESSENTIAL PROCESSES */
	go Elev_sensor_pooler(Button_Press_Chan)
	go Handle_new_order(Button_Press_Chan)
	//go Objective_dealer()
	go Orders_init()

	time.Sleep(time.Millisecond*200)

	fmt.Printf("\n\n\n####################################################\n")
	fmt.Printf("## The elevator has been succesfully initiated! #### \n") 
	fmt.Printf("####################################################\n\n")

	fmt.Printf("STATE: %d \n", Elevator.STATE)
	fmt.Printf("EVENT: %d \n", Elevator.EVENT)
	fmt.Printf("CURRENT_FLOOR: %d \n", Elevator.CURRENT_FLOOR)
	fmt.Printf("DIRECTION: %d \n\n\n", Elevator.DIRECTION)

	Print_all_orders()	
	
	
	for{
		select{
		case NewObjective := <- 	Objective_Chan:
			FSM_Start_Driving(NewObjective)
		case Arrival := <- 			Floor_Arrival_Chan:
			FSM_Should_stop_or_not(Arrival)
		case DoorClosed := <- 		Door_Close_Chan:
			FSM_Return_to_idle(Elevator)
		}









	/*
		//Print_all_orders()
		switch(Elevator.STATE) {
		case IDLE:
			//Sjekk etter ny event(nye ordre)
			//IF(ny event)
				//kjør i retning bestemt av event
				//state = driving
			Elevator.STATE = DRIVING
		case DRIVING:
			//Sjekk etter ny event(ankommer ny floor)
			//IF(floor = destination floor)
				//stop heisen
				//fix lys
				//start timer
				//state = door_timer
			Elevator.STATE = DOOR_TIMER
		case DOOR_TIMER:
			//Sjekker etter ny event(timer_timeout)
			//IF(timer_timeout)
				//lukk dør
				//fix kø
				//state = idle
			Elevator.STATE = IDLE
		}	
	}
	*/
}

















/*
type struct ElevatorState {
	state 		State
	floor 		int
	dirn 		Dirn
	orders 	    [][3]bool
}

func chooseDirection(e ElevatorState) bool {

}
*/

/*
func fsm(channels and stuff){

	var e 							ElevatorState
	var doorCloseCh 				<-chan time.Time
	var failedToArriveAtFloorCh 	<-chan time.Time 	// start time.After each time we start moving (or don't stop at that floor)

	func newOrder(floor int, type int){
		swtich(e.state){
		case Idle:			
			e.orders[o.floor][type] = true
			e.dirn = chooseDirection(e)
			if(e.dirn == Stop){
				e.state = DoorOpen
				doorCloseCh = time.After(doorOpenDuration)
					elel_set_door stuff
				} else {
					Elev_set_motor_direction(e.dirn)
					e.state = Moving
				}
			}
		case Moving:
			e.orders[o.floor][type] = true
		case DoorOpen:
			if(e.floor == o.floor){
				doorCloseCh = time.After(doorOpenDuration)
			} else {
				e.orders[o.floor][type] = true
			}
		}
	}

	for {
		select {
		case o := <-cabOrderCh:
			newOrder(o.floor, Cab)



		case o := <-hallOrderCh:
			newOrder(o.floor, o.type)

		case f := <-Floor_Arrival_Chan:
			if should Stop
				//what do if hall order? send network stuff? channel?

		case f := <-failedToArriveAtFloorCh:
			//tell network/stuff that we are "disconnected"

		case d := <-doorCloseCh:
		case getStateCh<-e:
		}
	}
}
*/

