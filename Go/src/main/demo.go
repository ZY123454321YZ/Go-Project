package main

import (  
    "fmt"  
    "os"  
)  
  
func main() {  
    who := "world!"  
    if len(os.Args[1:]) > 0 {  
        who = ""  
        for _, arg := range os.Args[1:] {  
            who += " " + arg  
        }  
    }  
    fmt.Println("hello", who)
    fmt.Println("test"); 
}  

