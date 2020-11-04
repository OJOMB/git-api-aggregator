package github

// ErrorResponse is the Data structure for an error response.
type ErrorResponse struct {
	StatusCode       int     `json:"status_code"`
	Message          string  `json:"message"`
	DocumentationURL string  `json:"documentation_url"`
	Errors           []Error `json:"errors"`
}

// Error is the Data structure for a Github error
type Error struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
