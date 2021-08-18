package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func main()  {
	url := `http://ns.hushijie.com.cn/bms/api/trainExe/app/findList?pageNum=1&pageSize=20`

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func(){
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI(url)
	//requestBody := []byte(`{"pageNum":"1","pageSize:20"}`)
	req.Header.Add("isApp","1")
	//token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIyMzU1MjAiLCJleHAiOjE2MjY0MjQyMDQsImlhdCI6MTYyNTgxOTQwNH0.hPwiWjy7Hv289MJ9BDKtH1YQi0iHgm48Tql4GnoXNrQ"
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIyMzU1MjAiLCJleHAiOjE2MjY2NjY1MjYsImlhdCI6MTYyNjA2MTcyNn0.nLN_j35PWTD7RKrtEX30fufqwbR-USJVfnSW-C08BG0"
	req.Header.Add("Authorization",token)
	//req.SetBody(requestBody)
	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	b := resp.Body()
	type ListData struct {
		ID int `json:"id"`
		Name string `json:"name"`
		SectionId int `json:"sectionId"`
		SectionName string `json:"sectionName"`
		PaperId int `json:"paperId"`
		StartTime int64 `json:"startTime"`
		EndTime int64 `json:"endTime"`
	}
	type ResData struct {
		PageNum int `json:"page_num"`
		PageSize int `json:"pageSize"`
		Size int `json:"size"`
		StartRow int `json:"startRow"`
		EndRow int `json:"endRow"`
		Total int `json:"total"`
		Pages int `json:"pages"`
		List []ListData `json:"list"`
	}
	type Res struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Data ResData `json:"data"`
		
	}
	//resMap := make(map[string]interface{})
	res := &Res{}
	err := json.Unmarshal(b, res)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	currentTime := time.Now().Unix()
	currentTime = currentTime*1000
	fmt.Println(currentTime)
	for i,v := range res.Data.List{
		if v.StartTime <= currentTime && currentTime <= v.EndTime {
			fmt.Println("符合数据====",i, v)
		}else {
			fmt.Println("不符合数据====", i, v)
		}

	}

}
