package main

import (
  "flag"
  "sync"
  "context"
  "io"
  "syscall"
  "encoding/json"
  "os"
  "net"
  "net/http"
  "log"
  "fmt"
  "os/signal"

  "google.golang.org/grpc"

  "github.com/bd878/gallery/server/api"
  config "github.com/bd878/gallery/server/users/config"
  httphandler "github.com/bd878/gallery/server/users/internal/handler/http"
  grpchandler "github.com/bd878/gallery/server/users/internal/handler/grpc"
  controller "github.com/bd878/gallery/server/users/internal/controller/users"
  sqlite "github.com/bd878/gallery/server/users/internal/repository/sqlite"
)

var (
  configPath = flag.String("config", "config/default.json", "config path")
  interactive = flag.Bool("interactive", true, "ignore logFile in config " + 
    "output log messages to stdout")
)

func main() {
  flag.Parse()

  serverCfg := loadConfig()

  if serverCfg.Debug {
    if *interactive {
      log.SetOutput(os.Stdout)
    } else {
      f := setLogOutput(serverCfg.LogFile)
      defer f.Close()
    }
  }

  c := make(chan os.Signal, 1)
  go trackConfig(c)
  defer close(c)

  var wg sync.WaitGroup
  wg.Add(2)

  go func() { httpRun(serverCfg); wg.Done() }()
  go func() { grpcRun(serverCfg); wg.Done() }()

  wg.Wait()
}

func httpRun(cfg *config.Config) {
  mem, err := sqlite.New(cfg.DBPath)
  if err != nil {
    panic(err)
  }
  ctrl := controller.New(mem)
  h := httphandler.New(ctrl, httphandler.Config{Domainname:cfg.Domainname})

  netCfg := net.ListenConfig{}
  l, err := netCfg.Listen(context.Background(), "tcp4", fmt.Sprintf(":%d", cfg.HttpPort))
  if err != nil {
    panic(err)
  }
  defer l.Close()

  http.Handle("/users/v1/signup", http.HandlerFunc(h.Register))
  http.Handle("/users/v1/login", http.HandlerFunc(h.Authenticate))
  http.Handle("/users/v1/auth", http.HandlerFunc(h.Auth))
  http.Handle("/users/v1/status", http.HandlerFunc(h.ReportStatus))

  log.Println("http server is listening on =", l.Addr())
  if err := http.Serve(l, nil); err != nil {
    panic(err)
  }
  log.Println("http server exited")
}

func grpcRun(cfg *config.Config) {
  mem, err := sqlite.New(cfg.DBPath)
  if err != nil {
    panic(err)
  }
  ctrl := controller.New(mem)
  h := grpchandler.New(ctrl)
  netCfg := net.ListenConfig{}
  l, err := netCfg.Listen(context.Background(), "tcp4", fmt.Sprintf(":%d", cfg.GrpcPort))
  if err != nil {
    panic(err)
  }
  defer l.Close()

  srv := grpc.NewServer()
  api.RegisterUserServiceServer(srv, h)
  log.Println("grpc server is listening on =", l.Addr())
  if err := srv.Serve(l); err != nil {
    panic(err)
  }
  log.Println("grpc server exited")
}

func loadConfig() *config.Config {
  f, err := os.Open(*configPath)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  var cfg config.Config
  if err := json.NewDecoder(f).Decode(&cfg); err != nil {
    panic(err)
  }

  return &cfg
}

func trackConfig(c chan os.Signal) {
  signal.Notify(c, syscall.SIGHUP)

  var f *os.File
  defer f.Close()

  for {
    switch <-c {
    case syscall.SIGHUP:
      log.Println("recieve sighup")

      cfg := loadConfig()
      if cfg.Debug {
        f = setLogOutput(cfg.LogFile)
      } else {
        f.Close()
        log.SetOutput(io.Discard)
      }
    }
  }
}

func setLogOutput(p string) *os.File {
  f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    panic(err)
  }

  log.SetOutput(f)
  return f
}
