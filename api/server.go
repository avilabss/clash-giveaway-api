package api

import (
	"github.com/gin-gonic/gin"
	"github.com/git-avilabs/clash-giveaway-api/clashofclans"
	"github.com/git-avilabs/clash-giveaway-api/utils"
)

type Server struct {
	Env             *utils.Env
	Router          *gin.Engine
	ClashOfClansApi *clashofclans.Api
}

func NewServer(env *utils.Env) (*Server, error) {
	// Init clash of clans api
	clashOfClansApi := clashofclans.Api{
		BaseUrl: env.ClashApiBaseUri,
		JWT:     env.ClashApiJwt,
	}

	// Init server
	server := Server{
		Env:             env,
		ClashOfClansApi: &clashOfClansApi,
	}

	server.SetupRouter()

	return &server, nil
}

func (server *Server) SetupRouter() {
	router := gin.Default()
	router.Use(server.CORS())

	server.Router = router

	// API routes
	apiRouter := server.Router.Group("/api")

	// v1 API routes
	v1ApiRouter := apiRouter.Group("/v1")

	server.ClansRoute(v1ApiRouter)
}
