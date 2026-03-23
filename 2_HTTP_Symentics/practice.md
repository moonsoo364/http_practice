# Http1.0의 Symentice 브라우저 기본 기능의 이면
이 장에서는 브라우저가 기본 요소들을 어떻게 응용하고 기능을 사용하는 지 알아봅시다.
curl 커맨드를 통해 브라우저의 동작 방식을 설명합니다.

# 2.1 단순한 폼 전송(x-www-form-urlencoded)
Http/1.0의 바디 수신은 클라이언트가 지정한 컨텐츠가 그대로 저장됩니다.
기본적으로 한 번 HTTP가 응답할 때마다 한 파일밖에 반환하지 못하기 때문입니다.
즉 응답의 본체를 지정한 바이트 수만큼 읽어오면 그만입니다.
Http/1.1에는 범위 엑세스라는 특수한 요청 방식이 있습니다. 이는 나중에 설명합니다.


```html
<form method="POST">
  <input name="title" />
  <input name="author" />
  <input name="submit" />
</form>
```
일반적인 웹에서 볼 수 있는 폼입니다. method에는 POST가 설정돼 있습니다. 다음처럼 curl 커맨드를 사용하면 폼과 같은 형식으로 전송할 수 있습니다.

```powershell
curl.exe --http1.0 -d title="The Art of Community" -d author="Jono Bacon" http://localhost:18888
```

curl 커맨드의 `-d` 옵션을 이용해 폼으로 전송할 데이터를 설정할 수 있습니다. 
curl 커맨드는 `-d` 옵션이 지정되면 브라우저와 똑같이 헤더로 **Content-Type:application/x-www-form-urlencoded**를 설정합니다.
이때 바디는 다음과 같은 형식이 됩니다. 키와 값이 '='로 연결되고, 각 항목이 &으로 연결된 문자열입니다.

```shell
title=The Art of Community&author=Jono Bacon
```

단 실제로는 이 커맨드가 생성하는 바디는 브라우저의 웹 폼에서 전송한 것과는 약간 차이가 있습니다. 
`-d` 옵션으로 보낼 경우 지정된 문자열을 그대로 연결합니다. 구분 문자인 &와 =이 있어도 그대로 연결해버리므로, 읽는 쪽에서 원래 데이터로 복원할 수 없습니다.
예를 들어 'Head First PHP & MySQL'이라는 서적명을 넣어보면, 어디서 구분해야 할지 알기 어려워집니다.

**curl**
```powershell
curl.exe --http1.0 -d title="Head First PHP & MySQL" -d author="Lynn Beifhley, Michael Morrison" http://localhost:18888
```
**출력**
```shell
title=Head First PHP & MySQL&author=Lynn Beifhley, Michael Morrison
```

브라우저는 RFC 1866에서 책정한 변환 포맷에 따라 변환을 실시합니다. 
이 포멧에서는 알파벳, 수치, 별표, 하이픈, 마침표, 언더스코어의 여섯 종류 문자 외에는 변환이 필요합니다. 공백은 +로 바뀌므로 실제로는 다음과 같이 됩니다.

```shell
title=Head+First+PHP+%26+MySQL&author=Lynn+Beighley%2C+Michael+Morrison
```

이 방식에서는 이름과 값 안에 포함되는 =와 &는 각각 '%3D'와 '%26'으로 변환됩니다.
curl에서는 이와 비슷한 기능을하는 --data-urlencode가 있습니다.
이를 -d 대신에 사용해서 변환할 수가 있는데 이때 RFC 3986에서 정의된 방법으로 변환됩니다.
RFC 1866과 다루는 문자 종류가 다소 다르며, 또한 공백이 +가 아니라 %20이 됩니다.

**curl**
```powershell
curl.exe --http1.0  --data-urlencode title="Head First PHP & MySQL" --data-urlencode author="Lynn Beifhley, Michael Morrison" http://localhost:18888
```
**출력**
```
title=Head+First+PHP+%26+MySQL&author=Lynn+Beifhley%2C+Michael+Morrison
```
다만 어떤 방법을 써도 같은 알고리즘으로 복원할 수 있으므로 문제가 되지 않습니다. 어느방식이건 'URL 인코딩'으로 부르며 하나로 취급하는 경우가 대부분입니다.

