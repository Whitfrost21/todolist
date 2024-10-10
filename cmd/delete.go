package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
var delt string
var deleteCmd=&cobra.Command{
  Use:"delete",
  Short:"delete the enterd task ",
  Args:cobra.MinimumNArgs(1),
  Run:func(cmd *cobra.Command,args []string){
    task:=args[0]
    file,err:=os.Open("task.json")
    if err!=nil{
      fmt.Println("error:",err)
    }
    defer file.Close()
    var tasklist[] Task

    databytes,err:=os.ReadFile("task.json")
    json.Unmarshal(databytes,&tasklist)

    for i,t:=range tasklist{
      if t.Title==task{
        tasklist=append(tasklist[:i],tasklist[i+1:]...)  
        break
      }
    }

    updateddata,_:=json.MarshalIndent(tasklist,"","")
    err=os.WriteFile("task.json",updateddata,0644)
    if err!=nil{
      fmt.Println("error at writing file")
    }
    
    fmt.Println("deleted task:",task)

     },
}
func init()  {
 deleteCmd.PersistentFlags().StringVarP(&delt,"delete a task","d","","remove the completed task from todolist") 
}
