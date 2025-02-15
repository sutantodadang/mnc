package payment

import (
	"context"
	"errors"
	"math/big"
	"mnc/internal/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type IPaymentService interface {
	MakePayment(ctx context.Context, req MakePaymentRequest) (data MakePaymentResponse, err error)
}

type PaymentService struct {
	repo repositories.Querier
}

// MakePayment implements IPaymentService.
func (r *PaymentService) MakePayment(ctx context.Context, req MakePaymentRequest) (data MakePaymentResponse, err error) {

	userData, err := r.repo.SelectOneUserById(ctx, pgtype.UUID{Bytes: uuid.MustParse(req.UserID), Valid: true})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	balance, err := userData.BalanceAmount.Float64Value()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	if balance.Float64 < req.Amount {
		err = errors.New("balance is not enough")
		return

	}

	pay, err := r.repo.InsertPayment(ctx, repositories.InsertPaymentParams{
		UserID:        pgtype.UUID{Bytes: uuid.MustParse(req.UserID), Valid: true},
		Remarks:       req.Remarks,
		PaymentAmount: pgtype.Numeric{Int: big.NewInt(int64(req.Amount)), Valid: true},
	},
	)

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	trx, err := r.repo.InsertTransaction(ctx, repositories.InsertTransactionParams{
		UserID:          pay.UserID,
		Status:          "success",
		TransactionType: repositories.TypeTransactionDEBIT,
		SourceID:        pay.PaymentID,
		SourceType:      repositories.TypeSourcePAYMENT,
	})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	newBalance := balance.Float64 - req.Amount

	err = r.repo.InsertBalanceHistory(ctx, repositories.InsertBalanceHistoryParams{
		BalanceID:           userData.BalanceID,
		TransactionID:       trx.TransactionID,
		BalanceAmountBefore: userData.BalanceAmount,
		BalanceAmountAfter:  pgtype.Numeric{Int: big.NewInt(int64(newBalance)), Valid: true},
	})

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = r.repo.UpdateBalance(ctx, repositories.UpdateBalanceParams{
		BalanceID:     userData.BalanceID,
		BalanceAmount: pgtype.Numeric{Int: big.NewInt(int64(newBalance)), Valid: true},
	})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	data = MakePaymentResponse{
		PaymentId:     pay.PaymentID.String(),
		Amount:        req.Amount,
		Remarks:       req.Remarks,
		BalanceBefore: balance.Float64,
		BalanceAfter:  newBalance,
		CreatedAt:     pay.CreatedAt.Time.String(),
	}

	return

}

func NewPaymentService(repo repositories.Querier) IPaymentService {
	return &PaymentService{
		repo: repo,
	}
}
