package main

import (
"fmt"
"github.com/gopherjs/gopherjs/js"
)

type Pet struct {
  name string
}

func New(name string) *js.Object {
  fmt.Printf("oeuoeu")
  return js.MakeWrapper(&Pet{name})
}

func (p *Pet) Name() string {
  return p.name
}

func (p *Pet) SetName(name string) {
  p.name = name
}

func main() {
  js.Module.Get("exports").Set("pet", map[string]interface{}{
    "New": New,
  })
}
