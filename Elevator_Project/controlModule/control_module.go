//need to include message.go
package controlModule
import(
"net"
."./network"
"strconv"
)
const port int = 40000;//setting the portnumber. Selected a random port

type elev struct {
	self_id int
	self_IP string
	elevators map[int]*Elevator
	external_orders [2][N_FLOORS]int
	master int 
}

//initializing of elevator
func Initialize_elev() elev{
	var e elev
	e.elevators = make(map[int]*Elevator) 


	//initialize message here ?
	//initialize driver
	addr,_ :=net.InterfaceAddrs()
	tempVar:=addr[1]
 	ip:=tempVar.String()
	e.self_IP=ip[0:15]
	e.self_id := int(addr[1].String()[12]-'0')*100 + int(addr[1].String()[13]-'0')*10 + int(addr[1].String()[14]-'0')//this will work for IP-addresses of format ###.###.###.###, but not with only for ###.###.###.##/
	e.elevators[e.self_id]=new(Elevator)



	e.set_master()


}

func Initialize_connections(){
 	var isMaster bool =false
 	var tempIP string= e.self_IP[0:12]
 	var masterIP string ="255.255.255.255"
 	masterIP=tempIP + strconv.Itoa(e.master)
 	if e.self_id==e.master
 		{	//the current elevator is master, needs master settings
 				//go masterUDPSEND
 				//go masterUDPLISTEN
 				isMaster=true
 				
 		}
 	UDP_initialize(isMaster,port,masterIP)
	
	
}

func Set_master(){
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
