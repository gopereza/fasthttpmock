package fasthttpmock

import "github.com/valyala/fasthttp"

type WrapClient struct {
	realClient *fasthttp.Client
	mockClient *Client
	mocked     bool
}

func NewWrapClient(realClient *fasthttp.Client) *WrapClient {
	return &WrapClient{realClient: realClient}
}

func (c *WrapClient) Do(request *fasthttp.Request, response *fasthttp.Response) error {
	if c.mocked {
		return c.mockClient.Do(request, response)
	}

	return c.realClient.Do(request, response)
}

func (c *WrapClient) SetMockClient(mockClient *Client) {
	c.mockClient = mockClient
	c.mocked = mockClient != nil
}
