package middlewares

import (
  "github.com/go-martini/martini"
  "MuShare/conf"
  "net/http"
  redis "gopkg.in/redis.v3"
)

func InjectRedis(r conf.Redis) martini.Handler {
  return func(c martini.Context, rw http.ResponseWriter) {
    client := redis.NewClient(&redis.Options{
      Addr: r.Addr,
      Password: r.Password,
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
