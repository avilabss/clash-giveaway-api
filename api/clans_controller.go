package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	chocolateclashgoapi "github.com/git-avilabs/chocolate-clash-go-api"
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

func (server *Server) getEligibleMembers(clanTag string) (*[]Member, error) {
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

	return &eligibleMembers, nil
}

func (server *Server) filterEligibleMembersBasedOnAttacks(eligibleMembers []Member) (*[]Member, error) {
	var newEligibleMembers []Member

	for x := 0; x < len(eligibleMembers); x++ {
		api, err := chocolateclashgoapi.Init(chocolateclashgoapi.FWALeague)

		if err != nil {
			return nil, err
		}

		member, err := api.GetMember(eligibleMembers[x].Tag, 0, 20, true)

		if err != nil {
			return nil, err
		}

		var eligibleAttacks []chocolateclashgoapi.Attack

		for x := 0; x < len(member.Attacks); x++ {
			fullTimeStr := fmt.Sprintf("%s %s", member.Attacks[x].Timestamp, "00:00:00")
			fullTime, _ := time.Parse("2006-01-02 15:04:05", fullTimeStr)

			nowTime := time.Now().UTC()
			xDaysBefore := nowTime.AddDate(0, 0, -30)

			if fullTime.After(xDaysBefore) {
				eligibleAttacks = append(eligibleAttacks, member.Attacks[x])
			}
		}

		isEligible := true

		for x := 0; x < len(eligibleAttacks); x++ {
			color := *eligibleAttacks[x].Color

			if color == "purple" || color == "red" {
				isEligible = false
				break
			}
		}

		if isEligible {
			newEligibleMembers = append(newEligibleMembers, eligibleMembers[x])
		}
	}

	return &newEligibleMembers, nil
}

func (server *Server) GetEligibleMembers(c *gin.Context) {
	clanTag := c.Param("clanTag")
	verifyAttacks := c.Query("verifyAttacks")

	if clanTag == "" {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, Response{Status: http.StatusNotAcceptable, Message: ERROR, Data: ErrClanTagRequired.Error()})
		return
	}

	eligibleMembersPointer, err := server.getEligibleMembers(clanTag)
	eligibleMembers := *eligibleMembersPointer

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Message: ERROR, Data: fmt.Sprintf("%s: %s", ErrFailedToGetEligibleMembers.Error(), err)})
		return
	}

	if strings.ToLower(verifyAttacks) == "true" {
		eligibleMembersPointer, err = server.filterEligibleMembersBasedOnAttacks(eligibleMembers)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Message: ERROR, Data: fmt.Sprintf("%s: %s", ErrFailedToGetEligibleMembersByVerifyingAttacks.Error(), err)})
			return
		}

		eligibleMembers = *eligibleMembersPointer
	}

	c.JSON(http.StatusOK, Response{Status: http.StatusOK, Message: SUCCESS, Data: EligibleMembers{Members: len(eligibleMembers), MemberList: eligibleMembers}})
}

func (server *Server) GetWinner(c *gin.Context) {
	clanTag := c.Param("clanTag")
	verifyAttacks := c.Query("verifyAttacks")

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

	eligibleMembersPointer, err := server.getEligibleMembers(clanTag)
	eligibleMembers := *eligibleMembersPointer

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Message: ERROR, Data: fmt.Sprintf("%s: %s", ErrFailedToGetEligibleMembers.Error(), err)})
		return
	}

	if strings.ToLower(verifyAttacks) == "true" {
		eligibleMembersPointer, err = server.filterEligibleMembersBasedOnAttacks(eligibleMembers)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, Response{Status: http.StatusInternalServerError, Message: ERROR, Data: fmt.Sprintf("%s: %s", ErrFailedToGetEligibleMembersByVerifyingAttacks.Error(), err)})
			return
		}

		eligibleMembers = *eligibleMembersPointer
	}

	if len(eligibleMembers) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, Response{Status: http.StatusNotFound, Message: ERROR, Data: fmt.Sprintf("%s: %s", ErrNoMemberEligible.Error(), err)})
		return
	}

	randomIndex := rand.Intn(len(eligibleMembers))
	winner := eligibleMembers[randomIndex]

	c.JSON(http.StatusOK, Response{Status: http.StatusOK, Message: SUCCESS, Data: winner})
}
