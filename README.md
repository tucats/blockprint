# blockprint

Blockprint text in a variety of fonts.

The `blockprint` command line interface (CLI) accepts a phrase or string
of text, and an optional font name, and prints it to the stdout as "block
print" text.

This makes use of the [go-figure package](https://github.com/common-nighthawk/go-figure)
by Daniel Deutsch.

## General Usage

```text
Usage:
   blockprint [options] [command]    Print large text, 1.0-3

Commands:
    config                           View or set application configuration   
    fonts                            List fonts                              
    help                             Show this help text                     
    version                          Show application version                
   *print                            Print text                              

(*) indicates the default subcommand if none given

Options:
   --help, -h                        Show this help text                    
   --profile, -p <string>            Name of profile to use [APP_PROFILE]   
   --version, -v                     Show application version     
```

The `blockprint` command accepts a command, which can show the preferences
configuration, list the avialable fonts, show the help text and version of
the CLI, and print a string of text.

You can get additional help by using the --help option after any of the
above subcommands.

## Printing Text

Use the `print` subcommand to print text. The subcommand is followed by a
parameter that is a quoted string containing the text to print. This can
also optionally include a font name if something other than the default
fault is desired.

```sh
blockprint print "Hello" --font italic
```

This prints the text "Hello" (without the quotation marks, of course) to
stdout, using the font named `italic`. An error is reported if the font
name is not valid.

## Listing Fonts

Use the `fonts` subcommand to print the list of available fonts. These fonts
are all built in to the `blockprint` command line tool, and there are
currently no supported means of extending this font set.

```sh
blockprint fonts
```

The output of the command is a listing of the font names.

## Default Font

Use the `config` subcommand to view the default fault, or to
set the name of the font.

```sh
blockprint config set default.font=italic
```

This sets the default font name to "italic". If a font is not
specified on the command line, this is the font that will be
used. The initial default font is "standard".

You can also use the `config show` command to see the current
default font name.

```sh
blockprint config show
```
