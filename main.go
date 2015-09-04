package main

import(
    "./lib"
    "fmt"
)

func main(){
    lib.SetProcessSetting()
    //lib.UpdateCode(1, "git@github.com:yubang/dockerManager.git")
    //lib.UpdateApplicationCode(1)
    containerId, _ := lib.BuildContainer("php5.3.3",8000,"/tmp/app/1")
    //lib.StartContainer(containerId)
    //lib.RemoveContainer(containerId)
    fmt.Println(containerId)
}