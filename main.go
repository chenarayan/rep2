package main

import (
	"os"
	"net/http"
	"github.com/labstack/echo/v4"
	"math/rand"
)

type quote struct {
	Title string
	Author string
}


var quotes []quote = []quote{
	
	{"Learn to Lead", "Sai Vidya"},
	{"Be the change you want to see","Gandhi"},
	{"Take my hand and I shall never let go","Sushant"},
	{}
}

func main() {

	//rand.Seed(time.Now().Unix())

	api:= echo.New()


	    //paricular end point
	api.GET("/quotes", getQuotes) //hello is a handler
	api.GET("/quotes/random", getRandomQuote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "32445"
	}

	//api.PUT("/home", hello)
	api.Start(":" + port)

}

func getQuotes (c echo.Context) error {
	return c.JSON(http.StatusOK,quotes)
}

func getRandomQuote (c echo.Context) error {
	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK,quotes[index])
}

/*func hello(c echo.Context) error{
	return c.JSON (200,"Hello Chethu")
}

func anotherHandler(c.echoContext) error {
	return c.NoContent(http.StatusOK)

}*/