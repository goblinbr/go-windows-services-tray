package main

import (
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/sys/windows/svc"

	"github.com/getlantern/systray"
)

const (
	textServicesNotRunning = "There's services that aren't running"
	iconServicesNotRunning = "assets/services-nok.ico"

	textAllServicesRunning = "All services are running"
	iconAllServicesRunning = "assets/services-ok.ico"

	scanServicesEvery = 5 * time.Second
)

func main() {
	systray.Run(onReady, onExit)
}

func onExit() {
}

func onReady() {
	configuration, err := ReadFile()
	if err != nil {
		println(err.Error())
		os.Exit(9)
	}

	services := configuration.Services

	mQuit := systray.AddMenuItem("Quit", "Quit")

	go monitorServices(services)
	go monitorSystrayMenu(mQuit)
}

func verifyIfAllServicesAreRunning(services []string) bool {
	for _, serviceName := range services {
		state, err := GetServiceState(serviceName)
		if err != nil {
			println(err.Error())
			return false
		} else if state != svc.Running {
			return false
		}
	}
	return true
}

func monitorServices(services []string) {
	allServicesAreRunning := false
	setIconAndTitleNotOk()

	for {
		allRunning := verifyIfAllServicesAreRunning(services)
		if allServicesAreRunning != allRunning {
			allServicesAreRunning = allRunning
			if allServicesAreRunning {
				setIconAndTitleOk()
			} else {
				setIconAndTitleNotOk()
			}
		}

		time.Sleep(scanServicesEvery)
	}
}

func setIconAndTitleNotOk() {
	systray.SetIcon(getIcon(iconServicesNotRunning))
	systray.SetTitle(textServicesNotRunning)
	systray.SetTooltip(textServicesNotRunning)
}

func setIconAndTitleOk() {
	systray.SetIcon(getIcon(iconAllServicesRunning))
	systray.SetTitle(textAllServicesRunning)
	systray.SetTooltip(textAllServicesRunning)
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		println(err.Error())
	}
	return b
}

func monitorSystrayMenu(mQuit *systray.MenuItem) {
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
			return
		}
	}
}
