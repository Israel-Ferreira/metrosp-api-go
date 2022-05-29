package exceptions

import "errors"

var ErrorValidation = errors.New("erro ao validar o objeto")
var ErrorNotFound = errors.New("erro: id não encontrado")

var ErrorLineExists = errors.New("erro: Linha já cadastrada")
