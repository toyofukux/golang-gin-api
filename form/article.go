package form

type ArticleJSON struct {
	Id int `json:id`
	Title string `json:title`
	Content string `json:content`
}
