package main
import "fmt"
func digitSama(awal int, akhir int){
    var b bool
    var i, x, y int
    for awal <= akhir{
        i = awal
        for i > 0{
            x = i%10
            y = (i/10)%10
            if x == y{
                b = true
            }
            i = i/10
        }
        if b == true{
            fmt.Println(awal)
        }
        awal++
        b = false
    }
}
func main(){
    var d1, d2 int
    fmt.Scan(&d1, &d2)
    digitSama(d1, d2)
}