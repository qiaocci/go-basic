注意需要设置：
```shell
 export GODEBUG=x509ignoreCN=0
```

否则会报错：
```shell
2021/07/17 21:00:35 client.Search err: rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0"
exit status 1

```