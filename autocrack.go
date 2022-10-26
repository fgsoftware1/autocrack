package main

import (
	"fmt"
	"log"
	"os/exec"
)

var (
	Inter string
	Net   string
	BSSID string
	SSID  string
)

func main() {
	setup()
	monitor()
}

func setup() {
	interfaces := []string{}

	fmt.Print("\x1b[38;5;4m Looking for up network interfaces...\n")

	scanCommand := "ip link | awk '/state UP/{print $2a;getline}' | awk '{ print substr( $0, 1, length($0)-1) }'"
	scan, err := exec.Command("sh", "-c", scanCommand).Output()
	scanOut := string(scan)

	if err != nil {
		log.Fatal(err)
	}

	interfaces = append(interfaces, scanOut)

	for index, element := range interfaces {
		fmt.Print("\x1b[38;5;11m\n")
		fmt.Println(index, element)
		fmt.Print("\x1b[38;5;13m")
		fmt.Println("choose interface index: ")
		var inter int
		fmt.Scanln(&inter)
		Inter = interfaces[inter]
		isPresent := arrayContains(interfaces, Inter)
		if !isPresent {
			panic("Selected index doesn't exist")
		}
	}

	//DEBUG
	fmt.Print("\x1b[38;5;124m")
	fmt.Print("DEBUG \n")
	fmt.Print("\x1b[38;5;141m")
	fmt.Print("interface selected: " + Inter + "\n")
}

func monitor() {
	Ssid := []string{}
	Bssid := []string{}

	fmt.Print("\x1b[38;5;4mLooking for available networks...\n")

	ssidCommand := "nmcli -f SSID device wifi | awk 'NR>1'"
	ssid, err := exec.Command("sh", "-c", ssidCommand).Output()
	ssidOut := string(ssid)
	if err != nil {
		log.Fatal(err)
	}

	bssidCommand := "nmcli -f BSSID device wifi | awk 'NR>1'"
	bssid, err := exec.Command("sh", "-c", bssidCommand).Output()
	bssidOut := string(bssid)
	if err != nil {
		log.Fatal(err)
	}

	Ssid = append(Ssid, ssidOut)
	Bssid = append(Bssid, bssidOut)

	for index, element := range Ssid {
		fmt.Print("\x1b[38;5;11m\n")
		fmt.Println(index, element)
		fmt.Print("\x1b[38;5;13m")
		fmt.Println("choose network index: ")
		var net int
		fmt.Scanln(&net)
		Net = Ssid[net]
		isPresent := arrayContains(Ssid, Net)
		if !isPresent {
			panic("Selected index doesn't exist")
		}
	}
	//DEBUG
	fmt.Print("\x1b[38;5;124m")
	fmt.Print("DEBUG \n")
	fmt.Print("\x1b[38;5;141m")
	fmt.Print("network selected: " + Net + "\n")
}
func arrayContains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}
