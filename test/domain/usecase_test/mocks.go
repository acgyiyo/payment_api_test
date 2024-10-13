package usecase_test

import (
	"context"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

// Mocking the BankService
type MockBankService struct {
	mock.Mock
}

func (m *MockBankService) ProcessPaymentInBank(payment *entity.Payment) (*entity.Payment, error) {
	args := m.Called(payment)
	return args.Get(0).(*entity.Payment), args.Error(1)
}

func (m *MockBankService) ProcessRefundInBank(refundReq *entity.Payment) (*entity.Payment, error) {
	args := m.Called(refundReq)
	return args.Get(0).(*entity.Payment), args.Error(1)
}

// Mocking the PaymentStore
type MockPaymentStore struct {
	mock.Mock
}

func (m *MockPaymentStore) SavePayment(ctx context.Context, payment *entity.Payment) error {
	return m.Called(ctx, payment).Error(0)
}

func (m *MockPaymentStore) SearchPaymentByTransactionID(ctx context.Context, transactionID string) (*entity.Payment, error) {
	args := m.Called(ctx, transactionID)
	return args.Get(0).(*entity.Payment), args.Error(1)
}

func (m *MockPaymentStore) UpdatePayment(ctx context.Context, payment *entity.Payment) error {
	return m.Called(ctx, payment).Error(0)
}
