package main

import (
	"crawler/frontend/controller"
	"net/http"
)

func main() {
	//文章展示的功能，只要不是search，就拿目录文件
	http.Handle("/",http.FileServer(http.Dir("src/crawler/frontend/view")))


	http.Handle("/search",controller.CreateSearchResultHandler("src/crawler/frontend/view/template.html"))

	//http.Handle("/search",controller.SearchResultHandler{})
	err := http.ListenAndServe(":8888", nil)
	if err !=nil{
		panic(err)
	}
}
