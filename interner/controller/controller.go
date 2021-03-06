package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Jangwooo/2022Hackathon/interner/domain/object/request"
	"github.com/Jangwooo/2022Hackathon/interner/domain/object/response"
	"github.com/Jangwooo/2022Hackathon/interner/middleware"
	"github.com/Jangwooo/2022Hackathon/interner/service"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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

	r.Use(cors.Middleware(cors.Config{
		ValidateHeaders: false,
		Origins:         "*",
		RequestHeaders: "Origin, Authorization, Content-Type, Referer, Accept, User-Agent, Accept-Encoding, " +
			"Accept-Language, Cache-Control, Connection, Host, Pragma, Sec-Fetch-Mode",
		ExposedHeaders: "",
		Methods:        "GET, POST",
		MaxAge:         50 * time.Second,
		Credentials:    true,
	}))

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

	r.Static("/image", os.Getenv("image_root"))

	u := r.Group("/user")
	{
		u.POST("/sign_up", ctl.SignUp)
		u.POST("/sign_in", ctl.SignIn)
	}

	p := r.Group("/post")
	{
		p.POST("/", middleware.Auth, ctl.CreatePost)
		p.GET("/", ctl.GetPosts)
		p.GET("/:post_id", ctl.GetPost)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// SignIn
// @Summary ?????????
// @Name sign In
// @Router /user/sign_in [POST]
// @Tags User
// @Accept json
// @Produce json
// @Param json body request.SingIn true "?????? ?????????, ????????????"
// @Success 200 {object} response.SingIn
// @failure 400 {object} response.Error
// @failure 500 {object} response.Error "???????????? ?????? ????????? ???????????? ??????.."
func (Controller) SignIn(c *gin.Context) {
	req := request.SingIn{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Massage: err.Error()})
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

// SignUp
// @Summary ????????????
// @Name sign up
// @Router /user/sign_up [POST]
// @Tags User
// @Accept json
// @Produce json
// @Param json body request.SignUp true "?????? ??????"
// @Success 201
// @failure 400 {object} response.Error
// @failure 500 {object} response.Error "???????????? ?????? ????????? ???????????? ??????.."
func (Controller) SignUp(c *gin.Context) {
	req := request.SignUp{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Massage: err.Error()})
		return
	}

	err := service.SignUp(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error{Massage: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.Message{Message: "success"})
}

// CreatePost
// @Summary ??? ??????
// @Name Create post
// @Router /post [POST]
// @Tags Post
// @Accept json
// @Produce json
// @Param json body request.CreatePost true "??? ??????"
// @Param Authorization header string true "API ??????"
// @Success 201 "??? ?????? ??????"
// @failure 400 {object} response.Error
// @failure 403 {object} response.Error "api ?????? ???????????? ??????"
// @failure 500 {object} response.Error "???????????? ?????? ????????? ???????????? ??????.."
func (Controller) CreatePost(c *gin.Context) {
	req := request.CreatePost{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Massage: err.Error()})
		return
	}

	err := service.CreatePost(req, c.GetString("userID"))

	switch err {
	case nil:
		c.JSON(http.StatusCreated, response.Message{Message: "success"})
	default:
		log.Panic(err.Error())
	}
}

// GetPosts
// @Summary ????????? ?????? ????????????
// @Name Get posts
// @Router /post [GET]
// @Tags Post
// @Accept json
// @Produce json
// @Param category query string false "1 - ????????????. 2 - ????????????. 3 - ????????????. ???????????? ?????? ?????? ??? ????????????"
// @Success 200 {object} response.Posts
// @failure 400 {object} response.Error
// @failure 500 {object} response.Error "???????????? ?????? ????????? ???????????? ??????.."
func (Controller) GetPosts(c *gin.Context) {
	req := request.GetPosts{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Massage: err.Error()})
		return
	}

	res := service.GetPosts(req)
	c.JSON(200, res)
}

// GetPost
// @Summary ????????? ???????????? ????????????
// @Name Get post
// @Router /post/:post_id [GET]
// @Tags Post
// @Accept json
// @Produce json
// @Param post_id path string true "????????? ID"
// @Success 200 {object} response.Post
// @failure 400 {object} response.Error
// @failure 500 {object} response.Error "???????????? ?????? ????????? ???????????? ??????.."
func (Controller) GetPost(c *gin.Context) {
	res := service.GetPost(c.Param("post_id"))
	c.JSON(200, res)
}
