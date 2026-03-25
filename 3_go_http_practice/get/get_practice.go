package main

import (
	"log"
	"net/http"
)

// 3.4 Get 메서드 송신과 바디 스테이터스 코드, 헤더 수신
// func main() {
// 	resp, err := http.Get("http://localhost:18888")
// 	if err != nil {
// 		panic(err) //panic: 오류 표시하고 프로그램 종료
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body) // 바디 내용 바이트로 받아옴
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(string(body))
// 	log.Println("Status:", resp.Status)
// 	log.Println("StatusCode:", resp.StatusCode)
// 	log.Println("Headers:", resp.Header)

// }

// 3.5 GET 메서드 + 쿼리 전송
// func main (){
// 	values := url.Values{
// 		"query": {"Hello", "World"},
// 	}
// 	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
// 	defer resp.Body.Close() // defer: 함수를 빠져나올 때 이 문을 실행한다, 여기서는 소켓에서 바디를 읽고 나서의 처리
// 	body, _ := io.ReadAll(resp.Body) // _ : 값을 할당하지 않고 버리겠다.
// 	// := 변수를 선언하고 값을 할당한다.
// 	log.Println(string(body))

// }

// 3.6 HEAD 메서드로 헤더 가져오기
func main (){
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
	
}