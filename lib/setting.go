package lib

/*
*处理程序自身的基本配置
*
*/

import(
    "os"
    "fmt"
)

/*
 *把参数写入配置文件
*/
func writeSettingFile(data map[string]string){
    
    fp, err := os.OpenFile("data/config/base.conf", os.O_CREATE|os.O_WRONLY, 0666)
    defer fp.Close()
    
    if err == nil{
        fp.WriteString("gitPath=" + data["gitPath"] + "\n")
        fp.WriteString("applicationPath=" + data["applicationPath"] + "\n")
        fmt.Println("配置文件创建成功！")
    }else{
        fmt.Println("创建配置文件失败！")
    }
    
}

/*
 *读取控制台输入
 */
func handleWrite(){
    var gitPath, applicationPath string
    data := make(map[string]string)
    
    fmt.Println("请输入git仓库地址")
    fmt.Scanln(&gitPath)
    fmt.Println("请输入代码仓库地址")
    fmt.Scanln(&applicationPath)
    
    data["gitPath"] = gitPath
    data["applicationPath"] = applicationPath
    
   writeSettingFile(data)
}

/*
*    设置管理程序的一些基本配置的入口函数（仅仅为了第一次使用调用）
*/
func SetProcessSetting(){
    _, err := os.Stat("data/config/base.conf")
    if os.IsNotExist(err){
        fmt.Println("第一次使用，需要配置一些参数")
        handleWrite();
    }
}


/*
 *获取配置参数
 */
func GetSettingValue(key string) (value string){
    return "/tmp/git"
}