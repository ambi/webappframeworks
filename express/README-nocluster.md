##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	1.9042 secs
  Slowest:	0.4098 secs
  Fastest:	0.0002 secs
  Average:	0.0091 secs
  Requests/sec:	5251.5496
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.041 [9950]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.082 [2]	|
  0.123 [10]	|
  0.164 [9]	|
  0.205 [5]	|
  0.246 [6]	|
  0.287 [5]	|
  0.328 [4]	|
  0.369 [3]	|
  0.410 [5]	|


Latency distribution:
  10% in 0.0052 secs
  25% in 0.0069 secs
  50% in 0.0079 secs
  75% in 0.0098 secs
  90% in 0.0113 secs
  95% in 0.0122 secs
  99% in 0.0159 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0001 secs, 0.0002 secs, 0.4098 secs
  DNS-lookup:	0.0001 secs, 0.0000 secs, 0.0226 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0023 secs
  resp wait:	0.0090 secs, 0.0001 secs, 0.3808 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	5.0618 secs
  Slowest:	0.1795 secs
  Fastest:	0.0012 secs
  Average:	0.0253 secs
  Requests/sec:	1975.5681
  
  Total data:	307862 bytes
  Size/request:	30 bytes

Response time histogram:
  0.001 [1]	|
  0.019 [4214]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.037 [5130]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.055 [320]	|■■
  0.073 [26]	|
  0.090 [181]	|■
  0.108 [50]	|
  0.126 [62]	|
  0.144 [11]	|
  0.162 [0]	|
  0.180 [5]	|


Latency distribution:
  10% in 0.0135 secs
  25% in 0.0150 secs
  50% in 0.0218 secs
  75% in 0.0303 secs
  90% in 0.0344 secs
  95% in 0.0394 secs
  99% in 0.1020 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0012 secs, 0.1795 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0015 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0017 secs
  resp wait:	0.0252 secs, 0.0011 secs, 0.1795 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0025 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	16.6262 secs
  Slowest:	0.4846 secs
  Fastest:	0.0336 secs
  Average:	0.0830 secs
  Requests/sec:	601.4595
  
  Total data:	3187755 bytes
  Size/request:	318 bytes

Response time histogram:
  0.034 [1]	|
  0.079 [5732]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.124 [3602]	|■■■■■■■■■■■■■■■■■■■■■■■■■
  0.169 [299]	|■■
  0.214 [88]	|■
  0.259 [94]	|■
  0.304 [89]	|■
  0.349 [47]	|
  0.394 [8]	|
  0.439 [7]	|
  0.485 [33]	|


Latency distribution:
  10% in 0.0611 secs
  25% in 0.0639 secs
  50% in 0.0745 secs
  75% in 0.0829 secs
  90% in 0.0977 secs
  95% in 0.1380 secs
  99% in 0.3005 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0336 secs, 0.4846 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0033 secs
  resp wait:	0.0829 secs, 0.0336 secs, 0.4845 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	6.2269 secs
  Slowest:	0.1674 secs
  Fastest:	0.0014 secs
  Average:	0.0310 secs
  Requests/sec:	1605.9372
  
  Total data:	14970000 bytes
  Size/request:	1497 bytes

Response time histogram:
  0.001 [1]	|
  0.018 [151]	|■
  0.035 [7034]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.051 [2547]	|■■■■■■■■■■■■■■
  0.068 [90]	|■
  0.084 [67]	|
  0.101 [10]	|
  0.118 [0]	|
  0.134 [64]	|
  0.151 [29]	|
  0.167 [7]	|


Latency distribution:
  10% in 0.0198 secs
  25% in 0.0221 secs
  50% in 0.0321 secs
  75% in 0.0351 secs
  90% in 0.0380 secs
  95% in 0.0410 secs
  99% in 0.1215 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0014 secs, 0.1674 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0019 secs
  resp wait:	0.0309 secs, 0.0014 secs, 0.1673 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0082 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	27.3421 secs
  Slowest:	0.3038 secs
  Fastest:	0.0681 secs
  Average:	0.1366 secs
  Requests/sec:	365.7359
  
  Total data:	3187901 bytes
  Size/request:	318 bytes

Response time histogram:
  0.068 [1]	|
  0.092 [0]	|
  0.115 [1667]	|■■■■■■■■■■■■
  0.139 [5463]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.162 [1806]	|■■■■■■■■■■■■■
  0.186 [532]	|■■■■
  0.210 [182]	|■
  0.233 [126]	|■
  0.257 [167]	|■
  0.280 [8]	|
  0.304 [48]	|


Latency distribution:
  10% in 0.1134 secs
  25% in 0.1174 secs
  50% in 0.1327 secs
  75% in 0.1409 secs
  90% in 0.1637 secs
  95% in 0.1905 secs
  99% in 0.2525 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0681 secs, 0.3038 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0013 secs
  resp wait:	0.1365 secs, 0.0680 secs, 0.3038 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0018 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	1.4564 secs
  Slowest:	0.1580 secs
  Fastest:	0.0002 secs
  Average:	0.0069 secs
  Requests/sec:	6866.1379
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.016 [9839]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.032 [34]	|
  0.047 [28]	|
  0.063 [79]	|
  0.079 [4]	|
  0.095 [3]	|
  0.111 [3]	|
  0.126 [3]	|
  0.142 [3]	|
  0.158 [3]	|


Latency distribution:
  10% in 0.0037 secs
  25% in 0.0057 secs
  50% in 0.0063 secs
  75% in 0.0070 secs
  90% in 0.0081 secs
  95% in 0.0093 secs
  99% in 0.0467 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.1580 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0018 secs
  resp wait:	0.0069 secs, 0.0001 secs, 0.1529 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0016 secs

Status code distribution:
  [200]	10000 responses



