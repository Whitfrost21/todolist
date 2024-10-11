package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)
var com string
var histoCmd=&cobra.Command{
  Use:"history",
  Short:"load history of todos",
  Run:func(cmd *cobra.Command,args []string){
    file,err:=os.Open("todos.txt")
    if err!=nil{
      fmt.Println("error:",err)
    }
    defer file.Close()
    scanner:=bufio.NewScanner(file)
    for scanner.Scan(){
        str:=scanner.Text()
        str=strings.Trim(str,".+{1234567890}")
        fmt.Println(str)
    }
    if scanner.Err();scanner.Err()!=nil{
      fmt.Println("error :",scanner.Err())
    }
  },
}

func init(){
  histoCmd.PersistentFlags().StringVarP(&com,"his","l","","show history of tasks")
}
