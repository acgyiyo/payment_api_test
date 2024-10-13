package repository

import (
	"context"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/gateway"
	"gorm.io/gorm"
)

type paymentStore struct {
	db *gorm.DB
}

func NewPaymentDataBase(db *gorm.DB) gateway.PaymentStore {
	return &paymentStore{
		db: db,
	}
}

func (r *paymentStore) SavePayment(ctx context.Context, payment *entity.Payment) error {
	if err := r.db.Create(payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *paymentStore) SearchPaymentByTransactionID(ctx context.Context, transactionID string) (*entity.Payment, error) {
	var payment entity.Payment
	if err := r.db.WithContext(ctx).Where("transaction_id = ?", transactionID).Take(&payment).Error; err != nil {
		return &entity.Payment{}, err
	}
	return &payment, nil
}

func (r *paymentStore) UpdatePayment(ctx context.Context, payment *entity.Payment) error {
	if err := r.db.Save(payment).Error; err != nil {
		return err
	}
	return nil
}
