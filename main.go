package main

import(
    "./lib"
)

func main(){
    lib.SetProcessSetting()
    //lib.UpdateCode(1, "git@github.com:yubang/dockerManager.git")
    //lib.UpdateApplicationCode(1)
    containerId, _ := lib.BuildContainer("centos:6")
    lib.StartContainer(containerId)
}