package response

type Response struct {
	Protocol string
	Status   string
	Headers  map[string]string
	Body     string
}
