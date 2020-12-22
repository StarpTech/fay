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
     script: url.js
     output: -

  scenarios: (100.00%) 1 scenario, 20 max VUs, 5m30s max duration (incl. graceful stop):
           * default: Up to 20 looping VUs for 5m0s over 1 stages (gracefulRampDown: 30s, gracefulStop: 30s)


running (5m01.1s), 00/20 VUs, 3949 complete and 0 interrupted iterations
default ✓ [======================================] 00/20 VUs  5m0s

    ✓ is status 200
    ✓ is not status 504
    ✓ is not status 500

    checks.....................: 100.00% ✓ 11847 ✗ 0   
    data_received..............: 785 MB  2.6 MB/s
    data_sent..................: 3.6 MB  12 kB/s
    http_req_blocked...........: avg=3.21µs   min=1.07µs   med=2.42µs   max=252.57µs p(90)=3.38µs  p(95)=3.77µs 
    http_req_connecting........: avg=438ns    min=0s       med=0s       max=116.65µs p(90)=0s      p(95)=0s     
    http_req_duration..........: avg=762.7ms  min=196.94ms med=702.44ms max=2.05s    p(90)=1.27s   p(95)=1.41s  
    http_req_receiving.........: avg=2.66ms   min=1.12ms   med=1.96ms   max=16.23ms  p(90)=4.73ms  p(95)=5.79ms 
    http_req_sending...........: avg=21.17µs  min=8.89µs   med=18.24µs  max=7.53ms   p(90)=23.83µs p(95)=25.93µs
    http_req_tls_handshaking...: avg=0s       min=0s       med=0s       max=0s       p(90)=0s      p(95)=0s     
    http_req_waiting...........: avg=760.02ms min=195.3ms  med=699.95ms max=2.05s    p(90)=1.26s   p(95)=1.4s   
    http_reqs..................: 3949    13.113461/s
    iteration_duration.........: avg=762.91ms min=197.12ms med=702.63ms max=2.05s    p(90)=1.27s   p(95)=1.41s  
    iterations.................: 3949    13.113461/s
    vus........................: 7       min=1   max=19
    vus_max....................: 20      min=20  max=20
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


running (5m00.3s), 00/20 VUs, 9024 complete and 0 interrupted iterations
default ✓ [======================================] 00/20 VUs  5m0s

    ✓ is not status 500
    ✓ is status 200
    ✓ is not status 504

    checks.....................: 100.00% ✓ 27072 ✗ 0   
    data_received..............: 1.6 GB  5.3 MB/s
    data_sent..................: 28 MB   94 kB/s
    http_req_blocked...........: avg=3.23µs   min=1.05µs   med=2.73µs   max=308.94µs p(90)=3.87µs   p(95)=4.42µs  
    http_req_connecting........: avg=230ns    min=0s       med=0s       max=164.97µs p(90)=0s       p(95)=0s      
    http_req_duration..........: avg=332.43ms min=136.55ms med=314.31ms max=1.22s    p(90)=487.62ms p(95)=575.57ms
    http_req_receiving.........: avg=144.72µs min=67.05µs  med=133.43µs max=6.3ms    p(90)=178.31µs p(95)=205.22µs
    http_req_sending...........: avg=23.67µs  min=8.09µs   med=19.92µs  max=6.81ms   p(90)=25.54µs  p(95)=27.78µs 
    http_req_tls_handshaking...: avg=0s       min=0s       med=0s       max=0s       p(90)=0s       p(95)=0s      
    http_req_waiting...........: avg=332.26ms min=136.45ms med=314.14ms max=1.22s    p(90)=487.46ms p(95)=575.45ms
    http_reqs..................: 9024    30.054112/s
    iteration_duration.........: avg=332.64ms min=136.76ms med=314.54ms max=1.22s    p(90)=487.85ms p(95)=575.77ms
    iterations.................: 9024    30.054112/s
    vus........................: 19      min=1   max=19
    vus_max....................: 20      min=20  max=20
```