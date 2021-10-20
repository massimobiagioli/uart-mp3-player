package models

type SD struct {
	Folders []Folder `json:"folders"`
}

type Folder struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   uint   `json:"year"`
	Songs  []Song `json:"songs"`
}

type Song struct {
	Id       string `json:"id"`
	Filename string `json:"filename"`
	Name     string `json:"name"`
	Author   string `json:"author"`
}
