package main

import "github.com/athunlal/Redis-Gin-CRUD-Go/pkg/di"

func main() {

	r := di.InitApi()

	r.Run(":8080")

}
