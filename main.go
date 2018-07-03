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

func verifyIfAllServicesAreRunning(services []string) (bool, error) {
	for _, serviceName := range services {
		state, err := GetServiceState(serviceName)
		if err != nil {
			println(err.Error())
			return false, err
		} else if state != svc.Running {
			return false, nil
		}
	}
	return true, nil
}

func monitorServices(services []string) {
	allServicesAreRunning := false
	setIconAndTitleNotOk(textServicesNotRunning)
	for {
		allRunning, err := verifyIfAllServicesAreRunning(services)
		if err != nil {
			setIconAndTitleNotOk("Error: " + err.Error() + "\nVerify if it's running with administrator privileges!")
		} else if allServicesAreRunning != allRunning {
			allServicesAreRunning = allRunning
			if allServicesAreRunning {
				setIconAndTitleOk()
			} else {
				setIconAndTitleNotOk(textServicesNotRunning)
			}
		}

		time.Sleep(scanServicesEvery)
	}
}

func setIconAndTitleNotOk(text string) {
	systray.SetIcon(getIcon(iconServicesNotRunning))
	systray.SetTitle(text)
	systray.SetTooltip(text)
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
