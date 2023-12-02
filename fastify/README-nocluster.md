##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.6028 secs
  Slowest:	0.1353 secs
  Fastest:	0.0001 secs
  Average:	0.0029 secs
  Requests/sec:	16588.1644
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.014 [9956]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.027 [7]	|
  0.041 [6]	|
  0.054 [4]	|
  0.068 [6]	|
  0.081 [5]	|
  0.095 [3]	|
  0.108 [4]	|
  0.122 [3]	|
  0.135 [5]	|


Latency distribution:
  10% in 0.0015 secs
  25% in 0.0019 secs
  50% in 0.0024 secs
  75% in 0.0033 secs
  90% in 0.0040 secs
  95% in 0.0042 secs
  99% in 0.0066 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.1353 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:	0.0028 secs, 0.0000 secs, 0.1353 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0012 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	1.2502 secs
  Slowest:	0.1229 secs
  Fastest:	0.0003 secs
  Average:	0.0062 secs
  Requests/sec:	7998.8187
  
  Total data:	307761 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.013 [9944]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.025 [17]	|
  0.037 [4]	|
  0.049 [5]	|
  0.062 [6]	|
  0.074 [4]	|
  0.086 [4]	|
  0.098 [5]	|
  0.111 [4]	|
  0.123 [6]	|


Latency distribution:
  10% in 0.0045 secs
  25% in 0.0046 secs
  50% in 0.0051 secs
  75% in 0.0070 secs
  90% in 0.0084 secs
  95% in 0.0100 secs
  99% in 0.0117 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.1229 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0015 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0008 secs
  resp wait:	0.0061 secs, 0.0003 secs, 0.1183 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	6.1191 secs
  Slowest:	0.1204 secs
  Fastest:	0.0096 secs
  Average:	0.0305 secs
  Requests/sec:	1634.2358
  
  Total data:	3187310 bytes
  Size/request:	318 bytes

Response time histogram:
  0.010 [1]	|
  0.021 [38]	|
  0.032 [7555]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.043 [2003]	|■■■■■■■■■■■
  0.054 [325]	|■■
  0.065 [23]	|
  0.076 [5]	|
  0.087 [0]	|
  0.098 [0]	|
  0.109 [0]	|
  0.120 [50]	|


Latency distribution:
  10% in 0.0265 secs
  25% in 0.0271 secs
  50% in 0.0280 secs
  75% in 0.0312 secs
  90% in 0.0385 secs
  95% in 0.0419 secs
  99% in 0.0517 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0096 secs, 0.1204 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0011 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0013 secs
  resp wait:	0.0305 secs, 0.0090 secs, 0.1204 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	5.6232 secs
  Slowest:	0.1406 secs
  Fastest:	0.0011 secs
  Average:	0.0279 secs
  Requests/sec:	1778.3408
  
  Total data:	14970000 bytes
  Size/request:	1497 bytes

Response time histogram:
  0.001 [1]	|
  0.015 [52]	|
  0.029 [7215]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.043 [2676]	|■■■■■■■■■■■■■■■
  0.057 [29]	|
  0.071 [5]	|
  0.085 [5]	|
  0.099 [5]	|
  0.113 [4]	|
  0.127 [3]	|
  0.141 [5]	|


Latency distribution:
  10% in 0.0239 secs
  25% in 0.0250 secs
  50% in 0.0269 secs
  75% in 0.0296 secs
  90% in 0.0341 secs
  95% in 0.0365 secs
  99% in 0.0408 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0011 secs, 0.1406 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0014 secs
  resp wait:	0.0279 secs, 0.0010 secs, 0.1365 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	12.5993 secs
  Slowest:	0.1287 secs
  Fastest:	0.0516 secs
  Average:	0.0628 secs
  Requests/sec:	793.6958
  
  Total data:	3187889 bytes
  Size/request:	318 bytes

Response time histogram:
  0.052 [1]	|
  0.059 [5933]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.067 [1551]	|■■■■■■■■■■
  0.075 [1233]	|■■■■■■■■
  0.082 [607]	|■■■■
  0.090 [315]	|■■
  0.098 [184]	|■
  0.106 [59]	|
  0.113 [13]	|
  0.121 [52]	|
  0.129 [52]	|


Latency distribution:
  10% in 0.0542 secs
  25% in 0.0557 secs
  50% in 0.0578 secs
  75% in 0.0671 secs
  90% in 0.0777 secs
  95% in 0.0860 secs
  99% in 0.1140 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0516 secs, 0.1287 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0628 secs, 0.0515 secs, 0.1286 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.4500 secs
  Slowest:	0.0575 secs
  Fastest:	0.0001 secs
  Average:	0.0022 secs
  Requests/sec:	22221.0307
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.006 [9962]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.012 [12]	|
  0.017 [5]	|
  0.023 [2]	|
  0.029 [4]	|
  0.035 [3]	|
  0.040 [3]	|
  0.046 [2]	|
  0.052 [3]	|
  0.058 [3]	|


Latency distribution:
  10% in 0.0014 secs
  25% in 0.0018 secs
  50% in 0.0021 secs
  75% in 0.0023 secs
  90% in 0.0026 secs
  95% in 0.0030 secs
  99% in 0.0047 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0575 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0018 secs
  resp wait:	0.0021 secs, 0.0000 secs, 0.0575 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0017 secs

Status code distribution:
  [200]	10000 responses



