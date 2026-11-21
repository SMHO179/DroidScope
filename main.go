package main

import (
	"fmt"
	"os/exec"
)

var version string = "0.1.0"
var build string = "2026-21-11"
var adbStatus bool = true
var connectedDevices int = 1
var condition string = "not OK"

func adb() {
	cmd := exec.Command("adb", "version")
	output, err := cmd.Output()
}

func start() {
	fmt.Println("DroidScope v" + version)
	fmt.Println("developed by: @SMHO179")
	fmt.Println("language: Go")
	fmt.Println("build: " + build)
	if condition == "OK" {
		adbStatus = true
	} else {
		adbStatus = false
		connectedDevices = 0
	}
	fmt.Printf("Status ADB: %t\n", adbStatus)
	fmt.Printf("Connected Devices: %d\n", connectedDevices)
	// fmt.Println()
}

func main() {
	start()
}
