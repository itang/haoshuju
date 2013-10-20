package main

import (
	"log"
	"os"
	"strconv"

	"github.com/astaxie/beego"
	//"github.com/itang/gotang"
	"github.com/itang/haoshuju/haoshuju.com/controllers"
)

func main() {
	appBootstrap()

	beego.Router("/", &controllers.MainController{})
	beego.Run()
}

func appBootstrap() {
	tryOverrideVarFromEnv(&beego.RunMode, "RunMode")
	tryOverrideVarFromEnv(&beego.HttpPort, "HttpPort")

	displayInfo()
}

func checkError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}

func tryOverrideVarFromEnv(target interface{}, key string) {
	v := os.Getenv(key)
	if v != "" {
		switch t := target.(type) {
		case *string:
			log.Printf("Try Override %s From Env: %s => %s", key, *t, v)
			*t = v
		case *int:
			log.Printf("Override %s From Env: %d => %s", key, *t, v)
			i, err := strconv.Atoi(v)
			checkError(err, "Can't atoi")
			*t = i
		default:
			panic("Only support these types: *int, *string")
		}
	}
}

func displayInfo() {
	log.Println("RunMode:", beego.RunMode)
	log.Println("HttpPort:", beego.HttpPort)
}
