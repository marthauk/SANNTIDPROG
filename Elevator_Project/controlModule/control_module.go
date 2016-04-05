//need to include message.go

import(
"net"
."./network"
)

type elev struct {
	self_id int
	self_IP string
	elevators map[int]*Elevator
	external_orders [2][N_FLOORS]int
	master int
}

//initializing of elevator
func initialize_elev() elev{
	var e elev
	e.elevators = make(map[int]*Elevator) 

	//initialize driver
	addr,_ :=net.InterfaceAddrs()
	tempVar:=addr[1]
 	ip:=tempVar.String()
	e.self_IP=ip[0:15]
	e.self_id := int(addr[1].String()[12]-'0')*100 + int(addr[1].String()[13]-'0')*10 + int(addr[1].String()[14]-'0')
	e.elevators[e.self_id]=new(Elevator)


	e.set_master()


}

func initialize_connections(){
 	if e.self_id==e.master
 		{	//the current elevator is master, needs master settings
 				//go masterUDPSEND
 				//go masterUDPLISTEN

 		}
	else
	 { //the current elevator is slave, needs slave settings for udp
	 				//go slaveUDPSEND
 				//go slaveUDPLISTEN
		}
}

func set_master(){
	// checking which elevator has the highest IP to determine master 
	max :=0
	for i,_ :=range(e.elevators){
		if max<i  {
			max=i  
				
		}
	}
	e.master=max
	fmt.Println("new master is", e.master)

}