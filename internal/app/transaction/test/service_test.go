package transaction_test

import (
	"context"
	"errors"
	"math/big"
	"mnc/internal/app/transaction"
	"mnc/internal/repositories"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) InsertBalanceHistory(ctx context.Context, arg repositories.InsertBalanceHistoryParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockRepo) InsertDefaultBalance(ctx context.Context, userID pgtype.UUID) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func (m *MockRepo) InsertPayment(ctx context.Context, arg repositories.InsertPaymentParams) (repositories.Payment, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repositories.Payment), args.Error(1)
}

func (m *MockRepo) InsertTopUp(ctx context.Context, arg repositories.InsertTopUpParams) (repositories.Topup, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repositories.Topup), args.Error(1)
}

func (m *MockRepo) InsertTransaction(ctx context.Context, arg repositories.InsertTransactionParams) (repositories.Transaction, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repositories.Transaction), args.Error(1)
}

func (m *MockRepo) InsertTransfer(ctx context.Context, arg repositories.InsertTransferParams) (repositories.Transfer, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repositories.Transfer), args.Error(1)
}

func (m *MockRepo) InsertUser(ctx context.Context, arg repositories.InsertUserParams) (repositories.User, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repositories.User), args.Error(1)
}

func (m *MockRepo) SelectOneUserById(ctx context.Context, userID pgtype.UUID) (repositories.SelectOneUserByIdRow, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(repositories.SelectOneUserByIdRow), args.Error(1)
}

func (m *MockRepo) SelectOneUserByPhoneNumber(ctx context.Context, phoneNumber string) (repositories.SelectOneUserByPhoneNumberRow, error) {
	args := m.Called(ctx, phoneNumber)
	return args.Get(0).(repositories.SelectOneUserByPhoneNumberRow), args.Error(1)
}

func (m *MockRepo) SelectTransactionByUserId(ctx context.Context, userID pgtype.UUID) ([]repositories.SelectTransactionByUserIdRow, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]repositories.SelectTransactionByUserIdRow), args.Error(1)
}

func (m *MockRepo) UpdateBalance(ctx context.Context, arg repositories.UpdateBalanceParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockRepo) UpdateUser(ctx context.Context, arg repositories.UpdateUserParams) (repositories.User, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repositories.User), args.Error(1)
}

func TestTransactionReport(t *testing.T) {
	mockRepo := new(MockRepo)
	service := transaction.NewTransactionService(mockRepo)
	ctx := context.Background()
	userID := uuid.New()
	trxID := uuid.New()
	srcID := uuid.New()

	mockTransactions := []repositories.SelectTransactionByUserIdRow{
		{TransactionID: pgtype.UUID{Bytes: trxID, Valid: true},
			UserID:          pgtype.UUID{Bytes: userID, Valid: true},
			Status:          "success",
			TransactionType: repositories.TypeTransactionDEBIT,
			SourceID:        pgtype.UUID{Bytes: srcID, Valid: true},
			SourceType:      repositories.TypeSourceTOPUP,
			CreatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
			TopUpAmount:     pgtype.Numeric{Int: big.NewInt(200), Exp: 0, Valid: true},
		},
	}

	mockRepo.On("SelectTransactionByUserId", ctx, pgtype.UUID{Bytes: userID, Valid: true}).Return(mockTransactions, nil)

	result, err := service.TransactionReport(ctx, userID.String())

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, 200.0, result[0].Amount)
	assert.Equal(t, "DEBIT", result[0].TransactionType)

	mockRepo.AssertExpectations(t)
}

func TestTransactionReport_ErrorFetchingData(t *testing.T) {
	mockRepo := new(MockRepo)
	service := transaction.NewTransactionService(mockRepo)
	ctx := context.Background()
	userID := uuid.New()

	mockRepo.On("SelectTransactionByUserId", ctx, pgtype.UUID{Bytes: userID, Valid: true}).Return([]repositories.SelectTransactionByUserIdRow{}, errors.New("database error"))

	result, err := service.TransactionReport(ctx, userID.String())

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
