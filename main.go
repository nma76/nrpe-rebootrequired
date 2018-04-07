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
	var rebootRequiredFile = "/var/run/reboot-required"

	if fileInfo, err := os.Stat(rebootRequiredFile); os.IsNotExist(err) {
		nrpeStatus.Code = OK
		nrpeStatus.Message = "Reboot not required."
	} else {
		//Check age of reboot-required. If older than a day, set status to CRITICAL, else set status to WARNING
		var rebootRequiredFileAge = time.Now().Sub(fileInfo.ModTime())
		if rebootRequiredFileAge.Hours() >= 48 {
			nrpeStatus.Code = CRITICAL
		} else {
			nrpeStatus.Code = WARNING
		}

		//The file exists, system needs reboot.
		nrpeStatus.Message = "System needs reboot!"

		//Get the reason why by reading the contents of reboot-required.pkgs
		var rebootRequiredPkgs = rebootRequiredFile + ".pkgs"
		if _, err := os.Stat(rebootRequiredPkgs); os.IsNotExist(err) {
			//File didn't exist, reason is unknown
			nrpeStatus.Message = "System needs reboot! Responsible packages: unknown"
		} else {
			//Read content of file to get the reason
			pkgs, err := ioutil.ReadFile(rebootRequiredPkgs)
			if err != nil {
			} else {
				pkgsString := strings.Replace(string(pkgs), "\n", ", ", -1)
				nrpeStatus.Message = fmt.Sprintf("System needs reboot! Responsible packages: %s", pkgsString)
			}
		}
	}
}
