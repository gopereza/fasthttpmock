# [fasthttp](https://github.com/valyala/fasthttp) mock [![Build Status](https://travis-ci.org/gopereza/fasthttpmock.svg?branch=master)](https://travis-ci.org/gopereza/fasthttpmock)

### Dep
```bash
dep ensure -add github.com/gopereza/fasthttpmock
```

### Example
```go
package main

import (
	"github.com/gopereza/fasthttpmock"
	"github.com/valyala/fasthttp"
)

func main() {
	pairs := fasthttpmock.NewRequestResponsePairs()

	{
		request := &fasthttp.Request{}
		request.Header.SetMethod("GET")
		request.SetRequestURI("http://example.com/test-1")

		response := &fasthttp.Response{}
		response.SetStatusCode(fasthttp.StatusOK)
		response.SetBodyString(`{"code":"1"}`)

		pairs.Add(request, response)
	}

	client := fasthttpmock.NewClient(pairs, fasthttpmock.Equal, fasthttpmock.Copy)

	request := fasthttp.AcquireRequest()
	request.Header.SetMethod("GET")
	request.SetRequestURI("http://example.com/test-1")
	response := fasthttp.AcquireResponse()

	_ = client.Do(request, response)
}
```

### Usage
```go
package main

import (
	"github.com/gopereza/fasthttpmock"
	"github.com/valyala/fasthttp"
)

func main() {
	fastClient := &fasthttp.Client{}

	client := fasthttpmock.NewWrapClient(fastClient)

	request := fasthttp.AcquireRequest()
	response := &fasthttp.Response{}
	_ = client.Do(request, response)

	mockClient := fasthttpmock.NewClient(fasthttpmock.NewRequestResponsePairs(), fasthttpmock.Equal, fasthttpmock.Copy)

	// switch to mock usage
	client.SetMockClient(mockClient)

	// switch to real usage
	client.SetMockClient(nil)
}
```