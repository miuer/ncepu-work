package model

type SearchResult struct {
	Title string `json:"title"`
	Href  string `json:"href"`
	Host  string `json:"host"`
}

type NovelChapter struct {
	Name string `json:"name"`
}

type NovelContent struct {
	Name string `json:"name"`
}
