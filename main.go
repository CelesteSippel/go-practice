package main

import (
	"net/http"

	"github.com/celestesippel/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
