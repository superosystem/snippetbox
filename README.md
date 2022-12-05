# Snippetbox

## Routing request
| Method | URL Pattern | Handler    | Action    |
| :---:   | :---: | :---: | :---: |
| GET | /   | home   | Display the home page   |
| GET | /snippet/view?id=1    | snipperView   | Display a specific snippet   |
| GET | /snippet/create   | snippetCreate   | Display a HTML form for creating a new snipper   |
| POST | /snippet/create   | snippetCreatePost   | Create a new snippet   |
| GET | /static   | http.FileServer   | Server a specific static file   |


## Security Improvements

* Make LTS with running generate_cert.go
```bash
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```