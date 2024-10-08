package entity

import "errors"

var (
	ErrAlreadyExists         = errors.New("already exists")
	ErrMissingToken          = errors.New("missing auth token")
	ErrInvalidToken          = errors.New("invalid token")
	ErrTokenExpired          = errors.New("token expired")
	ErrUsersPasswordNotMatch = errors.New("users password not match")
	ErrNotFound              = errors.New("not found")
	ErrUserNotFound          = errors.New("user not found")
	ErrNoData                = errors.New("no data")
	ErrOrderConflict         = errors.New("order already registered another user")
	ErrNotEnoughBonuses      = errors.New("not enough bonuses on balance")

	ErrThirdPartyOrderNotRegistered = errors.New("third party order not registered")
	ErrThirdPartyToManyRequests     = errors.New("third party to many requests")
	ErrThirdPartyInternal           = errors.New("third party internal error")
)