웹 폼의 경우 GET의 경우 바디가 아니라 쿼리로서 URL에 부여된다고 RFC 1866에 정의되어 있습니다.

# 2.5 쿠키
쿠키란 웹사이트의 정보를 브라우저 쪽에 저장하는 작은 파일입니다. 쿠키는 서버가 클라이언트에 '이 파일을 보관해줘'라고 쿠키 저장을 지시합니다.
쿠키도 HTTP 헤더 기반으로 구현됐습니다. 서버에서는 다음과 같이 응답 헤더를 보냅니다.
```shell
SET-Cookie: LAST_ACCESS_DATE=Jul/31/2016
SET-Cookie: LAST_ACCESS_TIME=12:04
```

다음 예제는 쿠키를 통해 첫 방문인지 아닌 지 확인합니다.
```go
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie","VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		// 쿠키가 있으면 1번은 다녀간 적이 있는 사람
		fmt.Fprintf(w, "<html><body>두 번째 이후</body></html>")
	} else {
		fmt.Fprint(w, "<html><body>첫 방문</body></html>")
	}
}
```

서버 프로그램이 볼 땐 마치 데이터베이스처럼 외부에 저장해두고 클라이언트가 액세스할 때마다 데이터를 로드하는 것과 같습니다.
HTTP는 Stateless(언제 누가 요청해도 결과가 같음)를 기본으로 설계되었지만 쿠키를 이용하면 서버가 상태를 유지하는 Stateful한 서비스를 제공할 수 있습니다.

쿠키는 헤더 기반으로 만들어졌으므로 curl 커맨드를 사용할 때 내용을 Cokkie에 넣고 재전송함으로써 구현할 수 있지만 쿠키를 위한 전용 옵션도 있습니다.
`-c cookie.txt` 옵션으로 지정한 파일에 쿠키를 수신한 쿠키를 저장하고 -b/--cookie 옵션으로 지정한 파일에서 쿠키를 읽어와 전송합니다.

```powershell
curl.exe -c .\2_HTTP_Symentics\cookie.txt http://localhost:18888
```

```powershell
curl.exe -b .\2_HTTP_Symentics\cookie.txt http://localhost:18888
```
## 2.5.1 쿠키의 잘못된 사용법
쿠키는 편리한 기능이지만 몇 가지 제약이 있습니다.
우선 영속성 문제가 있습니다. 쿠키는 어떤 상황에서도 확실하게 저장되는 것은 아닙니다.
비밀 모드 혹은 브라우저 보안 설정에 따라 세션이 끝나면 초기화되거나 쿠키를 저장하라는 서버의 지시를 무시하기도 합니다.
그래서 쿠키를 데이터베이스 대신 쓸 수는 없습니다.

또한 용량 문제도 있습니다. 쿠키의 최대 크기는 4Kb 사양으로 정해져 있습니다.
쿠키는 헤더로서 항상 통신에 부가되므로 통신량이 늘어나는데, 통신량 증가는 요청과 응답 속도 모두 영향을 미칩니다.
제한된 용량과 통신량 증가는 데이터베이스로 사용하는 데 제약이 됩니다.

마지막은 보안문제입니다. secure 속성을 이용하면 HTTPS 프로토콩로 암호화된 통신에서 쿠키가 전송되지만 HTTP에서는 쿠키가 평문으로 전송됩니다.
매 요청 시 쿠키가 송수신되는데 보여선 곤란한 비밀번호 등이 포함되면 노출될 위험성이 있습니다.
암호화된다고 해도 사용자가 자유롭게 접근할 수 있는 것도 문제입니다. 우너리상 사용자가 쿠키를 수정할 수 있으므로 
시스템에서 필요한 ID나 수정되면 오작동으로 이어지는 민감한 정보를 넣는데 적합하지 않습니다. 정보를 넣을 때는 서명이나 암호화 처리가 필요합니다.

기본적으로 인증 정보나 사라져도 상관없는 정보만 쿠키에 넣는 것이 좋습니다.