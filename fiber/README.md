##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.1343 secs
  Slowest:	0.0158 secs
  Fastest:	0.0001 secs
  Average:	0.0006 secs
  Requests/sec:	74471.6334
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9234]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [658]	|■■■
  0.005 [28]	|
  0.006 [23]	|
  0.008 [31]	|
  0.010 [14]	|
  0.011 [5]	|
  0.013 [2]	|
  0.014 [3]	|
  0.016 [1]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0004 secs
  75% in 0.0007 secs
  90% in 0.0014 secs
  95% in 0.0019 secs
  99% in 0.0035 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0158 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0064 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0034 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0040 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0066 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.6668 secs
  Slowest:	0.1654 secs
  Fastest:	0.0003 secs
  Average:	0.0031 secs
  Requests/sec:	14996.7133
  
  Total data:	307777 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.017 [9426]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.033 [543]	|■■
  0.050 [0]	|
  0.066 [2]	|
  0.083 [3]	|
  0.099 [0]	|
  0.116 [3]	|
  0.132 [10]	|
  0.149 [3]	|
  0.165 [9]	|


Latency distribution:
  10% in 0.0007 secs
  25% in 0.0012 secs
  50% in 0.0017 secs
  75% in 0.0023 secs
  90% in 0.0032 secs
  95% in 0.0181 secs
  99% in 0.0208 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.1654 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0014 secs
  resp wait:	0.0030 secs, 0.0002 secs, 0.1599 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0016 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	3.5302 secs
  Slowest:	0.2768 secs
  Fastest:	0.0011 secs
  Average:	0.0162 secs
  Requests/sec:	2832.6857
  
  Total data:	3187978 bytes
  Size/request:	318 bytes

Response time histogram:
  0.001 [1]	|
  0.029 [8708]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.056 [1151]	|■■■■■
  0.084 [7]	|
  0.111 [3]	|
  0.139 [92]	|
  0.167 [23]	|
  0.194 [0]	|
  0.222 [1]	|
  0.249 [0]	|
  0.277 [14]	|


Latency distribution:
  10% in 0.0049 secs
  25% in 0.0082 secs
  50% in 0.0107 secs
  75% in 0.0252 secs
  90% in 0.0291 secs
  95% in 0.0304 secs
  99% in 0.1302 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0011 secs, 0.2768 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0015 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:	0.0161 secs, 0.0010 secs, 0.2768 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0019 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	1.0630 secs
  Slowest:	0.2116 secs
  Fastest:	0.0003 secs
  Average:	0.0049 secs
  Requests/sec:	9407.0508
  
  Total data:	13080000 bytes
  Size/request:	1308 bytes

Response time histogram:
  0.000 [1]	|
  0.021 [9568]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.043 [330]	|■
  0.064 [63]	|
  0.085 [10]	|
  0.106 [8]	|
  0.127 [4]	|
  0.148 [14]	|
  0.169 [0]	|
  0.190 [0]	|
  0.212 [2]	|


Latency distribution:
  10% in 0.0007 secs
  25% in 0.0012 secs
  50% in 0.0023 secs
  75% in 0.0038 secs
  90% in 0.0127 secs
  95% in 0.0208 secs
  99% in 0.0429 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.2116 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0022 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:	0.0047 secs, 0.0002 secs, 0.2115 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0012 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	12.9579 secs
  Slowest:	0.2817 secs
  Fastest:	0.0090 secs
  Average:	0.0645 secs
  Requests/sec:	771.7276
  
  Total data:	3187924 bytes
  Size/request:	318 bytes

Response time histogram:
  0.009 [1]	|
  0.036 [37]	|
  0.064 [7007]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.091 [2274]	|■■■■■■■■■■■■■
  0.118 [393]	|■■
  0.145 [128]	|■
  0.173 [67]	|
  0.200 [51]	|
  0.227 [19]	|
  0.254 [18]	|
  0.282 [5]	|


Latency distribution:
  10% in 0.0510 secs
  25% in 0.0536 secs
  50% in 0.0573 secs
  75% in 0.0702 secs
  90% in 0.0810 secs
  95% in 0.1000 secs
  99% in 0.1700 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0090 secs, 0.2817 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0019 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0644 secs, 0.0090 secs, 0.2816 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1400 secs
  Slowest:	0.0164 secs
  Fastest:	0.0001 secs
  Average:	0.0007 secs
  Requests/sec:	71418.7825
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9171]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [681]	|■■■
  0.005 [89]	|
  0.007 [47]	|
  0.008 [4]	|
  0.010 [1]	|
  0.012 [1]	|
  0.013 [3]	|
  0.015 [1]	|
  0.016 [1]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0002 secs
  50% in 0.0004 secs
  75% in 0.0008 secs
  90% in 0.0015 secs
  95% in 0.0021 secs
  99% in 0.0040 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0164 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0057 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0075 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0032 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0074 secs

Status code distribution:
  [200]	10000 responses



