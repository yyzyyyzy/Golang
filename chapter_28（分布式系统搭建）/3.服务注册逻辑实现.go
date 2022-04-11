package main

type ServiceName string

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

const (
	LogService = ServiceName("LogService")
)
