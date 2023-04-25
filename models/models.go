package models

type UnsplashResponse struct {
	Description string `json:"description"`
	URLs        `json:"urls"`
}

type URLs struct {
	Regular string `json:"regular"`
}
