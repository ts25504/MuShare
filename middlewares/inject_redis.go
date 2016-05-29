package middlewares

import (
  "github.com/go-martini/martini"
  "MuShare/conf"
  "net/http"
  redis "gopkg.in/redis.v3"
)

func InjectRedis() martini.Handler {
  return func(c martini.Context, rw http.ResponseWriter, config *conf.Conf) {
    client := redis.NewClient(&redis.Options{
      Addr: config.Redis.Addr,
      Password: config.Redis.Password,
      DB: 0,
    })
    _, err := client.Ping().Result()
    if err != nil {
      panic(err.Error())
    }
    c.Map(client)
    c.Next()
    client.Close()
  }
}
