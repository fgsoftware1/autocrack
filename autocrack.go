package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var (
	Debug bool
	Inter string
	Net   string
	BSSID string
	SSID  string
	CH 	  string
)

func main() {
	debug := flag.Bool("debug", false, "define debug true/false false is default")
	flag.Parse()
	Debug = *debug

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

	interfaces = strings.Split(scanOut, "\n")
	interfaces[len(interfaces)-1] = ""
	interfaces = interfaces[:len(interfaces)-1]

	//DEBUG
	if Debug == true {
		fmt.Print("\x1b[38;5;124m")
		fmt.Print("DEBUG \n")
		fmt.Print("\x1b[38;5;141m")
		fmt.Print(interfaces, "\n")
	}

	for index, element := range interfaces {
		fmt.Print("\x1b[38;5;11m\n")
		fmt.Println(index, element)
	}
	fmt.Print("\x1b[38;5;13m")
	fmt.Println("choose interface index: ")
	var inter int
	fmt.Scanln(&inter)
	Inter = interfaces[inter]
	isPresent := arrayContains(interfaces, Inter)
	if !isPresent {
		panic("Selected index doesn't exist")
	}

	//DEBUG
	if Debug == true {
		fmt.Print("\x1b[38;5;124m")
		fmt.Print("DEBUG \n")
		fmt.Print("\x1b[38;5;141m")
		fmt.Print("interface selected: " + Inter + "\n")
	}
}

func monitor() {
	Ssid := []string{}
	Bssid := []string{}
	Channel := []string{}

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

	channelCommand := "nmcli -f CHAN device wifi | awk 'NR>1'"
	channel, err := exec.Command("sh", "-c", channelCommand).Output()
	channelOut := string(channel)
	if err != nil {
		log.Fatal(err)
	}

	Ssid = strings.Split(ssidOut, "\n")
	Ssid[len(Ssid)-1] = ""
	Ssid = Ssid[:len(Ssid)-1]

	Bssid = strings.Split(bssidOut, "\n")
	Bssid[len(Bssid)-1] = ""
	Bssid = Bssid[:len(Bssid)-1]

	Channel = strings.Split(channelOut, "\n")
	Channel[len(Channel)-1] = ""
	Channel = Channel[:len(Channel)-1]

	//DEBUG
	if Debug == true {
		fmt.Print("\x1b[38;5;124m")
		fmt.Print("DEBUG \n")
		fmt.Print("\x1b[38;5;141m")
		fmt.Print(Ssid, "\n")
		fmt.Print(Bssid, "\n")
		fmt.Print(Channel, "\n")
	}
	//SSID
	for index, element := range Ssid {
		fmt.Print("\x1b[38;5;11m\n")
		fmt.Println(index, element)
	}
	fmt.Print("\x1b[38;5;13m")
	fmt.Println("choose network index: ")
	var net int
	fmt.Scanln(&net)
	Net = Ssid[net]
	isPresent := arrayContains(Ssid, Net)
	if !isPresent {
		panic("Selected index doesn't exist")
	}

	SSID 	= Ssid[net]
	BSSID 	= Bssid[net]
	CH 		= Channel[net]

	//DEBUG
	if Debug == true {
		fmt.Print("\x1b[38;5;124m")
		fmt.Print("DEBUG \n")
		fmt.Print("\x1b[38;5;141mSSID\n", "\x1b[38;5;81m"+SSID, "\n")
		fmt.Print("\x1b[38;5;141mBSSID\n", "\x1b[38;5;81m"+BSSID, "\n")
		fmt.Print("\x1b[38;5;141mCH\n", "\x1b[38;5;81m"+CH, "\n")
	}
}

func arrayContains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}
