package main

import (
   "os"
   "fmt"
   "github.com/alanxoc3/concards/file"
   "github.com/alanxoc3/concards/core"
)

func main() {
   c := file.GenConfig()
   d := core.NewDeck()

   for _, f := range c.Files {
      if err := file.ReadCardsToDeck(f, d); err != nil {
         fmt.Printf("Error: File \"%s\" does not exist!\n", f)
         os.Exit(1)
      }
   }

   if c.IsPrint {
      for i := 0; i < d.Len(); i++ {
         fmt.Printf("@> %s\n", d.GetCard(i).String())
      }

      fmt.Printf("<@\n")
      return
   }
}
