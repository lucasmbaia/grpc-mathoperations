package server

import (
  "golang.org/x/net/context"
  "github.com/lucasmbaia/grpc-mathoperations/proto"
  empty "github.com/golang/protobuf/ptypes/empty"
)

type MathOperationsServer struct {}

func NewMathOperationsServer() MathOperationsServer {
  return MathOperationsServer{}
}

func (m MathOperationsServer) Double(ctx context.Context, v *mathoperations.Number) (*mathoperations.Result, error) {
  var double = v.Value * v.Value

  return &mathoperations.Result{Value: double}, nil
}

func (m MathOperationsServer) Health(ctx context.Context, emp *empty.Empty) (*empty.Empty, error) {
  select {
  case <-ctx.Done():
    return nil, ctx.Err()
  default:
    return new(empty.Empty), nil
  }
}
