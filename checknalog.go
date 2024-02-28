package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type AnswerOnRequest struct {
	Success   bool `json:"success"`
	UserToken string
	ID        string
}
type AnswerOnCheck struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
}

func main() {
	fmt.Println("Hello World")
	client := &http.Client{}
	//var data = strings.NewReader(`type=request&date=28.12.2023&time=17%3A08&operationtype=1&summ=74&fn=7281440500467899&fd=125616&fp=3754267781`)
	datestr := "19.02.2024"
	timestr := "11%3A10"
	opertypestr := "1"
	summstr := "112"
	fnstr := "7281440500467899"
	fdstr := "147387"
	fpstr := "3956843965"
	reqstr := fmt.Sprintf("type=request&date=%v&time=%v&operationtype=%v&summ=%v&fn=%v&fd=%v&fp=%v", datestr, timestr, opertypestr, summstr, fnstr, fdstr, fpstr)
	//var data = strings.NewReader(`type=request&date=28.12.2023&time=17%3A08&operationtype=1&summ=74&fn=7281440500467899&fd=125616&fp=3754267781`)
	fmt.Println("post", reqstr)
	var data = strings.NewReader(reqstr)
	req, err := http.NewRequest("POST", "https://kkt-online.nalog.ru/openapikkt.html", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "kkt-online.nalog.ru")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	//req.Header.Set("cookie", "_ym_uid=167939142713020985; _ym_d=1709109532; _ym_isad=2; kkt-nalog-ru-cookie=49928640cd314165961aa12bf4463f03; session-cookie=17b7fc033d74a7820592c1d4beb261f5cf6455c83442daee13d8edf0f6d02d7bb8432e6f2385b7c84084f914dcee73ce; sputnik_session=1709114274628|0")
	req.Header.Set("origin", "https://kkt-online.nalog.ru")
	req.Header.Set("referer", "https://kkt-online.nalog.ru/")
	req.Header.Set("sec-ch-ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var bodyText2 AnswerOnRequest
	err = json.Unmarshal(bodyText, &bodyText2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", bodyText)
	token := bodyText2.UserToken
	id := bodyText2.ID
	fmt.Println("Пауза 2 секункды...")
	sec2 := time.Second * 2
	time.Sleep(sec2)
	fmt.Println("Полчаем ответ по чеку")
	codeofcheck := getresponse(token, id)
	fmt.Println("Код чека:", codeofcheck)
	if codeofcheck != "200" {
		fmt.Println("Чек не прошел проверку")
	} else {
		fmt.Println("Чек прошел проверку")
	}
}

func getresponse(token, id string) string {
	client := &http.Client{}
	respstr := fmt.Sprintf("type=poll&UserToken=%v&id=%v", token, id)
	//var data = strings.NewReader(`type=poll&UserToken=49928640cd314165961aa12bf4463f03&id=8c9f7abb-339c-4481-a21a-e8af0c140eaf`)
	var data = strings.NewReader(respstr)
	fmt.Println("post", respstr)
	req, err := http.NewRequest("POST", "https://kkt-online.nalog.ru/openapikkt.html", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "kkt-online.nalog.ru")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	//req.Header.Set("cookie", "_ym_uid=167939142713020985; _ym_d=1709109532; _ym_isad=2; kkt-nalog-ru-cookie=49928640cd314165961aa12bf4463f03; session-cookie=17b7fc033d74a7820592c1d4beb261f5cf6455c83442daee13d8edf0f6d02d7bb8432e6f2385b7c84084f914dcee73ce; sputnik_session=1709114274628|0")
	req.Header.Set("origin", "https://kkt-online.nalog.ru")
	req.Header.Set("referer", "https://kkt-online.nalog.ru/")
	req.Header.Set("sec-ch-ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	var bodyText2 AnswerOnCheck
	err = json.Unmarshal(bodyText, &bodyText2)
	return bodyText2.Code
}
