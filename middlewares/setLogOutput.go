package middlewares

import (
  "os"
  "log"
  "github.com/go-martini/martini"
)

func SetLogOutput(logger *log.Logger, c martini.Context) {
  if martini.Env == "test" {
    return
  }
  f, err := os.OpenFile("log/log_" + martini.Env, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
  } else {
    logger.SetOutput(f)
    c.Next()
    f.Close()
  }

}
