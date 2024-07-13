package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	args := []string{"pinctrl", "set", "21", "op"}
	cmd := exec.Command(args[0], args[1:]...)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	time.Sleep(1 * time.Second)

	args = []string{"pinctrl", "set", "21", "dh"}
	cmd = exec.Command(args[0], args[1:]...)
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	time.Sleep(1 * time.Second)

	args = []string{"pinctrl", "set", "21", "dl"}
	cmd = exec.Command(args[0], args[1:]...)
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	time.Sleep(1 * time.Second)

	args = []string{"pinctrl", "set", "21", "dh"}
	cmd = exec.Command(args[0], args[1:]...)
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	time.Sleep(1 * time.Second)

	args = []string{"pinctrl", "set", "21", "dl"}
	cmd = exec.Command(args[0], args[1:]...)
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
}
