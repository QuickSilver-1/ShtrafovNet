package service

// MockPaymentService представляет моковую реализацию платежного сервиса
type MockPaymentService struct{}

// Payment выполняет моковую платежную транзакцию
func (m *MockPaymentService) Payment(transaction int, amount float64) error {
    // В данной реализации всегда возвращаем nil, имитируя успешную транзакцию
    return nil
}