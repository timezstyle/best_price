package shop

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/context/ctxhttp"
)

func search(ctx context.Context, method, url, body string, header *http.Header) (ret []byte, respHeader http.Header, err error) {
	var (
		req  *http.Request
		resp *http.Response
	)
	req, err = http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return
	}
	if header != nil {
		req.Header = *header
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36")

	resp, err = ctxhttp.Do(ctx, http.DefaultClient, req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respHeader = resp.Header

	ret, err = ioutil.ReadAll(resp.Body)
	return
}
