package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/git-avilabs/clash-giveaway-api/clashofclans"
)

func (server *Server) GetClanInfo(c *gin.Context) {
	clanTag := c.Param("clanTag")

	if clanTag == "" {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, Response{Status: http.StatusNotAcceptable, Message: ERROR, Data: ErrClanTagRequired.Error()})
		return
	}

	clan, err := server.ClashOfClansApi.GetClanInfo(clanTag)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Message: ERROR, Data: fmt.Sprintf("%s: %s", ErrFailedToGetClanInfo.Error(), err)})
		return
	}

	c.JSON(http.StatusOK, Response{Status: http.StatusOK, Message: SUCCESS, Data: clan})
}

func (server *Server) getEligibleMembers(clanTag string) ([]Member, error) {
	clan, err := server.ClashOfClansApi.GetClanInfo(clanTag)

	if err != nil {
		return nil, err
	}

	var eligibleMembers []Member

	for x := 0; x < len(clan.MemberList); x++ {
		if clan.MemberList[x].Donations > 1000 {
			newMember := Member{
				Name: clan.MemberList[x].Name,
				Tag:  clan.MemberList[x].Tag,
			}

			eligibleMembers = append(eligibleMembers, newMember)
		}
	}

	return eligibleMembers, nil
}

func (server *Server) GetEligibleMembers(c *gin.Context) {
	clanTag := c.Param("clanTag")

	if clanTag == "" {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, Response{Status: http.StatusNotAcceptable, Message: ERROR, Data: ErrClanTagRequired.Error()})
		return
	}

	eligibleMembers, err := server.getEligibleMembers(clanTag)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Message: ERROR, Data: fmt.Sprintf("%s: %s", ErrFailedToGetEligibleMembers.Error(), err)})
		return
	}

	c.JSON(http.StatusOK, Response{Status: http.StatusOK, Message: SUCCESS, Data: EligibleMembers{Members: len(eligibleMembers), MemberList: eligibleMembers}})
}

func (server *Server) GetWinner(c *gin.Context) {
	clanTag := c.Param("clanTag")

	if clanTag == "" {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, Response{Status: http.StatusNotAcceptable, Message: ERROR, Data: ErrClanTagRequired.Error()})
		return
	}

	goldPass, err := server.ClashOfClansApi.GetGoldPassInfo()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Message: ERROR, Data: fmt.Sprintf("%s: %s", ErrFailedToGetGoldPassInfo.Error(), err)})
		return
	}

	goldPassEndTime, err := time.Parse(clashofclans.TimeFormat, goldPass.EndTime)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Message: ERROR, Data: fmt.Sprintf("%s: %s", ErrFailedToGetGoldPassInfo.Error(), err)})
		return
	}

	year, month, date := goldPassEndTime.AddDate(0, 0, -2).Date()
	now_year, now_month, now_date := time.Now().Date()

	if now_year != year || now_month != month || now_date != date {
		c.AbortWithStatusJSON(http.StatusForbidden, Response{Status: http.StatusForbidden, Message: ERROR, Data: fmt.Sprintf("Gold pass winner can only be generated 2 days before gold pass season ends. Be back on: %v-%v-%v", date, month, year)})
		return
	}

	eligibleMembers, err := server.getEligibleMembers(clanTag)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Message: ERROR, Data: fmt.Sprintf("%s: %s", ErrFailedToGetEligibleMembers.Error(), err)})
		return
	}

	randomIndex := rand.Intn(len(eligibleMembers))
	winner := eligibleMembers[randomIndex]

	c.JSON(http.StatusOK, Response{Status: http.StatusOK, Message: SUCCESS, Data: winner})
}
