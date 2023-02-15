package core

type BaseResponse struct {
	Object   string `json:"object"`
	LiveMode bool   `json:"livemode"`
}

type ListResponse struct {
	Object  string        `json:"object"`
	HasMore bool          `json:"has_more"`
	Data    []interface{} `json:"data"`
}

type DeleteResponse struct {
	Object  string `json:"object"`
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}
