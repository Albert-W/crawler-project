package view

import (
	"html/template"
	"io"
	"crawler/frontend/view/model"
)

//包装view
type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView{
	return SearchResultView{
		template:template.Must(template.ParseFiles(filename)),
	}
}

//执行写操作
func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error{
	return s.template.Execute(w, data)
}
