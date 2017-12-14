package main

import (
  "log"
  "testing"
  "github.com/lucasmbaia/grpc-mathoperations/client"
  "github.com/lucasmbaia/grpc-mathoperations/proto"
  "golang.org/x/net/context"
  "github.com/lucasmbaia/grpc-base/zipkin"
  "github.com/lucasmbaia/grpc-base/base"
)

func TestClientCalcDouble(t *testing.T) {
  var (
    err	    error
    value   mathoperations.Result
    conf    client.Config
    number  mathoperations.Number
    collector zipkin.Collector
  )

  if collector, err = newCollector() ; err != nil {
    log.Fatalf("Error to calc fibonacci: ", err)
  }

  number = mathoperations.Number{Value: 10}

  conf = client.Config {
    base.Config {
      Collector:	collector,
    },
  }

  if value, err = conf.CalcDouble(context.Background(), &number); err != nil {
    log.Fatalf("Error to calc double number: ", err)
  }

  log.Println(value.Value)
}

func newCollector() (zipkin.Collector, error) {
  return zipkin.NewCollector(
    "http://172.16.95.113:9411/api/v1/spans",
    "192.168.75.128:9090",
    "fibonacci",
    true,
    false,
  )
}
