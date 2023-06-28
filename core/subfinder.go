package core

import (
	"bytes"
	"context"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
	"io"
)

type Subfinder struct {
	Domain string `json:"domain"`
}

func (s *Subfinder) Run() (content string, err error) {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	options := &runner.Options{
		Threads:            10, // Thread controls the number of threads to use for active enumerations
		Timeout:            30, // Timeout is the seconds to wait for sources to respond
		MaxEnumerationTime: 10, // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		//Resolvers:          resolve.DefaultResolvers, // Use the default list of resolvers by marshaling it to the config
		All: true, // All specifies whether to use all (slow) sources.
	}
	r, err := runner.NewRunner(options)
	if err != nil {
		return
	}

	output := &bytes.Buffer{}
	// To run subdomain enumeration on a single domain
	if err = r.EnumerateSingleDomainWithCtx(context.Background(), s.Domain, []io.Writer{output}); err != nil {
		return
	}

	content = output.String()
	return
}
