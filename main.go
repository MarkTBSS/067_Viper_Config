package main

import (
	"fmt"

	"github.com/MarkTBSS/067_Viper_Config/config"
)

func main() {
	conf := config.ConfigGetting()
	fmt.Println(conf)
}
