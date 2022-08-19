package model

import (
	"go_blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid int `gorm:"type:int;not null" json:"cid"`
	Desc string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img string `gorm:"type:varchar(100)" json:"img"`
}

// CreateArt 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

//todo 查询文章列表
func GetArticleList(pagesize int, pagenum int)([]Article,int,int64){
	var articleList []Article
	var total int64
	err = db.Preload("Category").Limit(pagesize).Offset((pagenum - 1) * pagesize).Find(&articleList).Error
	db.Model(&articleList).Count(&total)
	if err != nil {
		return nil,errmsg.ERROR,0
	}
	return articleList,errmsg.SUCCESS,total
}
//查询单个列表的信息
func GetArticleInfo(id int)(Article,int){
	var art Article
	err := db.Preload("Category").Where("id = ?",id).First(&art).Error
	if err != nil {
		return art,errmsg.ERROR_ART_NOT_EXIST
	}
	return art,errmsg.SUCCESS
}
//查询分类下的所有文章
func GetArticleByCategory(id int,pageSize int, pageNum int)([]Article,int,int64){
	var cateArtList []Article
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&cateArtList).Error
	db.Model(&cateArtList).Where("cid =?", id).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}



// EditArt 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// DeleteArt 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}