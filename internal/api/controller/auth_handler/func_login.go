package auth_handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niceBackend/common/global"
	"niceBackend/common/transform/request"
	"niceBackend/common/transform/response"
	"niceBackend/internal/middleware"
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/repository/db_repo/user_repo"
	"time"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (h *handler) Login() core.HandlerFunc {
	//var l request.Login
	//_ = c.ShouldBindJSON(&l)
	//if err := pkg.Verify(l, pkg.LoginVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//account := l.Account
	//password := l.Password
	//u := &model.User{Account: account, Password: password}
	//if err, user := service.Login(u); err != nil {
	//	global.NiceLog.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
	//	response.FailWithCode(4002, c)
	//} else {
	//	// 颁发token
	//	tokenNext(c, *user)
	//}
	var result map[string]interface{}
	result = make(map[string]interface{})
	var data [1]int
	result["data"] = data
	return func(c core.Context) {}

}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user user_repo.User) {
	j := &middleware.JWT{SigningKey: []byte(global.NiceConfig.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		Account:    user.Account,
		BufferTime: global.NiceConfig.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.NiceConfig.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                              // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.NiceLog.Error("获取token失败!", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
	return
}