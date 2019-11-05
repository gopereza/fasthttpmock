package fasthttpmock

import (
	"bytes"
	"github.com/valyala/fasthttp"
)

func Equal(a *fasthttp.Request, b *fasthttp.Request) bool {
	return bytes.Equal(a.RequestURI(), b.RequestURI()) &&
		bytes.Equal(a.Header.Method(), b.Header.Method()) &&
		bytes.Equal(a.Host(), b.Host()) &&
		bytes.Equal(a.Body(), b.Body())
}

func Copy(from, to *fasthttp.Response) {
	to.SetStatusCode(from.StatusCode())
	to.SetBody(from.Body())
}
