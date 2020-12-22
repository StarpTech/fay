# Load testing

You may wonder how `fay` behaves under load.

In order to help you having an idea, we created a test-load scenario for the
[k6](https://docs.k6.io/docs) load testing tool.

## Load an external page

The URL scenario is quite simple:

* Ramp up from 0 to 20 virtual users, each one uploading as many time as possible the [Header.html](../example/header.html) and load the example page from `http://localhost:3000/page.html`
* Stop when at least one response is not an HTTP 200 code.

Test machine:

```
OS: Manjaro Linux x86_64
CPU: Intel i7-10510U (8) @ 4.900GHz
Memory: 31773MiB
```

```bash
❯ k6 run --env MAX_VUS=20 --env BASE_URL=http://localhost:3000 url.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: html.js
     output: -

  scenarios: (100.00%) 1 scenario, 20 max VUs, 5m30s max duration (incl. graceful stop):
           * default: Up to 20 looping VUs for 5m0s over 1 stages (gracefulRampDown: 30s, gracefulStop: 30s)


running (5m00.6s), 00/20 VUs, 3141 complete and 0 interrupted iterations
default ✓ [======================================] 00/20 VUs  5m0s

    ✓ is status 200
    ✓ is not status 504
    ✓ is not status 500

    checks.....................: 100.00% ✓ 9423 ✗ 0   
    data_received..............: 71 MB   237 kB/s
    data_sent..................: 2.9 MB  9.6 kB/s
    http_req_blocked...........: avg=7.27µs   min=1.88µs   med=5.04µs   max=686.1µs  p(90)=7.12µs  p(95)=8.73µs 
    http_req_connecting........: avg=1.34µs   min=0s       med=0s       max=496.64µs p(90)=0s      p(95)=0s     
    http_req_duration..........: avg=956.65ms min=255.16ms med=945.49ms max=2.18s    p(90)=1.28s   p(95)=1.42s  
    http_req_receiving.........: avg=16.7ms   min=1.77ms   med=14.06ms  max=135.96ms p(90)=33.67ms p(95)=39.97ms
    http_req_sending...........: avg=37.09µs  min=14.2µs   med=31.54µs  max=5.54ms   p(90)=44.07µs p(95)=48.99µs
    http_req_tls_handshaking...: avg=0s       min=0s       med=0s       max=0s       p(90)=0s      p(95)=0s     
    http_req_waiting...........: avg=939.91ms min=251.45ms med=930.29ms max=2.14s    p(90)=1.26s   p(95)=1.4s   
    http_reqs..................: 3141    10.448383/s
    iteration_duration.........: avg=956.98ms min=255.46ms med=945.79ms max=2.18s    p(90)=1.28s   p(95)=1.42s  
    iterations.................: 3141    10.448383/s
    vus........................: 19      min=1  max=19
    vus_max....................: 20      min=20 max=20
```

## Send static HTML

The HTML scenario is quite simple:

* Ramp up from 0 to 20 virtual users, each one uploading as many time as possible the [Header.html](../example/header.html) and [page.html](../example/page.html).
* Stop when at least one response is not an HTTP 200 code.

Test machine:

```
OS: Manjaro Linux x86_64
CPU: Intel i7-10510U (8) @ 4.900GHz
Memory: 31773MiB
```

```bash
❯ k6 run --env MAX_VUS=20 --env BASE_URL=http://localhost:3000 html.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: html.js
     output: -

  scenarios: (100.00%) 1 scenario, 20 max VUs, 5m30s max duration (incl. graceful stop):
           * default: Up to 20 looping VUs for 5m0s over 1 stages (gracefulRampDown: 30s, gracefulStop: 30s)


running (5m00.6s), 00/20 VUs, 4096 complete and 0 interrupted iterations
default ✓ [======================================] 00/20 VUs  5m0s

    ✓ is not status 500
    ✓ is status 200
    ✓ is not status 504

    checks.....................: 100.00% ✓ 12288 ✗ 0   
    data_received..............: 725 MB  2.4 MB/s
    data_sent..................: 13 MB   43 kB/s
    http_req_blocked...........: avg=7.1µs    min=1.79µs   med=5µs      max=2.81ms  p(90)=6.64µs  p(95)=7.98µs 
    http_req_connecting........: avg=1.47µs   min=0s       med=0s       max=2.69ms  p(90)=0s      p(95)=0s     
    http_req_duration..........: avg=733.68ms min=180.54ms med=696.93ms max=1.74s   p(90)=1.17s   p(95)=1.29s  
    http_req_receiving.........: avg=15.83ms  min=165.64µs med=12.19ms  max=106.3ms p(90)=33.3ms  p(95)=41.98ms
    http_req_sending...........: avg=45.77µs  min=13.19µs  med=31.11µs  max=11.2ms  p(90)=42.56µs p(95)=47.41µs
    http_req_tls_handshaking...: avg=0s       min=0s       med=0s       max=0s      p(90)=0s      p(95)=0s     
    http_req_waiting...........: avg=717.79ms min=178.26ms med=682.23ms max=1.71s   p(90)=1.14s   p(95)=1.26s  
    http_reqs..................: 4096    13.625701/s
    iteration_duration.........: avg=734.02ms min=180.74ms med=697.3ms  max=1.74s   p(90)=1.17s   p(95)=1.29s  
    iterations.................: 4096    13.625701/s
    vus........................: 19      min=1   max=19
    vus_max....................: 20      min=20  max=20
```