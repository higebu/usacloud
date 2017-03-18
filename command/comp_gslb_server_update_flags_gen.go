// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func GSLBServerUpdateCompleteFlags(ctx Context, params *ServerUpdateGSLBParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "enabled":
		comp = define.Resources["GSLB"].Commands["server-update"].Params["enabled"].CompleteFunc
	case "id":
		comp = define.Resources["GSLB"].Commands["server-update"].Params["id"].CompleteFunc
	case "index":
		comp = define.Resources["GSLB"].Commands["server-update"].Params["index"].CompleteFunc
	case "ipaddress":
		comp = define.Resources["GSLB"].Commands["server-update"].Params["ipaddress"].CompleteFunc
	case "weight":
		comp = define.Resources["GSLB"].Commands["server-update"].Params["weight"].CompleteFunc
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
