package lib

import(
    "os/exec"
    "fmt"
)


/*
*创建容器
*@param imageName 镜像名称
*@return 新建容器的id（string）， 创建结果（bool）
*/
func BuildContainer(imageName string)(containerId string, result bool){
    
    containerId = "0"
    result = true
    
    command := "docker run -d centos:6 /bin/bash -c 'while true;do sleep 500;done'"
    cmd := exec.Command("/bin/bash", "-c", command)
    data,_ := cmd.Output()
    containerId = string(data)
    cmd.Start()
    fmt.Println("执行命令：" + command)
    
    return containerId, result
}


/*
*启动容器
*
*/
func StartContainer(containerId string){
    
    command := "docker start " + containerId
    cmd := exec.Command("/bin/bash", "-c", command)
    cmd.Start()
    fmt.Println("执行命令：" + command)
}