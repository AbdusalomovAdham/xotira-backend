package gallery

type Create struct {
	Image string `form:"images"`
}

type Delete struct {
	Url string `json:"url"`
}
