package main

type Request struct {
	method          string
	path            string
	protocolVersion string
	headers         map[string]string
}
