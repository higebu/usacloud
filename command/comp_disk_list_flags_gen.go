// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func DiskListCompleteFlags(ctx Context, params *ListDiskParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "from":
		comp = define.Resources["Disk"].Commands["list"].Params["from"].CompleteFunc
	case "id":
		comp = define.Resources["Disk"].Commands["list"].Params["id"].CompleteFunc
	case "max":
		comp = define.Resources["Disk"].Commands["list"].Params["max"].CompleteFunc
	case "name":
		comp = define.Resources["Disk"].Commands["list"].Params["name"].CompleteFunc
	case "scope":
		comp = define.Resources["Disk"].Commands["list"].Params["scope"].CompleteFunc
	case "sort":
		comp = define.Resources["Disk"].Commands["list"].Params["sort"].CompleteFunc
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
