# atk (Access Token Keeper)

- [x] common api.
- [x] auto refresh access token.
- [ ] auto check.

# App types

| type   | des                  |
| ------ | -------------------- |
| `wxmp` | wechat mini program  |

# api

`/app/:appType/:appId/token?refresh=false`

`/pong`

# References

- https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
- https://developers.weixin.qq.com/doc/offiaccount/Message_Management/API_Call_Limits.html


```bash
ab -c100 -n10000 http://127.0.0.1:8080/app/wxmp/xxx/token
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /app/wxmp/xxx/token
Document Length:        265 bytes

Concurrency Level:      100
Time taken for tests:   2.304 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      3890000 bytes
HTML transferred:       2650000 bytes
Requests per second:    4341.18 [#/sec] (mean)
Time per request:       23.035 [ms] (mean)
Time per request:       0.230 [ms] (mean, across all concurrent requests)
Transfer rate:          1649.14 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   2.5      1      71
Processing:     4   22  50.7     14     516
Waiting:        2   22  50.7     13     516
Total:          4   23  50.8     14     517

Percentage of the requests served within a certain time (ms)
  50%     14
  66%     17
  75%     18
  80%     20
  90%     26
  95%     62
  98%     82
  99%    505
 100%    517 (longest request)
```


## ip allow ref

- https://docs.nginx.com/nginx/admin-guide/security-controls/denylisting-ip-addresses/
- 