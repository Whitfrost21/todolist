package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
var task string
var ts []Task
var addCmd = &cobra.Command{
  Use:"add",
  Short:"add the task to the list",
  Args:cobra.MinimumNArgs(1),
  Run:func(cmd *cobra.Command,args []string){

    task:=args[0]
    
  
    file,_:=os.Open("task.json")
    byteval,_:=os.ReadFile("task.json")
    json.Unmarshal(byteval,&ts)

    nextid:=1
    if len(ts)>0{
      nextid=ts[len(ts)-1].Id
    }

    container:=Task{
      Id: nextid,
      Title:task,
      Status:false,
    } 
    defer file.Close()
    ts=append(ts,container)
    udata,_:=json.MarshalIndent(ts,"","")

    err := os.WriteFile("task.json",udata,0644)
    if err!=nil{
      fmt.Println("error:",err)
    }
    fmt.Println("added:",task)
   
  },
}

func init(){
  addCmd.PersistentFlags().StringVarP(&task,"add a task","n","","add the task to todolist")
}


