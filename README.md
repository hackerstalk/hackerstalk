해커들의 수다
---------

## Local 개발 환경

- [Postgres.app](https://postgresapp.com/)

```
$ createdb ht
$ psql -f sql/schema.sql ht
```

## Build

### Client
```
$ npm run build # 한번만 생성
$ npm run dev   # watch 하면서 계속 업데이트
```

### Server
```
$ ./build.sh
$ export GITHUB_CLIENT_ID=<Client Id>
$ export GITHUB_CLIENT_SECRET=<Client Secret>
$ ./app
```
