package main

import (
	"assignment3/model"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main(){
	ginEngine:= gin.Default()
	ginEngine.LoadHTMLFiles("template.html")
	
	go updateValue()

	ginEngine.GET("/", getValue)


	err:= ginEngine.Run("localhost:8081")
	if err != nil{
		panic(err)
	}
}

func getValue(c *gin.Context) {
	byte,err:= os.ReadFile("data.json")
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,err.Error())	
		return
	}


	var data model.Data
	err = json.Unmarshal(byte, &data)
	if err!= nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,err.Error())
		return
	}

	var water string
	var wind string
	
	switch {
	case data.Status.Water<= 5:
		water = "aman"
	case data.Status.Water <= 8 && data.Status.Water >= 6:
		water = "siaga"
	default:
		water = "bahaya"
	}

	switch {
	case data.Status.Wind<= 6:
		wind = "aman"
	case data.Status.Wind <= 15 && data.Status.Water >= 7:
		wind = "siaga"
	default:
		wind = "bahaya"
	}
		

	c.HTML(200, "template.html", gin.H{
		"Water": water,
		"Wind": wind,
	})
}

func updateValue() {
	for {
		data := model.Data{
			Status: model.Status{
				Water: rand.Intn(100),
				Wind:  rand.Intn(100),
			},
		}
	
		bytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
	
		err = os.WriteFile("data.json", bytes, 0644)

		if err != nil {
			panic(err)
		}
		
		time.Sleep(15 * time.Second)
	}
}

