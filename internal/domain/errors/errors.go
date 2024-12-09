// Пакет для кастомных ошибок доменного слоя
package errors

type NonValidData struct {
	Code int
	Err  string
}

func (e *NonValidData) Error() string {
	return e.Err
}

type JsonCodeError struct {
	Code int
	Err  string
}

func (e *JsonCodeError) Error() string {
	return e.Err
}
