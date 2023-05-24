package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/git-avilabs/clash-giveaway-api/api"
	"github.com/git-avilabs/clash-giveaway-api/clashofclans"
	"github.com/git-avilabs/clash-giveaway-api/utils"
	"github.com/stretchr/testify/require"
)

func TestGetClanInfo(t *testing.T) {
	env, err := utils.LoadEnv(".")

	if err != nil {
		log.Fatal(fmt.Errorf("%s: %s", api.ErrFailedToLoadEnv.Error(), err))
	}

	clashOfClansApi := clashofclans.Api{
		BaseUrl: env.ClashApiBaseUri,
		JWT:     env.ClashApiJwt,
	}

	_, err = clashOfClansApi.GetClanInfo("#8QCCY8U")

	require.NoError(t, err)
}

func TestGetGoldPassInfo(t *testing.T) {
	env, err := utils.LoadEnv(".")

	if err != nil {
		log.Fatal(fmt.Errorf("%s: %s", api.ErrFailedToLoadEnv.Error(), err))
	}

	clashOfClansApi := clashofclans.Api{
		BaseUrl: env.ClashApiBaseUri,
		JWT:     env.ClashApiJwt,
	}

	goldPass, err := clashOfClansApi.GetGoldPassInfo()

	require.NoError(t, err)

	goldPassEndsOn, err := time.Parse(clashofclans.TimeFormat, goldPass.EndTime)

	require.NoError(t, err)

	year, month, date := goldPassEndsOn.AddDate(0, 0, -2).Date()
	now_year, now_month, now_date := time.Now().Date()

	if now_year == year && now_month == month && now_date == date {
		t.Log("Today is the day!")
	} else {
		t.Logf("Winner will be decided on: %v-%v-%v", date, month, year)
	}
}
