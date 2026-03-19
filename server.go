package main

import (
	"fmt"
	"log" // 줄바꿈이 필요합니다.
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 요청 내용을 덤프하여 출력 (디버깅용)
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError) // '.' 누락 수정
		return // retrun 오타 수정
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>Hello</body></html>\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler) // '.' 누락 및 함수명(handler) 일치
	
	log.Println("start http listening :18888")
	
	httpServer.Addr = ":18888"
	// httpServer 객체를 사용하거나 http 패키지를 직접 사용합니다.
	log.Fatal(httpServer.ListenAndServe()) // HttpServer 대소문자 및 에러 처리 수정
}