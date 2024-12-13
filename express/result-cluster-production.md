##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.3948 secs
  Slowest:	0.1089 secs
  Fastest:	0.0001 secs
  Average:	0.0017 secs
  Requests/sec:	25332.1573
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.011 [9849]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.022 [72]	|
  0.033 [39]	|
  0.044 [5]	|
  0.054 [10]	|
  0.065 [8]	|
  0.076 [1]	|
  0.087 [1]	|
  0.098 [1]	|
  0.109 [13]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0004 secs
  50% in 0.0008 secs
  75% in 0.0016 secs
  90% in 0.0029 secs
  95% in 0.0040 secs
  99% in 0.0178 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.1089 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0035 secs
  resp wait:	0.0016 secs, 0.0000 secs, 0.1089 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0044 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.7943 secs
  Slowest:	0.1557 secs
  Fastest:	0.0001 secs
  Average:	0.0037 secs
  Requests/sec:	12589.4262
  
  Total data:	307735 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.016 [9846]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.031 [99]	|
  0.047 [12]	|
  0.062 [20]	|
  0.078 [1]	|
  0.093 [6]	|
  0.109 [1]	|
  0.125 [4]	|
  0.140 [5]	|
  0.156 [5]	|


Latency distribution:
  10% in 0.0005 secs
  25% in 0.0017 secs
  50% in 0.0033 secs
  75% in 0.0042 secs
  90% in 0.0052 secs
  95% in 0.0064 secs
  99% in 0.0205 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.1557 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0036 secs, 0.0001 secs, 0.1543 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	2.2875 secs
  Slowest:	0.4674 secs
  Fastest:	0.0004 secs
  Average:	0.0110 secs
  Requests/sec:	4371.5900
  
  Total data:	3187726 bytes
  Size/request:	318 bytes

Response time histogram:
  0.000 [1]	|
  0.047 [9867]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.094 [5]	|
  0.141 [117]	|
  0.187 [5]	|
  0.234 [0]	|
  0.281 [3]	|
  0.327 [0]	|
  0.374 [0]	|
  0.421 [1]	|
  0.467 [1]	|


Latency distribution:
  10% in 0.0031 secs
  25% in 0.0068 secs
  50% in 0.0089 secs
  75% in 0.0119 secs
  90% in 0.0154 secs
  95% in 0.0184 secs
  99% in 0.1318 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.4674 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0035 secs
  resp wait:	0.0109 secs, 0.0004 secs, 0.4674 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0048 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.7827 secs
  Slowest:	0.3872 secs
  Fastest:	0.0001 secs
  Average:	0.0033 secs
  Requests/sec:	12775.9519
  
  Total data:	14970000 bytes
  Size/request:	1497 bytes

Response time histogram:
  0.000 [1]	|
  0.039 [9939]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.078 [20]	|
  0.116 [8]	|
  0.155 [13]	|
  0.194 [1]	|
  0.232 [10]	|
  0.271 [6]	|
  0.310 [0]	|
  0.349 [0]	|
  0.387 [2]	|


Latency distribution:
  10% in 0.0006 secs
  25% in 0.0013 secs
  50% in 0.0023 secs
  75% in 0.0032 secs
  90% in 0.0043 secs
  95% in 0.0051 secs
  99% in 0.0221 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.3872 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0006 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0038 secs
  resp wait:	0.0033 secs, 0.0001 secs, 0.3872 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0051 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	4.3218 secs
  Slowest:	0.0365 secs
  Fastest:	0.0022 secs
  Average:	0.0214 secs
  Requests/sec:	2313.8378
  
  Total data:	3187817 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.006 [32]	|
  0.009 [40]	|
  0.012 [165]	|■
  0.016 [490]	|■■■■
  0.019 [1122]	|■■■■■■■■■
  0.023 [5205]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.026 [2190]	|■■■■■■■■■■■■■■■■■
  0.030 [582]	|■■■■
  0.033 [163]	|■
  0.036 [10]	|


Latency distribution:
  10% in 0.0174 secs
  25% in 0.0200 secs
  50% in 0.0215 secs
  75% in 0.0231 secs
  90% in 0.0255 secs
  95% in 0.0271 secs
  99% in 0.0303 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0022 secs, 0.0365 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:	0.0214 secs, 0.0022 secs, 0.0364 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.2473 secs
  Slowest:	0.0597 secs
  Fastest:	0.0000 secs
  Average:	0.0011 secs
  Requests/sec:	40435.1292
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.006 [9838]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.012 [85]	|
  0.018 [9]	|
  0.024 [0]	|
  0.030 [47]	|
  0.036 [7]	|
  0.042 [3]	|
  0.048 [8]	|
  0.054 [0]	|
  0.060 [2]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0003 secs
  50% in 0.0007 secs
  75% in 0.0012 secs
  90% in 0.0018 secs
  95% in 0.0025 secs
  99% in 0.0085 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0597 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0038 secs
  resp wait:	0.0010 secs, 0.0000 secs, 0.0597 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0037 secs

Status code distribution:
  [200]	10000 responses



