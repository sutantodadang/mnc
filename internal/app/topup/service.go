package topup

import (
	"context"
	"math/big"
	"mnc/internal/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type ITopUpService interface {
	AddTopUp(ctx context.Context, req AddTopup) (resp AddTopupResponse, err error)
}

type TopUpService struct {
	repo repositories.Querier
}

// AddWarehouse implements IWarehouseService.
func (t *TopUpService) AddTopUp(ctx context.Context, req AddTopup) (resp AddTopupResponse, err error) {

	data, err := t.repo.InsertTopUp(ctx, repositories.InsertTopUpParams{
		UserID:      pgtype.UUID{Bytes: uuid.MustParse(req.Userid), Valid: true},
		TopUpAmount: pgtype.Numeric{Int: big.NewInt(int64(req.Amount)), Valid: true},
	})

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	trx, err := t.repo.InsertTransaction(ctx, repositories.InsertTransactionParams{
		UserID:          data.UserID,
		Status:          "success",
		TransactionType: repositories.TypeTransactionCREDIT,
		SourceID:        data.TopUpID,
		SourceType:      repositories.TypeSourceTOPUP,
	},
	)

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	userData, err := t.repo.SelectOneUserById(ctx, trx.UserID)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	topUpVal, err := data.TopUpAmount.Float64Value()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	balanceVal, err := userData.BalanceAmount.Float64Value()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	balanceAfter := balanceVal.Float64 + topUpVal.Float64

	err = t.repo.InsertBalanceHistory(ctx, repositories.InsertBalanceHistoryParams{
		BalanceID:           userData.BalanceID,
		TransactionID:       trx.TransactionID,
		BalanceAmountBefore: userData.BalanceAmount,
		BalanceAmountAfter: pgtype.Numeric{
			Int:   big.NewInt(int64(balanceAfter)),
			Valid: true,
		},
	},
	)

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = t.repo.UpdateBalance(ctx, repositories.UpdateBalanceParams{
		BalanceID: userData.BalanceID,
		BalanceAmount: pgtype.Numeric{
			Int:   big.NewInt(int64(balanceAfter)),
			Valid: true,
		},
	})

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	resp = AddTopupResponse{
		TopUpId:       data.TopUpID.String(),
		AmountTopup:   req.Amount,
		BalanceBefore: balanceVal.Float64,
		BalanceAfter:  balanceAfter,
		CreatedAt:     data.CreatedAt.Time.String(),
	}

	return

}

func NewTopUpService(repo repositories.Querier) ITopUpService {
	return &TopUpService{
		repo: repo,
	}
}
