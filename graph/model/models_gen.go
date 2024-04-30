// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Books []*Book `json:"books"`
}

type Book struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Price  int       `json:"price"`
	Author []*Author `json:"Author"`
}

type CreateBook struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Mutation struct {
}

type Query struct {
}
