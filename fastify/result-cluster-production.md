##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.2199 secs
  Slowest:	0.0256 secs
  Fastest:	0.0000 secs
  Average:	0.0010 secs
  Requests/sec:	45475.8364
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.003 [9466]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [431]	|■■
  0.008 [26]	|
  0.010 [2]	|
  0.013 [21]	|
  0.015 [1]	|
  0.018 [19]	|
  0.020 [10]	|
  0.023 [3]	|
  0.026 [20]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0003 secs
  50% in 0.0006 secs
  75% in 0.0012 secs
  90% in 0.0020 secs
  95% in 0.0027 secs
  99% in 0.0052 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0256 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0024 secs
  resp wait:	0.0009 secs, 0.0000 secs, 0.0256 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0032 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.6448 secs
  Slowest:	0.0900 secs
  Fastest:	0.0001 secs
  Average:	0.0031 secs
  Requests/sec:	15507.7058
  
  Total data:	307787 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.009 [9809]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.018 [67]	|
  0.027 [50]	|
  0.036 [49]	|
  0.045 [0]	|
  0.054 [0]	|
  0.063 [0]	|
  0.072 [18]	|
  0.081 [3]	|
  0.090 [3]	|


Latency distribution:
  10% in 0.0005 secs
  25% in 0.0016 secs
  50% in 0.0027 secs
  75% in 0.0034 secs
  90% in 0.0042 secs
  95% in 0.0057 secs
  99% in 0.0218 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0900 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0030 secs, 0.0001 secs, 0.0887 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0018 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.9519 secs
  Slowest:	0.3838 secs
  Fastest:	0.0004 secs
  Average:	0.0093 secs
  Requests/sec:	5123.1446
  
  Total data:	3188266 bytes
  Size/request:	318 bytes

Response time histogram:
  0.000 [1]	|
  0.039 [9854]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.077 [37]	|
  0.115 [1]	|
  0.154 [85]	|
  0.192 [0]	|
  0.230 [0]	|
  0.269 [19]	|
  0.307 [2]	|
  0.345 [0]	|
  0.384 [1]	|


Latency distribution:
  10% in 0.0015 secs
  25% in 0.0053 secs
  50% in 0.0075 secs
  75% in 0.0096 secs
  90% in 0.0129 secs
  95% in 0.0149 secs
  99% in 0.1285 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.3838 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0006 secs
  resp wait:	0.0092 secs, 0.0004 secs, 0.3838 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.4843 secs
  Slowest:	0.2196 secs
  Fastest:	0.0001 secs
  Average:	0.0018 secs
  Requests/sec:	20650.2097
  
  Total data:	14970000 bytes
  Size/request:	1497 bytes

Response time histogram:
  0.000 [1]	|
  0.022 [9937]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.044 [12]	|
  0.066 [2]	|
  0.088 [5]	|
  0.110 [19]	|
  0.132 [1]	|
  0.154 [14]	|
  0.176 [2]	|
  0.198 [1]	|
  0.220 [6]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0005 secs
  75% in 0.0012 secs
  90% in 0.0026 secs
  95% in 0.0036 secs
  99% in 0.0170 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.2196 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0006 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0017 secs, 0.0001 secs, 0.2180 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	4.0191 secs
  Slowest:	0.0603 secs
  Fastest:	0.0019 secs
  Average:	0.0198 secs
  Requests/sec:	2488.1408
  
  Total data:	3187945 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.008 [58]	|■
  0.014 [766]	|■■■■■■■
  0.019 [4124]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.025 [4134]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.031 [651]	|■■■■■■
  0.037 [146]	|■
  0.043 [44]	|
  0.049 [21]	|
  0.054 [30]	|
  0.060 [25]	|


Latency distribution:
  10% in 0.0141 secs
  25% in 0.0175 secs
  50% in 0.0195 secs
  75% in 0.0215 secs
  90% in 0.0248 secs
  95% in 0.0276 secs
  99% in 0.0384 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0019 secs, 0.0603 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0198 secs, 0.0019 secs, 0.0602 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0842 secs
  Slowest:	0.0038 secs
  Fastest:	0.0000 secs
  Average:	0.0004 secs
  Requests/sec:	118736.0542
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [5392]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [3485]	|■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [889]	|■■■■■■■
  0.002 [175]	|■
  0.002 [15]	|
  0.002 [15]	|
  0.003 [20]	|
  0.003 [2]	|
  0.003 [1]	|
  0.004 [5]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0004 secs
  75% in 0.0006 secs
  90% in 0.0008 secs
  95% in 0.0010 secs
  99% in 0.0013 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0038 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0010 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0008 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0015 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [200]	10000 responses



