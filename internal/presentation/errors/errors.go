// Пакет для кастомных ошибок слоя презентации
package errors

type ServerError struct {
	Code int
	Err  string
}

func (e *ServerError) Error() string {
	return e.Err
}

type NeedAuntification struct {
	Code int
	Err  string
}

func (e *NeedAuntification) Error() string {
	return e.Err
}

type InvalidDate struct {
	Code int
	Err  string
}

func (e *InvalidDate) Error() string {
	return e.Err
}

type RightsError struct {
	Code int
	Err  string
}

func (e *RightsError) Error() string {
	return e.Err
}