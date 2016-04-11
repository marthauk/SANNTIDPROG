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
	Door_Open_Req_Chan := 		make(chan int)
	
	/* STARTS ESSENTIAL PROCESSES */
	Orders_init()
	go Order_handler(Button_Press_Chan)
	go FSM_safekill()
	go FSM_sensor_pooler(Button_Press_Chan)
	go FSM_floor_tracker(&e, Location_Chan, Floor_Arrival_Chan, Motor_Direction_Chan, Destination_Chan, State_Chan)
	go FSM_objective_dealer(&e, State_Chan, Destination_Chan, Objective_Chan, Motor_Direction_Chan, Location_Chan)
	//go FSM_elevator_updater(&e, Motor_Direction_Chan, Location_Chan, Destination_Chan, State_Chan)
	time.Sleep(time.Millisecond*200)

	/* STARTUP TEXT */
	fmt.Printf("\n\n\n####################################################\n")
	fmt.Printf("## The elevator has been succesfully initiated! #### \n") 
	fmt.Printf("####################################################\n\n")
	fmt.Printf("STATE: %d , ", e.STATE)
	fmt.Printf("CURRENT_FLOOR: %d , ", e.CURRENT_FLOOR)
	fmt.Printf("DESTINATION_FLOOR: %d , ", e.DESTINATION_FLOOR)
	fmt.Printf("DIRECTION: %d \n\n\n", e.DIRECTION)

	Print_all_orders()

	for{
		select{
		case newObjective := 		<- Objective_Chan:
			FSM_Start_Driving(newObjective, &e, State_Chan, Motor_Direction_Chan, Location_Chan, Destination_Chan)

		case newFloorArrival := 	<- Floor_Arrival_Chan:
			FSM_should_stop_or_not(newFloorArrival, &e, State_Chan, Motor_Direction_Chan, Door_Open_Req_Chan, Location_Chan, Destination_Chan)
		
		case doorReq := 			<- Door_Open_Req_Chan:
			fmt.Println("DoorOpen fuck")
			FSM_door_opener(doorReq, &e, State_Chan, Motor_Direction_Chan, Location_Chan, Destination_Chan)

		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
