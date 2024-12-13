##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0773 secs
  Slowest:	0.0063 secs
  Fastest:	0.0000 secs
  Average:	0.0004 secs
  Requests/sec:	129283.3881
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8324]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1339]	|■■■■■■
  0.002 [262]	|■
  0.003 [24]	|
  0.003 [0]	|
  0.004 [0]	|
  0.004 [0]	|
  0.005 [0]	|
  0.006 [17]	|
  0.006 [33]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0002 secs
  75% in 0.0005 secs
  90% in 0.0009 secs
  95% in 0.0012 secs
  99% in 0.0016 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0063 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0016 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0015 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.2001 secs
  Slowest:	0.0214 secs
  Fastest:	0.0000 secs
  Average:	0.0009 secs
  Requests/sec:	49966.1791
  
  Total data:	307831 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9948]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [11]	|
  0.006 [1]	|
  0.009 [1]	|
  0.011 [3]	|
  0.013 [7]	|
  0.015 [12]	|
  0.017 [5]	|
  0.019 [5]	|
  0.021 [6]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0005 secs
  50% in 0.0011 secs
  75% in 0.0012 secs
  90% in 0.0014 secs
  95% in 0.0015 secs
  99% in 0.0017 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0214 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:	0.0009 secs, 0.0000 secs, 0.0202 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.0697 secs
  Slowest:	0.6286 secs
  Fastest:	0.0003 secs
  Average:	0.0048 secs
  Requests/sec:	9348.4351
  
  Total data:	3188215 bytes
  Size/request:	318 bytes

Response time histogram:
  0.000 [1]	|
  0.063 [9967]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.126 [22]	|
  0.189 [1]	|
  0.252 [0]	|
  0.314 [3]	|
  0.377 [2]	|
  0.440 [0]	|
  0.503 [3]	|
  0.566 [0]	|
  0.629 [1]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0005 secs
  50% in 0.0040 secs
  75% in 0.0053 secs
  90% in 0.0076 secs
  95% in 0.0117 secs
  99% in 0.0325 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.6286 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0048 secs, 0.0002 secs, 0.6286 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.2547 secs
  Slowest:	0.1816 secs
  Fastest:	0.0001 secs
  Average:	0.0009 secs
  Requests/sec:	39268.1143
  
  Total data:	19750000 bytes
  Size/request:	1975 bytes

Response time histogram:
  0.000 [1]	|
  0.018 [9968]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.036 [0]	|
  0.055 [0]	|
  0.073 [15]	|
  0.091 [0]	|
  0.109 [0]	|
  0.127 [0]	|
  0.145 [0]	|
  0.163 [0]	|
  0.182 [16]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0004 secs
  75% in 0.0005 secs
  90% in 0.0007 secs
  95% in 0.0012 secs
  99% in 0.0057 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.1816 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0009 secs, 0.0001 secs, 0.1804 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	3.0038 secs
  Slowest:	0.0409 secs
  Fastest:	0.0024 secs
  Average:	0.0149 secs
  Requests/sec:	3329.1548
  
  Total data:	3187655 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.006 [33]	|
  0.010 [165]	|■
  0.014 [2715]	|■■■■■■■■■■■■■■■■■
  0.018 [6459]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.022 [371]	|■■
  0.025 [180]	|■
  0.029 [56]	|
  0.033 [7]	|
  0.037 [10]	|
  0.041 [3]	|


Latency distribution:
  10% in 0.0127 secs
  25% in 0.0137 secs
  50% in 0.0148 secs
  75% in 0.0155 secs
  90% in 0.0168 secs
  95% in 0.0189 secs
  99% in 0.0247 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0024 secs, 0.0409 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:	0.0149 secs, 0.0024 secs, 0.0409 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0553 secs
  Slowest:	0.0036 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	180764.2735
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [8127]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1372]	|■■■■■■■
  0.001 [346]	|■■
  0.001 [58]	|
  0.002 [52]	|
  0.002 [35]	|
  0.003 [8]	|
  0.003 [0]	|
  0.003 [0]	|
  0.004 [1]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0002 secs
  75% in 0.0003 secs
  90% in 0.0006 secs
  95% in 0.0007 secs
  99% in 0.0014 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0036 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0011 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0002 secs, 0.0000 secs, 0.0015 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0018 secs

Status code distribution:
  [200]	10000 responses



