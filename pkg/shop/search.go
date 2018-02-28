package shop

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/context/ctxhttp"
)

func search(ctx context.Context, method, url, contentType, body string) (ret []byte, err error) {
	var (
		req  *http.Request
		resp *http.Response
	)
	req, err = http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.167 Safari/537.36")
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err = ctxhttp.Do(ctx, http.DefaultClient, req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	ret, err = ioutil.ReadAll(resp.Body)
	return
}
