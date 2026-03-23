// package main

// import (
// 	"fmt"
// 	"log" // 줄바꿈이 필요합니다.
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Set-Cookie","VISIT=TRUE")
// 	if _, ok := r.Header["Cookie"]; ok {
// 		// 쿠키가 있으면 1번은 다녀간 적이 있는 사람
// 		fmt.Fprintf(w, "<html><body>두 번째 이후</body></html>")
// 	} else {
// 		fmt.Fprint(w, "<html><body>첫 방문</body></html>")
// 	}
// }

// func main() {
// 	var httpServer http.Server
// 	http.HandleFunc("/", handler) // '.' 누락 및 함수명(handler) 일치
	
// 	log.Println("start http listening :18888")
	
// 	httpServer.Addr = ":18888"
// 	// httpServer 객체를 사용하거나 http 패키지를 직접 사용합니다.
// 	log.Fatal(httpServer.ListenAndServe()) // HttpServer 대소문자 및 에러 처리 수정
// }