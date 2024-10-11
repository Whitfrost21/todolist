package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
var margin string 
var taskl[] Task
var didCmd = &cobra.Command{
	Use:   "mark",
	Short: "mark task as complete",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
    task:=args[0]
    file,err:=os.Open("task.json")
    
    if err!=nil{
      fmt.Println(err)
    }
    defer file.Close()

    databytes,_:=os.ReadFile("task.json")
    json.Unmarshal(databytes,&taskl)

    for i,t:=range taskl{
      if t.Title==task{
        t.Status=true
        
        st:=fmt.Sprintf("+%v",t)
        histo,err:=os.OpenFile("todos.txt",os.O_WRONLY|os.O_APPEND,0644)
        if err!=nil{
          fmt.Println(err)
        }
        _,err=histo.WriteString(st+"\n")
        if err!=nil{
          fmt.Println("error:",err)
        }
        defer histo.Close()
        taskl=append(taskl[:i],taskl[i+1:]...)
        for i := range taskl {        taskl[i].Id=i+1
        }
        updatedlist,_:=json.MarshalIndent(taskl,"","")
        os.WriteFile("task.json",updatedlist,0644)
        fmt.Println("marked as completed:",task)
      }
    }
	},
}


func init(){
  didCmd.PersistentFlags().StringVarP(&margin,"mark","m","","mark task as completed,the task will be removed from list and showed in history")
}
