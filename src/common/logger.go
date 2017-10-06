package common

import (
  "go.uber.org/zap"
)

var Log *zap.Logger;

func InitLog(logFilepath string){
  
  config := zap.NewProductionConfig()
  config.OutputPaths = []string{ logFilepath }
  
  zapLog, _ := config.Build()

  Log = zapLog
  
  defer Log.Sync()
  Log.Info("logger set")
}
