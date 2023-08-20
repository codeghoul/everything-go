package patterns

import "fmt"

type DB interface {
	Store(string) error
}

type Store struct {
}

func (s *Store) Store(value string) error {
	fmt.Println("storing into db", value)
	return nil
}

func MyExecuteFunc(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println("my ex func", s)
		db.Store(s)
	}
}

// This is coming from a third party lib.
type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("FOO BAR BAZ")
}
