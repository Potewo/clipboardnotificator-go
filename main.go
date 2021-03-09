package main
import "github.com/gen2brain/beeep"

func main() {
  err := beeep.Notify("Title", "Hello from go", "Icon.jpeg");
  if err != nil {
    panic(err)
  }
}


