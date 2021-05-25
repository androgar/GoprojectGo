// You can edit this code!
// Click here and start typing.
package main

import ("fmt" 
        "math"
		)

func main() {
     var a,b,c,d float64
	 	 fmt.Println("a=")
	 fmt.Scanf("%v \n", &a)
	 	 fmt.Println("b=")
	 fmt.Scanf("%v \n", &b)
	 	 fmt.Println("c=")
	 fmt.Scanf("%v \n", &c)
	  	 d=b*b-4*a*c
	 	 fmt.Println("\n")
	 if d>0 {
	   fmt.Println("koren № 1=", (-1*b-math.Pow(d,0.5))/(2*a))
	   fmt.Println("koren № 2=", (-1*b+math.Pow(d,0.5))/(2*a))
	 }else if d==0 {
	   fmt.Println("koren =", (-1*b-math.Pow(d,0.5))/(2*a))
	 }else if d<0 {
	   fmt.Println("kornei net ")
	 }
	     
}