package fasthttpmock

import "github.com/valyala/fasthttp"

type requestResponsePair struct {
	request  *fasthttp.Request
	response *fasthttp.Response
}

type RequestResponsePairs struct {
	pairs []requestResponsePair
}

func NewRequestResponsePairs() *RequestResponsePairs {
	return &RequestResponsePairs{}
}

func (p *RequestResponsePairs) Add(request *fasthttp.Request, response *fasthttp.Response) {
	p.pairs = append(p.pairs, requestResponsePair{
		request:  request,
		response: response,
	})
}
