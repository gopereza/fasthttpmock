package fasthttpmock

import "github.com/valyala/fasthttp"

type Client struct {
	pairs *RequestResponsePairs
	equal func(*fasthttp.Request, *fasthttp.Request) bool
	copy  func(*fasthttp.Response, *fasthttp.Response)
}

func NewClient(pairs *RequestResponsePairs, equal func(*fasthttp.Request, *fasthttp.Request) bool, copy func(*fasthttp.Response, *fasthttp.Response)) *Client {
	return &Client{pairs: pairs, equal: equal, copy: copy}
}

func (c *Client) Do(request *fasthttp.Request, response *fasthttp.Response) error {
	for _, pair := range c.pairs.pairs {
		if c.equal(pair.request, request) {
			c.copy(pair.response, response)

			return nil
		}
	}

	return ErrNoMatch
}
