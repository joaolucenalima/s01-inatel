package main
import (
    "fmt"
    "math/rand"
)

func main() {
    number := rand.Intn(10)
    var fatorial = 1
    
    for i := number; i >= 1; i-- {
           fatorial *= i
    }
    
    fmt.Println("Fatoria de ", number, " é igual a ", fatorial)
}
