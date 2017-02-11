해커스톡
------

## Local 개발 환경

- [Postgres.app](https://postgresapp.com/)

```
$ createdb ht
$ psql -f sql/schema.sql ht
```

## Build

### Client
```
$ npm install
$ npm run build # 한번만 생성
$ npm run dev   # watch 하면서 계속 업데이트
```

### Server
```
$ ./run.sh
```

## Contribute

### npm package 추가할 때 주의 사항

``node_modules``를 저장소에서 뺐기 때문에 매번 빌드할 때 다른 버전의 npm 패키지를
받아올 수 있는 문제가 생긴다.
``npm install --save --save-exact``를 이용하면 우리가 직접 참조하는 패키지의 버전을
고정할 수 있기 때문에 문제를 줄일 수 있지만, 참조된 패키지가 내부적으로 참조하는
패키지의 버전은 고정할 수 없기 때문에 여전히 문제가 될 수 있다. 예를 들면 ``A@1.0``
이라는 패키지를 추가했는데 ``A`` 패키지가 내부적으로 ``SUB_A@^1.0``라는 패키지 의존성을  가질 경우, 설치할
때 마다 ``A``는 1.0으로 고정되지만 ``SUB_A``는 1.X 로 달라질 수 있다.

이를 해결하기 위해 ``npm shrinkwrap`` 기능을 사용한다. ``npm install --save``로
패키지를 추가한 후에는 ``npm shrinkwrap``으로 ``npm-shrinkwrap.json`` 파일도 업데이트 해준다.
[shrinkwrap 문서 참고](https://docs.npmjs.com/cli/shrinkwrap#building-shrinkwrapped-packages)
