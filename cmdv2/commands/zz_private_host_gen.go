// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"errors"
	"sync"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// privateHostCmd represents the command to manage SAKURA Cloud PrivateHost
func privateHostCmd() *cobra.Command {
	return &cobra.Command{
		Use: "private-host",

		Short: "A manage commands of PrivateHost",
		Long:  `A manage commands of PrivateHost`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}
}

func privateHostListCmd() *cobra.Command {
	privateHostListParam := params.NewListPrivateHostParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find", "select"},
		Short:        "List PrivateHost",
		Long:         `List PrivateHost`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return privateHostListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, privateHostListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if privateHostListParam.GenerateSkeleton {
				return generateSkeleton(ctx, privateHostListParam)
			}

			return funcs.PrivateHostList(ctx, privateHostListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&privateHostListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &privateHostListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&privateHostListParam.Tags, "tags", "", []string{}, "set filter by tags(AND) (aliases: selector)")
	fs.IntVarP(&privateHostListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&privateHostListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&privateHostListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&privateHostListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&privateHostListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&privateHostListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&privateHostListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(privateHostListNormalizeFlagNames)
	buildFlagsUsage(cmd, privateHostListFlagOrder(cmd))
	return cmd
}

func privateHostCreateCmd() *cobra.Command {
	privateHostCreateParam := params.NewCreatePrivateHostParam()
	cmd := &cobra.Command{
		Use: "create",

		Short:        "Create PrivateHost",
		Long:         `Create PrivateHost`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return privateHostCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, privateHostCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if privateHostCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, privateHostCreateParam)
			}

			// confirm
			if !privateHostCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.PrivateHostCreate(ctx, privateHostCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&privateHostCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&privateHostCreateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.StringSliceVarP(&privateHostCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &privateHostCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&privateHostCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&privateHostCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&privateHostCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&privateHostCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&privateHostCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(privateHostCreateNormalizeFlagNames)
	buildFlagsUsage(cmd, privateHostCreateFlagOrder(cmd))
	return cmd
}

func privateHostReadCmd() *cobra.Command {
	privateHostReadParam := params.NewReadPrivateHostParam()
	cmd := &cobra.Command{
		Use: "read",

		Short:        "Read PrivateHost",
		Long:         `Read PrivateHost`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return privateHostReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, privateHostReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if privateHostReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, privateHostReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPrivateHostReadTargets(ctx, privateHostReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				privateHostReadParam.SetId(id)
				go func(p *params.ReadPrivateHostParam) {
					err := funcs.PrivateHostRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(privateHostReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&privateHostReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&privateHostReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&privateHostReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&privateHostReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&privateHostReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostReadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(privateHostReadNormalizeFlagNames)
	buildFlagsUsage(cmd, privateHostReadFlagOrder(cmd))
	return cmd
}

func privateHostUpdateCmd() *cobra.Command {
	privateHostUpdateParam := params.NewUpdatePrivateHostParam()
	cmd := &cobra.Command{
		Use: "update",

		Short:        "Update PrivateHost",
		Long:         `Update PrivateHost`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return privateHostUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, privateHostUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if privateHostUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, privateHostUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPrivateHostUpdateTargets(ctx, privateHostUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !privateHostUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				privateHostUpdateParam.SetId(id)
				go func(p *params.UpdatePrivateHostParam) {
					err := funcs.PrivateHostUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(privateHostUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&privateHostUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&privateHostUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&privateHostUpdateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.StringSliceVarP(&privateHostUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &privateHostUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&privateHostUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&privateHostUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&privateHostUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&privateHostUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&privateHostUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostUpdateParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(privateHostUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, privateHostUpdateFlagOrder(cmd))
	return cmd
}

func privateHostDeleteCmd() *cobra.Command {
	privateHostDeleteParam := params.NewDeletePrivateHostParam()
	cmd := &cobra.Command{
		Use:          "delete",
		Aliases:      []string{"rm"},
		Short:        "Delete PrivateHost",
		Long:         `Delete PrivateHost`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return privateHostDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, privateHostDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if privateHostDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, privateHostDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPrivateHostDeleteTargets(ctx, privateHostDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !privateHostDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				privateHostDeleteParam.SetId(id)
				go func(p *params.DeletePrivateHostParam) {
					err := funcs.PrivateHostDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(privateHostDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&privateHostDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&privateHostDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&privateHostDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&privateHostDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&privateHostDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&privateHostDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostDeleteParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(privateHostDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, privateHostDeleteFlagOrder(cmd))
	return cmd
}

func privateHostServerInfoCmd() *cobra.Command {
	privateHostServerInfoParam := params.NewServerInfoPrivateHostParam()
	cmd := &cobra.Command{
		Use:          "server-info",
		Aliases:      []string{"server-list"},
		Short:        "ServerInfo PrivateHost",
		Long:         `ServerInfo PrivateHost`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return privateHostServerInfoParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, privateHostServerInfoParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if privateHostServerInfoParam.GenerateSkeleton {
				return generateSkeleton(ctx, privateHostServerInfoParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPrivateHostServerInfoTargets(ctx, privateHostServerInfoParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				privateHostServerInfoParam.SetId(id)
				go func(p *params.ServerInfoPrivateHostParam) {
					err := funcs.PrivateHostServerInfo(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(privateHostServerInfoParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&privateHostServerInfoParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&privateHostServerInfoParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostServerInfoParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostServerInfoParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostServerInfoParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostServerInfoParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostServerInfoParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&privateHostServerInfoParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&privateHostServerInfoParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostServerInfoParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&privateHostServerInfoParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostServerInfoParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostServerInfoParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostServerInfoParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(privateHostServerInfoNormalizeFlagNames)
	buildFlagsUsage(cmd, privateHostServerInfoFlagOrder(cmd))
	return cmd
}

func privateHostServerAddCmd() *cobra.Command {
	privateHostServerAddParam := params.NewServerAddPrivateHostParam()
	cmd := &cobra.Command{
		Use: "server-add",

		Short:        "ServerAdd PrivateHost",
		Long:         `ServerAdd PrivateHost`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return privateHostServerAddParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, privateHostServerAddParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if privateHostServerAddParam.GenerateSkeleton {
				return generateSkeleton(ctx, privateHostServerAddParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPrivateHostServerAddTargets(ctx, privateHostServerAddParam)
			if err != nil {
				return err
			}

			// confirm
			if !privateHostServerAddParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("server-add", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				privateHostServerAddParam.SetId(id)
				go func(p *params.ServerAddPrivateHostParam) {
					err := funcs.PrivateHostServerAdd(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(privateHostServerAddParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &privateHostServerAddParam.ServerId), "server-id", "", "set server ID")
	fs.StringSliceVarP(&privateHostServerAddParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&privateHostServerAddParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&privateHostServerAddParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostServerAddParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostServerAddParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostServerAddParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostServerAddParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostServerAddParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&privateHostServerAddParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&privateHostServerAddParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostServerAddParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&privateHostServerAddParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostServerAddParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostServerAddParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostServerAddParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(privateHostServerAddNormalizeFlagNames)
	buildFlagsUsage(cmd, privateHostServerAddFlagOrder(cmd))
	return cmd
}

func privateHostServerDeleteCmd() *cobra.Command {
	privateHostServerDeleteParam := params.NewServerDeletePrivateHostParam()
	cmd := &cobra.Command{
		Use: "server-delete",

		Short:        "ServerDelete PrivateHost",
		Long:         `ServerDelete PrivateHost`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return privateHostServerDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, privateHostServerDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if privateHostServerDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, privateHostServerDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPrivateHostServerDeleteTargets(ctx, privateHostServerDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !privateHostServerDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("server-delete", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				privateHostServerDeleteParam.SetId(id)
				go func(p *params.ServerDeletePrivateHostParam) {
					err := funcs.PrivateHostServerDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(privateHostServerDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &privateHostServerDeleteParam.ServerId), "server-id", "", "set server ID")
	fs.StringSliceVarP(&privateHostServerDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&privateHostServerDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&privateHostServerDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostServerDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostServerDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostServerDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostServerDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostServerDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&privateHostServerDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&privateHostServerDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostServerDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&privateHostServerDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostServerDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostServerDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostServerDeleteParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(privateHostServerDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, privateHostServerDeleteFlagOrder(cmd))
	return cmd
}
