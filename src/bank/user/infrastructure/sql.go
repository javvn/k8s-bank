package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jawncorp.com/bank/user/utils"
)

type (
	DB struct {
		Host     string
		Username string
		Password string
		DBName   string
		Connect  *gorm.DB
	}

	SQLHandler struct {
		Conn *gorm.DB
	}
)

func NewSQLHandler() *SQLHandler {
	c := utils.NewConfigDB()

	USER := c.MySQL.Username
	PASS := c.MySQL.Password
	PROTOCOL := c.MySQL.Protocol
	DBNAME := c.MySQL.DBName

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FSeoul"
	conn, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		// panic(err.Error())
		fmt.Print("Error: No database connection established.")
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

func (handler SQLHandler) Scan(out interface{}, order string) interface{} {
	return handler.Conn.Order(order).Find(out)
}

func (handler *SQLHandler) Where(out interface{}, query interface{}, args ...interface{}) interface{} {
	if query == "" {
		return handler.Conn.Find(out)
	}

	return handler.Conn.Where(query, args).Find(out)
}

func (handler *SQLHandler) Count(out *int, model interface{}, query interface{}, args ...interface{}) interface{} {
	out64 := int64(*out)
	return handler.Conn.Model(model).Where(query, args).Count(&out64)
}

func (handler *SQLHandler) Create(input interface{}) (result interface{}, err error) {
	fmt.Println(input)
	response := handler.Conn.Create(input)

	return nil, response.Error
}

func (handler *SQLHandler) Update(out interface{}, value interface{}, query interface{}, args ...interface{}) interface{} {
	return handler.Conn.Model(out).Where(query, args).Updates(value)
}
