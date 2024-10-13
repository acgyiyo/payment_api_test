package usecase_test

import (
	"context"
	"testing"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/usecase"
	"github.com/stretchr/testify/assert"
)

func TestRetrievePayment_SearchPaymentByTransactionID_Success(t *testing.T) {
	ctx := context.TODO()
	payment := entity.Payment{TransactionID: "txn-001", Amount: 100.00, Status: "success"}

	mockStore := new(MockPaymentStore)

	// Setup expectations
	mockStore.On("SearchPaymentByTransactionID", ctx, "txn-001").Return(&payment, nil)

	// Create the use case
	retrievePayment := usecase.NewRetrievePayment(mockStore)

	// Execute the test
	resp, err := retrievePayment.SearchPaymentByTransactionID(ctx, "txn-001")

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, "txn-001", resp.TransactionID)
	mockStore.AssertExpectations(t)
}
