package store

//Store ...
type Store interface {
	Product() ProductRepo
	Order() OrderRepo
}
