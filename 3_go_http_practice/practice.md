# Go 언어를 활용한 HTTP 클라이언트 구현
이 장 에서는 Go 언어를 활용한 HTTP 클라이언트 예제를 실습합니다.

Go 언어를 사용하는 이유
- 다른 언어보다 간결한 언어 사양과 풍부한 표준 라이브러리
- 컴파일이 동적 스크립트 언어를 실행하는 만큼 빠르다.
- 아웃풋이 단일 바이너리로 나온다.

# 3.4 GET 메서드 송신과 바디. 스테이터스 코드 헤더 수신 

# 3.4.1 io.Reader
Go 언어에서는 데이터의 순차적인 입출력을 io.Reader, io Writer 인터페이스로 추상화 했습니다.
이 인터페이스는 파일이나 소켓 등 다양한 곳에서 사용합니다.
http.Response형 변수 resp의 Body도 io.Reader 인터페이스를 가진 오브젝트입니다.

Go언어는 상대가 파일이든 소켓이든 모두 다룰 수 있는 다양한 높은 수준의 기능을 제공합니다.

1. `io.Reader`의 내용을 모아서 바이트 배열로 읽어온다.
: `func ioutil.ReadeAll (reader io.Reader) ([]byte, error)`

2. `io.Reader`에서 `io.Writer`를 통째로 복사한다.
: `func Copy(dst io.Writer, src io.Reader ) (written int 64, err error)`

3. `io.Reader`를 래핑해 버퍼 기능을 추가하고, 편리한 메서드를 다수 추가하는 오브젝트를 만든다.
: `bufio.NewReader(reader io.Reader) *bufio.Reader`

# 3.5 GET 메서드 + 쿼리 전송

# 3.6 HEAD 메서드로 헤더 가져오기

# 3.7 x-www-form-urlencoded 형식의 POST 메서드 전송

Go에서 `http.PostForm()` 메서드를 호출할 경우 RFC 3986에 따라 인코딩되어 요청이 전송된다.

# 3.8 POST 메서드로 임의의 바디 전송

HTTP 프로토콜에서 chunk 방식으로 데이터를 보낼 때는 각 데이터 덩어리 앞에 해당 덩어리의 크기를 알려줘야 합니다.

'Hello World' 라는 문자열이 담긴 `main.txt` 보내면 서버로그에 다음과 같이 표시됩니다.
앞에 `b`는 16진수로 숫자 11을 의미합니다. `Hello World`라는 문자열의 글자 수(공백 포함)이 11글자이기 때문에
`b`로 표시한 것 입니다.

`0` '이제 보낼 데이터가 더 이상 없어'라는 뜻의 **종료 신호(End of Stream)** 입니다.

```powershell
b
Hello World
0
```
`Contet-Type` 헤더는 `http.Post()`의 두번째 인수로 전달합니다. 이 때 전송할 내용은 텍스트화하지 않고
`io.Reader` 형식으로 전달합니다. `os.Open()` 함수에서 생성되는 `os.File` 오브젝트는 `io.Reader` 인터페이스를 만족하므로 그대로 `http.Post()`에 넘길 수 있습니다.
파일이 아니라 프로그램의 안에서 생성한 텍스트의 경우 `strings.Reader`를 통해 문자열을 `io.Reader` 인터페이스화 합니다.
```go
reader := strings.NewReader("텍스트")
resp, err := http.Post("http://localhost:18888", "text/plain", reader)
```

# 3.9 multipart/form-data 형식으로 파일 전송

`io.Reader`인 `bytes.Buffer`에 멀티파트 폼이 전송할 컨텐츠를 작성하고 있습니다.
이 컨텐츠를 만드는 데 사용하는 것이 `multipart.Writer` 오브젝트입니다.
이 오브젝트를 통해 폼의 항목이나 파일을 써넣으면 `multipart.NewWriter`의 인수로 전달한
`bytes.Buffer`에 기록됩니다. 이 `bytes.Buffer`는 `io.Reader`인 동시에 `io.Writer`이기도 합니다.

```go
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
```


`Content-type`에는 경계 문자열을 넣어야 합니다. `multipart.Writer` 오브젝트가 내부에서 난수를 생성합니다.
`Boundary` 메서드로 취득할 수 있으므로 다음과 같이 써서 `Content-Type`을 만들어 낼 수도 있습니다.
`FormDataContentType` 메서드는 이 코드를 간단하게 사용하는 방식입니다.

```go
"multipart/form-data; boundary=" +writer.Boundary()
```
