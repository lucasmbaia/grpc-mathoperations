package client

import (
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/lucasmbaia/grpc-base/config"
  "github.com/lucasmbaia/grpc-base/base"
  "github.com/lucasmbaia/grpc-mathoperations/proto"
)

type Config struct {
  base.Config
}

func init() {
  config.LoadConfig()
}

func (c Config) CalcDouble(ctx context.Context, n *mathoperations.Number) (mathoperations.Result, error) {
  var (
    conn  *grpc.ClientConn
    cM    mathoperations.MathOperationsServiceClient
    err   error
    r     *mathoperations.Result
    value mathoperations.Result
  )

  if conn, err = c.ClientConnect(); err != nil {
    return value, err
  }
  defer conn.Close()

  cM = mathoperations.NewMathOperationsServiceClient(conn)

  if r, err = cM.Double(ctx, &mathoperations.Number{Value: n.Value}); err != nil {
    return value, err
  }

  return mathoperations.Result{Value: r.Value}, nil
}
