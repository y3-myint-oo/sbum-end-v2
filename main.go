package main

import (
	router "ymo.me/sbum-end-v2/routers"
	driver "ymo.me/sbum-end-v2/utils"
)

func main() {
	//fmt.Println("Hello World")
	driver.InitDriver()
	router.InitConnection()
}
