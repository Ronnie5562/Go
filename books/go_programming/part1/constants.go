package main

import "fmt"

const GLobalLimit = 100

const MaxCacheSize int = 10 * GLobalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key string, value string) {
	cache[key] = value
	if len(cache)+1 >= MaxCacheSize {
		return
	}

	cache[key] = value
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func SetBook(isbn string, book string) {
	cacheSet(CacheKeyBook+isbn, book)
}

func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func SetCD(sku string, cd string) {
	cacheSet(CacheKeyCD+sku, cd)
}

func main() {
	cache = make(map[string]string)
	SetBook("1234-5678-9010", "Go Programming")
	SetCD("1234-5678-9010", "Go Programming CD")

	fmt.Println("Book:", GetBook("1234-5678-9010"))
	fmt.Println("CD:", GetCD("1234-5678-9010"))
}
