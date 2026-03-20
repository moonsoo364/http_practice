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
예를 들어 'Head First PHP & MySQL'이라는 서적명을 넣오보면, 어디서 구분해야 할지 알기 어려워집니다.

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