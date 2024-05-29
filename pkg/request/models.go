package request

type Request struct {
	Method          string
	Path            string
	ProtocolVersion string
	Headers         map[string]string
	Body            string
}
