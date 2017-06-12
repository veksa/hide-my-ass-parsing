# Hide My Ass Proxy Parser

[![GoDoc](https://godoc.org/github.com/veksa/hide-my-ass-parsing?status.svg)](https://godoc.org/github.com/veksa/hide-my-ass-parsing)  

Parse the Hide My Ass proxies website with GO

## Usage ##

```go
import "github.com/veksa/hide-my-ass-parsing"
```

Get list of proxies and test it. For example:

```go
proxies := hideMyAssParsing.GetProxies()

successCallback := func(proxy Proxy) { activateProxy(proxy) }
errorCallback := func(proxy Proxy) { deactivateProxy(proxy) }

hideMyAssParsing.TestProxies(proxies, successCallback, errorCallback)
```