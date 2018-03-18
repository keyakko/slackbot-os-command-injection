package main

import "os"
import "io/ioutil"

func main() {
  content := []byte("hello world\n")
  ioutil.WriteFile("test.stdout", content, os.ModePerm)
}

