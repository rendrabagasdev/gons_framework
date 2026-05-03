package models

var ModelRegistry []interface{}

func RegisterModel(model interface{}) {
	ModelRegistry = append(ModelRegistry, model)
}
