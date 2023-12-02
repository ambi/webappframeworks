##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.2875 secs
  Slowest:	0.0214 secs
  Fastest:	0.0001 secs
  Average:	0.0014 secs
  Requests/sec:	34777.2309
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [8401]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [1127]	|■■■■■
  0.006 [208]	|■
  0.009 [52]	|
  0.011 [84]	|
  0.013 [47]	|
  0.015 [34]	|
  0.017 [13]	|
  0.019 [15]	|
  0.021 [18]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0004 secs
  50% in 0.0007 secs
  75% in 0.0014 secs
  90% in 0.0030 secs
  95% in 0.0042 secs
  99% in 0.0118 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0214 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0032 secs
  req write:	0.0001 secs, 0.0000 secs, 0.0106 secs
  resp wait:	0.0008 secs, 0.0001 secs, 0.0124 secs
  resp read:	0.0004 secs, 0.0000 secs, 0.0140 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.9366 secs
  Slowest:	0.1290 secs
  Fastest:	0.0004 secs
  Average:	0.0045 secs
  Requests/sec:	10676.7143
  
  Total data:	307780 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.013 [9146]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.026 [738]	|■■■
  0.039 [28]	|
  0.052 [29]	|
  0.065 [1]	|
  0.078 [18]	|
  0.090 [21]	|
  0.103 [0]	|
  0.116 [0]	|
  0.129 [18]	|


Latency distribution:
  10% in 0.0011 secs
  25% in 0.0019 secs
  50% in 0.0025 secs
  75% in 0.0033 secs
  90% in 0.0058 secs
  95% in 0.0196 secs
  99% in 0.0333 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.1290 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0026 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0023 secs
  resp wait:	0.0044 secs, 0.0003 secs, 0.1289 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0018 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	3.7159 secs
  Slowest:	0.3817 secs
  Fastest:	0.0010 secs
  Average:	0.0174 secs
  Requests/sec:	2691.1460
  
  Total data:	3188156 bytes
  Size/request:	318 bytes

Response time histogram:
  0.001 [1]	|
  0.039 [9604]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.077 [74]	|
  0.115 [52]	|
  0.153 [159]	|■
  0.191 [41]	|
  0.229 [8]	|
  0.268 [36]	|
  0.306 [22]	|
  0.344 [0]	|
  0.382 [3]	|


Latency distribution:
  10% in 0.0044 secs
  25% in 0.0059 secs
  50% in 0.0090 secs
  75% in 0.0223 secs
  90% in 0.0285 secs
  95% in 0.0338 secs
  99% in 0.1640 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0010 secs, 0.3817 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0026 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0022 secs
  resp wait:	0.0173 secs, 0.0009 secs, 0.3817 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.9264 secs
  Slowest:	0.1465 secs
  Fastest:	0.0003 secs
  Average:	0.0042 secs
  Requests/sec:	10794.0140
  
  Total data:	14440000 bytes
  Size/request:	1444 bytes

Response time histogram:
  0.000 [1]	|
  0.015 [9175]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.030 [763]	|■■■
  0.044 [23]	|
  0.059 [12]	|
  0.073 [1]	|
  0.088 [0]	|
  0.103 [0]	|
  0.117 [1]	|
  0.132 [12]	|
  0.146 [12]	|


Latency distribution:
  10% in 0.0010 secs
  25% in 0.0015 secs
  50% in 0.0023 secs
  75% in 0.0032 secs
  90% in 0.0064 secs
  95% in 0.0192 secs
  99% in 0.0231 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.1465 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0023 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0024 secs
  resp wait:	0.0040 secs, 0.0003 secs, 0.1463 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0020 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	13.5332 secs
  Slowest:	0.4537 secs
  Fastest:	0.0120 secs
  Average:	0.0674 secs
  Requests/sec:	738.9246
  
  Total data:	3187982 bytes
  Size/request:	318 bytes

Response time histogram:
  0.012 [1]	|
  0.056 [4369]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.100 [4927]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.144 [480]	|■■■■
  0.189 [116]	|■
  0.233 [57]	|
  0.277 [10]	|
  0.321 [2]	|
  0.365 [3]	|
  0.410 [3]	|
  0.454 [32]	|


Latency distribution:
  10% in 0.0496 secs
  25% in 0.0529 secs
  50% in 0.0573 secs
  75% in 0.0733 secs
  90% in 0.0932 secs
  95% in 0.1160 secs
  99% in 0.1912 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0120 secs, 0.4537 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0027 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:	0.0673 secs, 0.0118 secs, 0.4537 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1445 secs
  Slowest:	0.0064 secs
  Fastest:	0.0001 secs
  Average:	0.0007 secs
  Requests/sec:	69201.8583
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [7329]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1445]	|■■■■■■■■
  0.002 [714]	|■■■■
  0.003 [348]	|■■
  0.003 [98]	|■
  0.004 [14]	|
  0.004 [1]	|
  0.005 [25]	|
  0.006 [11]	|
  0.006 [14]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0005 secs
  75% in 0.0007 secs
  90% in 0.0015 secs
  95% in 0.0020 secs
  99% in 0.0029 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0064 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0019 secs
  resp wait:	0.0004 secs, 0.0001 secs, 0.0021 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0030 secs

Status code distribution:
  [200]	10000 responses



