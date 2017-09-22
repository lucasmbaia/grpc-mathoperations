package main

import (
  "log"
  "reflect"

  "github.com/lucasmbaia/grpc-base/config"
  "github.com/lucasmbaia/grpc-base/base"
  "github.com/lucasmbaia/grpc-mathoperations/proto"
  "github.com/lucasmbaia/grpc-mathoperations/server"
)

func init() {
  config.LoadConfig()
}

func main() {
  var (
    configCMD base.ConfigCMD
    errChan   = make(chan error, 1)
  )

  go func() {
    configCMD = base.ConfigCMD {
      SSL:		true,
      RegisterConsul:	true,
      ServiceServer:	reflect.Indirect(reflect.ValueOf(mathoperations.RegisterMathOperationsServiceServer)),
      HandlerEndpoint:  reflect.Indirect(reflect.ValueOf(mathoperations.RegisterMathOperationsServiceHandlerFromEndpoint)),
      ServerConfig:     server.NewMathOperationsServer(),
    }

    errChan <- configCMD.Run()
  }()

  select {
  case e := <-errChan:
    log.Fatalf("Error grpc server: ", e)
  }
}
