package fasthttpmock

import (
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestClient_Do_EmptyPairs(t *testing.T) {
	client := NewClient(NewRequestResponsePairs(), Equal, Copy)

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()

	assert.Equal(t, ErrNoMatch, client.Do(request, response))
}

func TestClient_Do(t *testing.T) {
	pairs := NewRequestResponsePairs()

	{
		request := &fasthttp.Request{}
		request.Header.SetMethod("GET")
		request.SetRequestURI("http://example.com/test-1")

		response := &fasthttp.Response{}
		response.SetStatusCode(fasthttp.StatusOK)
		response.SetBodyString(`{"code":"1"}`)

		pairs.Add(request, response)
	}

	{
		request := &fasthttp.Request{}
		request.Header.SetMethod("POST")
		request.SetRequestURI("http://example.com/test-2")

		response := &fasthttp.Response{}
		response.SetStatusCode(fasthttp.StatusOK)
		response.SetBodyString(`{"code":"2"}`)

		pairs.Add(request, response)
	}

	client := NewClient(pairs, Equal, Copy)

	{
		request := fasthttp.AcquireRequest()
		request.Header.SetMethod("GET")
		request.SetRequestURI("http://example.com/test-1")
		response := fasthttp.AcquireResponse()

		assert.NoError(t, client.Do(request, response))
		assert.Equal(t, fasthttp.StatusOK, response.StatusCode())
		assert.Equal(t, []byte(`{"code":"1"}`), response.Body())

		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}

	{
		request := fasthttp.AcquireRequest()
		request.Header.SetMethod("POST")
		request.SetRequestURI("http://example.com/test-2")
		response := fasthttp.AcquireResponse()

		assert.NoError(t, client.Do(request, response))
		assert.Equal(t, fasthttp.StatusOK, response.StatusCode())
		assert.Equal(t, []byte(`{"code":"2"}`), response.Body())

		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}

	{
		request := fasthttp.AcquireRequest()
		request.Header.SetMethod("POST")
		request.SetRequestURI("http://example.com/test-3")
		response := fasthttp.AcquireResponse()

		assert.Equal(t, ErrNoMatch, client.Do(request, response))

		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}
}
