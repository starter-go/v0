package wuclient

import (
	"net/http"
	"time"

	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

type WebUnitTestSessionRunner struct {

	//starter:component

}

func (inst *WebUnitTestSessionRunner) _impl() {

}

func (inst *WebUnitTestSessionRunner) Life() *application.Life {

	l := &application.Life{
		Order:  0,
		OnLoop: inst.run,
	}
	return l
}

func (inst *WebUnitTestSessionRunner) run() error {

	time.Sleep(time.Second * 2)

	baseUrl := "http://localhost:45138/test/web-unit/sessions/"

	urllist := make([]string, 0)

	urllist = append(urllist, baseUrl+"fetch")
	urllist = append(urllist, baseUrl+"insert")

	urllist = append(urllist, baseUrl+"fetch")
	urllist = append(urllist, baseUrl+"update")

	urllist = append(urllist, baseUrl+"fetch")
	urllist = append(urllist, baseUrl+"remove")

	urllist = append(urllist, baseUrl+"fetch")

	for _, l := range urllist {
		err := inst.doHttpGet(l)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *WebUnitTestSessionRunner) doHttpGet(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	code := resp.StatusCode
	vlog.Debug("http.status.code = %d", code)

	return nil
}
