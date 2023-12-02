##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.2359 secs
  Slowest:	0.0111 secs
  Fastest:	0.0001 secs
  Average:	0.0011 secs
  Requests/sec:	42390.9562
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [7088]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [2333]	|■■■■■■■■■■■■■
  0.003 [477]	|■■■
  0.005 [46]	|
  0.006 [7]	|
  0.007 [6]	|
  0.008 [12]	|
  0.009 [15]	|
  0.010 [7]	|
  0.011 [8]	|


Latency distribution:
  10% in 0.0005 secs
  25% in 0.0008 secs
  50% in 0.0009 secs
  75% in 0.0013 secs
  90% in 0.0020 secs
  95% in 0.0024 secs
  99% in 0.0034 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0111 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0021 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0028 secs
  resp wait:	0.0010 secs, 0.0001 secs, 0.0041 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0018 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	1.7672 secs
  Slowest:	0.0899 secs
  Fastest:	0.0004 secs
  Average:	0.0088 secs
  Requests/sec:	5658.5317
  
  Total data:	307836 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.009 [8013]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.018 [207]	|■
  0.027 [1606]	|■■■■■■■■
  0.036 [24]	|
  0.045 [9]	|
  0.054 [33]	|
  0.063 [27]	|
  0.072 [39]	|
  0.081 [32]	|
  0.090 [9]	|


Latency distribution:
  10% in 0.0037 secs
  25% in 0.0041 secs
  50% in 0.0051 secs
  75% in 0.0079 secs
  90% in 0.0219 secs
  95% in 0.0229 secs
  99% in 0.0594 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.0899 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0015 secs
  resp wait:	0.0087 secs, 0.0003 secs, 0.0898 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	10.5756 secs
  Slowest:	0.1284 secs
  Fastest:	0.0181 secs
  Average:	0.0528 secs
  Requests/sec:	945.5699
  
  Total data:	3187810 bytes
  Size/request:	318 bytes

Response time histogram:
  0.018 [1]	|
  0.029 [378]	|■■■
  0.040 [222]	|■■
  0.051 [5143]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.062 [2814]	|■■■■■■■■■■■■■■■■■■■■■■
  0.073 [571]	|■■■■
  0.084 [408]	|■■■
  0.095 [243]	|■■
  0.106 [183]	|■
  0.117 [26]	|
  0.128 [11]	|


Latency distribution:
  10% in 0.0433 secs
  25% in 0.0465 secs
  50% in 0.0500 secs
  75% in 0.0549 secs
  90% in 0.0712 secs
  95% in 0.0829 secs
  99% in 0.1025 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0181 secs, 0.1284 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0019 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0527 secs, 0.0181 secs, 0.1283 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	2.5059 secs
  Slowest:	0.0997 secs
  Fastest:	0.0011 secs
  Average:	0.0125 secs
  Requests/sec:	3990.5494
  
  Total data:	16050000 bytes
  Size/request:	1605 bytes

Response time histogram:
  0.001 [1]	|
  0.011 [6906]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.021 [700]	|■■■■
  0.031 [2087]	|■■■■■■■■■■■■
  0.041 [129]	|■
  0.050 [28]	|
  0.060 [49]	|
  0.070 [9]	|
  0.080 [0]	|
  0.090 [14]	|
  0.100 [77]	|


Latency distribution:
  10% in 0.0051 secs
  25% in 0.0058 secs
  50% in 0.0078 secs
  75% in 0.0149 secs
  90% in 0.0248 secs
  95% in 0.0282 secs
  99% in 0.0624 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0011 secs, 0.0997 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0019 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0033 secs
  resp wait:	0.0124 secs, 0.0011 secs, 0.0996 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0017 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	29.2900 secs
  Slowest:	0.4187 secs
  Fastest:	0.0954 secs
  Average:	0.1464 secs
  Requests/sec:	341.4131
  
  Total data:	3187969 bytes
  Size/request:	318 bytes

Response time histogram:
  0.095 [1]	|
  0.128 [3126]	|■■■■■■■■■■■■■■■■■■■■■■■■■
  0.160 [4920]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.192 [987]	|■■■■■■■■
  0.225 [456]	|■■■■
  0.257 [329]	|■■■
  0.289 [83]	|■
  0.322 [23]	|
  0.354 [18]	|
  0.386 [22]	|
  0.419 [35]	|


Latency distribution:
  10% in 0.1103 secs
  25% in 0.1214 secs
  50% in 0.1401 secs
  75% in 0.1530 secs
  90% in 0.1912 secs
  95% in 0.2255 secs
  99% in 0.2848 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0954 secs, 0.4187 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0010 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.1463 secs, 0.0953 secs, 0.4186 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.2597 secs
  Slowest:	0.0087 secs
  Fastest:	0.0001 secs
  Average:	0.0012 secs
  Requests/sec:	38506.5482
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [5523]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [2610]	|■■■■■■■■■■■■■■■■■■■
  0.003 [1112]	|■■■■■■■■
  0.004 [427]	|■■■
  0.004 [214]	|■■
  0.005 [53]	|
  0.006 [14]	|
  0.007 [20]	|
  0.008 [14]	|
  0.009 [12]	|


Latency distribution:
  10% in 0.0005 secs
  25% in 0.0008 secs
  50% in 0.0009 secs
  75% in 0.0015 secs
  90% in 0.0025 secs
  95% in 0.0030 secs
  99% in 0.0045 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0087 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0024 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0026 secs
  resp wait:	0.0010 secs, 0.0001 secs, 0.0058 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0029 secs

Status code distribution:
  [200]	10000 responses



