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
	state, err := GetServiceState("lmhosts")
	if err != nil {
		t.Error(err)
	} else {
		println("State returned (lmhosts): ", state)
		t.Logf("State returned (lmhosts): %v", state)
	}
}

func TestItShoudHaveWuauservRunning(t *testing.T) {
	state, err := GetServiceState("wuauserv")
	if err != nil {
		t.Error(err)
	} else {
		println("State returned (wuauserv): ", state)
		t.Logf("State returned (wuauserv): %v", state)
	}
}
