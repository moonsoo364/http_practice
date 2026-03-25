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