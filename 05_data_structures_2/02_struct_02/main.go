package main

import (
	"fmt"
	"go4noobs/query"
)

func main() {
	sql := query.Table("user")

	// SELECT * from user
	fmt.Printf("%s\n", sql.Get())

	sql = query.Table("user")

	// SELECT * from user WHERE name = 'Rafael Breno' AND age > '20'%
	fmt.Printf("%s", sql.Where("name", "=", "Rafael Breno").Where("age", ">", "20").Get())
}
