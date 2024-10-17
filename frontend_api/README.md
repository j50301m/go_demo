# Frontend Api

## Introduction
`Frontend Api`是對前端(Web)公開的統一接口。

## Rule
1. 接口風格盡量符合RESTful所定義的規範。
2. 一般情況Response Body採用json格式回傳，並符合與前端約定之格式。
3. Token參數由Header傳入與傳出。

## 結構目錄
```
.
├── api
│   └── v1
├── config
├── doc
├── infrastructure
│   ├── client
│   └── httpserver
├── middleware
│   └── auth
├── route
└── test
```
資料夾說明:
- root: 會有一個`main.go`當作服務的啟動點, `.env` 或 `docker-compose.yaml`...與專案無直接相關的檔案先放在跟資料夾下。是否要將build相關的檔案另開資料夾，可以討論
- /api: 放置各版本的的api實作內容(handler / controller)
- /config: 放置設定相關的檔案
- /doc: 裡面放置OpenApi或其他相關資料的文檔
- /infrastructure: 放置基礎設施 如: gRPC的Client DI的啟動函數 Repo的實作...等
- /middleware: 放置中間件
- /route: 放置路由綁定的檔案
- /test: 放置測試,如果有private function 或小型的uni_test要想做測試 也可將測試檔建立在該資料夾底下(沒有硬性規定 但盡量還是放在/test內)

# Specifications
為了確保大家的api是可被驗證、擁有統一的格式，規範大家的handler寫法，一般情況下的handler 須滿足。(Streaming 與 SSE 或 Websocket 長連線的的回覆不在此限內)
1. 在處理邏輯前 開啟一個Trace
2. 回覆時使用responder封裝成統一格式

範例：
``` go

//  test struct  just for this example
type test struct {
    Msg   string            `json:"msg"`
    Array []string          `json:"array"`
    H     map[string]string `json:"map"`
}

func fooHandler(c *gin.Context) {
    // Start trace before you start
    _, span := kgsotel.StartTrace(c.Request.Context())
    defer span.End()

    // Do some logic...

    // Use responder for unified reply
    responder.Ok(test{
        Msg:   "Hello, World!",
        Array: []string{"a", "b", "c"},
        H: map[string]string{
            "key1": "value1",
            "key2": "value2",
        },
    }).WithContext(c)
}













