package main
import "fmt"
 
func main(){
 var j,k,sum int = 1,0,0;
 
 for j < 4000000 {
   j,k = j+k,j;
 
   if j%2 == 0 {
     sum += j;
   }
 }
 
 fmt.Println(sum)
 
}
