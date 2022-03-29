package main

func main() {
  // Load the config file.
  import_config()

  // Start game loop
  for true {
    proceed := kgit_gameloop()
    if !proceed {
      break
    }
  }
}
