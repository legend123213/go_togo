package api

import (
	"fmt"

	"github.com/legend123213/go_togo/tree/master/apidemo/internal/api/routers"
)

func Run(port string) {
	web := routers.Setup()
	fmt.Print("server is running in port 8000 \n")
	web.Run(port)
}
