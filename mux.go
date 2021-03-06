package gosrv

import (
  "net/http"
  "time"
  "os"
  "sync"
)


// The Mux struct handles HTTP logging and graceful server shutdown.
type Mux struct {
  *http.ServeMux
  Logger    HttpLogger
  conns     *sync.WaitGroup
  stopped   bool
  rwlock    sync.RWMutex
}


func NewMux() *Mux {
  return &Mux{http.NewServeMux(), NewHttpLogger(os.Stdout),
    &sync.WaitGroup{}, false, sync.RWMutex{}}
}


func (m *Mux) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
  m.conns.Add(1)
  res := NewResponse(wr, m)

  stime := time.Now()
  m.ServeMux.ServeHTTP(res, req)
  m.Logger.Log(stime, res, req)
  
  if m.conns != nil { m.conns.Done() }
}