package main
import (
    "fmt"
    "math"
)

func main() {
    var a float64
    var b float64
    var c float64
    
    fmt.Println("Digite o valor de A:")
    fmt.Scan(&a)
    
    fmt.Println("Digite o valor de B:")
    fmt.Scan(&b)
    
    fmt.Println("Digite o valor de C:")
    fmt.Scan(&c)
    
    var delta float64 = (b*b) - 4 * a * c
    
    if delta < 0 {
        fmt.Println("Delta menor que zero! Não existem raízes reais")
        return
    }
    
    if delta == 0 {
        resultado := (-b) / (2 * a)
        fmt.Println("Delta igual a zero. Raiz é igual a: ", resultado)
        return
    }
    
    resultado1 := (-b + math.Sqrt(delta)) / (2 * a)
    resultado2 := (-b - math.Sqrt(delta)) / (2 * a)
    
    fmt.Println("Delta positivo! As raízes são: ", resultado1, " ", resultado2)
}
