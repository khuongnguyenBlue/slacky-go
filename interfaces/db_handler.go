package interfaces

type IDbHandler interface {
	Create(model interface{}) (error)
}