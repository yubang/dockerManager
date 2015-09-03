package lib

import(
    "fmt"
    "strconv"
    "os"
    "os/exec"
)

/*
 *更新代码仓库
 *@param appId 应用id
 *@param gitUrl git地址
*/
func UpdateCode(appId int, gitUrl string){
    
    var command string
    gitPath := GetSettingValue("gitPath") + "/" + strconv.Itoa(appId)
    
    //判断目录是否存在
    _,  err := os.Stat(gitPath)
    if err == nil{
        command = "cd " + gitPath + " && git pull origin master"
    }else{
        command = " git clone " + gitUrl + " " + gitPath
    }
        
    fmt.Println(command)
    cmd := exec.Command("/bin/bash", "-c", command)
    err = cmd.Start()
    
     fmt.Println(err)
    
}