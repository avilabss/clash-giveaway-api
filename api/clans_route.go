package api

import "github.com/gin-gonic/gin"

func (server *Server) ClansRoute(router *gin.RouterGroup) {
	clans := router.Group("/clans")

	clans.GET(":clanTag", server.GetClanInfo)
	clans.GET(":clanTag/eligible", server.GetEligibleMembers)
	clans.GET(":clanTag/winner", server.GetWinner)
}
