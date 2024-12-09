package odt

//Структура для передачи данных аунтификации
type UserOdt struct {
	Email string
	Pass  string
}

//Структура для передачи результатов аукциона
type FinalAuction struct {
	Winner string
	Other  []string
}