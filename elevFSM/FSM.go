package main

import (
	."./FSM_Controller"
	"fmt"
	"time"
)




func main(){
	/* INITIALIZATION */
	FSM_setup_elevator()
	
	/* SETS INITIAL STATE VARIABLES */
	e := 						FSM_create_elevator()

	/* CHANNELS FOR UPDATING THE ELEVATOR VARIABLES */
	Button_Press_Chan :=		make(chan Button, 10)
	Location_Chan :=			make(chan int)
	Motor_Direction_Chan :=		make(chan int)
	Destination_Chan := 		make(chan int)
	State_Chan :=				make(chan int)
	
	/* EVENT CHANNELS */
	Objective_Chan :=			make(chan Button)
	Floor_Arrival_Chan :=		make(chan int)
	//Door_Close_Chan :=		make(chan int)
	
	/* STARTS ESSENTIAL PROCESSES */
	Orders_init()
	go Order_handler(Button_Press_Chan)
	go FSM_sensor_pooler(Button_Press_Chan)
	go FSM_floor_tracker(e, Location_Chan, Floor_Arrival_Chan,)
	go FSM_objective_dealer(e, State_Chan, Destination_Chan, Objective_Chan)
	go FSM_elevator_updater(e, Motor_Direction_Chan, Floor_Arrival_Chan,Destination_Chan, State_Chan)
	//Elevator_Stop_Chan:=		make(chan int)
	//Door_Close_Chan :=		make(chan int)
	
	/* STARTS ESSENTIAL PROCESSES */
	go Order_handler(Button_Press_Chan)
	go FSM_sensor_pooler(Button_Press_Chan)
	go FSM_floor_tracker(e, Location_Chan, Floor_Arrival_Chan)
	go FSM_objective_dealer(e, State_Chan, Destination_Chan, Objective_Chan)
	go FSM_elevator_updater(e, Motor_Direction_Chan, Floor_Arrival_Chan,Destination_Chan, State_Chan)
	go Orders_init()
	time.Sleep(time.Millisecond*200)

	/* STARTUP TEXT */
	fmt.Printf("\n\n\n####################################################\n")
	fmt.Printf("## The elevator has been succesfully initiated! #### \n") 
	fmt.Printf("####################################################\n\n")
	fmt.Printf("STATE: %d \n", e.STATE)
	fmt.Printf("CURRENT_FLOOR: %d \n", e.CURRENT_FLOOR)
	fmt.Printf("DESTINATION_FLOOR: %d \n", e.DESTINATION_FLOOR)
	fmt.Printf("DIRECTION: %d \n\n\n", e.DIRECTION)
	fmt.Printf("STATE: %d \n", e.STATE)
	fmt.Printf("CURRENT_FLOOR: %d \n", e.CURRENT_FLOOR)
	fmt.Printf("DIRECTION: %d \n\n\n", e.DIRECTION)

	Print_all_orders()	
	
	for{
		select{
		case newObjective := 		<- Objective_Chan:
			FSM_Start_Driving(newObjective, e, State_Chan, Motor_Direction_Chan, Floor_Arrival_Chan)
			fmt.Printf("STATE: %d \n", e.STATE)
			fmt.Printf("CURRENT_FLOOR: %d \n", e.CURRENT_FLOOR)
			fmt.Printf("DESTINATION_FLOOR: %d \n", e.DESTINATION_FLOOR)
			fmt.Printf("DIRECTION: %d \n\n\n", e.DIRECTION)
		
		case newFloorArrival := 	<- Floor_Arrival_Chan:
			FSM_should_stop_or_not(newFloorArrival, e, State_Chan, Motor_Direction_Chan)
/*
		case doorClosed := 			<- Door_Close_Chan:
			FSM_Return_to_idle(Elevator)
*/

		/*
		case Arrival := <- Floor_Arrival_Chan:
			FSM_Should_stop_or_not(Arrival)
		case DoorClosed := <- Door_Close_Chan:
			FSM_Return_to_idle(Elevator)
		*/
		}
	}
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

