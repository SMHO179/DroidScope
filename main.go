package main

import (
	"fmt"
	"os/exec"
)

var version = "0.1.5"
var build = "2026-11-21"

func adb() (string, error) {
	cmd := exec.Command("adb", "version")

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("ADB Error: %w", err)
	}

	return string(output), nil
}

func start() {
	fmt.Println("DroidScope v" + version)
	fmt.Println("Developed by: @SMHO179")
	fmt.Println("Language: Go")
	fmt.Println("Build: " + build)

	output, err := adb()

	if err != nil {
		fmt.Println("Status ADB: Not Available")
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Status ADB: Available")
	fmt.Println("ADB Version:")
	fmt.Print(output)
}

func main() {
	start()
}
