package main

import (
	"net/http"
	"github.com/kdcllc/go-apiexample/controllers"
)

func main ()  {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}