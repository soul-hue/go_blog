package model

import (
	"go_blog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null " json:"password" validate:"required,min=4,max=12" label:"密码"`
	Role int `gorm:"type:int " json:"role" validate:"required,gte=2" label:"角色码"`
}
//CheckUser 查询用户是否存在
func CheckUser(name string)(code int){
	var user User
	db.Select("id").Where("username = ?",name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//CreateUser 添加用户
func CreateUser(data *User)int{

	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表
func GetUsers(pageSize int,pageNum int)([]User,int64){
	var users []User
	var total int64
	err := db.Limit(pageSize).Offset((pageNum - 1)*pageSize).Find(&users).Error
	db.Model(&users).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,0
	}
	return users,total
}
//查找用户

//删除用户
func DelUser(id int)int{
	var user User
	err := db.Where("id = ?",id).Delete(&user).Error
	if err!= nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//编辑用户
func EditUser(id int,data *User)int{
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&data).Where("id = ?",id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// BeforeCreate 密码加密&权限控制
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	//u.Role = 2

	return
}

//func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
//	u.Password = ScryptPw(u.Password)
//	return nil
//}

// ScryptPw 生成密码
func ScryptPw(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}

// CheckLogin 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	return user, errmsg.SUCCESS
}

// CheckLoginFront 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESS
}