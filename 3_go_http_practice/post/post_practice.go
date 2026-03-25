package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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
func main(){
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	fileWriter, err := writer.CreateFormFile("thumnail", "photo.jpg")
	if err != nil{
		panic(err)
	}
	readFile, err := os.Open("request_file/photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter,readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:",resp.Status)
}