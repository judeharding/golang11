package config

// GORM = A GOLANG Object-Relational Mapper (ORM) is a code
// library that automates the transfer of data stored in
// relational database tables into objects that are more
// commonly used in application code.
// GORM = A GOLANG Object-Relational Mapper (ORM) is a code
// library that automates the transfer of data stored in
// relational database tables into objects that are more
// commonly used in application code.

import (
	"github.com/jinzhu/gorm"                  // orm pkg
	_ "github.com/jinzhu/gorm/dialects/mysql" // NEEDS the blank space at the front of the line
)

var (
	db * gorm.DB
)


// opens the connection to the db 
func Connect(){   	
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")  // d == db
	// if there is an error  - ps -- panic is a keyword
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
} 