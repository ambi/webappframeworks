##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.2356 secs
  Slowest:	0.0473 secs
  Fastest:	0.0000 secs
  Average:	0.0012 secs
  Requests/sec:	42449.4762
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.005 [9898]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.009 [2]	|
  0.014 [5]	|
  0.019 [8]	|
  0.024 [6]	|
  0.028 [51]	|
  0.033 [17]	|
  0.038 [4]	|
  0.043 [4]	|
  0.047 [4]	|


Latency distribution:
  10% in 0.0004 secs
  25% in 0.0006 secs
  50% in 0.0008 secs
  75% in 0.0011 secs
  90% in 0.0014 secs
  95% in 0.0019 secs
  99% in 0.0086 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0473 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0265 secs
  resp wait:	0.0011 secs, 0.0000 secs, 0.0427 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0012 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.3801 secs
  Slowest:	0.0427 secs
  Fastest:	0.0001 secs
  Average:	0.0019 secs
  Requests/sec:	26310.5995
  
  Total data:	307810 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.004 [9869]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.009 [80]	|
  0.013 [0]	|
  0.017 [6]	|
  0.021 [9]	|
  0.026 [8]	|
  0.030 [7]	|
  0.034 [6]	|
  0.038 [7]	|
  0.043 [7]	|


Latency distribution:
  10% in 0.0013 secs
  25% in 0.0014 secs
  50% in 0.0015 secs
  75% in 0.0019 secs
  90% in 0.0025 secs
  95% in 0.0029 secs
  99% in 0.0048 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0427 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0006 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0005 secs
  resp wait:	0.0019 secs, 0.0001 secs, 0.0412 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.7375 secs
  Slowest:	0.0204 secs
  Fastest:	0.0018 secs
  Average:	0.0087 secs
  Requests/sec:	5755.2820
  
  Total data:	3187706 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.004 [6]	|
  0.006 [14]	|
  0.007 [42]	|
  0.009 [9696]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.011 [207]	|■
  0.013 [3]	|
  0.015 [6]	|
  0.017 [4]	|
  0.019 [8]	|
  0.020 [13]	|


Latency distribution:
  10% in 0.0083 secs
  25% in 0.0085 secs
  50% in 0.0087 secs
  75% in 0.0088 secs
  90% in 0.0090 secs
  95% in 0.0091 secs
  99% in 0.0094 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0018 secs, 0.0204 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:	0.0087 secs, 0.0018 secs, 0.0188 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.4788 secs
  Slowest:	0.0296 secs
  Fastest:	0.0001 secs
  Average:	0.0024 secs
  Requests/sec:	20884.1732
  
  Total data:	14970000 bytes
  Size/request:	1497 bytes

Response time histogram:
  0.000 [1]	|
  0.003 [9809]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [135]	|■
  0.009 [5]	|
  0.012 [0]	|
  0.015 [25]	|
  0.018 [5]	|
  0.021 [4]	|
  0.024 [4]	|
  0.027 [4]	|
  0.030 [8]	|


Latency distribution:
  10% in 0.0021 secs
  25% in 0.0022 secs
  50% in 0.0022 secs
  75% in 0.0024 secs
  90% in 0.0025 secs
  95% in 0.0028 secs
  99% in 0.0045 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0296 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:	0.0024 secs, 0.0001 secs, 0.0282 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	3.8174 secs
  Slowest:	0.0529 secs
  Fastest:	0.0136 secs
  Average:	0.0191 secs
  Requests/sec:	2619.5531
  
  Total data:	3188053 bytes
  Size/request:	318 bytes

Response time histogram:
  0.014 [1]	|
  0.018 [471]	|■■
  0.021 [8809]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.025 [427]	|■■
  0.029 [242]	|■
  0.033 [0]	|
  0.037 [0]	|
  0.041 [0]	|
  0.045 [1]	|
  0.049 [2]	|
  0.053 [47]	|


Latency distribution:
  10% in 0.0178 secs
  25% in 0.0181 secs
  50% in 0.0185 secs
  75% in 0.0190 secs
  90% in 0.0200 secs
  95% in 0.0231 secs
  99% in 0.0282 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0136 secs, 0.0529 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0191 secs, 0.0136 secs, 0.0529 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1685 secs
  Slowest:	0.0085 secs
  Fastest:	0.0000 secs
  Average:	0.0008 secs
  Requests/sec:	59363.5780
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [3986]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [5857]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [123]	|■
  0.003 [15]	|
  0.004 [6]	|
  0.005 [2]	|
  0.006 [1]	|
  0.007 [2]	|
  0.008 [3]	|
  0.008 [4]	|


Latency distribution:
  10% in 0.0004 secs
  25% in 0.0006 secs
  50% in 0.0009 secs
  75% in 0.0010 secs
  90% in 0.0011 secs
  95% in 0.0011 secs
  99% in 0.0019 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0085 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0008 secs, 0.0000 secs, 0.0077 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [200]	10000 responses



