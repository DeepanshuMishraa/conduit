package types

type JobRequest struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type JobResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
