package main


type MainItem struct {
	Id int
	Title string
	Pic string
	Url string
}
type MainData struct {
	Date string
	MainPages []MainItem
}


type FinalData struct {
	MainData []MainData
	PageData []string
}
