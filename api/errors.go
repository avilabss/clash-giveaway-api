package api

import "errors"

var (
	ErrFailedToLoadEnv            = errors.New("failed to load env variables")
	ErrFailedToSetupServer        = errors.New("failed to setup server")
	ErrClanTagRequired            = errors.New("clan tag is required")
	ErrFailedToGetClanInfo        = errors.New("failed to get clan info")
	ErrFailedToGetGoldPassInfo    = errors.New("failed to get gold pass info")
	ErrFailedToGetEligibleMembers = errors.New("failed to get eligible members")
)
