package user

import (
	"context"
	"errors"
	"mnc/internal/repositories"
	"mnc/internal/utils"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	RegisterUser(ctx context.Context, req RegisterUserRequest) (data RegisterUserResponse, err error)
	LoginUser(ctx context.Context, req LoginUserRequest) (data LoginUserResponse, err error)
	UpdateUser(ctx context.Context, req UpdateUserRequest) (data UpdateUserResponse, err error)
}

type UserService struct {
	repo repositories.Querier
}

// UpdateUser implements IUserService.
func (u *UserService) UpdateUser(ctx context.Context, req UpdateUserRequest) (data UpdateUserResponse, err error) {

	currentUser, err := u.repo.SelectOneUserById(ctx, pgtype.UUID{Bytes: uuid.MustParse(req.UserID), Valid: true})
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Error().Err(err).Send()
		return
	}

	if req.LastName == "" {
		req.LastName = currentUser.LastName.String
	}

	result, err := u.repo.UpdateUser(ctx, repositories.UpdateUserParams{
		FirstName: req.FirstName,
		LastName:  pgtype.Text{String: req.LastName, Valid: true},
		Address:   req.Address,
		UserID:    pgtype.UUID{Bytes: uuid.MustParse(req.UserID), Valid: true},
	})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	data = UpdateUserResponse{
		UserID:    result.UserID.String(),
		FirstName: result.FirstName,
		LastName:  result.LastName.String,
		Address:   result.Address,
		UpdatedAt: result.UpdatedAt.Time.String(),
	}

	return

}

// RegisterUser implements IUserService.
func (u *UserService) RegisterUser(ctx context.Context, req RegisterUserRequest) (data RegisterUserResponse, err error) {

	currentUser, err := u.repo.SelectOneUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Error().Err(err).Send()
		return
	}

	if currentUser.UserID.Valid {
		err = errors.New("phone number already registered")
		return
	}

	hashPin, err := bcrypt.GenerateFromPassword([]byte(req.Pin), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	dataUser, err := u.repo.InsertUser(ctx, repositories.InsertUserParams{
		FirstName:   req.FirstName,
		LastName:    pgtype.Text{String: req.LastName, Valid: true},
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Pin:         string(hashPin),
	},
	)

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	data = RegisterUserResponse{
		UserID:      dataUser.UserID.String(),
		FirstName:   dataUser.FirstName,
		LastName:    dataUser.LastName.String,
		PhoneNumber: dataUser.PhoneNumber,
		Address:     dataUser.Address,
		CreatedAt:   dataUser.CreatedAt.Time.String(),
	}

	err = u.repo.InsertDefaultBalance(ctx, dataUser.UserID)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return
}

// LoginUser implements IUserService.
func (u *UserService) LoginUser(ctx context.Context, req LoginUserRequest) (data LoginUserResponse, err error) {

	currentUser, err := u.repo.SelectOneUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentUser.Pin), []byte(req.Pin))
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	accessToken, err := utils.GenerateJwt(currentUser, time.Now().Add(time.Hour*24).Unix())
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	refreshToken, err := utils.GenerateJwt(currentUser, time.Now().AddDate(0, 0, 7).Unix())
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	data = LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}

func NewUserService(repo repositories.Querier) IUserService {
	return &UserService{
		repo: repo,
	}
}
