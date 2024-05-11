package main

type Request struct {
	method          string
	path            string
	protocolVersion string
	headers         map[string]string
}

type Response struct {
	Protocol string
	Status   string
	Headers  map[string]string
	Body     string
}
