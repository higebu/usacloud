// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

func init() {
	listParam := NewListPriceParam()

	cliCommand := &cli.Command{
		Name:    "price",
		Aliases: []string{"public-price"},
		Usage:   "A manage commands of Price",
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List Price",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "from",
						Usage:       "set offset",
						Destination: &listParam.From,
					},
					&cli.Int64SliceFlag{
						Name:  "id",
						Usage: "set filter by id(s)",
					},
					&cli.IntFlag{
						Name:        "max",
						Usage:       "set limit",
						Destination: &listParam.Max,
					},
					&cli.StringSliceFlag{
						Name:  "name",
						Usage: "set filter by name(s)",
					},
					&cli.StringSliceFlag{
						Name:  "sort",
						Usage: "set field(s) for sort",
					},
					&cli.StringFlag{
						Name:        "output-type",
						Aliases:     []string{"out"},
						Usage:       "Output type [json/csv/tsv]",
						Destination: &listParam.OutputType,
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:        "quiet",
						Aliases:     []string{"q"},
						Usage:       "Only display IDs",
						Destination: &listParam.Quiet,
					},
					&cli.StringFlag{
						Name:        "format",
						Aliases:     []string{"fmt"},
						Usage:       "Output format(see text/template package document for detail)",
						Destination: &listParam.Format,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, listParam)

					// Set option values for slice
					listParam.Id = c.Int64Slice("id")
					listParam.Name = c.StringSlice("name")
					listParam.Sort = c.StringSlice("sort")
					listParam.Column = c.StringSlice("column")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								PriceListCompleteArgs(ctx, listParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										PriceListCompleteArgs(ctx, listParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									PriceListCompleteFlags(ctx, listParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							PriceListCompleteArgs(ctx, listParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					listParam.Id = c.Int64Slice("id")
					listParam.Name = c.StringSlice("name")
					listParam.Sort = c.StringSlice("sort")
					listParam.Column = c.StringSlice("column")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Validate specific for each command params
					if errors := listParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), listParam)

					// Run command with params
					return PriceList(ctx, listParam)
				},
			},
		},
	}

	// build Category-Resource mapping
	appendResourceCategoryMap("price", &schema.Category{
		Key:         "information",
		DisplayName: "Service/Product informations",
		Order:       90,
	})

	// build Category-Command mapping

	appendCommandCategoryMap("price", "list", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})

	// build Category-Param mapping

	appendFlagCategoryMap("price", "list", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("price", "list", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("price", "list", "from", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("price", "list", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("price", "list", "max", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("price", "list", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("price", "list", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("price", "list", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483646,
	})
	appendFlagCategoryMap("price", "list", "sort", &schema.Category{
		Key:         "default",
		DisplayName: "Common options",
		Order:       2147483646,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
