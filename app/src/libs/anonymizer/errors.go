package anonymizer

import "errors"

var (
	ErrRecordNotStruct  = errors.New("o registro não é uma struct")
	ErrMethodNotAllowed = errors.New("o método fornecido não é permitido")
	ErrMaskEmpty        = errors.New("a máscara não pode estar vazia")
)

func ErrFieldNotFound(fieldName string) error {
	return errors.New("campo " + fieldName + " não encontrado")
}

func ErrFieldNotString(fieldName string) error {
	return errors.New("campo " + fieldName + " não é uma string")
}
