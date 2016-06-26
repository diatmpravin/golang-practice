package main
import "fmt"

func main() {
    var mealCost float64  
    var tipPercent int 
    var taxPercent int

    fmt.Scan(&mealCost)
    fmt.Scan(&tipPercent)
    fmt.Scan(&taxPercent)
    
    fmt.Printf("The total meal cost is %.f dollars. \n", float64(mealCost) + float64(mealCost) * float64(tipPercent) * .01 + float64(mealCost) * float64(taxPercent) * .01)
}