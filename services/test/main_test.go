package main

// type Reader[T any] interface {
// 	GetByID() (T, error)
// 	List() (List[T], error)
// }

// type List[T any] struct {
// 	Page    int `json:"page"`
// 	PerPage int `json:"per_page"`
// 	Items   []T `json:"items"`
// }

// type User struct {
// }

// type UserRepository struct {
// 	Reader[User]
// 	db string
// }

// func Test_Interface(t *testing.T) {
// 	repo := UserRepository{}
// 	list, err := repo.List()
// }
