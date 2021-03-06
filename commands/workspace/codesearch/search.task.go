package main

import (
	"github.com/walterjwhite/go-application/libraries/utils/workspace/task/plugins"
	"github.com/walterjwhite/go-application/libraries/utils/workspace/task/plugins/codesearch"

	"flag"
	"github.com/walterjwhite/go-application/libraries/application"
)

func doSearchTask() {
	codesearch.Search(application.Context, plugins.Initialize(application.Context), flag.Args()[0])
}
