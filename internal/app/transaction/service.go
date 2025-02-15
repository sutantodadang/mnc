package transaction

import (
	"context"
	"mnc/internal/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type ITransactionService interface {
	TransactionReport(ctx context.Context, userId string) (data []TransactionReportResponse, err error)
}

type TransactionService struct {
	repo repositories.Querier
}

// TransactionReport implements ITransactionService.
func (t *TransactionService) TransactionReport(ctx context.Context, userId string) (data []TransactionReportResponse, err error) {

	trxData, err := t.repo.SelectTransactionByUserId(ctx, pgtype.UUID{Bytes: uuid.MustParse(userId), Valid: true})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	for _, v := range trxData {

		amountBefore, errF := v.BalanceAmountBefore.Float64Value()
		if errF != nil {
			log.Error().Err(errF).Send()
			err = errF
			return
		}

		amountAfter, errF := v.BalanceAmountAfter.Float64Value()
		if errF != nil {
			log.Error().Err(errF).Send()
			err = errF
			return
		}

		switch v.SourceType {
		case repositories.TypeSourcePAYMENT:

			pay, errF := v.PaymentAmount.Float64Value()
			if errF != nil {
				log.Error().Err(errF).Send()
				err = errF
				return
			}

			data = append(data, TransactionReportResponse{
				PaymentId:       v.PaymentID.String(),
				UserId:          v.UserID.String(),
				Status:          v.Status,
				TransactionType: string(v.TransactionType),
				Amount:          pay.Float64,
				Remarks:         v.PaymentRemarks.String,
				BalanceBefore:   amountBefore.Float64,
				BalanceAfter:    amountAfter.Float64,
				CreatedAt:       v.CreatedAt.Time.String(),
			})

		case repositories.TypeSourceTRANSFER:

			trans, errF := v.TransferAmount.Float64Value()
			if errF != nil {
				log.Error().Err(errF).Send()
				err = errF
				return
			}

			data = append(data, TransactionReportResponse{
				TransferId:      v.TransferID.String(),
				UserId:          v.UserID.String(),
				Status:          v.Status,
				TransactionType: string(v.TransactionType),
				Amount:          trans.Float64,
				Remarks:         v.TransferRemarks.String,
				BalanceBefore:   amountBefore.Float64,
				BalanceAfter:    amountAfter.Float64,
				CreatedAt:       v.CreatedAt.Time.String(),
			})

		case repositories.TypeSourceTOPUP:

			topup, errF := v.TopUpAmount.Float64Value()
			if errF != nil {
				log.Error().Err(errF).Send()
				err = errF
				return
			}

			data = append(data, TransactionReportResponse{
				TopUpId:         v.TopUpID.String(),
				UserId:          v.UserID.String(),
				Status:          v.Status,
				TransactionType: string(v.TransactionType),
				Amount:          topup.Float64,
				Remarks:         "",
				BalanceBefore:   amountBefore.Float64,
				BalanceAfter:    amountAfter.Float64,
				CreatedAt:       v.CreatedAt.Time.String(),
			})

		}
	}

	return

}

func NewTransactionService(repo repositories.Querier) ITransactionService {
	return &TransactionService{
		repo: repo,
	}
}
