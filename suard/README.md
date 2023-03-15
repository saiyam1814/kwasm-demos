Spin Up and Running

```bash
$ spin build && spin up --listen 127.0.0.1:8080
Executing the build command for component suard: tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm main.go
Working directory: "/Users/c.voigt/go/src/github.com/voigt/spin_bugbash/suard/./"
Successfully ran the build command for the Spin components.
Serving http://127.0.0.1:8080
Available Routes:
  suard: http://127.0.0.1:8080 (wildcard)
```

Test

```bash
$ spin_routers curl -s http://127.0.0.1:8080/    
Welcome!
$ spin_routers git:(main) ✗ curl -s http://127.0.0.1:8080/hello/name 
hello, name!
$ curl -s http://127.0.0.1:8080/env/api
{"Host": "127.0.0.1:8080","User-Agent": "curl/7.86.0","Accept": "*/*","Spin-Matched-Route": "/...","Spin-Raw-Component-Route": "/...","Spin-Base-Path": "/","Spin-Component-Route": "","Spin-Path-Info": "/env/api","Spin-Full-Url": "http://127.0.0.1:8080/env/api"}%
```
