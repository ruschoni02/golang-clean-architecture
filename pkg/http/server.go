package http

import (
	"github.com/golang-clean-architecture/app/adapters"
	"github.com/golang-clean-architecture/pkg/config"
	"github.com/golang-clean-architecture/pkg/validator"
	"net/http"
	"time"

	"github.com/Depado/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	binding "github.com/gin-gonic/gin/binding"
)

type Endpoint interface {
	Handler(gin.IRouter, *config.Config, *Server) error
}

type Server struct {
	*http.Server
	Router *gin.Engine
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func NewServer(logger adapters.LoggerAdapter, addr string) *Server {
	gin.SetMode(gin.ReleaseMode)
	binding.Validator = validator.New("binding")

	router := gin.New()
	router.Use(ginzap.Ginzap(logger.ZapLogger, time.RFC3339, true))
	router.Use(requestid.New())
	router.Use(cors.Default())

	return &Server{
		Router: router,
		Server: &http.Server{Addr: addr},
	}
}

func (s *Server) Debug() {
	pprof.Register(s.Router)
}

func (s *Server) Load(prefix string, config *config.Config, endpoints ...Endpoint) (gin.IRouter, error) {
	router := s.Router.Group(prefix)
	for _, endpoint := range endpoints {
		err := endpoint.Handler(router, config, s)
		if err != nil {
			return nil, err
		}
	}
	return router, nil
}

func (s *Server) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.Router.Use(middleware...)
}

func (s *Server) Group(path string) gin.IRoutes {
	return s.Router.Group(path)
}

func (s *Server) Prometheus(name, path string) {
	p := ginprom.New(
		ginprom.Engine(s.Router),
		ginprom.Subsystem(name),
		ginprom.Path(path),
	)
	s.Router.Use(p.Instrument())
}

func (s *Server) Start() error {
	s.Router.NoRoute(NotFound)
	s.Server.Handler = s.Router
	return s.Server.ListenAndServe()
}

func (s *Server) Stop() {
	s.Close()
}
