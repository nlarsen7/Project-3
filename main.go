package main

import "fmt"

func well() (radius float64, depth float64){
  var r,d float64
  fmt.Println("Enter the radius of the well casing in inches")
  fmt.Scanln(&r)
  fmt.Println("enter the depth of the well in feet")
  fmt.Scanln(&d)
  r=r/12.0
  return r,d
}
func gallons(){
r,d:=well()
gallons:=(r*r*d*3.14)*7.5
fmt.Println("The amount of gallons of water the well holds is",gallons,"gallons of water")
}
func main(){
gallons()
}