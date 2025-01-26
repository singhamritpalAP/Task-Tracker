package relationaldatabase

type RelationalDatabase interface {
	Create(model interface{}, values interface{}) error
	Read(model interface{}, id string) error
	Update(model interface{}, value map[string]interface{}, condition string, args ...interface{}) error
	Delete(model interface{}, id string, condition string, args ...interface{}) error
	ReadAll(models interface{}) error
}
