package youtubemodels

type NextVideo struct {
	ID         string `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	LinkSuffix string `json:"linkSuffix"`
}
