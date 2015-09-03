package main

import(
    "./lib"
)

func main(){
    lib.SetProcessSetting()
    lib.UpdateCode(1, "git@github.com:yubang/dockerManager.git")
}