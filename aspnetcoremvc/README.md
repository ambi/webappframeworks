##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.2589 secs
  Slowest:	0.0245 secs
  Fastest:	0.0002 secs
  Average:	0.0013 secs
  Requests/sec:	38629.8833


Response time histogram:
  0.000 [1]	|
  0.003 [9803]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [142]	|■
  0.007 [3]	|
  0.010 [1]	|
  0.012 [0]	|
  0.015 [0]	|
  0.017 [0]	|
  0.020 [0]	|
  0.022 [22]	|
  0.025 [28]	|


Latency distribution:
  10% in 0.0006 secs
  25% in 0.0008 secs
  50% in 0.0011 secs
  75% in 0.0013 secs
  90% in 0.0018 secs
  95% in 0.0022 secs
  99% in 0.0032 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.0245 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0025 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0021 secs
  resp wait:	0.0010 secs, 0.0002 secs, 0.0167 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0041 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	1.1320 secs
  Slowest:	0.2321 secs
  Fastest:	0.0007 secs
  Average:	0.0055 secs
  Requests/sec:	8834.1284


Response time histogram:
  0.001 [1]	|
  0.024 [9864]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.047 [34]	|
  0.070 [51]	|
  0.093 [0]	|
  0.116 [0]	|
  0.140 [1]	|
  0.163 [0]	|
  0.186 [0]	|
  0.209 [42]	|
  0.232 [7]	|


Latency distribution:
  10% in 0.0027 secs
  25% in 0.0033 secs
  50% in 0.0039 secs
  75% in 0.0046 secs
  90% in 0.0056 secs
  95% in 0.0074 secs
  99% in 0.0556 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0007 secs, 0.2321 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0014 secs
  resp wait:	0.0054 secs, 0.0007 secs, 0.2320 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0213 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	5.0700 secs
  Slowest:	0.1744 secs
  Fastest:	0.0034 secs
  Average:	0.0249 secs
  Requests/sec:	1972.3810


Response time histogram:
  0.003 [1]	|
  0.021 [2370]	|■■■■■■■■■■■■■
  0.038 [7129]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.055 [367]	|■■
  0.072 [73]	|
  0.089 [24]	|
  0.106 [9]	|
  0.123 [8]	|
  0.140 [14]	|
  0.157 [4]	|
  0.174 [1]	|


Latency distribution:
  10% in 0.0173 secs
  25% in 0.0207 secs
  50% in 0.0236 secs
  75% in 0.0272 secs
  90% in 0.0322 secs
  95% in 0.0376 secs
  99% in 0.0619 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0034 secs, 0.1744 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0017 secs
  resp wait:	0.0248 secs, 0.0034 secs, 0.1744 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0136 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	1.0367 secs
  Slowest:	0.0748 secs
  Fastest:	0.0006 secs
  Average:	0.0051 secs
  Requests/sec:	9646.0492


Response time histogram:
  0.001 [1]	|
  0.008 [9487]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.015 [363]	|■■
  0.023 [29]	|
  0.030 [13]	|
  0.038 [5]	|
  0.045 [1]	|
  0.053 [4]	|
  0.060 [28]	|
  0.067 [66]	|
  0.075 [3]	|


Latency distribution:
  10% in 0.0030 secs
  25% in 0.0036 secs
  50% in 0.0041 secs
  75% in 0.0050 secs
  90% in 0.0064 secs
  95% in 0.0081 secs
  99% in 0.0452 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0006 secs, 0.0748 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0049 secs, 0.0005 secs, 0.0738 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0111 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	8.6386 secs
  Slowest:	0.2438 secs
  Fastest:	0.0049 secs
  Average:	0.0428 secs
  Requests/sec:	1157.6007


Response time histogram:
  0.005 [1]	|
  0.029 [746]	|■■■■
  0.053 [7899]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.077 [1115]	|■■■■■■
  0.100 [145]	|■
  0.124 [23]	|
  0.148 [20]	|
  0.172 [1]	|
  0.196 [0]	|
  0.220 [0]	|
  0.244 [50]	|


Latency distribution:
  10% in 0.0302 secs
  25% in 0.0346 secs
  50% in 0.0398 secs
  75% in 0.0471 secs
  90% in 0.0556 secs
  95% in 0.0634 secs
  99% in 0.0986 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0049 secs, 0.2438 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0427 secs, 0.0049 secs, 0.2398 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0051 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1632 secs
  Slowest:	0.0091 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	61266.0733


Response time histogram:
  0.000 [1]	|
  0.001 [8226]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1698]	|■■■■■■■■
  0.003 [25]	|
  0.004 [0]	|
  0.005 [0]	|
  0.006 [0]	|
  0.006 [7]	|
  0.007 [33]	|
  0.008 [4]	|
  0.009 [6]	|


Latency distribution:
  10% in 0.0005 secs
  25% in 0.0006 secs
  50% in 0.0007 secs
  75% in 0.0009 secs
  90% in 0.0013 secs
  95% in 0.0015 secs
  99% in 0.0018 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0091 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0006 secs, 0.0001 secs, 0.0059 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0014 secs

Status code distribution:
  [200]	10000 responses

