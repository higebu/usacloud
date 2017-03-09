// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func BillCsvCompleteFlags(ctx Context, params *CsvBillParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "bill-output", "file":
		comp = define.Resources["Bill"].Commands["csv"].Params["bill-output"].CompleteFunc
	case "id":
		comp = define.Resources["Bill"].Commands["csv"].Params["id"].CompleteFunc
	case "no-header":
		comp = define.Resources["Bill"].Commands["csv"].Params["no-header"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}