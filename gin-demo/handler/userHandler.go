package handler

// 此handler可以理解是controller
import (
	"fmt"
	"gin-demo/utils"
	"os"
	"path/filepath"
	"strconv"
	"time"

	// "fmt"
	"log"

	// "fmt"
	"gin-demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 可以理解是一个action
func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

// 用户查询
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.Query("age")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}

// 注册
func UserRegister(context *gin.Context) {
	// email := context.PostForm("email")
	// password := context.DefaultPostForm("password", "Wa123456")
	// passwordAgain := context.DefaultPostForm("password-again", "Wa123456")
	// fmt.Println("email", email, "password", password, "password again", passwordAgain)

	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		return
	}

	id := user.Save()
	log.Println("id is ", id)
	// context.Redirect(http.StatusMovedPermanently, "/")
	// fmt.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
}

// 用户登录
func UserLogin(context *gin.Context) {
	var user model.UserModel
	if e := context.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}
	u := user.QueryByEmail()
	if u.Password == user.Password {
		context.String(http.StatusOK, "登录成功"+u.Email)
	} else {
		context.String(http.StatusOK, "登录失败")
	}
}

// 用户上传页面
func UserProfile(context *gin.Context) {
	id := context.Query("id")
	var user model.UserModel
	i, err := strconv.Atoi(id) // 将string转int
	u, e := user.QueryById(i)
	if e != nil || err != nil {
		// 将变量信息渲染给模板
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	context.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})
}

// 用户上传操作
func UpdateUserProfile(context *gin.Context) {
	var user model.UserModel
	// 将参数信息绑定到user上面
	if err := context.ShouldBind(&user); err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		log.Panicln("绑定发生错误 ", err.Error())
	}
	// file表单
	file, e := context.FormFile("avatar-file")
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("文件上传错误", e.Error())
	}

	// 拼路径并创建路径
	path := utils.RootPath()
	path = filepath.Join(path, "avatar")
	e = os.MkdirAll(path, os.ModePerm)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}

	// 拼文件名
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	fmt.Println(fileName)

	// 保存文件
	e = context.SaveUploadedFile(file, path+"/"+fileName)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}

	// 保存
	avatarUrl := "http://localhost:9000/avatar/" + fileName
	user.Avatar = avatarUrl
	e = user.Update(user.Id)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("数据无法更新", e.Error())
	}

	// 跳转  将int转string 才能使用+
	context.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(user.Id))
}
