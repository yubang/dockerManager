package lib


import(
    "net/http"
    "net/url"
    "fmt"
    "strconv"
    "encoding/json"
)

type Msg struct{
    Code int `json:"code"`
    Message string `json:"message"`
}

/*
*输出结果
*@param w http.ResponseWriter对象
*@param r http.Request指针对象
*/
func outputData(w http.ResponseWriter, r *http.Request, code int, message string){
    m := Msg{code, message}
    data, _ := json.Marshal(m)
    w.Write(data)
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
        outputData(w, r, 0, "ok")
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