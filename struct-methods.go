type rect struct {
  width  int
  height i
}

//struct method 
func (r rect) area() int {
  return r.width* r.height
}

r:= rect{
  width: 5,
  height: 10,
}
fmt.Println(r.area())
