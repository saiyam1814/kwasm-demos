# Testing Go HTTP Routers

Usage

```bash
$ spin build && spin up --listen 127.0.0.1:8080
Executing the build command for component httprouter: tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm main.go
Working directory: "/Users/c.voigt/go/src/github.com/voigt/spin_bugbash/spin-routers/./api/httprouter"
Executing the build command for component chirouter: tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm main.go
Working directory: "/Users/c.voigt/go/src/github.com/voigt/spin_bugbash/spin-routers/./api/chirouter"
Successfully ran the build command for the Spin components.
Serving http://127.0.0.1:8080
Available Routes:
  httprouter: http://127.0.0.1:8080/httprouter (wildcard)
  chirouter: http://127.0.0.1:8080/chi (wildcard)
```

Test

```bash
$ spin_routers curl -s http://127.0.0.1:8080/httprouter/                
Welcome!
$ spin_routers curl -s http://127.0.0.1:8080/chi/         
Welcome!
$ spin_routers curl -s http://127.0.0.1:8080/chi/hello/chi  
hello, chi!
$ spin_routers curl -s http://127.0.0.1:8080/httprouter/hello/httprouter
hello, httprouter!
```
