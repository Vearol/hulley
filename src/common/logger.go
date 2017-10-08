package common

import (
  "go.uber.org/zap"
)

var Log *zap.Logger;

func InitLog(logFilepaths []string) {
  
  config := zap.NewProductionConfig()
  
  config.OutputPaths = logFilepaths
  
  zapLog, _ := config.Build()

  Log = zapLog
  
  defer Log.Sync()

  Log.Info("Logging initialized to " + logFilepaths[0])
}
