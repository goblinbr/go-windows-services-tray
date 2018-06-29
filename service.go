package main

import (
	"fmt"

	"golang.org/x/sys/windows/svc"

	"golang.org/x/sys/windows/svc/mgr"
)

type WinService struct {
	Manager *mgr.Mgr
	Service *mgr.Service
}

func (m *WinService) Close() error {
	return m.Manager.Disconnect()
}

func HasService(name string) (bool, error) {
	service, err := openService(name)
	if service != nil {
		defer service.Close()
	}
	if err != nil {
		return false, err
	}

	return service != nil, nil
}

func GetServiceStatus(name string) (svc.State, error) {
	service, err := openService(name)
	if service != nil {
		defer service.Close()
	}
	if err != nil {
		return svc.Stopped, err
	}

	status, err := service.Service.Query()
	if err != nil {
		return svc.Stopped, err
	}
	return status.State, nil
}

func openService(name string) (*WinService, error) {
	manager, err := mgr.Connect()
	if err != nil {
		return nil, err
	}
	defer manager.Disconnect()
	service, err := manager.OpenService(name)
	if err != nil {
		return nil, fmt.Errorf("could not access service: %v", err)
	}

	return &WinService{Manager: manager, Service: service}, nil
}
