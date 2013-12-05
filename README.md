# GoSrv

GoSrv is a thin wrapper around GoLang's HTTP server, to provide basic
command-line functionality, and env-specific configuration.

### Command Line

```Bash
$ myserver -h

Usage: myserver [options]

-a: ":9000"  Server address
-pid: "myserver.pid" Server PID File
-c: "myserver.cfg"  Config file
-e: "dev" Environment to run server in
-d: false Run server as daemon
-stop: false  Stop running server and exit
-restart: false Stop running server and boot daemon
```


### Config File

```ini
[DEFAULT]
pidFile: path/to/file.pid
readTimeout: 5s
writeTimeout: 500ms
certFile: path/to/myserver.cert
keyFile: path/to/myserver.key

customThing: foobar

[dev]
customThing: baz

[prod]
readTimeout: 2s

```


### Go Code


```Go
package main

include (
  "fmt"
  "gosrv"
)


func main() {
  s := gosrv.NewServerFromFlag()

  customThing, err := s.Config.String("customThing")
  if err != nil { fmt.Println("Custom thing is: "+customThing) }

  s.OnStop(func(){
    // Do something when ^C is received, before shut-down
  })

  s.HandleFunc("/", handler)

  panic( s.ListenAndServe() )
}


func handler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello World!"))
}
```
