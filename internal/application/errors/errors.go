// Пакет для кастомных ошибок слоя приложения
package errors

type InvalidPassword struct {
	Code int
	Err  string
}

func (e *InvalidPassword) Error() string {
	return e.Err
}

type JWTError struct {
	Code int
	Err  string
}

func (e *JWTError) Error() string {
	return e.Err
}