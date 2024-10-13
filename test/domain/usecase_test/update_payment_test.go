package usecase_test

import (
	"context"
	"testing"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/usecase"
	"github.com/stretchr/testify/assert"
)

func TestUpdatePayment_UpdatePayment_Success(t *testing.T) {
	ctx := context.TODO()
	payment := entity.Payment{TransactionID: "txn-001", Amount: 100.00, Status: "success"}

	mockBank := new(MockBankService)
	mockStore := new(MockPaymentStore)

	// Setup expectations
	mockStore.On("SearchPaymentByTransactionID", ctx, "txn-001").Return(&payment, nil)
	mockBank.On("ProcessRefundInBank", &payment).Return(&payment, nil)
	mockStore.On("UpdatePayment", ctx, &payment).Return(nil)

	// Create the use case
	updatePayment := usecase.NewUpdatePayment(mockStore, mockBank)

	// Execute the test
	resp, err := updatePayment.UpdatePayment(ctx, payment)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, "txn-001", resp.TransactionID)
	mockBank.AssertExpectations(t)
	mockStore.AssertExpectations(t)
}
