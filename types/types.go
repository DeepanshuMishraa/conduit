package types

type JobRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type JobResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
