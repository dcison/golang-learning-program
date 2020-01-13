package main

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
	"github.com/goinggo/mapstructure"
	"github.com/joho/godotenv"
)

type User struct {
	Host     string `db:"host"`
	User     string `db:"user"`
	Password string `db:"password"`
}

func initMariadbConfig() (string, string) {

	configErr := godotenv.Load()
	if configErr != nil {
		fmt.Println("读取环境配置出错", configErr)
		return "", ""
	}

	dbCfg := make(map[string]string)
	dbCfg["dbtype"] = os.Getenv("DB_CONNECTION")
	dbCfg["host"] = os.Getenv("MATRIX_DB1_HOST")
	dbCfg["user"] = os.Getenv("MATRIX_DB1_USER")
	dbCfg["pass"] = os.Getenv("MATRIX_DB1_PASS")
	dbCfg["port"] = os.Getenv("MATRIX_DB1_PORT")
	dbCfg["database"] = os.Getenv("DB_DATABASE")

	cfg := fmt.Sprintf("%s:%s@(%s:%s)/%s", dbCfg["user"], dbCfg["pass"], dbCfg["host"], dbCfg["port"], dbCfg["database"])

	return dbCfg["dbtype"], cfg
}

func initDBbyGorose() (*gorose.Engin, error) {
	dbtype, dbcfg := initMariadbConfig()
	engin, err := gorose.Open(&gorose.Config{Driver: dbtype, Dsn: dbcfg})
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("数据库连接出错")
	}
	return engin, nil
}

var engin, err = initDBbyGorose()

func main() {
	var result string
	gorosesql := "user,host,password"
	result, err := queryByGorose(gorosesql)
	if err != nil {
		fmt.Println("数据库连接出错")
	}
	fmt.Println(result)
}

func db() gorose.IOrm {
	return engin.NewOrm()
}

func queryByGorose(order string) (string, error) {
	res, e := db().Table("user").Fields(order).Get()
	if e != nil {
		fmt.Println("数据库查询语句出错")
	}
	var user User
	_res := []User{}
	for _, data := range res {
		//将 map 转换为指定的结构体
		if err := mapstructure.Decode(data, &user); err != nil {
			fmt.Println(err)
			return "", err
		}
		_res = append(_res, user)
	}
	data, err := json.Marshal(_res)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
