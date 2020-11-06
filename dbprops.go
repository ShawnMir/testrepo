package dbutil

import "fmt"

// Info ....
type Info struct {
	ID   int
	Name string
}

//DB Driver ...
const DbDriver = "mysql"

//DB User ...
const User = "root"

//DB Password ...
const Password = "golangtest"

//DB Name ...
const DbName = "go_db1"

//Table Name ...
const TableName = "person"

// Data Source Name ...
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8", User, Password, DbName)
