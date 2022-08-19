package v1

import (
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
	"strconv"
)
// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code := model.CreateArt(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArticleList(c *gin.Context){
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code,total := model.GetArticleList(pageSize, pageNum)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
			"total" : total,
		},
	)
}

//查询单个文章
func GetArticleInfo(c *gin.Context){

	id,_ := strconv.Atoi(c.Param("id"))
	data,code := model.GetArticleInfo(id)
	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"msg" : errmsg.GetErrMsg(code),
		"data" : data,
	})
}
//根据分类查询文章
func GetArticleByCategory(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("cid"))
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNum,_ := strconv.Atoi(c.Query("pagenum"))
	data,code,total := model.GetArticleByCategory(id,pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"total" : total,
		"msg" : errmsg.GetErrMsg(code),
	})
}
// EditArt 编辑文章
func EditArt(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.EditArt(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteArt 删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
