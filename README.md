# http_practice
real http world 도서의 읽고 실습한 내용을 기록했습니다.
각 장에 `*_practice.md` 확장자 파일에 curl 테스트 했던 코드와 책의 개념 설명을 작성했습니다.

# 기타 
## Go 설치
[Go 설치 웹사이트](https://go.dev/dl/)
## Go 프로젝트 실행
go 모듈 설치

```shell
go mod init http_practice
```

```shell
go run .\2_HTTP_Symentics\server_cookie.go
```


```powershell
go run .\1_HTTP_Syntax\server.go
```

파일 빌드
```popwershell
go build -o .\build\server_cookie.exe .\2_HTTP_Symentics\server_cookie.go
```

## Window에서 curl
`powershell` 에서 curl.exe 명령어를 사용하면 cli에서 curl 명령어 사용가능

```powershell
curl.exe --http1.0 http://localhost:18888/greeting
```

