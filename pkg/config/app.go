package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


/* 
var (

) 

means the variable become global package means it can be used any where in the same package.
*/

var (
	db *gorm.DB
)

func Connect() {
	d,err := gorm.Open("mysql", "root:root@/books")
	if err != nil {
		panic(err)
	}
	db = d
}



func GetDB() *gorm.DB {
	return db
}