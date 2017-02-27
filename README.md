해커스톡
------

지금의 [해커스톡](https://www.hackerstalk.com)은 한국판 [Hacker News](https://news.ycombinator.com/). 

한국에는 오래된 개발자 커뮤니티 사이트/포럼은 몇개 있지만, 최신기술에 대한 뉴스를 볼수 있거나 괜찮은 오픈소스를 발견할 수 있는 사이트가 없다. 그리고 요즘 Github에서 인기 있는 프로젝트들을 보고 있으면 점점 중국에서 만든 프로젝트가 많이 눈에 띄는데, 한국에서 만든 괜찮은 오픈소스가 많이 생겼으면 하는 바람에서 해커스톡을 만들게 되었다.

아직까지는 뉴스, 좋은 블로그 글등을 공유하는 링크 공유 기능만 있지만 앞으로 다른 기능을 점점 채울 예정이다.

## 현재 구현
 - Server: Golang [GIN](https://github.com/gin-gonic/gin)
 - Client: React
 - DB: Postgres


## Build

### Local 개발 환경(Mac)

- [Postgres.app](https://postgresapp.com/)

```
$ createdb ht
$ psql -f sql/schema.sql ht
```


### Client
```
$ npm run dev
```

### Server
```
$ ./run.sh
```

## Deploy
현재 Heroku에 호스팅. `app.json`파일을 참조.


## Maintainers
 - [bbirec](http://github.com/bbirec)
 - [jngbng](https://github.com/jngbng)
