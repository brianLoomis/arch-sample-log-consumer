package main 

import (
	"github.com/chef/chef_go_server_utilities"
	"fmt"
)

func main () {
	myLog := logging.StdOutLogger{Val: 7}

	a := myLog.Log()
	fmt.Println(a)
}
