package lib


import(
    "net/http"
    "net/url"
    "fmt"
    "strconv"
    "encoding/json"
)

type Obj interface{}

type ContainerData struct{
    ContainerId string `json:"containerId"`
    Port int `json:"port"`
}

type Msg struct{
    Code int `json:"code"`
    Message string `json:"message"`
    Result Obj `json:"result"`
}

/*
*输出结果
*@param w http.ResponseWriter对象
*@param r http.Request指针对象
*@param 
*/
func outputData(w http.ResponseWriter, r *http.Request, code int, message string, ownData interface{}){
    m := Msg{code, message, ownData}
    data, _ := json.Marshal(m)
    w.Write(data)
}

/*
*处理容器操作
*@param w http.ResponseWriter对象
*@param r http.Request指针对象
*/
func handleContainer(w http.ResponseWriter, r *http.Request){
    
    appId, _ := strconv.Atoi(r.FormValue("appId"))
    token := r.FormValue("token")
    option := r.FormValue("option")
    //校验密码
    if !CheckToken(token){
        outputData(w, r, 2, "token error", nil)
        return
    }
    
    //创建容器
    if option == "create"{
        imageName := r.FormValue("imageName")
        
        //检测参数
        if imageName == ""{
            outputData(w, r, 6, "Parameter missing", nil)
            return
        }
        
        port := GetAbleUsePort()
        containerId, _ := BuildContainer(imageName, port, appId)
        obj := ContainerData{containerId, port}
        outputData(w, r, 0, "ok", obj)
        return
    }
    
    containerId := r.FormValue("containerId")
    if containerId == ""{
        //参数缺失
        outputData(w, r, 6, "Parameter missing", nil)
        return
    }
    //基本容器操作
    if option == "start"{
        StartContainer(containerId)
    }else if option == "stop"{
        StopContainer(containerId)
    }else if option == "remove"{
        RemoveContainer(containerId)
    }else if option == "restart"{
       RestartContainer(containerId)
    }
    outputData(w, r, 0, "ok", nil)
}


/*
*处理git操作
*@param w http.ResponseWriter对象
*@param r http.Request指针对象
*/
func handleGitOption(w http.ResponseWriter, r *http.Request){
    
    appIdValue := r.FormValue("appId")
    appId, err := strconv.Atoi(appIdValue)
    gitUrl := r.FormValue("gitUrl")
    token := r.FormValue("token")

    if CheckToken(token) && err == nil && gitUrl != ""{
        gitUrl, _ = url.QueryUnescape(gitUrl)
        UpdateCode(appId, gitUrl)
        UpdateApplicationCode(appId)
        outputData(w, r, 0, "ok", nil)
    }
    
}


/*
*路由处理
*@param w http.ResponseWriter对象
*@param r http.Request指针对象
*/
func handleHttp(w http.ResponseWriter, r *http.Request){
   if r.URL.Path == "/git"{
        handleGitOption(w, r);
    }else if r.URL.Path == "/container"{
        handleContainer(w, r)
    }else{
        outputData(w, r, 1, "not found", nil)
    }
}


/*
*服务器监听
*
*/
func StartServer(){
    fmt.Println("服务器启动！")
    http.HandleFunc("/", handleHttp)
    http.ListenAndServe("0.0.0.0:8000", nil)
}