package prepare

import (
	"math/rand"

	"github.com/wada1001/algorithm/src/pkg/util"
)

type Book struct {
	Id int
	Title string
	Author string
	IsAvailable bool
}

func MakeRandomBook() *Book {
	return &Book{
		Id : rand.Int(),
		Title: util.MakeRandomStr(5),
		Author: util.MakeRandomStr(5),
		IsAvailable: util.MakeRandomBool(),
	}
}

func MakeRandomBookArr(length int) []*Book {
	books := make([]*Book, length)
	i := 0
	for i < length {
		books = append(books, MakeRandomBook())
	}
	
	return books
}