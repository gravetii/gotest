package http

// Response : a typical HTTP response.
type Response struct {
	Message string
	Code    int
}

// Request : a typical HTTP request.
type Request struct {
	Header string
	Body   string
}
