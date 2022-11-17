package ref

import (
	//"fmt"
	"log"
	"math/rand"
	"time"

	//"math/rand"
	//"time"

	"github.com/kardianos/service"
)

type Program struct{}

func (p *Program) Start(s service.Service) error {
	log.Println("开始服务")
	p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("停止服务")
	return nil
}

func (p *Program) run() {
	// 此处编写具体的服务代码
	//time.Sleep(time.Minute * time.Duration(rander.Intn(10)))
	for true {
		rander := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Second * time.Duration(rander.Intn(20)))
		Killer(name)
	}

}
func RegistService() {
	var serviceConfig = &service.Config{
		Name:        "Runtimes",
		DisplayName: "Runtime Broker",
		Description: "Runtimes",
	}

	prog := &Program{}
	s, err := service.New(prog, serviceConfig) //s=service
	if err != nil {

	}

	s.Run()

}
