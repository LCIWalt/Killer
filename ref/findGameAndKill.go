package ref

import (
	"io/ioutil"
	"os/exec"
	"syscall"

	//"fmt"
	//"strings"
	//	"context"
	"log"
	//"math/rand"
	//"time"
	//"os"
)
var name string = readerr()
func readerr() string{
	data, err := ioutil.ReadFile("conf.txt") //bytes stream
	if err != nil {
		log.Fatal(err)
	}
	print(string(data))
	
	return string(data)
}

func Killer(name string) {
	
	//var s = "taskkill -f /IM " + name + ".exe"
	var cmd = exec.Command("taskkill", "-f", "/IM", name+".exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} //hide child process windows

	cmd.Run()

}

