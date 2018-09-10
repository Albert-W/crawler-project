# Crawler-website
It's crawler website using Go language.

This is the mainPage
![image](https://raw.githubusercontent.com/Albert-W/crawler-website/master/image/MainPage.jpg)

This is the homePage
![image](https://raw.githubusercontent.com/Albert-W/crawler-website/master/image/HomePage.jpg)

## Features
- Go language
- Docker
- Elastic Search
- MVC pattern
- Microservices
- Singleton -> Concurrent -> Distribute

## Installation and go package
- go language
- docker
- elasticsearch
- go get golang.org/x/text
- go get -v github.com/gpmgo/gopm
- gopm get -g -v golang.org/x/text
- gopm get -g -v golang.org/x/net/html
- go get gopkg.in/olivere/elastic.v5

## Usage for Singleton 
- Start Docker.
- Run Script "docker run -d -p 9200:9200 elasticsearch"
- Run "src/crawler/main.go", to start the singleton crawler.
- Run "src/crawler/frontend/starter.go", to view the result in the website.
- Visit "http://localhost:8888/" in your browser
- Type in query string with REST format. such as "女 && Age>20"

## Usage for Concurrent
- Start Docker.
- Run Script "docker run -d -p 9200:9200 elasticsearch"
- run "src/crawler_distributed/persist/server/ItemSaver.go"
- run "src/crawler_distributed/worker/server/worker.go"
- run "src/crawler_distributed/main.go"
- Run "src/crawler/frontend/starter.go", to view the result in the website.
- Visit "http://localhost:8888/" in your browser
- Type in query string with REST format. such as "男 && 已购车"

## Reference
- Google资深工程师深度讲解Go语言 @ https://coding.imooc.com/class/180.html
