##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0729 secs
  Slowest:	0.0047 secs
  Fastest:	0.0000 secs
  Average:	0.0004 secs
  Requests/sec:	137264.0411
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [7646]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1918]	|■■■■■■■■■■
  0.001 [320]	|■■
  0.002 [44]	|
  0.002 [18]	|
  0.003 [3]	|
  0.003 [0]	|
  0.004 [0]	|
  0.004 [33]	|
  0.005 [17]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0003 secs
  75% in 0.0005 secs
  90% in 0.0007 secs
  95% in 0.0009 secs
  99% in 0.0015 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0047 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0014 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0020 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0019 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.2017 secs
  Slowest:	0.0201 secs
  Fastest:	0.0000 secs
  Average:	0.0010 secs
  Requests/sec:	49586.4901
  
  Total data:	307787 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9904]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [55]	|
  0.006 [6]	|
  0.008 [5]	|
  0.010 [2]	|
  0.012 [4]	|
  0.014 [4]	|
  0.016 [10]	|
  0.018 [5]	|
  0.020 [4]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0005 secs
  50% in 0.0011 secs
  75% in 0.0013 secs
  90% in 0.0014 secs
  95% in 0.0016 secs
  99% in 0.0020 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0201 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0010 secs, 0.0000 secs, 0.0181 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.0590 secs
  Slowest:	0.3778 secs
  Fastest:	0.0002 secs
  Average:	0.0051 secs
  Requests/sec:	9442.7818
  
  Total data:	3187963 bytes
  Size/request:	318 bytes

Response time histogram:
  0.000 [1]	|
  0.038 [9804]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.076 [0]	|
  0.114 [24]	|
  0.151 [168]	|■
  0.189 [0]	|
  0.227 [0]	|
  0.265 [2]	|
  0.302 [0]	|
  0.340 [0]	|
  0.378 [1]	|


Latency distribution:
  10% in 0.0020 secs
  25% in 0.0023 secs
  50% in 0.0026 secs
  75% in 0.0030 secs
  90% in 0.0041 secs
  95% in 0.0055 secs
  99% in 0.1277 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.3778 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:	0.0051 secs, 0.0002 secs, 0.3778 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.3088 secs
  Slowest:	0.2647 secs
  Fastest:	0.0001 secs
  Average:	0.0012 secs
  Requests/sec:	32382.2749
  
  Total data:	13080000 bytes
  Size/request:	1308 bytes

Response time histogram:
  0.000 [1]	|
  0.027 [9951]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.053 [18]	|
  0.079 [0]	|
  0.106 [0]	|
  0.132 [5]	|
  0.159 [22]	|
  0.185 [1]	|
  0.212 [0]	|
  0.238 [0]	|
  0.265 [2]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0005 secs
  50% in 0.0007 secs
  75% in 0.0009 secs
  90% in 0.0010 secs
  95% in 0.0011 secs
  99% in 0.0074 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.2647 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0011 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0006 secs
  resp wait:	0.0012 secs, 0.0001 secs, 0.2637 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	2.9874 secs
  Slowest:	0.0442 secs
  Fastest:	0.0015 secs
  Average:	0.0148 secs
  Requests/sec:	3347.4346
  
  Total data:	3187951 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.006 [21]	|
  0.010 [55]	|
  0.014 [3364]	|■■■■■■■■■■■■■■■■■■■■■■
  0.019 [6247]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.023 [234]	|■
  0.027 [70]	|
  0.031 [5]	|
  0.036 [2]	|
  0.040 [0]	|
  0.044 [1]	|


Latency distribution:
  10% in 0.0129 secs
  25% in 0.0139 secs
  50% in 0.0148 secs
  75% in 0.0155 secs
  90% in 0.0164 secs
  95% in 0.0174 secs
  99% in 0.0224 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0015 secs, 0.0442 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0011 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0148 secs, 0.0015 secs, 0.0442 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0545 secs
  Slowest:	0.0037 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	183547.5603
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [8404]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1067]	|■■■■■
  0.001 [393]	|■■
  0.002 [44]	|
  0.002 [66]	|
  0.002 [8]	|
  0.003 [7]	|
  0.003 [0]	|
  0.003 [5]	|
  0.004 [5]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0002 secs
  75% in 0.0003 secs
  90% in 0.0006 secs
  95% in 0.0008 secs
  99% in 0.0014 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0037 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0002 secs, 0.0000 secs, 0.0020 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0019 secs

Status code distribution:
  [200]	10000 responses



