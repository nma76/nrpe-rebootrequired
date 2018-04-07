package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Main function
func main() {
	//Initialize variable that holds status
	var nrpeStatus NrpeStatus
	//Call function that does all work with
	checkNrpe(&nrpeStatus)

	//Print status message and exit with correct code
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

	if _, err := os.Stat(rebootRequiredFile); os.IsNotExist(err) {
		nrpeStatus.Code = OK
		nrpeStatus.Message = "Reboot not required."
	} else {
		//The file exists, system needs reboot.
		nrpeStatus.Message = "System needs reboot!"
		nrpeStatus.Code = WARNING

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
				pkgsString := string(pkgs)
				nrpeStatus.Message = fmt.Sprintf("System needs reboot! Responsible packages: %s", pkgsString)
			}
		}
	}
}
