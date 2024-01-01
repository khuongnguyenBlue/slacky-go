package interfaces

type DbHandler interface {
	FindByID(id int) (interface{}, error)
}