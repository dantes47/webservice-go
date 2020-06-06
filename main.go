package main

import (
	"net/http"

	"github.com/dantes47/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
