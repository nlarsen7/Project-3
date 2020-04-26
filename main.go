//Nicholas Larsen
//4/25/2020
//asks user for dimensions on a well and then tells the user how many gallons of water can be stored in that well
package main

import "fmt"

func well() (radius float64, depth float64){
  var r,d float64
  fmt.Println("Enter the radius of the well casing in inches")
  fmt.Scanln(&r)
  fmt.Println("enter the depth of the well in feet")
  fmt.Scanln(&d)
  //finds the radius and depth of the well
  r=r/12.0
  //divides radius by 12 to put it into feet
  return r,d
  //returns the radius and the depth
}
func gallons(){
r,d:=well()
//recalls the radius and the depth
gallons:=(r*r*d*3.14)*7.5
//finds the volume of a cylinder then multiplies that by 7.5 since that's how many gallons are in a cubic foot
fmt.Println("The amount of gallons of water the well holds is",gallons,"gallons of water")
//Prints out the solution
}
func main(){
gallons()
//recalls from gallons function
}