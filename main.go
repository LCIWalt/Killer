package main // 声明 main 包，表明当前是一个可执行程序

import (
	"fmt"
	"os"

	//"os/user"
	"killer/ref"
	//"os/exec"
	//"log"
	//"github.com/kardianos/service"
)

func setHostname() string {
	var aa string
	aa, err := os.Hostname()
	if err != nil {
		fmt.Errorf("failed to get hostname")
	}
	return aa
}

func main() {
	var hostname string
	hostname = setHostname()
	fmt.Print(hostname) //certificate the hostname
	fmt.Printf("\n")
	ref.RegistService()
}
