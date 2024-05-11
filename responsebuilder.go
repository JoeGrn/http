package main

func NewResponse() *Response {
	return &Response{
		Headers: make(map[string]string),
	}
}

func (r *Response) SetStatus(status string) {
	r.Status = status
}

func (r *Response) SetProtocol(protocol string) {
	r.Protocol = protocol
}

func (r *Response) SetHeader(key, value string) {
	r.Headers[key] = value
}

func (r *Response) SetBody(body string) {
	r.Body = body
}

func (r *Response) String() string {
	res := r.Protocol + r.Status + SEPARATOR
	for k, v := range r.Headers {
		res += k + ": " + v + SEPARATOR
	}
	res += SEPARATOR + r.Body
	return res
}
