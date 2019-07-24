package main

import (
	"lt_go_web/config"
	"lt_go_web/model"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/spf13/viper"
	"lt_go_web/router"
)

func main() {
	if err := config.Init();err != nil {
		panic(err)
	}

	if err := model.Init();err != nil {
		panic(err)
	}
	//g := gin.Default()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	router.InitRouter(g)
	log.Printf("Start to listening the incoming requests on http address: %s\n", viper.GetString("addr"))
	//log.Println(http.ListenAndServe(viper.GetString("addr"), g).Error())
	if err := g.Run(viper.GetString("addr"));err != nil {log.Fatal("ListenAndServe:", err)}

}
