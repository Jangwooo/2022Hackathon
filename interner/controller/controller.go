package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Jangwooo/2022Hackathon/interner/domain/object/request"
	"github.com/Jangwooo/2022Hackathon/interner/domain/object/response"
	"github.com/Jangwooo/2022Hackathon/interner/middleware"
	"github.com/Jangwooo/2022Hackathon/interner/service"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func SetUp() *gin.Engine {
	var f *os.File
	var err error

	if f, err = os.OpenFile(".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	ctl := Controller{}
	r := gin.New()

	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Massage: fmt.Sprint(err)})
	}))
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		var body map[string]interface{}
		err = json.NewDecoder(params.Request.Body).Decode(&body)

		return fmt.Sprintf("[GIN] %v | %v | %v | %v | %v | %v |\nHeader: %v \nBody:   %v \n",
			params.TimeStamp.Format("2006-01-02:15:04:05"), params.StatusCode, params.Latency, params.ClientIP,
			params.Method, params.Path, params.Request.Header, body)
	}))

	u := r.Group("/user")
	{
		u.POST("/sign_up", ctl.SignUp)
		u.POST("/sign_in", ctl.SignIn)
	}

	p := r.Group("/post").Use(middleware.Auth)
	{
		p.POST("/", ctl.CreatePost)
	}

	return r
}

func (Controller) SignIn(c *gin.Context) {
	req := request.SingIn{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := service.SignIn(req)

	switch err {
	case nil:
		c.JSON(http.StatusOK, res)
	case service.ErrNoMatchingUsers, service.ErrPasswordMismatch:
		c.JSON(http.StatusBadRequest, response.Error{Massage: err.Error()})
	default:
		log.Panic("unknown error")
	}
}

func (Controller) SignUp(c *gin.Context) {
	req := request.SignUp{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := service.SignUp(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error{Massage: err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (Controller) CreatePost(c *gin.Context) {
	req := request.CreatePost{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := service.CreatePost(req, c.GetString("userID"))

	switch err {
	case nil:
		c.Status(http.StatusCreated)
	default:
		log.Panic("unknown error")
	}
}
