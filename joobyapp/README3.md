##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.6979 secs
  Slowest:	0.1811 secs
  Fastest:	0.0002 secs
  Average:	0.0031 secs
  Requests/sec:	14328.7095
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.018 [9722]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.036 [150]	|■
  0.054 [34]	|
  0.073 [11]	|
  0.091 [7]	|
  0.109 [12]	|
  0.127 [27]	|
  0.145 [31]	|
  0.163 [1]	|
  0.181 [4]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0004 secs
  50% in 0.0006 secs
  75% in 0.0018 secs
  90% in 0.0043 secs
  95% in 0.0081 secs
  99% in 0.0474 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.1811 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0015 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0075 secs
  resp wait:	0.0028 secs, 0.0002 secs, 0.1811 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0113 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	2.0894 secs
  Slowest:	0.1545 secs
  Fastest:	0.0006 secs
  Average:	0.0101 secs
  Requests/sec:	4785.9527
  
  Total data:	307766 bytes
  Size/request:	30 bytes

Response time histogram:
  0.001 [1]	|
  0.016 [7915]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.031 [1588]	|■■■■■■■■
  0.047 [315]	|■■
  0.062 [59]	|
  0.078 [31]	|
  0.093 [30]	|
  0.108 [0]	|
  0.124 [0]	|
  0.139 [49]	|
  0.154 [12]	|


Latency distribution:
  10% in 0.0018 secs
  25% in 0.0035 secs
  50% in 0.0050 secs
  75% in 0.0107 secs
  90% in 0.0235 secs
  95% in 0.0313 secs
  99% in 0.0764 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0006 secs, 0.1545 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0014 secs
  resp wait:	0.0100 secs, 0.0005 secs, 0.1544 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0012 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	7.7514 secs
  Slowest:	0.4383 secs
  Fastest:	0.0020 secs
  Average:	0.0381 secs
  Requests/sec:	1290.0893
  
  Total data:	3187253 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.046 [8041]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.089 [1213]	|■■■■■■
  0.133 [114]	|■
  0.176 [457]	|■■
  0.220 [110]	|■
  0.264 [24]	|
  0.307 [25]	|
  0.351 [0]	|
  0.395 [1]	|
  0.438 [14]	|


Latency distribution:
  10% in 0.0072 secs
  25% in 0.0152 secs
  50% in 0.0315 secs
  75% in 0.0421 secs
  90% in 0.0616 secs
  95% in 0.1473 secs
  99% in 0.1831 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0020 secs, 0.4383 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0023 secs
  resp wait:	0.0380 secs, 0.0019 secs, 0.4382 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	1.5526 secs
  Slowest:	0.2548 secs
  Fastest:	0.0004 secs
  Average:	0.0068 secs
  Requests/sec:	6440.6105
  
  Total data:	21300000 bytes
  Size/request:	2130 bytes

Response time histogram:
  0.000 [1]	|
  0.026 [9637]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.051 [240]	|■
  0.077 [8]	|
  0.102 [9]	|
  0.128 [4]	|
  0.153 [79]	|
  0.178 [2]	|
  0.204 [2]	|
  0.229 [1]	|
  0.255 [17]	|


Latency distribution:
  10% in 0.0006 secs
  25% in 0.0011 secs
  50% in 0.0024 secs
  75% in 0.0054 secs
  90% in 0.0185 secs
  95% in 0.0233 secs
  99% in 0.1280 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.2548 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0018 secs
  resp wait:	0.0067 secs, 0.0004 secs, 0.2543 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0015 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	15.1056 secs
  Slowest:	0.2517 secs
  Fastest:	0.0077 secs
  Average:	0.0751 secs
  Requests/sec:	662.0040
  
  Total data:	3187620 bytes
  Size/request:	318 bytes

Response time histogram:
  0.008 [1]	|
  0.032 [48]	|
  0.056 [1666]	|■■■■■■■■■■■■■
  0.081 [5210]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.105 [2421]	|■■■■■■■■■■■■■■■■■■■
  0.130 [299]	|■■
  0.154 [131]	|■
  0.179 [31]	|
  0.203 [94]	|■
  0.227 [65]	|
  0.252 [34]	|


Latency distribution:
  10% in 0.0542 secs
  25% in 0.0586 secs
  50% in 0.0685 secs
  75% in 0.0845 secs
  90% in 0.0981 secs
  95% in 0.1148 secs
  99% in 0.2029 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0077 secs, 0.2517 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0020 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0025 secs
  resp wait:	0.0750 secs, 0.0076 secs, 0.2516 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1395 secs
  Slowest:	0.0103 secs
  Fastest:	0.0001 secs
  Average:	0.0007 secs
  Requests/sec:	71661.2838
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8686]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1046]	|■■■■■
  0.003 [199]	|■
  0.004 [18]	|
  0.005 [14]	|
  0.006 [16]	|
  0.007 [10]	|
  0.008 [1]	|
  0.009 [3]	|
  0.010 [6]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0004 secs
  50% in 0.0005 secs
  75% in 0.0007 secs
  90% in 0.0014 secs
  95% in 0.0018 secs
  99% in 0.0027 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0103 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0036 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0029 secs
  resp wait:	0.0005 secs, 0.0001 secs, 0.0039 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0029 secs

Status code distribution:
  [200]	10000 responses



