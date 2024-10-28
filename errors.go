func main() {
  user, err :- getUser()
  if err!= nil {
    fmt.Println(err)
    return
  }
  //user user here
}

//NOTE: Javascript error handling

// function main() {
//   try {
// const user = getUser()
// // user user here
//   } catch (err) {
// console.log(err)
//   }
// }
