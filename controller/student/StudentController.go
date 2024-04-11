package student

import (
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/service/student"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	if userObj, exists := ctx.Get("user"); exists {
		if userDto, ok := userObj.(model.User); ok { //类型断言 检查userObj变量是否是dto.UserDto类型。
			// 现在可以安全地使用userDto.Telephone等字段
			tel := userDto.Telephone
			if tel != "15572261989" { //说明不是管理员 这里应该给加个字段 获取当前token，根据唯一标识来判断是否是管理员
				response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
				return
			}
			//获取student
			studentService := student.Service{}
			students, err := studentService.GetStudentListByID(userDto.ID)
			if err == nil {
				response.Success(ctx, gin.H{
					"studentList": students,
				}, "成功查询")

			}
		}
	}
}

type UserForm struct {
	Name      string
	StudentId string //学生卡号
	Teacher   string //user的id
	Grade     string //年级 1 2 3 4 5
}

func Add(ctx *gin.Context) {

	name := ctx.PostForm("name")
	studentId := ctx.PostForm("studentId")
	//teacher := ctx.PostForm("teacher") 这个字段不用写，会根据解析的token获取当前添加的人
	grade := ctx.PostForm("grade")

	//逻辑表单判断
	if len(name) == 0 || len(studentId) == 0 || len(grade) == 0 {
		response.Fail(ctx, gin.H{}, "表单填充错误，存在空白字段")
	}

	if userObj, exists := ctx.Get("user"); exists {
		if userDto, ok := userObj.(model.User); ok { //类型断言 检查userObj变量是否是dto.UserDto类型。
			newStudent := model.Student{StudentId: studentId, Name: name,
				Grade: grade, Teacher: strconv.Itoa(int(userDto.ID))}
			studentService := student.Service{}
			err := studentService.Add(newStudent)
			if err != nil {
				response.Fail(ctx, gin.H{}, err.Error())
				return
			}
			response.Success(ctx, gin.H{}, "添加成功")

		}
	}
}
