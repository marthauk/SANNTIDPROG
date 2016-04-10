package FSM_Controller

import(
	."./elevDrivers"
	"time"
	"fmt"
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
	STATE 				int
	CURRENT_FLOOR 		int
	DESTINATION_FLOOR	int
	DIRECTION 			int
}

func FSM_setup_elevator(){
	Elev_init()
	Elev_set_motor_direction(DIRN_DOWN)
	for{
		if (Elev_get_floor_sensor_signal() != -1){
			Elev_set_motor_direction(DIRN_STOP)
			break
		}
	}
}


func FSM_create_elevator() Elevator{
	e := Elevator{STATE: IDLE, CURRENT_FLOOR: Elev_get_floor_sensor_signal(), DESTINATION_FLOOR: Elev_get_floor_sensor_signal(), DIRECTION: DIRN_STOP}
	return e
}


func FSM_Start_Driving(NewObjective Button, e Elevator, State_Chan chan int, Motor_Direction_Chan chan int, Location_Chan chan int){
	if e.CURRENT_FLOOR > NewObjective.Floor{
		Elev_set_motor_direction(-1)
		Motor_Direction_Chan <- -1
		State_Chan <- DRIVING
	}
	if e.CURRENT_FLOOR < NewObjective.Floor{
		Elev_set_motor_direction(1)
		Motor_Direction_Chan <- 1
		State_Chan <- DRIVING
	}
	if e.CURRENT_FLOOR == NewObjective.Floor{
		Location_Chan <- e.CURRENT_FLOOR
	}
	//Print_all_orders()
	time.Sleep(time.Millisecond*200)
}

func FSM_objective_dealer(e Elevator, State_Chan chan int, Destination_Chan chan int, Objective_Chan chan Button){
	for{
		nextOrder := Next_order()
		//fmt.Println(nextOrder.Floor)
		if (e.STATE == IDLE && nextOrder.Floor != -1){
			Objective_Chan <- nextOrder
			State_Chan <- DRIVING
			Destination_Chan <- nextOrder.Floor
			time.Sleep(time.Millisecond*200)
			fmt.Println("fuck")
		}
	}
}

func FSM_elevator_updater(e Elevator, Motor_Direction_Chan chan int, Location_Chan chan int, Destination_Chan chan int, State_Chan chan int) {
	for{
		select{
			case NewDirection := <- 						Motor_Direction_Chan:
				e.DIRECTION = NewDirection

			case NewFloor := <- 							Location_Chan:
				e.CURRENT_FLOOR = NewFloor

			case NewDestination := <- 						Destination_Chan:
				e.DESTINATION_FLOOR = NewDestination

			case NewState := <- 							State_Chan:
				e.STATE = NewState
		}
	}	
}

func FSM_floor_tracker(e Elevator, Location_Chan chan int, Floor_Arrival_Chan chan int){
	for{
		if Elev_get_floor_sensor_signal() != -1 && Elev_get_floor_sensor_signal() != e.CURRENT_FLOOR{
			NewFloor := Elev_get_floor_sensor_signal()
			Location_Chan <- NewFloor
			Floor_Arrival_Chan <- NewFloor
		}
	}
}

func FSM_sensor_pooler(Button_Press_Chan chan Button){
	for{
		for button := B_UP; button <= B_COMMAND; button++ {
			for floor:= 0; floor < N_FLOORS; floor++{
				if button == B_UP && floor == N_FLOORS-1 { continue }
				if button == B_DOWN && floor == 0 { continue }
				button_signal := Elev_get_button_signal(button, floor);
				if button_signal == 1 {
					Button_Press_Chan <- Button{Button_type : button, Floor : floor}
				}
			}
		}
		time.Sleep(time.Millisecond*100)
	}
}

func FSM_should_stop_or_not(Arrival_floor int, e Elevator, State_Chan chan int, Motor_Direction_Chan chan int){
	if Arrival_floor == e.DESTINATION_FLOOR{
		Elev_set_motor_direction(0)
		Motor_Direction_Chan <- 0
		State_Chan <- DOOR_TIMER
	}
}
