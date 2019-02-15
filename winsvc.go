package winsvc

import (
	"fmt"
  "log"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
  
)

func StartService(name string) {
	m, err := mgr.Connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer s.Close()
	err = s.Start("is", "manual-started")
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func StopService(name string) {
	c := svc.Stop
	m, err := mgr.Connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer s.Close()
	status, err := s.Control(c)
	if err != nil {
		log.Println(fmt.Sprintf("Info: %v", status))
		log.Println(err.Error())
		return
	}
}

// Potentially dangerous function
func DeleteService(name string) {
	m, err := mgr.Connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer s.Close()
	err = s.Delete()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
