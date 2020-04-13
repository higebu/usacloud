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

func isoImageCommandOrder(cmd *cobra.Command) []*commandSet {
	var commands []*commandSet
	{
		set := &commandSet{
			title: "Basics",
		}
		set.commands = append(set.commands, lookupCmd(cmd, "list"))
		set.commands = append(set.commands, lookupCmd(cmd, "create"))
		set.commands = append(set.commands, lookupCmd(cmd, "read"))
		set.commands = append(set.commands, lookupCmd(cmd, "update"))
		set.commands = append(set.commands, lookupCmd(cmd, "delete"))
		commands = append(commands, set)
	}
	{
		set := &commandSet{
			title: "Upload/Download(SFTP)",
		}
		set.commands = append(set.commands, lookupCmd(cmd, "upload"))
		set.commands = append(set.commands, lookupCmd(cmd, "download"))
		set.commands = append(set.commands, lookupCmd(cmd, "ftp-open"))
		set.commands = append(set.commands, lookupCmd(cmd, "ftp-close"))
		commands = append(commands, set)
	}

	return commands
}

func isoImageListFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("scope"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("limit-offset", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("from"))
		fs.AddFlag(cmd.LocalFlags().Lookup("max"))
		sets = append(sets, &flagSet{
			title: "Limit/Offset options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("sort", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("sort"))
		sets = append(sets, &flagSet{
			title: "Sort options",
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

func isoImageCreateFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("ISO-Image", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("size"))
		fs.AddFlag(cmd.LocalFlags().Lookup("iso-file"))
		sets = append(sets, &flagSet{
			title: "ISO-Image options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("common", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		sets = append(sets, &flagSet{
			title: "Common options",
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

func isoImageReadFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
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
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func isoImageUpdateFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("common", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		sets = append(sets, &flagSet{
			title: "Common options",
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
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func isoImageDeleteFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
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
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func isoImageUploadFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("ISO-Image", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("iso-file"))
		sets = append(sets, &flagSet{
			title: "ISO-Image options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
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
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func isoImageDownloadFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("ISO-Image", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("file-destination"))
		sets = append(sets, &flagSet{
			title: "ISO-Image options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
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
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func isoImageFTPOpenFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
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
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func isoImageFTPCloseFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
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
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}
