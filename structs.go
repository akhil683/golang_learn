package main

type car struct {
 Make string
 Model string
}

func main()  {
  test(messageToSend{
    Make:  'Akhil',
    Model:  "i123",
  })
}

func test(m car) {
  fmt.Printf("Sending message: "%s" to: %v\n", m.Make, m.Model)
}


//NOTE: Nested Struct

// type car struct {
//  Make string
//  Model string
//  Height int
//  Width int
//  FrontWheel Wheel
//  BackWheel Wheel
// }
// type Wheel struct {
//   Radius int
//   Material string
// }
// myCar := car{}
// myCar.FrontWheel.Radius = 5
//
