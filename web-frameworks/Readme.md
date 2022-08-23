## Web Frameworks Demo

### Gin

Web framework which comes with a default set of middleware, also supports custom middleware. 
It has it's own context which is different from golang's ctx. 

```
▶ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /albums                   --> main.getAlbums (3 handlers)
[GIN-debug] GET    /albums/:id               --> main.getAlbumByID (3 handlers)
[GIN-debug] POST   /albums                   --> main.postAlbums (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on localhost:8080
[GIN] 2022/08/22 - 20:49:59 | 200 |     235.017µs |       127.0.0.1 | GET      "/albums"
[GIN] 2022/08/22 - 20:57:35 | 200 |     198.638µs |       127.0.0.1 | GET      "/albums"
[GIN] 2022/08/22 - 21:00:44 | 400 |    3.021096ms |       127.0.0.1 | POST     "/albums"
[GIN] 2022/08/22 - 21:02:49 | 201 |    7.352149ms |       127.0.0.1 | POST     "/albums"
[GIN] 2022/08/22 - 21:02:56 | 201 |      72.507µs |       127.0.0.1 | POST     "/albums"
[GIN] 2022/08/22 - 21:02:57 | 201 |      62.234µs |       127.0.0.1 | POST     "/albums"
[GIN] 2022/08/22 - 21:03:04 | 201 |      92.161µs |       127.0.0.1 | POST     "/albums"
[GIN] 2022/08/22 - 21:03:10 | 200 |      91.815µs |       127.0.0.1 | GET      "/albums"
[GIN] 2022/08/22 - 21:05:48 | 200 |     359.086µs |       127.0.0.1 | GET      "/albums"
[GIN] 2022/08/22 - 21:06:31 | 201 |      74.555µs |       127.0.0.1 | POST     "/albums"
[GIN] 2022/08/22 - 21:06:37 | 200 |      60.259µs |       127.0.0.1 | GET      "/albums"

```

```
▶ curl http://localhost:8080/albums
[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    },
    {
        "id": "4",
        "title": "The Modern Sound of Betty Carter",
        "artist": "Betty Carter",
        "price": 49.99
    },
    {
        "id": "4",
        "title": "The Modern Sound of Betty Carter",
        "artist": "Betty Carter",
        "price": 49.99
    },
    {
        "id": "4",
        "title": "The Modern Sound of Betty Carter",
        "artist": "Betty Carter",
        "price": 49.99
    },
    {
        "id": "5",
        "title": "The Modern Sound of Betty Carter",
        "artist": "Betty Carter",
        "price": 49.99
    },
    {
        "id": "1",
        "title": "Blue Car",
        "artist": "John Coltrane",
        "price": 56.99
    }
]
```


### Gorilla Mux

A router library which integrates with golang's native http package. 


### Echo 

A webframework very similar to gin, it has it's own context and comes with a collection of middleware along with supporting custom ones. 
https://echo.labstack.com/guide/

```
golang-experiments/web-frameworks/echo-demo  main ✗                                                                                    39m ⚑ ◒  
▶ go run main.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.8.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
{"time":"2022-08-22T22:57:32.637865-07:00","id":"","remote_ip":"127.0.0.1","host":"localhost:8080","method":"GET","uri":"/albums","user_agent":"Thunder Client (https://www.thunderclient.com)","status":200,"error":"","latency":158776,"latency_human":"158.776µs","bytes_in":95,"bytes_out":232}
{"time":"2022-08-22T22:57:46.508593-07:00","id":"","remote_ip":"127.0.0.1","host":"localhost:8080","method":"GET","uri":"/albums/2","user_agent":"Thunder Client (https://www.thunderclient.com)","status":200,"error":"","latency":67317,"latency_human":"67.317µs","bytes_in":95,"bytes_out":66}
{"time":"2022-08-22T22:58:08.078913-07:00","id":"","remote_ip":"127.0.0.1","host":"localhost:8080","method":"POST","uri":"/albums","user_agent":"Thunder Client (https://www.thunderclient.com)","status":201,"error":"","latency":146143,"latency_human":"146.143µs","bytes_in":82,"bytes_out":66}
{"time":"2022-08-22T22:58:21.930263-07:00","id":"","remote_ip":"127.0.0.1","host":"localhost:8080","method":"GET","uri":"/albums","user_agent":"Thunder Client (https://www.thunderclient.com)","status":200,"error":"","latency":42115,"latency_human":"42.115µs","bytes_in":82,"bytes_out":298}

```


## Useful tools

### API testing
- Insomnia - external client, similar to postman
- Thunder Client Plugin (VS Code)



## References
- Gin tutorial: https://go.dev/doc/tutorial/web-service-gin
- Insomnia REST Client: https://www.codemag.com/Article/2107051/Test-Your-REST-APIs-Using-Insomnia-REST-Client
- Thunder Client: https://marketplace.visualstudio.com/items?itemName=rangav.vscode-thunder-client
- Framework Comparisons:
  - https://pkg.go.dev/github.com/mingrammer/go-web-framework-stars#readme-top-go-web-frameworks