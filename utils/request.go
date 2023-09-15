package utils

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	client  *http.Client
	request *http.Request
}

func NewClient(method, url string, body io.Reader) (*Client, error) {
	request, err := http.NewRequest(method, url, body)
	return &Client{
		client:  &http.Client{},
		request: request,
	}, err
}

func (c *Client) Do() (*http.Response, error) {
	return c.client.Do(c.request)
}

func (c *Client) Struct(result interface{}) error {
	response, err := c.client.Do(c.request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	reader := response.Body
	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(response.Body)
	}

	bytes, err := io.ReadAll(reader)
	err = json.Unmarshal(bytes, result)
	return err
}

func (c *Client) SetReferer() *Client {
	c.request.Header.Set("Referer", "https://www.bilibili.com")
	return c
}

func (c *Client) SetHeader() *Client {
	c.request.Header.Set("Accept-Encoding", "gzip")
	c.request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36 Edg/116.0.1938.76")
	c.request.Header.Set("Cookie", "buvid3=10F8A50E-8860-9DF6-E49C-F7C8E2E9EFB319957infoc; b_nut=1694406419; i-wanna-go-back=-1; b_ut=7; _uuid=B7C5B1C4-2949-BA84-C1E1-BCC8A75215E924847infoc; buvid_fp=eef5f289d5920f1365470dc05538e71c; home_feed_column=5; browser_resolution=1912-1124; buvid4=D4992931-39F1-56CE-9EDF-8F191E74007126039-023091112-3YsJwcGcPHuh%2F6bk4pHLeA%3D%3D; SESSDATA=55cac37d%2C1709958451%2C90973%2A91CjAL7Lu10nIFzV-smSBziwclfsBGCoKLgagiRTxWXam5ZsoHSBYvB5GkRjZkMjJnxPsSVldwUFhRUVYwNkg2alo0V2RaZ3cwSDRVQ05xM2xQWi14eWgtSEVMV1M5RHpVVy1RczZJcDVDNXh1NWd6c3Q2RVpQTU9lWDkyS0ZQSDdkVjJDdHFKdk1BIIEC; bili_jct=75821628171d88066b5ccf7d1e02f179; DedeUserID=29396805; DedeUserID__ckMd5=098d448652db5561; sid=67800uhi; header_theme_version=CLOSE; hit-new-style-dyn=1; hit-dyn-v2=1; CURRENT_BLACKGAP=0; rpdid=|(J|)Rm|mYYR0J'uYmRR|Ylkm; CURRENT_FNVAL=4048; PVID=1; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ3NDczMjEsImlhdCI6MTY5NDQ4ODEyMSwicGx0IjotMX0.yOcd1hrjgCGmZVbDYgpwFUq1b-Z55u0yc-f2CO0binY; bili_ticket_expires=1694747321; CURRENT_QUALITY=16; bp_video_offset_29396805=840374183294664723")
	return c
}

func (c *Client) SetCookie() *Client {
	c.request.AddCookie(&http.Cookie{Name: "SESSDATA", Value: "55cac37d%2C1709958451%2C90973%2A91CjAL7Lu10nIFzV-smSBziwclfsBGCoKLgagiRTxWXam5ZsoHSBYvB5GkRjZkMjJnxPsSVldwUFhRUVYwNkg2alo0V2RaZ3cwSDRVQ05xM2xQWi14eWgtSEVMV1M5RHpVVy1RczZJcDVDNXh1NWd6c3Q2RVpQTU9lWDkyS0ZQSDdkVjJDdHFKdk1BIIEC"})
	return c
}
