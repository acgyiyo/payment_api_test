package usecase_test

import (
	"context"
	"testing"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/usecase"
	"github.com/stretchr/testify/assert"
)

func TestRegisterPayment_SavePayment_Success(t *testing.T) {
	ctx := context.TODO()
	payment := entity.Payment{TransactionID: "txn-001", Amount: 100.00, Status: "success"}

	mockBank := new(MockBankService)
	mockStore := new(MockPaymentStore)

	// Setup expectations
	mockBank.On("ProcessPaymentInBank", &payment).Return(&payment, nil)
	mockStore.On("SavePayment", ctx, &payment).Return(nil)

	// Create the use case
	registerPayment := usecase.NewRegisterPayment(mockStore, mockBank)

	// Execute the test
	resp, err := registerPayment.SavePayment(ctx, payment)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, "txn-001", resp.TransactionID)
	mockBank.AssertExpectations(t)
	mockStore.AssertExpectations(t)
}
