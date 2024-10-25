package main

import (
	"bufio"
	"fmt"
	"os"
)

type messageToSend struct{
    message string
    sender int
    recipient bool
}

func main(){
    reader:= bufio.NewReader(os.Stdin)
    
    fmt.Print("Enter your message:")

    message,_ := reader.ReadString('\n')

    fmt.Printf("You wished: %s\n",message)

}
