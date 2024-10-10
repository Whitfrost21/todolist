package cmd

import (
	
	"os"
  "fmt"
	"github.com/spf13/cobra"
)
var enter string
var truncCmd = &cobra.Command{
  Use:"trunc",
  Short:"truncate/clear the todolist",
  Run:func(cmd *cobra.Command,args []string){
    _,err:=os.OpenFile("task.json",os.O_TRUNC|os.O_WRONLY,0644)
    if err!=nil{
      fmt.Println("error:",err)
    }
fmt.Println("truncated succesfully!")
        
      
  },
}

func init(){
  truncCmd.PersistentFlags().StringVarP(&enter,"trunc","t","","truncate/clear the todolist")
}
