package main

// 结构体标签在JSON中的应用

type Movie struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Price int      `json:"price"`
	Actor []string `json:"actor"`
}
