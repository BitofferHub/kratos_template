// routes.go

package interfaces

import (
	"bytes"
	"context"
	"github.com/BitofferHub/pkg/constant"
	engine "github.com/BitofferHub/pkg/middlewares/gin"
	"github.com/BitofferHub/pkg/middlewares/log"
	"github.com/bitstormhub/bitstorm/userX/internal/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

type Handler struct {
	userXService *service.UserXService
}

func NewHandler(s *service.UserXService) *Handler {
	return &Handler{
		userXService: s,
	}
}

func NewRouter(h *Handler) *gin.Engine {
	r := engine.NewEngine(engine.WithLogger(false))
	// 使用gin中间件
	r.Use(InfoLog())
	r.GET("/get_user_info", h.GetUserInfo)
	r.POST("/create_user", h.CreateUserInfo)
	return r
}

// InfoLog
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: gin middleware for log request and reply
//	@return gin.HandlerFunc
func InfoLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		beginTime := time.Now()
		// ***** 1. get request body ****** //
		traceID := c.Request.Header.Get(constant.TraceID)
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close() //  must close
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		// ***** 2. set requestID for goroutine ctx ****** //
		// duration := float64(time.Since(beginTime)) / float64(time.Second)
		ctx := context.WithValue(context.Background(), constant.TraceID, traceID)
		log.InfoContextf(ctx, "ReqPath[%s]-Cost[%v]\n", c.Request.URL.Path, time.Since(beginTime))
	}
}
