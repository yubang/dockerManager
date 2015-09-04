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
    if os.IsExist(err){
        command = "cd " + gitPath + " && git pull origin master"
    }else{
        command = " git clone " + gitUrl + " " + gitPath
    }
        
    fmt.Println("执行命令：" + command)
    cmd := exec.Command("/bin/bash", "-c", command)
    cmd.Run()
    
}


/*
 *更新应用代码
 *@param appId 应用ID
 */
func UpdateApplicationCode(appId int){
    
    gitPath := GetSettingValue("gitPath") + "/" + strconv.Itoa(appId)
    applicationPath := GetSettingValue("applicationPath") +  "/" + strconv.Itoa(appId)   +  "/code/"
    
    command := "mkdir -p " + applicationPath
    cmd := exec.Command("/bin/bash", "-c", command)
    cmd.Run()
    fmt.Println("执行命令：" + command)
    
    command = "rm -rf " + applicationPath
    command =  command + " && cp -r " + gitPath + " " + applicationPath
    cmd = exec.Command("/bin/bash", "-c", command)
    cmd.Run()
    fmt.Println("执行命令：" + command)
}