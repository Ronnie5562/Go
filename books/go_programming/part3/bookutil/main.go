package main

import (
	"fmt"
	"bookutil/author"
)

func main() {
	// Create an author instance
	author := author.NewAuthor("John Doe", "johndoe@example.com")
	fmt.Println(author)

	chapterTitle := "Introduction to Go Modules"
	chapterContent := "Go modules provide a structured way to manage dependencies and improve code maintainability"

	author.WriteChapter(chapterTitle, chapterContent)
	author.ReviewChapter(chapterTitle, "This chapter looks great, but let's add some more examples.")
	author.PublishChapter(chapterTitle, chapterContent)
}
