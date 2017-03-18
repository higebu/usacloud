// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func SwitchUpdateCompleteFlags(ctx Context, params *UpdateSwitchParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "description", "desc":
		comp = define.Resources["Switch"].Commands["update"].Params["description"].CompleteFunc
	case "icon-id":
		comp = define.Resources["Switch"].Commands["update"].Params["icon-id"].CompleteFunc
	case "id":
		comp = define.Resources["Switch"].Commands["update"].Params["id"].CompleteFunc
	case "name":
		comp = define.Resources["Switch"].Commands["update"].Params["name"].CompleteFunc
	case "tags":
		comp = define.Resources["Switch"].Commands["update"].Params["tags"].CompleteFunc
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
