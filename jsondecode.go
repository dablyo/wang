package main

/*func main(){
  const jsonStream = `
    { "Name":"Ed", "Text": "Knock knock"}
    { "Name":"Sam","Text": "Who's there"}
  `
  type Message struct {
    Name,Text string
  }
  dec := json.NewDecoder(strings.NewReader(jsonStream))
  for {
    var m Message
    if err := dec.Decode(&m); err == io.EOF{
      break
    } else if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%s:%s\n",m.Name,m.Text)
  }
}*/
