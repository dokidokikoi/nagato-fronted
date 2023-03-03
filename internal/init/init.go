package inittask

import (
	"fmt"
	"nagato/internal/model"

	conf "github.com/dokidokikoi/go-common/config"
)

func Init(configFile string) {
	c := model.GetConfig()
	conf.Parse(configFile, &c)
	model.SetConfig(c)
	fmt.Printf("%+v", c)
}
