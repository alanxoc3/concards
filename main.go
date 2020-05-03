package main

import (
   "fmt"
   "github.com/akamensky/argparse"
   "os"
)

func main() {
   // Create new parser object
   parser := argparse.NewParser("concards", "Concards is a simple CLI based SRS flashcard app.")

   // Create flags
   f_review   := parser.Flag("r", "review",   &argparse.Options{Help: "Show cards available to be reviewed"})
   f_memorize := parser.Flag("m", "memorize", &argparse.Options{Help: "Show cards available to be memorized"})
   f_done     := parser.Flag("d", "done",     &argparse.Options{Help: "Show cards not available to be reviewed or memorized"})
   f_print    := parser.Flag("p", "print",    &argparse.Options{Help: "Prints all cards, one line per card"})
   f_number   := parser.Int("n", "number",    &argparse.Options{Help: "Limit the number of cards in the program to \"#\""})
   f_editor   := parser.String("E", "editor", &argparse.Options{Help: "Which editor to use. Defaults to \"$EDITOR\""})
   f_meta     := parser.String("M", "meta",   &argparse.Options{Help: "Path to meta file. Defaults to \"$CONCARDS_META\" or \"~/.concards-meta\""})

   parser.HelpFunc = func(c *argparse.Command, msg interface{}) string {
      var help string
      help += fmt.Sprintf("%s\n\nUsage:\n  %s [OPTIONS]... [FILE|FOLDER]...\n\nOptions:\n", c.GetDescription(), c.GetName())

      for _, arg := range c.GetArgs() {
         help += fmt.Sprintf("  -%s  --%-9s %s.\n", arg.GetSname(), arg.GetLname(), arg.GetOpts().Help)
      }

      return help
   }

   // Parse input
   err := parser.Parse(os.Args)
   if err != nil {
      fmt.Print(parser.Help(nil))
      os.Exit(1)
	}

   if *f_review || *f_memorize || *f_done || *f_print || f_editor != nil || f_number != nil || f_meta != nil {
      println("Hello World")
   }
}
