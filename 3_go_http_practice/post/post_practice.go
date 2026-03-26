package main

import (
	"fmt"

	"golang.org/x/net/idna"
)

// 3.7 x-www-form-urlencoded 형식의 POST 메서드 전송
// func main() {
// 	values := url.Values{
// 		"test": {"value"},
// 	}

// 	resp, err := http.PostForm("http://localhost:18888", values)
// 	if err != nil {
// 		// 전송 실패
// 		panic(err)
// 	}
// 	log.Println("Status:", resp.Status)
// }

// 3.8 POST 메서드로 임의의 바디 전송 -1
// func main(){
// 	file, err := os.Open("request_file/main.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := http.Post("http://localhost:18888", "text/plain", file)
// 	if err != nil {
// 		//전송 실패
// 		panic(err)
// 	}
// 	log.Println("Status:", resp.Status)
// }

// 3.8 POST 메서드로 임의의 바디 전송 -2
// func main(){
// 	reader := strings.NewReader("텍스트")
// 	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println("Status:",resp.Status)
// }

// 3.9 multipart/form-data 형식으로 파일 전송
// func main(){
// 	var buffer bytes.Buffer
// 	writer := multipart.NewWriter(&buffer)
// 	writer.WriteField("name", "Michael Jackson")

// 	fileWriter, err := writer.CreateFormFile("thumnail", "photo.jpg")
// 	if err != nil{
// 		panic(err)
// 	}
// 	readFile, err := os.Open("request_file/photo.jpg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer readFile.Close()
// 	io.Copy(fileWriter,readFile)
// 	writer.Close()

// 	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println("Status:",resp.Status)
// }

// 3.10 쿠키 송수신
// func main(){
// 	jar, err := cookiejar.New(nil)
// 	if (err != nil){
// 		panic(err)
// 	}
// 	client := http.Client{
// 		Jar: jar,
// 	}
// 	for i := 0; i < 2; i++ {
// 		resp, err := client.Get("http://localhost:18888/cookie")
// 		if err != nil {
// 			panic(err)
// 		}
// 		dump, err := httputil.DumpResponse(resp, true)
// 		if err != nil {
// 			panic(err)
// 		}
// 		log.Println(string(dump))
// 	}
// }

// 3.11 프록시 이용
// func main (){
// 	proxyUrl, err := url.Parse("http://localhost:18888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	client := http.Client{
// 		Transport: &http.Transport{
// 			Proxy: http.ProxyURL(proxyUrl),
// 		},
// 	}
// 	resp, err := client.Get("http://github.com")
// 	if err != nil {
// 		panic(err)
// 	}
// 	dump, err := httputil.DumpResponse(resp, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(string(dump))
// }

// 3.12 파일 시스템 액세스
// func main (){
// 	transport := &http.Transport{}
// 	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
// 	client := http.Client{
// 		Transport : transport,
// 	}
// 	resp, err := client.Get("file://./request_file/main.txt")
// 	if err != nil{
// 		panic(err)
// 	}
// 	dump, err := httputil.DumpResponse(resp, true)
// 	if err != nil{
// 		panic(err)
// 	}
// 	log.Println(string(dump))
// }

// 3.13 자유로운 메서드 전송
// func main(){
// 	// 3.13 postForm 전송 시 사용
// 	values := url.Values{"test": {"values"}}
// 	reader := strings.NewReader(values.Encode())
// 	client := &http.Client{}
// 	request, err := http.NewRequest("DELETE", "http://localhost:18888", reader)
// 	// 3.14 헤더 전송
// 	request.Header.Add("Content-Type","application/json")
// 	// 3.14 쿠키 전송
// 	request.AddCookie(&http.Cookie{Name: "test", Value:"test value"})
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := client.Do(request)
// 	if err != nil {
// 		panic(err)
// 	}
// 	dump, err := httputil.DumpResponse(resp, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(string(dump))
// }

// 3.15 국제화 도메인
func main() {
	src := "남한강휴게소.com"
	ascii, err := idna.ToASCII(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -> %s\n",src, ascii)
}