package main

import (
	"strings"
	"testing"
)

func TestItShoudHaveLmhostsService(t *testing.T) {
	hasService, err := HasService("lmhosts")
	if err != nil {
		t.Error(err)
	} else if !hasService {
		t.Error("It should have lmhosts service")
	} else {
		t.Log("lmhosts service found")
	}
}

func TestItShoudNotHaveKryptonxkService(t *testing.T) {
	hasService, err := HasService("Kryptonxk")
	if err != nil && !strings.HasPrefix(err.Error(), "could not access service") {
		t.Error(err)
	} else if hasService {
		t.Error("It shouldn't have Kryptonxk service")
	} else {
		t.Log("Kryptonxk service not found")
	}
}

func TestItShoudHaveLmhostsRunning(t *testing.T) {
	status, err := GetServiceStatus("lmhosts")
	if err != nil {
		t.Error(err)
	} else {
		println("Status returned (lmhosts): ", status)
		t.Logf("Status returned (lmhosts): %v", status)
	}
}

func TestItShoudHaveWuauservRunning(t *testing.T) {
	status, err := GetServiceStatus("wuauserv")
	if err != nil {
		t.Error(err)
	} else {
		println("Status returned (wuauserv): ", status)
		t.Logf("Status returned (wuauserv): %v", status)
	}
}
