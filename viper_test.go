package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("/opt/gopath/src/github.com/hyperledger/fabric/")
	AddConfigPath(v, OfficialPath)
	viper.SetConfigName(configName)

}
