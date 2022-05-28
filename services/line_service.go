package services


type LineService interface {
	FindAll() ([]any, error)
	FindById(id string) (any, error)
	Create(any) (any, error)
	Delete(id string) error
	Update(string, any) error
}
