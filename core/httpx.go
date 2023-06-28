package core

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/httpx/runner"
)

type Httpx struct {
	Url string `json:"url"`
}

func (h *Httpx) Run() (content string, err error) {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	options := runner.Options{
		Methods:         "GET",
		InputTargetHost: goflags.StringSlice{h.Url},
		FollowRedirects: true,
		Silent:          true,
		OnResult: func(r runner.Result) {
			// handle error
			if r.Err != nil {
				err = r.Err
				return
			}
			content = fmt.Sprintf("Url:[%s] Host:[%s] Title:[%s] Status:[%d] Length:[%d] Favicon:[%s] Technologies:%v",
				r.URL, r.Host, r.Title, r.StatusCode, r.ContentLength, r.FavIconMMH3, r.Technologies)
		},
	}

	if err = options.ValidateOptions(); err != nil {
		return
	}

	r, err := runner.New(&options)
	if err != nil {
		return
	}
	defer r.Close()

	r.RunEnumeration()

	return
}
