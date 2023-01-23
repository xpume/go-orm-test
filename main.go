package main

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	gORM()
	goFrameORM()
}

type Employee struct {
	EmpNo     int       `gorm:"column:emp_no;type:int(11);primary_key" json:"emp_no"`
	BirthDate time.Time `gorm:"column:birth_date;type:date;NOT NULL" json:"birth_date"`
	FirstName string    `gorm:"column:first_name;type:varchar(14);NOT NULL" json:"first_name"`
	LastName  string    `gorm:"column:last_name;type:varchar(16);NOT NULL" json:"last_name"`
	Gender    string    `gorm:"column:gender;type:enum('M','F');NOT NULL" json:"gender"`
	HireDate  time.Time `gorm:"column:hire_date;type:date;NOT NULL" json:"hire_date"`
}

func gORM() {
	db, err := gorm.Open(mysql.Open("root:123@tcp(localhost:3306)/employees?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	var employees []Employee
	now := time.Now()
	err = db.Limit(6000).Find(&employees).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("GORM: ", time.Since(now))
}

func goFrameORM() {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host: "localhost",
				Port: "3306",
				User: "root",
				Pass: "123",
				Name: "employees",
				Type: "mysql",
			},
		},
	})
	var employees []Employee
	now := time.Now()
	err := g.DB().Model("employees").Limit(6000).Scan(&employees)
	if err != nil {
		panic(err)
	}
	fmt.Println("GoFrameORM: ", time.Since(now))
}
