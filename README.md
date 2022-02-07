# hello

Demo to use `tao-hello`

```
go mod tidy
go run main.go
go run main.go -f conf/config.json -a 100
go run main.go -f conf/config.yml -a 50
```

> uncomment `_ "github.com/taouniverse/hello/conf"` in main.go to use config in code instead of config file.

