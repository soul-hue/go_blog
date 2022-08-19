package v1

import (
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/utils/errmsg"
	"go_blog/utils/validate"
	"net/http"
	"strconv"
)

//查询用户是否存在


//添加用户
func AddUser(c *gin.Context){
//  添加用户
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	 if code == errmsg.SUCCESS {
		model.CreateUser(&data)

	 }

	msg, validCode := validate.Validate(&data)
	if validCode != errmsg.SUCCESS {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}
	 c.JSON(http.StatusOK,gin.H{
		 "status" : "ok",
		 "Data" : data,
		 "msg" : errmsg.GetErrMsg(code),
	 })


}
//查询用户列表
func GetUserList(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNum,_ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data,total := model.GetUsers(pageSize,pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" :data,
		"msg" : errmsg.GetErrMsg(code),
		"total" :total,
	})

}
//编辑用户
func EditUser(c *gin.Context){
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ =c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id,&data)
	}
	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"msg" : errmsg.GetErrMsg(code),
	})

}
//删除用户
func DelUser(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	code := model.DelUser(id)
	c.JSON(http.StatusOK,gin.H{
		"msg" : errmsg.GetErrMsg(code),
		"status" : code,
	})

}
