package web

import (
	"fmt"
	"github.com/GZIY/gbook/gbook/internal/domain"
	"github.com/GZIY/gbook/gbook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`

	userIdKey = "userId"
	bizLogin  = "login"
)

type UserHandler struct {
	svc              *service.UserService
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc:              svc,
		emailRegexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

// 注册 用户功能的 路由
func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", u.SignUp)

	ug.POST("/login", u.Login)

	ug.POST("/edit", u.Edit)

	ug.GET("/profile", u.Profile)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	// 接收请求
	var req SignUpReq

	// bind 方法根据 Content-Type 解析请求数据到 req里
	// 解析错了，就会直接写回一个 400 错误
	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v", err)
		return
	}

	// 正则校验邮箱
	isEmail, err := u.emailRegexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "邮箱不正确")
		return
	}

	// 校验两次密码
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次输入的密码不相同")
		return
	}

	// 正则校验密码
	isPassword, err := u.passwordRegexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isPassword {
		ctx.String(http.StatusOK,
			"密码必须包含数字、特殊字符，并且长度不能小于 8 位")
		return
	}

	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
	}

	ctx.String(http.StatusOK, "hello,注册成功")
}

func (u *UserHandler) Login(ctx *gin.Context) {

}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello")
}
