package execise

import (
  //"os"
  //"fmt"
  "github.com/wang/common/flogging"
)

var logger = flogging.MustGetLogger("log")


func Uselog() {
  //logger.ResetLevels()
  //logger.Apply(flogging.Config{LogSpec:"DEBUG"})
  //loging, _ := flogging.New(flogging.Config{LogSpec: "DEBUG"})
  flogging.Init(flogging.Config{LogSpec: "debug"})
  logger.Debug("this is a debug")
  logger.Debugf("this is %dth debug",10)
  logger.Info("info,aaaaaaaaaaa")
  //flogging.Reset()
  //logger.Info("after reset")
  logger.Debugf("this is %dth debug",11)
  logger.Info("info,bbbbbbbbbbb")
  flogging.Init(flogging.Config{})
  logger.Info("after init")
  logger.Debugf("this is %dth debug",11)
  logger.Info("info,bbbbbbbbbbb")
}
