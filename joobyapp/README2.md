##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.4863 secs
  Slowest:	0.1404 secs
  Fastest:	0.0002 secs
  Average:	0.0020 secs
  Requests/sec:	20562.4318
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.014 [9810]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.028 [101]	|
  0.042 [22]	|
  0.056 [21]	|
  0.070 [18]	|
  0.084 [8]	|
  0.098 [6]	|
  0.112 [4]	|
  0.126 [1]	|
  0.140 [8]	|


Latency distribution:
  10% in 0.0004 secs
  25% in 0.0004 secs
  50% in 0.0006 secs
  75% in 0.0014 secs
  90% in 0.0036 secs
  95% in 0.0055 secs
  99% in 0.0260 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.1404 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0020 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0030 secs
  resp wait:	0.0019 secs, 0.0002 secs, 0.1403 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0201 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	1.5479 secs
  Slowest:	0.1462 secs
  Fastest:	0.0005 secs
  Average:	0.0073 secs
  Requests/sec:	6460.4673
  
  Total data:	307733 bytes
  Size/request:	30 bytes

Response time histogram:
  0.001 [1]	|
  0.015 [8516]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.030 [1315]	|■■■■■■
  0.044 [89]	|
  0.059 [53]	|
  0.073 [2]	|
  0.088 [0]	|
  0.103 [0]	|
  0.117 [0]	|
  0.132 [19]	|
  0.146 [5]	|


Latency distribution:
  10% in 0.0017 secs
  25% in 0.0033 secs
  50% in 0.0042 secs
  75% in 0.0063 secs
  90% in 0.0208 secs
  95% in 0.0230 secs
  99% in 0.0364 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0005 secs, 0.1462 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0015 secs
  resp wait:	0.0072 secs, 0.0005 secs, 0.1462 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0023 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	6.9316 secs
  Slowest:	0.4150 secs
  Fastest:	0.0013 secs
  Average:	0.0336 secs
  Requests/sec:	1442.6646
  
  Total data:	3187421 bytes
  Size/request:	318 bytes

Response time histogram:
  0.001 [1]	|
  0.043 [7724]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.084 [1721]	|■■■■■■■■■
  0.125 [160]	|■
  0.167 [224]	|■
  0.208 [102]	|■
  0.250 [1]	|
  0.291 [42]	|
  0.332 [12]	|
  0.374 [4]	|
  0.415 [9]	|


Latency distribution:
  10% in 0.0037 secs
  25% in 0.0131 secs
  50% in 0.0284 secs
  75% in 0.0414 secs
  90% in 0.0543 secs
  95% in 0.0927 secs
  99% in 0.1783 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0013 secs, 0.4150 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0019 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0025 secs
  resp wait:	0.0335 secs, 0.0013 secs, 0.4150 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	1.2725 secs
  Slowest:	0.1784 secs
  Fastest:	0.0004 secs
  Average:	0.0058 secs
  Requests/sec:	7858.3480
  
  Total data:	21300000 bytes
  Size/request:	2130 bytes

Response time histogram:
  0.000 [1]	|
  0.018 [8955]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.036 [849]	|■■■■
  0.054 [119]	|■
  0.072 [20]	|
  0.089 [19]	|
  0.107 [16]	|
  0.125 [1]	|
  0.143 [17]	|
  0.161 [0]	|
  0.178 [3]	|


Latency distribution:
  10% in 0.0012 secs
  25% in 0.0022 secs
  50% in 0.0028 secs
  75% in 0.0039 secs
  90% in 0.0187 secs
  95% in 0.0206 secs
  99% in 0.0461 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.1784 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0024 secs
  resp wait:	0.0057 secs, 0.0003 secs, 0.1784 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0046 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	13.1403 secs
  Slowest:	0.2109 secs
  Fastest:	0.0089 secs
  Average:	0.0652 secs
  Requests/sec:	761.0173
  
  Total data:	3187718 bytes
  Size/request:	318 bytes

Response time histogram:
  0.009 [1]	|
  0.029 [23]	|
  0.049 [431]	|■■■
  0.069 [6720]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.090 [2076]	|■■■■■■■■■■■■
  0.110 [429]	|■■■
  0.130 [154]	|■
  0.150 [73]	|
  0.170 [14]	|
  0.191 [46]	|
  0.211 [33]	|


Latency distribution:
  10% in 0.0510 secs
  25% in 0.0540 secs
  50% in 0.0589 secs
  75% in 0.0724 secs
  90% in 0.0842 secs
  95% in 0.1018 secs
  99% in 0.1454 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0089 secs, 0.2109 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0015 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0651 secs, 0.0088 secs, 0.2108 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1402 secs
  Slowest:	0.0091 secs
  Fastest:	0.0001 secs
  Average:	0.0007 secs
  Requests/sec:	71307.0439
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8579]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1122]	|■■■■■
  0.003 [220]	|■
  0.004 [27]	|
  0.005 [0]	|
  0.005 [1]	|
  0.006 [7]	|
  0.007 [2]	|
  0.008 [23]	|
  0.009 [18]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0004 secs
  50% in 0.0005 secs
  75% in 0.0007 secs
  90% in 0.0012 secs
  95% in 0.0016 secs
  99% in 0.0026 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0091 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0034 secs
  resp wait:	0.0005 secs, 0.0001 secs, 0.0047 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0031 secs

Status code distribution:
  [200]	10000 responses



