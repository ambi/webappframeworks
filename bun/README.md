##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.2251 secs
  Slowest:	0.0121 secs
  Fastest:	0.0001 secs
  Average:	0.0011 secs
  Requests/sec:	44426.6337
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [7570]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [2048]	|■■■■■■■■■■■
  0.004 [258]	|■
  0.005 [29]	|
  0.006 [34]	|
  0.007 [10]	|
  0.008 [0]	|
  0.010 [0]	|
  0.011 [0]	|
  0.012 [50]	|


Latency distribution:
  10% in 0.0005 secs
  25% in 0.0007 secs
  50% in 0.0009 secs
  75% in 0.0013 secs
  90% in 0.0018 secs
  95% in 0.0023 secs
  99% in 0.0044 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0121 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0011 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0022 secs
  resp wait:	0.0009 secs, 0.0001 secs, 0.0120 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0023 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	1.4975 secs
  Slowest:	0.0321 secs
  Fastest:	0.0005 secs
  Average:	0.0075 secs
  Requests/sec:	6677.9064
  
  Total data:	307740 bytes
  Size/request:	30 bytes

Response time histogram:
  0.001 [1]	|
  0.004 [1170]	|■■■■■■■
  0.007 [6823]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [486]	|■■■
  0.013 [20]	|
  0.016 [0]	|
  0.019 [0]	|
  0.023 [928]	|■■■■■
  0.026 [467]	|■■■
  0.029 [86]	|■
  0.032 [19]	|


Latency distribution:
  10% in 0.0036 secs
  25% in 0.0039 secs
  50% in 0.0048 secs
  75% in 0.0062 secs
  90% in 0.0217 secs
  95% in 0.0228 secs
  99% in 0.0259 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0005 secs, 0.0321 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0024 secs
  resp wait:	0.0074 secs, 0.0005 secs, 0.0283 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	10.5046 secs
  Slowest:	0.1799 secs
  Fastest:	0.0234 secs
  Average:	0.0524 secs
  Requests/sec:	951.9660
  
  Total data:	3187616 bytes
  Size/request:	318 bytes

Response time histogram:
  0.023 [1]	|
  0.039 [559]	|■■■
  0.055 [7290]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.070 [1313]	|■■■■■■■
  0.086 [487]	|■■■
  0.102 [150]	|■
  0.117 [47]	|
  0.133 [27]	|
  0.149 [77]	|
  0.164 [26]	|
  0.180 [23]	|


Latency distribution:
  10% in 0.0434 secs
  25% in 0.0461 secs
  50% in 0.0494 secs
  75% in 0.0536 secs
  90% in 0.0643 secs
  95% in 0.0790 secs
  99% in 0.1386 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0234 secs, 0.1799 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0022 secs
  resp wait:	0.0523 secs, 0.0233 secs, 0.1799 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	3.0808 secs
  Slowest:	0.2281 secs
  Fastest:	0.0005 secs
  Average:	0.0153 secs
  Requests/sec:	3245.9087
  
  Total data:	16050000 bytes
  Size/request:	1605 bytes

Response time histogram:
  0.001 [1]	|
  0.023 [7792]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.046 [1948]	|■■■■■■■■■■
  0.069 [56]	|
  0.092 [68]	|
  0.114 [18]	|
  0.137 [58]	|
  0.160 [9]	|
  0.183 [3]	|
  0.205 [19]	|
  0.228 [28]	|


Latency distribution:
  10% in 0.0054 secs
  25% in 0.0062 secs
  50% in 0.0081 secs
  75% in 0.0228 secs
  90% in 0.0287 secs
  95% in 0.0341 secs
  99% in 0.1251 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0005 secs, 0.2281 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0034 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0026 secs
  resp wait:	0.0152 secs, 0.0005 secs, 0.2280 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	31.8912 secs
  Slowest:	0.3952 secs
  Fastest:	0.0900 secs
  Average:	0.1594 secs
  Requests/sec:	313.5662
  
  Total data:	3187655 bytes
  Size/request:	318 bytes

Response time histogram:
  0.090 [1]	|
  0.121 [1304]	|■■■■■■■■■■■■■■
  0.151 [3862]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.182 [2520]	|■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.212 [1574]	|■■■■■■■■■■■■■■■■
  0.243 [333]	|■■■
  0.273 [210]	|■■
  0.304 [142]	|■
  0.334 [4]	|
  0.365 [12]	|
  0.395 [38]	|


Latency distribution:
  10% in 0.1173 secs
  25% in 0.1364 secs
  50% in 0.1504 secs
  75% in 0.1770 secs
  90% in 0.2018 secs
  95% in 0.2368 secs
  99% in 0.2942 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0900 secs, 0.3952 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.1593 secs, 0.0899 secs, 0.3951 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.3598 secs
  Slowest:	0.0156 secs
  Fastest:	0.0001 secs
  Average:	0.0017 secs
  Requests/sec:	27793.2234
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [6286]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [2466]	|■■■■■■■■■■■■■■■■
  0.005 [770]	|■■■■■
  0.006 [282]	|■■
  0.008 [162]	|■
  0.009 [16]	|
  0.011 [6]	|
  0.013 [8]	|
  0.014 [2]	|
  0.016 [1]	|


Latency distribution:
  10% in 0.0006 secs
  25% in 0.0008 secs
  50% in 0.0012 secs
  75% in 0.0021 secs
  90% in 0.0036 secs
  95% in 0.0047 secs
  99% in 0.0071 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0156 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0041 secs
  resp wait:	0.0015 secs, 0.0001 secs, 0.0125 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0055 secs

Status code distribution:
  [200]	10000 responses



