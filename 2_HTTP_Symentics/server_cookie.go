package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 1. 클라이언트에게 전달할 쿠키 생성 및 설정 (Set-Cookie)
	cookie := &http.Cookie{
		Name:  "my_session_id",
		Value: "abcdefg-12345",
	}
	http.SetCookie(w, cookie)

	// 2. 클라이언트의 요청 내용을 덤프하여 출력 (디버깅용)
	// 클라이언트가 쿠키를 다시 보내면 여기서 Cookie 헤더를 확인할 수 있습니다.
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>Hello, Cookie!</body></html>\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	
	log.Println("start http listening :18888")
	
	httpServer.Addr = ":18888"
	log.Fatal(httpServer.ListenAndServe())
}