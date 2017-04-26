package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func sendBack(newTab [20][20]int, tab []string)  {
  var x = 0
  var b = 0
  for x < 19 {
    var y = 0
    for y < 19 {
      tab[b] = strconv.Itoa(newTab[x][y])
      y++
      b++
    }
    x++
  }
  file, err := os.Create("tabToSend")
  if err != nil {
    fmt.Print(err)
  }
  file.WriteString(strings.Join(tab,""))
return
}

func main () {

  newTab := [20][20]int{}
  var x = 0
  var b = 0

  if 2 == len(os.Args) {
    tab := strings.Split(os.Args[1],",")
    for x < 19 {
      var y = 0
      for y < 19 {
        newTab[x][y], _= strconv.Atoi(tab[b])
        y++
        b++
      }
      x++
    }
    newTab = algo(newTab)
    sendBack(newTab, tab)
   return
 }
  return
}
