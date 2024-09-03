package server

import "errors"

var (
	ErrAlreadyExists         = errors.New("already exists")
	ErrInvalidToken          = errors.New("invalid token")
	ErrTokenExpired          = errors.New("token expired")
	ErrUsersPasswordNotMatch = errors.New("users password not match")
	ErrNotFound              = errors.New("not found")
	ErrUserNotFound          = errors.New("user not found")
	ErrOrderConflict         = errors.New("order already registered another user")
	ErrNotEnoughBonuses      = errors.New("not enough bonuses on balance")

	ErrThirdPartyOrderNotRegistered = errors.New("third party order not registered")
	ErrThirdPartyToManyRequests     = errors.New("third party to many requests")
	ErrThirdPartyInternal           = errors.New("third party internal error")
)
