##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0726 secs
  Slowest:	0.0068 secs
  Fastest:	0.0000 secs
  Average:	0.0004 secs
  Requests/sec:	137822.5197
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8783]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1074]	|■■■■■
  0.002 [58]	|
  0.003 [22]	|
  0.003 [14]	|
  0.004 [15]	|
  0.005 [16]	|
  0.005 [3]	|
  0.006 [2]	|
  0.007 [12]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0002 secs
  75% in 0.0004 secs
  90% in 0.0007 secs
  95% in 0.0010 secs
  99% in 0.0016 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0068 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0002 secs, 0.0000 secs, 0.0019 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0019 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.1696 secs
  Slowest:	0.0190 secs
  Fastest:	0.0000 secs
  Average:	0.0008 secs
  Requests/sec:	58960.4967
  
  Total data:	307756 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9949]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [1]	|
  0.006 [7]	|
  0.008 [11]	|
  0.010 [11]	|
  0.011 [4]	|
  0.013 [4]	|
  0.015 [4]	|
  0.017 [4]	|
  0.019 [4]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0007 secs
  50% in 0.0009 secs
  75% in 0.0009 secs
  90% in 0.0011 secs
  95% in 0.0011 secs
  99% in 0.0012 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0190 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:	0.0008 secs, 0.0000 secs, 0.0172 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	0.7640 secs
  Slowest:	0.2220 secs
  Fastest:	0.0002 secs
  Average:	0.0034 secs
  Requests/sec:	13089.4900
  
  Total data:	3187931 bytes
  Size/request:	318 bytes

Response time histogram:
  0.000 [1]	|
  0.022 [9917]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.045 [14]	|
  0.067 [15]	|
  0.089 [14]	|
  0.111 [13]	|
  0.133 [21]	|
  0.155 [2]	|
  0.178 [0]	|
  0.200 [0]	|
  0.222 [3]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0020 secs
  50% in 0.0030 secs
  75% in 0.0037 secs
  90% in 0.0040 secs
  95% in 0.0048 secs
  99% in 0.0108 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.2220 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:	0.0034 secs, 0.0002 secs, 0.2220 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.2938 secs
  Slowest:	0.0048 secs
  Fastest:	0.0001 secs
  Average:	0.0014 secs
  Requests/sec:	34035.9247
  
  Total data:	13080000 bytes
  Size/request:	1308 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [96]	|■
  0.001 [3404]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [788]	|■■■■■■■
  0.002 [4517]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1020]	|■■■■■■■■■
  0.003 [133]	|■
  0.003 [22]	|
  0.004 [6]	|
  0.004 [0]	|
  0.005 [13]	|


Latency distribution:
  10% in 0.0008 secs
  25% in 0.0009 secs
  50% in 0.0017 secs
  75% in 0.0018 secs
  90% in 0.0020 secs
  95% in 0.0022 secs
  99% in 0.0026 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0048 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0014 secs, 0.0001 secs, 0.0035 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	2.4087 secs
  Slowest:	0.0456 secs
  Fastest:	0.0017 secs
  Average:	0.0120 secs
  Requests/sec:	4151.6293
  
  Total data:	3187873 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.006 [37]	|
  0.010 [1205]	|■■■■■■
  0.015 [8440]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.019 [124]	|■
  0.024 [143]	|■
  0.028 [0]	|
  0.032 [0]	|
  0.037 [0]	|
  0.041 [12]	|
  0.046 [38]	|


Latency distribution:
  10% in 0.0103 secs
  25% in 0.0111 secs
  50% in 0.0117 secs
  75% in 0.0124 secs
  90% in 0.0131 secs
  95% in 0.0141 secs
  99% in 0.0218 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0017 secs, 0.0456 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:	0.0119 secs, 0.0017 secs, 0.0456 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0044 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0527 secs
  Slowest:	0.0033 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	189801.9417
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [7971]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1517]	|■■■■■■■■
  0.001 [419]	|■■
  0.001 [33]	|
  0.002 [24]	|
  0.002 [8]	|
  0.002 [8]	|
  0.003 [1]	|
  0.003 [3]	|
  0.003 [15]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0002 secs
  75% in 0.0003 secs
  90% in 0.0005 secs
  95% in 0.0007 secs
  99% in 0.0010 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0033 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0006 secs
  resp wait:	0.0002 secs, 0.0000 secs, 0.0010 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0013 secs

Status code distribution:
  [200]	10000 responses



