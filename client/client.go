package client

import (
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/lucasmbaia/grpc-base/config"
  "github.com/lucasmbaia/grpc-mathoperations/proto"
  "google.golang.org/grpc/credentials"
)

type Config struct {
  SSL bool
}

func init() {
  config.LoadConfig()
}

func (c Config) CalcDouble(n *mathoperations.Number) (mathoperations.Result, error) {
  var (
    conn  *grpc.ClientConn
    cM    mathoperations.MathOperationsServiceClient
    err   error
    r     *mathoperations.Result
    value mathoperations.Result
  )

  if conn, err = c.connect(); err != nil {
    return value, err
  }
  defer conn.Close()

  cM = mathoperations.NewMathOperationsServiceClient(conn)

  if r, err = cM.Double(context.Background(), &mathoperations.Number{Value: n.Value}); err != nil {
    return value, err
  }

  return mathoperations.Result{Value: r.Value}, nil
}

func (c Config) connect() (*grpc.ClientConn, error) {
  var (
    opts  []grpc.DialOption
    creds credentials.TransportCredentials
    err   error
  )

  if config.EnvConfig.GrpcSSL {
    if creds, err = credentials.NewClientTLSFromFile(config.EnvConfig.CAFile, config.EnvConfig.ServerNameAuthority); err != nil {
      return new(grpc.ClientConn), err
    }

    opts = []grpc.DialOption{
      grpc.WithTransportCredentials(creds),
    }
  } else {
    opts = []grpc.DialOption{
      grpc.WithInsecure(),
    }
  }

  return grpc.Dial(config.EnvLocal.LinkerdURL, opts...)
}
