# What is this?
This is a Book Store Management API, just for learn and experiment golang.

## How to run?
`$ go build && ./book-store`

The server will be listing in port 8080
These are valid urls:
- http://localhost:8000/movies

## What libraries have we used?
In this project we used the following libraries:
- [gorilla/mux](https://github.com/gorilla/mux) - implements a request router and dispatcher for matching incoming requests to their respective handler;
- [GORM](https://github.com/jinzhu/gorm) - Full-Featured ORM
- [go to env](https://github.com/joho/godotenv)

## Commands
`go get github.com/jinzhu/gorm`

`go get github.com/jinzhu/gorm/dialects/mysql`

`go get github.com/gorilla/mux`

`go get github.com/joho/godotenv`