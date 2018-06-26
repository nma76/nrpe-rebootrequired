package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Main function
func main() {
	//Initialize variable that holds status
	var nrpeStatus NrpeStatus

	//Call function that does all work with
	checkNrpe(&nrpeStatus)

	//Print status message and exit with correct code
	// fmt.Println(nrpeStatus.Code)
	fmt.Println(nrpeStatus.Message)
	os.Exit(int(nrpeStatus.Code))
}

// This function does all work and sets the status message and code
func checkNrpe(nrpeStatus *NrpeStatus) {
	//Set default return values
	nrpeStatus.Message = "unknown"
	nrpeStatus.Code = UNKNOWN

	//Check if reboot-required file exists
	//	var rebootRequiredFile = "/var/run/reboot-required"
	var rebootRequiredFile = "c:\\Temp\\reboot-required"

	if fileInfo, err := os.Stat(rebootRequiredFile); os.IsNotExist(err) {
		//Reboot not required
		nrpeStatus.Code = OK
		nrpeStatus.Message = "Reboot not required."
	} else {
		//Check age of reboot-required. If older than two days, set status to CRITICAL, else set status to WARNING
		var rebootRequiredFileAge = time.Now().Sub(fileInfo.ModTime())
		if rebootRequiredFileAge.Hours() >= 0 {
			nrpeStatus.Code = CRITICAL
		} else {
			nrpeStatus.Code = WARNING
		}

		//Get the reason why by reading the contents of reboot-required.pkgs
		var rebootRequiredPkgs = rebootRequiredFile + ".pkgs"
		if _, err := os.Stat(rebootRequiredPkgs); os.IsNotExist(err) {
			//File didn't exist, reason is unknown
			nrpeStatus.Message = fmt.Sprintf("System needs reboot! Reboot pending for %.0f hours. Responsible packages: unknown", rebootRequiredFileAge.Hours())
		} else {
			//Read content of file to get the reason
			pkgs, err := ioutil.ReadFile(rebootRequiredPkgs)
			if err != nil {
			} else {
				pkgsString := strings.Replace(string(pkgs), "\n", ", ", -1)
				nrpeStatus.Message = fmt.Sprintf("System needs reboot! Reboot pending for %.0f hours. Responsible packages: %s", rebootRequiredFileAge.Hours(), pkgsString)
			}
		}
	}
}
