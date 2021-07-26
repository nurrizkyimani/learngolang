package main

import (
  "fmt"
  "encoding/json"

)

type Book struct {
  Title string `json:"title"`
  Author Author `json:author`
}

type Author struct {
  Name string `json:"name"`
  Age int `json:"age"`
  Developer bool `json:is_developer`
}

type SensorReading struct {
  Name string `json:"name"`
  Capacity int `json:"capacity"`
}


func main() {
  fmt.Println("Hello World")

  auth2 := Author{Name:"Ell", Age: 23, Developer: true}

  book := Book{Title : "Testing title", Author: auth2 }

  fmt.Printf("%+v\n", book)

  byteArray, err := json.MarshalIndent(book, "", "")
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(byteArray))

  jsonCap := `{"name" : "b sensor", "capacity": 40, "info" : {"desc": "a sensor reading"}}`

  var reading map[string]interface{}

  err2 := json.Unmarshal([]byte(jsonCap), &reading)

  fmt.Println("print the reading")

  if err2 != nil {
    fmt.Println(err)
  }

  fmt.Printf("%+v\n", reading)
}