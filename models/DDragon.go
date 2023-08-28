package models

type Champions struct {
	Type    string              `json:"type"`
	Format  string              `json:"format"`
	Version string              `json:"version"`
	Data    map[string]Champion `json:"data"`
}

type Champion struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Blurb string `json:"blurb"`
}
