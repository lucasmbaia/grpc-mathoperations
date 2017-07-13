package main

import (
  "log"
  "testing"
  "github.com/lucasmbaia/grpc-mathoperations/client"
  "github.com/lucasmbaia/grpc-mathoperations/proto"
)

func TestClientCalcDouble(t *testing.T) {
  var (
    err	    error
    value   mathoperations.Result
    conf    client.Config
    number  mathoperations.Number
  )

  number = mathoperations.Number{Value: 10}

  if value, err = conf.CalcDouble(&number); err != nil {
    log.Fatalf("Error to calc double number: ", err)
  }

  log.Println(value.Value)
}
