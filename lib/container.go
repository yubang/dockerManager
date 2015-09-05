package lib

import(
    "os/exec"
    "fmt"
    "strconv"
    "os"
)


/*
*创建容器
*@param imageName 镜像名称
*@param port 端口映射名称
*@param shareDirPath 镜像名称
*@return 新建容器的id（string）， 创建结果（bool）
*/
func BuildContainer(imageName string, port int, appId int)(containerId string, result bool){
    
    containerId = "0"
    result = true
    
    portData := strconv.Itoa(port) + ":80"
    shareDirPath := GetSettingValue("applicationPath") + "/" + strconv.Itoa(appId)
    
    //判断共享目录是否有文件，如果没有则需要创建必须文件
    if !checkDirExist(shareDirPath + "/code"){
        os.MkdirAll(shareDirPath+"/code", 0777)
    }
    if !checkDirExist(shareDirPath + "/log"){
        os.MkdirAll(shareDirPath+"/log", 0777)
    }
    
    shareData := "  -v  " + shareDirPath + ":/data "
    
    command := "docker run -d -p " + portData + shareData +  imageName + "  /bin/bash /var/script/start.sh"
    cmd := exec.Command("/bin/bash", "-c", command)
    data,_ := cmd.Output()
    containerId = string(data)
    cmd.Start()
    fmt.Println("执行命令：" + command)
    
    return containerId, result
}


/*
*启动容器
*@param containerId 容器ID
*/
func StartContainer(containerId string){
    command := "docker start " + containerId
    cmd := exec.Command("/bin/bash", "-c", command)
    cmd.Start()
    fmt.Println("执行命令：" + command)
}

/*
*停止容器
*@param containerId 容器ID
*/
func StopContainer(containerId string){
    command := "docker stop " + containerId
    cmd := exec.Command("/bin/bash", "-c", command)
    cmd.Run()
    fmt.Println("执行命令：" + command)
}

/*
*重启容器
*@param containerId 容器ID
*/
func RestartContainer(containerId string){
    command := "docker restart " + containerId
    cmd := exec.Command("/bin/bash", "-c", command)
    cmd.Start()
    fmt.Println("执行命令：" + command)
}

/*
*删除容器
*@param containerId 容器ID
*/
func RemoveContainer(containerId string){
    StopContainer(containerId)//删除容器前，先停止容器
    command := "docker rm " + containerId
    cmd := exec.Command("/bin/bash", "-c", command)
    cmd.Start()
    fmt.Println("执行命令：" + command)
}