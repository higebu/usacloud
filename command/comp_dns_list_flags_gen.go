// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func DNSListCompleteFlags(ctx Context, params *ListDNSParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "from":
		comp = define.Resources["DNS"].Commands["list"].Params["from"].CompleteFunc
	case "id":
		comp = define.Resources["DNS"].Commands["list"].Params["id"].CompleteFunc
	case "max":
		comp = define.Resources["DNS"].Commands["list"].Params["max"].CompleteFunc
	case "name":
		comp = define.Resources["DNS"].Commands["list"].Params["name"].CompleteFunc
	case "sort":
		comp = define.Resources["DNS"].Commands["list"].Params["sort"].CompleteFunc
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
