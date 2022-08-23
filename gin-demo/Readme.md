## Simple API setup using Gin


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

## Useful tools

### API testing
- Insomnia - external client, similar to postman
- Thunder Client Plugin (VS Code)



## References
- Gin tutorial: https://go.dev/doc/tutorial/web-service-gin
- Insomnia REST Client: https://www.codemag.com/Article/2107051/Test-Your-REST-APIs-Using-Insomnia-REST-Client
- Thunder Client: https://marketplace.visualstudio.com/items?itemName=rangav.vscode-thunder-client
