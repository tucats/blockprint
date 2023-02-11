package main

import "github.com/tucats/gopackages/app-cli/cli"

var grammar = []cli.Option{
	{
		LongName:    "fonts",
		Description: "List fonts",
		Action:      FontsAction,
		OptionType:  cli.Subcommand,
	},
	{
		LongName:             "print",
		Description:          "Print text",
		ParametersExpected:   1,
		Action:               PrintAction,
		OptionType:           cli.Subcommand,
		DefaultVerb:          true,
		ParameterDescription: "\"text to print\"",
		Value: []cli.Option{
			{
				LongName:    "font",
				ShortName:   "f",
				Description: "font name",
				OptionType:  cli.StringType,
			},
		},
	},
}
