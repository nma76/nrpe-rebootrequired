package main

type ReturnStatus int

// All available return codes
const (
	OK       ReturnStatus = 0
	WARNING  ReturnStatus = 1
	CRITICAL ReturnStatus = 2
	UNKNOWN  ReturnStatus = 3
)

// Return value struct
type NrpeStatus struct {
	Message string
	Code    ReturnStatus
}
