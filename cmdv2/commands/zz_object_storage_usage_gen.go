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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-usage'; DO NOT EDIT

package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func objectStorageCommandOrder(cmd *cobra.Command) []*commandSet {
	var commands []*commandSet
	{
		set := &commandSet{
			title: "Basics",
		}
		set.commands = append(set.commands, lookupCmd(cmd, "list"))
		set.commands = append(set.commands, lookupCmd(cmd, "put"))
		set.commands = append(set.commands, lookupCmd(cmd, "get"))
		set.commands = append(set.commands, lookupCmd(cmd, "delete"))
		commands = append(commands, set)
	}

	return commands
}

func objectStorageListFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("auth", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("access-key"))
		fs.AddFlag(cmd.LocalFlags().Lookup("secret-key"))
		fs.AddFlag(cmd.LocalFlags().Lookup("bucket"))
		sets = append(sets, &flagSet{
			title: "Auth options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("column"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &flagSet{
			title: "Output options",
			flags: fs,
		})
	}

	return sets
}

func objectStoragePutFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("auth", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("access-key"))
		fs.AddFlag(cmd.LocalFlags().Lookup("secret-key"))
		fs.AddFlag(cmd.LocalFlags().Lookup("bucket"))
		sets = append(sets, &flagSet{
			title: "Auth options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("operation", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("content-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("recursive"))
		sets = append(sets, &flagSet{
			title: "Operation options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}

	return sets
}

func objectStorageGetFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("auth", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("access-key"))
		fs.AddFlag(cmd.LocalFlags().Lookup("secret-key"))
		fs.AddFlag(cmd.LocalFlags().Lookup("bucket"))
		sets = append(sets, &flagSet{
			title: "Auth options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("operation", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("recursive"))
		sets = append(sets, &flagSet{
			title: "Operation options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}

	return sets
}

func objectStorageDeleteFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("auth", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("access-key"))
		fs.AddFlag(cmd.LocalFlags().Lookup("secret-key"))
		fs.AddFlag(cmd.LocalFlags().Lookup("bucket"))
		sets = append(sets, &flagSet{
			title: "Auth options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("operation", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("recursive"))
		sets = append(sets, &flagSet{
			title: "Operation options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}

	return sets
}
