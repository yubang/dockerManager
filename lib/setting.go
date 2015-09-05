package lib

/*
*处理程序自身的基本配置
*
*/

import(
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

/*
 *获取应用常量
 *
 */
func GetApplicationValueFromKey(key string)(value string){
    
    return "123"
}

/*
*信息校验秘钥判定
*@param token 需要检测的密码
*@param result(bool) 校验结果
*/
func CheckToken(token string)(result bool){
    return token == GetSettingValue("token")
}

/*
 *把参数写入配置文件
*/
func writeSettingFile(data map[string]string){
    
    fp, err := os.OpenFile("data/config/base.conf", os.O_CREATE|os.O_WRONLY, 0666)
    defer fp.Close()
    
    if err == nil{
        fp.WriteString("gitPath=" + data["gitPath"] + "\n")
        fp.WriteString("applicationPath=" + data["applicationPath"] + "\n")
        fp.WriteString("token=" + data["token"] + "\n")
        fmt.Println("配置文件创建成功！")
    }else{
        fmt.Println("创建配置文件失败！")
    }
    
}

/*
 *读取控制台输入
 */
func handleWrite(){
    var gitPath, applicationPath, token string
    data := make(map[string]string)
    
    fmt.Println("请输入git仓库地址")
    fmt.Scanln(&gitPath)
    fmt.Println("请输入代码仓库地址")
    fmt.Scanln(&applicationPath)
    fmt.Println("请输入校验密码")
    fmt.Scanln(&token)
    
    data["gitPath"] = gitPath
    data["applicationPath"] = applicationPath
    data["token"] = token
    
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
    fp, err := os.Open("data/config/base.conf")
    defer fp.Close()
    if err == nil{
        reader := bufio.NewReader(fp)
        for{
            text, err := reader.ReadString('\n')
            if err == nil{
                arrs := strings.Split(text,"=")
                if arrs[0] == key{
                    return strings.Replace(arrs[1], "\n", "", -1)
                }
            }else{
                break
            }
        }
    }
    return ""
}


/*
*判断文件夹是否存在
*
*/
func checkDirExist(dirPath string)(result bool){
    fi, err := os.Stat(dirPath)
    if err != nil {
        return os.IsExist(err)
    } else {
        return fi.IsDir()
    }
}


/*
*获取一个可用的端口号
*
*/
func GetAbleUsePort()(port int){
    port = 20000
    
    _, err := os.Stat("data/config/port.txt")
    if os.IsNotExist(err){
        fi, _ := os.Create("data/config/port.txt")
        defer fi.Close()
        fi.Write([]byte(strconv.Itoa(port)))
    }else{
        fi, _ := os.Open("data/config/port.txt")
        reader := bufio.NewReader(fi)
        text, _ := reader.ReadString('\n')
        fi.Close()
        
        fi, _ = os.Create("data/config/port.txt")
        port, _ = strconv.Atoi(text)
        port ++
        fi.WriteString(strconv.Itoa(port))
        fi.Close()
    }
    
    return port
}