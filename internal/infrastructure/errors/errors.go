// Пакет для кастомных ошибок слоя инфраструктуры
package errors

type ErrDatabaseConnection struct {
	Code int
	Err  string
}

func (e *ErrDatabaseConnection) Error() string {
	return e.Err
}

type QueryErr struct {
	Code int
	Err  string
}

func (e *QueryErr) Error() string {
	return e.Err
}

type NoDataInBase struct {
	Code int
	Err  string
}

func (e *NoDataInBase) Error() string {
	return e.Err
}

type JsonCodeError struct {
	Code int
	Err  string
}

func (e *JsonCodeError) Error() string {
	return e.Err
}

type DublicateError struct {
	Code int
	Err  string
}

func (e *DublicateError) Error() string {
	return e.Err
}

type NeedMoreBid struct {
	Code int
	Err  string
}

func (e *NeedMoreBid) Error() string {
	return e.Err
}

type NotSingleBid struct {
	Code int
	Err  string
}

func (e *NotSingleBid) Error() string {
	return e.Err
}

type NoMoney struct {
	Code int
	Err  string
}

func (e *NoMoney) Error() string {
	return e.Err
}

type NotificatorError struct {
	Code int
	Err  string
}

func (e *NotificatorError) Error() string {
	return e.Err
}

type InvalidDate struct {
	Code int
	Err  string
}

func (e *InvalidDate) Error() string {
	return e.Err
}