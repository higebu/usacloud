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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// ipv6Cmd represents the command to manage SAKURA Cloud IPv6
func ipv6Cmd() *cobra.Command {
	return &cobra.Command{
		Use: "ipv6",

		Short: "A manage commands of IPv6",
		Long:  `A manage commands of IPv6`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}
}

func ipv6ListCmd() *cobra.Command {
	ipv6ListParam := params.NewListIPv6Param()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find"},
		Short:        "List IPv6",
		Long:         `List IPv6`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipv6ListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, ipv6ListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if ipv6ListParam.GenerateSkeleton {
				return generateSkeleton(ctx, ipv6ListParam)
			}

			return funcs.IPv6List(ctx, ipv6ListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&ipv6ListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &ipv6ListParam.Id), "id", "", "set filter by id(s)")
	fs.VarP(newIDValue(0, &ipv6ListParam.IPv6netId), "ipv6net-id", "", "set filter by ipv6net-id")
	fs.VarP(newIDValue(0, &ipv6ListParam.InternetId), "internet-id", "", "set filter by internet-id")
	fs.IntVarP(&ipv6ListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&ipv6ListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&ipv6ListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&ipv6ListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&ipv6ListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&ipv6ListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&ipv6ListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&ipv6ListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&ipv6ListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&ipv6ListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&ipv6ListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&ipv6ListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&ipv6ListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&ipv6ListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&ipv6ListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(ipv6ListNormalizeFlagNames)
	buildFlagsUsage(cmd, ipv6ListFlagOrder(cmd))
	return cmd
}

func ipv6PtrAddCmd() *cobra.Command {
	ipv6PtrAddParam := params.NewPtrAddIPv6Param()
	cmd := &cobra.Command{
		Use: "ptr-add",

		Short:        "PtrAdd IPv6",
		Long:         `PtrAdd IPv6`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipv6PtrAddParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, ipv6PtrAddParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if ipv6PtrAddParam.GenerateSkeleton {
				return generateSkeleton(ctx, ipv6PtrAddParam)
			}

			// confirm
			if !ipv6PtrAddParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("ptr-add", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.IPv6PtrAdd(ctx, ipv6PtrAddParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&ipv6PtrAddParam.Hostname, "hostname", "", "", "set server hostname")
	fs.BoolVarP(&ipv6PtrAddParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&ipv6PtrAddParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&ipv6PtrAddParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&ipv6PtrAddParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&ipv6PtrAddParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&ipv6PtrAddParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&ipv6PtrAddParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&ipv6PtrAddParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&ipv6PtrAddParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&ipv6PtrAddParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&ipv6PtrAddParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&ipv6PtrAddParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&ipv6PtrAddParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(ipv6PtrAddNormalizeFlagNames)
	buildFlagsUsage(cmd, ipv6PtrAddFlagOrder(cmd))
	return cmd
}

func ipv6PtrReadCmd() *cobra.Command {
	ipv6PtrReadParam := params.NewPtrReadIPv6Param()
	cmd := &cobra.Command{
		Use: "ptr-read",

		Short:        "PtrRead IPv6",
		Long:         `PtrRead IPv6`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipv6PtrReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, ipv6PtrReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if ipv6PtrReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, ipv6PtrReadParam)
			}

			return funcs.IPv6PtrRead(ctx, ipv6PtrReadParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&ipv6PtrReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&ipv6PtrReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&ipv6PtrReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&ipv6PtrReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&ipv6PtrReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&ipv6PtrReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&ipv6PtrReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&ipv6PtrReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&ipv6PtrReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&ipv6PtrReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&ipv6PtrReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&ipv6PtrReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(ipv6PtrReadNormalizeFlagNames)
	buildFlagsUsage(cmd, ipv6PtrReadFlagOrder(cmd))
	return cmd
}

func ipv6PtrUpdateCmd() *cobra.Command {
	ipv6PtrUpdateParam := params.NewPtrUpdateIPv6Param()
	cmd := &cobra.Command{
		Use: "ptr-update",

		Short:        "PtrUpdate IPv6",
		Long:         `PtrUpdate IPv6`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipv6PtrUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, ipv6PtrUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if ipv6PtrUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, ipv6PtrUpdateParam)
			}

			// confirm
			if !ipv6PtrUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("ptr-update", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.IPv6PtrUpdate(ctx, ipv6PtrUpdateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&ipv6PtrUpdateParam.Hostname, "hostname", "", "", "set server hostname")
	fs.BoolVarP(&ipv6PtrUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&ipv6PtrUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&ipv6PtrUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&ipv6PtrUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&ipv6PtrUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&ipv6PtrUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&ipv6PtrUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&ipv6PtrUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&ipv6PtrUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&ipv6PtrUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&ipv6PtrUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&ipv6PtrUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&ipv6PtrUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(ipv6PtrUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, ipv6PtrUpdateFlagOrder(cmd))
	return cmd
}

func ipv6PtrDeleteCmd() *cobra.Command {
	ipv6PtrDeleteParam := params.NewPtrDeleteIPv6Param()
	cmd := &cobra.Command{
		Use: "ptr-delete",

		Short:        "PtrDelete IPv6",
		Long:         `PtrDelete IPv6`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipv6PtrDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, ipv6PtrDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if ipv6PtrDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, ipv6PtrDeleteParam)
			}

			// confirm
			if !ipv6PtrDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("ptr-delete", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.IPv6PtrDelete(ctx, ipv6PtrDeleteParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&ipv6PtrDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&ipv6PtrDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&ipv6PtrDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&ipv6PtrDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&ipv6PtrDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&ipv6PtrDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&ipv6PtrDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&ipv6PtrDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&ipv6PtrDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&ipv6PtrDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&ipv6PtrDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&ipv6PtrDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&ipv6PtrDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(ipv6PtrDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, ipv6PtrDeleteFlagOrder(cmd))
	return cmd
}
