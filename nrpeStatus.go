package main

//ReturnStatus is a custom type used to mimic a enumeration
type ReturnStatus int

// All available return codes
const (
	OK       ReturnStatus = 0
	WARNING  ReturnStatus = 1
	CRITICAL ReturnStatus = 2
	UNKNOWN  ReturnStatus = 3
)

//NrpeStatus is the return value
type NrpeStatus struct {
	Message string
	Code    ReturnStatus
}
