package user

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/service/user"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "手机号必须11位")
		return
	}
	if len(password) < 8 {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "密码必须大于8位")
		return
	}
	log.Println(name, telephone, password)

	//判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "手机号已注册")
		return
	}

	//创建用户
	hashdPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //密码hash化
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashdPassword),
	}
	if err := db.Create(&newUser).Error; err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, err.Error())
		return
	}

	response.Success(ctx, nil, "注册成功")

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	return user.ID != 0
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据校验
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "手机号必须11位")
		return
	}
	if len(password) < 8 {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "密码必须大于8位")
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 442, nil, "手机号不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "密码错误")
		return
	}

	//发送token
	token, err := common.GetToken(user)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统token获取异常")
		return
	}

	response.Success(ctx, gin.H{"token": token}, "登录成功")

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": user}, "")
}

func List(ctx *gin.Context) {
	if userObj, exists := ctx.Get("user"); exists {
		if userDto, ok := userObj.(model.User); ok { //类型断言 检查userObj变量是否是dto.UserDto类型。
			// 现在可以安全地使用userDto.Telephone等字段
			tel := userDto.Telephone
			if tel != "15572261989" { //说明不是管理员 这里应该给加个字段 获取当前token，根据唯一标识来判断是否是管理员
				response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
				return
			}
			//获取用户的list
			//userList=
			userService := user.Service{}
			users, err := userService.GetAllUsers()
			if err != nil {
				response.Response(ctx, http.StatusInternalServerError, 500, nil, "内部服务器错误")
				return
			}
			response.Success(ctx, gin.H{
				"list": users,
			}, "获取用户列表成功")
		}
	}
}
