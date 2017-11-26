# translate

[google-cloud-go/translate](https://github.com/GoogleCloudPlatform/google-cloud-go/tree/master/translate)

[go test で出来ること - Qiita](https://qiita.com/taizo/items/82930518430f940721a0)


## init

```
dep ensure
```

```
export GOOGLE_APPLICATION_CREDENTIALS=server/service-account.json
```

## test

```
go test -v examples_test.go
go test -v -run TestTranslateJA examples_test.go
```

