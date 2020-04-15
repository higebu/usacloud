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
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/spf13/cobra"
)

// priceCmd represents the command to manage SAKURA Cloud Price
func priceCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "price",
		Aliases: []string{"public-price"},
		Short:   "A manage commands of Price",
		Long:    `A manage commands of Price`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDefaultCmd(cmd, args, "list")
		},
	}
}

func priceListCmd() *cobra.Command {
	priceListParam := params.NewListPriceParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find"},
		Short:        "List Price (default)",
		Long:         `List Price (default)`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return priceListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, priceListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if priceListParam.GenerateSkeleton {
				return generateSkeleton(ctx, priceListParam)
			}

			return funcs.PriceList(ctx, priceListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&priceListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &priceListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&priceListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&priceListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&priceListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&priceListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&priceListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&priceListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&priceListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&priceListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&priceListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&priceListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&priceListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&priceListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&priceListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&priceListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&priceListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(priceListNormalizeFlagNames)
	buildFlagsUsage(cmd, priceListFlagOrder(cmd))
	return cmd
}
