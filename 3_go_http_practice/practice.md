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

# 3.10 쿠키 송수신
지금까지는 HTTP 전송을 1회 요청하는 코드 였습니다, HTTP는 스테이트리스므로 각 전송끼리 사용하는 함수끼리 영향이 없습니다. 하지만 쿠키는 브라우저 내부에서 상태를 유지해야만 합니다. 이 경우는 지금까지 소개했던 함수가 아니라 `http.Client` 구조체를 이용합니다.

**NOTE** Go에서 오브젝트 생성하는 문법은 크게 3가지로 나뉨, 이책에서는 주로 초깃값을 지정해서 생성하는 법을 이용
```go
// 초기값을 지정
a := Struct{
	Member: "Value",
}
// new 함수로 초기화
a := new(Struct)

// make 함수로 초기화
// 배열의 슬라이스, map, 채널 전용
a := make(map[string]string)
```

# 3.11 프록시 이용
이번에 사용할 것은 `Transport`입니다. `Transport`는 실제 통신을 하는 백앤드입니다.
아래는 실제 동작하는 curl 입니다.

```powershell
curl.exe -x http://localhost:18888 http://github.com
```

`client.Get`의 대상은 외부 사이트이지만, 프록시의 방향은 로컬 테스트 서버입니다.
이 코드를 실행하면 외부로 직접 요청을 날리지 않고, 로컬 서버가 일단 요청을 받습니다.
그러나 로컬 사버가 직접 응답을 반환하므로 `github.com`에 대한 액세스가 일어나지 않습니다.

```go
func main (){
	proxyUrl, err := url.Parse("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	resp, err := client.Get("http://github.com")
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
```

# 3.12 파일 시스템 액세스
file 스키마는 로컬 파일에 액세스할 때 사용하는 스키마 입니다.
curl에서는 다음 명령을 실행하면 작업 폴더 내 해당 파일 내용을 콘솔에 출력 할 수 있습니다.

```powershell
curl.exe file://main.go
```

`http.Transport`에는 이 밖의 스키마용 트랜스포트를 추가하는 `Register Protocol` 메서드가 있습니다.
`http.NewFileTransport()`는 로컬 파일에 액세스할 수 있는 메서드 입니다.

# 3.13 자유로운 메서드 전송
지금까지의 살펴본 코드는 http 모듈 함수나 `http.Client` 구조체의 메서드를 사용했습니다.
이들 메서드가 지원하는 것은 `GET`,`HEAD`,`POST` 뿐입니다. 다른 메서드를 요청할 때는
`http.Request` 구조체의 오브젝트를 사용해야 합니다.

다음은 `DELETE` 메서드를 구현한 예제 입니다.
`http.Request` 구조체는 `http.NewRequest`라는 빌더 함수를 사용하여 생성합니다.
함수의 인수는 HTTP 메서드, URI, 바디 입니다. 바디에는 Post와 마찬가지로 `io.Reader`를 사용할 수 있습니다.
```go
func main(){
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", "http://localhost:18888",nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
```

# 3.14 헤더 전송
임의의 메서드를 전송할 때 사용한 `http.Request` 구조체에는 `Header` 라는 필드가 있습니다.
이는 `http.Response` 의 헤더와 같습니다. GET 메서드 예제에서는 `Get()` 메서드로 헤더를 가져왔습니다.
헤더를 추가할 때는 `Add()` 메서드를 사용합니다.

```go
	request.Header.Add("Content-Type","application/json")
```

쿠키는 헤더에 설정하지 않아도 `http.Client`의 `Jar`에 `cookie.Jar`의 인스턴스를 설정해 송수신하게 됩니다.
수동으로 헤더에 설정하면 받지 않은 쿠키도 자유롭게 이용할 수 있습니다.

```go
	request.AddCookie(&http.Cookie{Name: "test", Value:"test value"})
```


# 3.15 국제화 도메인
Go 언어로도 URL을 변환할 수 있습니다.
변환 처리는 `idna.ToASCII()`와 `idna.ToUnicode()` 함수로 실행합니다.
도메인 이름을 `idna.ToASCII()`로 변환함으로써 지금까지 설명해온 API로 한글 도메인 정보를 가져올 수 있습니다.
