package api

import (
	"github.com/chaowen112/gin-lib/prometheus"

	"github.com/chaowen112/gin-template/app/api/handlers/validation"

	v1 "github.com/chaowen112/gin-template/app/api/handlers/v1"

	"github.com/chaowen112/gin-template/pkg/validators"

	"github.com/chaowen112/gin-lib/middlewares"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	pprof.Register(router)

	version := "0.0.0"

	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(version),
		middlewares.Jsonifier(version),
	)

	router.NoRoute(middlewares.Default404)

	v, err := validators.Init()
	if err != nil {
		log.WithError(err).Fatal("init_validators_failed")
	}
	if err = validation.RegisterCustomValidation(v); err != nil {
		log.WithError(err).Fatal("init_validators_custom_failed")
	}

	// Expose prometheus metric.
	p := prometheus.NewPrometheus()
	p.ReqCntURLLabelMappingFn = prometheus.PromURLLabelMappingFn
	p.Use(router)

	router.LoadHTMLGlob("templates/*")

	api := router.Group("/apis/")
	// v1 API endpoints.
	v1API := api.Group("/v1")
	v1API.GET("/", v1.ImplementMe)
	return &Router{
		router,
	}
}

func (r *Router) Run() {
	err := r.Engine.Run()
	if err != nil {
		log.WithError(err).Fatal("router_start_failed")
	}
}

func Run() {
	NewRouter().Run()
}
