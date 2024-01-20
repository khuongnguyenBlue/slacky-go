package interfaces

type IDbHandler interface {
	Create(model interface{}) (error)
	FindByField(model interface{}, field string, value string) (error)
}