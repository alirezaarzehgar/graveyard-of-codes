package storage

import "github.com/BaseMax/JokeGoServiceAPI/internal/book"

type Storage interface {
	Create(book.Book) error
}
