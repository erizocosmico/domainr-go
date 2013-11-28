domainr-go [![GoDoc](http://godoc.org/github.com/mvader/domainr-go?status.png)](http://godoc.org/github.com/mvader/domainr-go)
==========

Domai.nr API wrapper for Go(lang).

```go
// Get the url to register a domain
domainr.Register("myfancydomain.com", "")
// Get the url of a specific registrar to register a domain
domainr.Register("myfancydomain.com", "gandi.net")
// Get information about a domain
domainr.Json(domainr.METHOD_INFO, "github.io", "")
// Get information about the domains matching your query
domainr.Json(domainr.METHOD_SEARCH, "myfancydomain", "")
// The optional third parameter at domainr.Json is a callback as defined in
// the API documentation https://domai.nr/api/docs/json
```

Install
------
```bash
go get github.com/mvader/domainr-go
```
