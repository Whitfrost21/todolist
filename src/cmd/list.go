package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
var listen string
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
    file,err:=os.Open("task.json")    
    if err!=nil{
      panic(err)
    }
    var tasklist[] Task

    defer file.Close()

    databytes,err:=os.ReadFile("task.json")
    json.Unmarshal(databytes,&tasklist)

    if err!=nil{
      fmt.Println(err)
    }

    for _,t:=range tasklist{
      fmt.Println(t.Title," status:",t.Status)
    }

    
    },
}

func init() {
  listCmd.PersistentFlags().StringVarP(&listen,"list all task","l","","list all todos")   
}
 
