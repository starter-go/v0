package core

import (
	"net/http"
	"testing"
)

func TestUrlResolver(t *testing.T) {

	urlist := []string{
		"http://apple.com/path/2/file",
		"https://bilibili.com/tv666",
		"/a/b/c",
		"/users/{{username}}/info?service=fetch",
		"x/y/z",
	}

	query := map[string]string{
		"id": "666",
	}
	params := map[string]string{
		"username": "lilei",
	}

	resolver := new(DefaultUrlResolver)
	service := new(ServiceImpl)
	method := http.MethodGet

	service.BaseURL = "http://example.com/foo/bar"
	service.init()

	const bar = "--------------------------------------------------------"

	for _, url := range urlist {

		t.Log(bar)

		t.Logf("Resolve: URL = %s", url)
		reqctx := service.NewTransaction(method, url)

		reqctx.Want.Params = params
		reqctx.Want.Query = query

		location, err := resolver.Resolve(reqctx)
		if err != nil {
			t.Error(err)
			return
		}

		t.Logf("       base = %s", reqctx.AC.BaseURL)
		t.Logf("        raw = %s", reqctx.Want.RawURL)
		t.Logf("  result(s) = %s", reqctx.Want.URL)
		t.Logf("  result(l) = %s", location.String())

		t.Log(bar)
	}

}
