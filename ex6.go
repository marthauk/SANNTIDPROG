/* going to use UDP
my IP-address : 129.241.187.161

*/
package main

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

//start å sende imAlive før du initialiserer ny master.

//master sends numbers
func master(ipAdr string, port string, counter int) {
	fmt.Println("Im alive")
	serverAddr, err := net.ResolveUDPAddr("udp", ipAdr+":"+port)
	conn, err := net.DialUDP("udp", nil, serverAddr)
	CheckError(err)
	countSendS := ""
	countSend := make([]byte, 1024)
	for {
		counter++
		countSendS = strconv.Itoa(counter)
		countSend = []byte(countSendS)
		_, err = conn.WriteToUDP(countSend, serverAddr)
		fmt.Print(counter)
		time.Sleep(1)

	}
}

//backup listens for number, saves numbers in string.
func backup(ipAdr string, port string, counter int) {
	fmt.Println("Im backup")
	Backup := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run ex6.go")
	Backup.Run()
	serverAddr, err := net.ResolveUDPAddr("udp", ipAdr+":"+port)
	psock, err := net.ListenUDP("udp4", serverAddr)
	psock.SetDeadline(time.Now().Add(3 * time.Second)) //using timer on socket to check if master is alive

	if err != nil {
		return
	}
	buf := make([]byte, 1024)

	for {
		n, remoteAddr, err := psock.ReadFromUDP(buf) // if psock times out, err vil get a value
		fmt.Println(n)

		if err != nil {
			fmt.Println("Im alive")
			return
		}

		//else{
		//	fmt.Println("Im alive")

		//}
		//if remoteAddr.IP.String() != MY_IP {
		//		fmt.Printf("%s\n", buf)
		//	}

	}

}

func main() {

	counter := 0
	backup("localhost", "30000", counter) //
	master("localhost", "30000", counter) //
	//LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	//CheckError(err)

}
