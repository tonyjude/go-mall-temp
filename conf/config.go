package conf

import (
	"go-mall-temp/dao"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode  string
	HttpPort string

	DB         string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	// 本地读取环境变量

	file, err := ini.Load("../conf/config.ini")
	if err != nil {
		panic(err)
	}

	LoadServer(file)
	LoadMysql(file)
	LoadRedis(file)
	LoadPhotoPath(file)

	//mysql 读 (80%)
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	//mysql 写 (20%)
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")

	dao.Database(pathRead, pathWrite)

}

func LoadServer(file *ini.File) {

	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()

}

func LoadMysql(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func LoadPhotoPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}
