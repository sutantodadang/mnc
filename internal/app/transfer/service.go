package transfer

import (
	"context"
	"errors"
	"math/big"
	"mnc/infrastructure"
	"mnc/internal/app/kafka"
	"mnc/internal/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type ITransferService interface {
	MakeTransfer(ctx context.Context, req MakeTransferRequest) (data MakeTransferResponse, err error)
	ReceivedTransfer(ctx context.Context, msg kafka.TransferMessage) (err error)
}

type TransferService struct {
	repo repositories.Querier
	kp   *kafka.KafkaProducer
}

// ReceivedTransfer implements ITransferService.
func (p *TransferService) ReceivedTransfer(ctx context.Context, msg kafka.TransferMessage) (err error) {

	userData, err := p.repo.SelectOneUserById(ctx, pgtype.UUID{Bytes: uuid.MustParse(msg.To), Valid: true})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	balance, err := userData.BalanceAmount.Float64Value()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	newBalance := balance.Float64 + msg.Amount

	err = p.repo.InsertBalanceHistory(ctx, repositories.InsertBalanceHistoryParams{
		BalanceID:           userData.BalanceID,
		TransactionID:       pgtype.UUID{Bytes: uuid.MustParse(msg.TransactionID), Valid: true},
		BalanceAmountBefore: userData.BalanceAmount,
		BalanceAmountAfter:  pgtype.Numeric{Int: big.NewInt(int64(newBalance)), Valid: true},
	})

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = p.repo.UpdateBalance(ctx, repositories.UpdateBalanceParams{
		BalanceID:     userData.BalanceID,
		BalanceAmount: pgtype.Numeric{Int: big.NewInt(int64(newBalance)), Valid: true},
	})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return

}

// MakeTransfer implements IProductService.
func (p *TransferService) MakeTransfer(ctx context.Context, req MakeTransferRequest) (data MakeTransferResponse, err error) {

	userData, err := p.repo.SelectOneUserById(ctx, pgtype.UUID{Bytes: uuid.MustParse(req.UserId), Valid: true})
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

	transfer, err := p.repo.InsertTransfer(ctx, repositories.InsertTransferParams{
		SourceUserID:   pgtype.UUID{Bytes: uuid.MustParse(req.UserId), Valid: true},
		TargetUserID:   pgtype.UUID{Bytes: uuid.MustParse(req.TargetUser), Valid: true},
		Remarks:        req.Remarks,
		TransferAmount: pgtype.Numeric{Int: big.NewInt(int64(req.Amount)), Valid: true},
	},
	)

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	trx, err := p.repo.InsertTransaction(ctx, repositories.InsertTransactionParams{
		UserID:          transfer.SourceUserID,
		Status:          "success",
		TransactionType: repositories.TypeTransactionDEBIT,
		SourceID:        transfer.TransferID,
		SourceType:      repositories.TypeSourceTRANSFER,
	})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	newBalance := balance.Float64 - req.Amount

	err = p.repo.InsertBalanceHistory(ctx, repositories.InsertBalanceHistoryParams{
		BalanceID:           userData.BalanceID,
		TransactionID:       trx.TransactionID,
		BalanceAmountBefore: userData.BalanceAmount,
		BalanceAmountAfter:  pgtype.Numeric{Int: big.NewInt(int64(newBalance)), Valid: true},
	})

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = p.repo.UpdateBalance(ctx, repositories.UpdateBalanceParams{
		BalanceID:     userData.BalanceID,
		BalanceAmount: pgtype.Numeric{Int: big.NewInt(int64(newBalance)), Valid: true},
	})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	data = MakeTransferResponse{
		TransferId:    transfer.TransferID.String(),
		Amount:        req.Amount,
		Remarks:       req.Remarks,
		BalanceBefore: balance.Float64,
		BalanceAfter:  newBalance,
		CreatedAt:     transfer.CreatedAt.Time.String(),
	}

	msg := kafka.TransferMessage{
		TransferID:    transfer.TransferID.String(),
		TransactionID: trx.TransactionID.String(),
		Amount:        req.Amount,
		Remarks:       req.Remarks,
		Status:        "success",
		From:          userData.UserID.String(),
		To:            req.TargetUser,
	}

	err = p.kp.PublishTransfer(msg)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return
}

func NewTransferService(repo repositories.Querier, kpCfg *infrastructure.KafkaConfig) ITransferService {
	return &TransferService{
		repo: repo,
		kp:   kafka.NewKafkaProducer(kpCfg.Producer),
	}
}
