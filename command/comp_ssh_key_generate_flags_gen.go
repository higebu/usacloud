// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func SSHKeyGenerateCompleteFlags(ctx Context, params *GenerateSSHKeyParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "description", "desc":
		comp = define.Resources["SSHKey"].Commands["generate"].Params["description"].CompleteFunc
	case "name":
		comp = define.Resources["SSHKey"].Commands["generate"].Params["name"].CompleteFunc
	case "pass-phrase":
		comp = define.Resources["SSHKey"].Commands["generate"].Params["pass-phrase"].CompleteFunc
	case "private-key-output", "file":
		comp = define.Resources["SSHKey"].Commands["generate"].Params["private-key-output"].CompleteFunc
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
