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

	p := r.Group("/post").Use(middleware.Auth)
	{
		p.POST("/", ctl.CreatePost)
	}

	return r
}

// SignIn
// @Summary 로그인
// @Name sign In
// @Router /user/sign_in [POST]
// @Tags User
// @Accept json
// @Produce json
// @Param json body request.SingIn true "유저 아이디, 비밀번호"
// @Success 200 {object} response.SingIn
// @failure 400 {object} response.Error
// @failure 500 {object} response.Error "서버에서 뭔가 큰일이 일어나고 있음.."
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

// SignUp
// @Summary 회원가입
// @Name sign up
// @Router /user/sign_up [POST]
// @Tags User
// @Accept json
// @Produce json
// @Param json body request.SignUp true "유저 정보"
// @Success 201
// @failure 400 {object} response.Error
// @failure 500 {object} response.Error "서버에서 뭔가 큰일이 일어나고 있음.."
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

// CreatePost
// @Summary 글 작성
// @Name Create post
// @Router /post [POST]
// @Tags Post
// @Accept json
// @Produce json
// @Param json body request.CreatePost true "글 정보"
// @Param Authorization header string true "API 토큰"
// @Success 201 "글 생성 성공"
// @failure 400 {object} response.Error
// @failure 403 {object} response.Error "api 키가 올바르지 않음"
// @failure 500 {object} response.Error "서버에서 뭔가 큰일이 일어나고 있음.."
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
		log.Panic(err.Error())
	}
}

// GetPosts
// @Summary 게시글 목록 가져오기
// @Name Get posts
// @Router /post [GET]
// @Tags Post
// @Accept json
// @Produce json
// @Param category query string false "1 - 벼락치기. 2 - 결정장애. 3 - 설득하기. 아무것도 주지 않을 시 전체검색"
// @Success 200 {object} response.Posts
// @failure 400 {object} response.Error
// @failure 500 {object} response.Error "서버에서 뭔가 큰일이 일어나고 있음.."
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
// @Summary 게시글 상세정보 가져오기
// @Name Get post
// @Router /post/:post_id [GET]
// @Tags Post
// @Accept json
// @Produce json
// @Param post_id path string true "게시글 ID"
// @Success 200 {object} response.Post
// @failure 400 {object} response.Error
// @failure 500 {object} response.Error "서버에서 뭔가 큰일이 일어나고 있음.."
func (Controller) GetPost(c *gin.Context) {
	res := service.GetPost(c.Param("post_id"))
	c.JSON(200, res)
}
