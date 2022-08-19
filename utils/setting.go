package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string
	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string
	JwtKey string

	AccessKey  string
	SecretKey string
	Bucket string
	QiNiuSever string
)

func init(){
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件出错",err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiNiu(file)

}
func LoadServer(file *ini.File){
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("sever").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("sever").Key("JwtKey").MustString("eko8817")
}

func LoadData(file *ini.File){

	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginBlog")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("ginBlog")
}

func LoadQiNiu(file *ini.File){
	AccessKey = file.Section("qiNiu").Key("AccessKey").String()
	SecretKey = file.Section("qiNiu").Key("SecretKey").String()
	Bucket = file.Section("qiNiu").Key("Bucket").String()
	QiNiuSever = file.Section("qiNiu").Key("QiNiuSever").String()
}
