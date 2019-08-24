package redis_api

import (
   // "os"
    //"net/http"
)

var connections redisConnections
var handler *Message

func init() {
    connections := make(redisConnections)
    for _,conf := range RedisHosts {
        hval := conf.GetHval()
        connections[hval] = GetRedis(hval)
        conn := connections[hval]
        if conn.Err == nil {
            conn.initKeys()
        }
    }

    //root := os.Getenv("GOPATH")+"/src/github.com/go-libraries/redis-api"
    handler = &Message{
        //Url:"http://127.0.0.1:10003",
        //Root:root,
        //FileHandler:http.FileServer(http.Dir(root+"/resources/app")),
    }
}
