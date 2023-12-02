##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.1571 secs
  Slowest:	0.0096 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	63660.9919
  
  Total data:	280000 bytes
  Size/request:	28 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [7956]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1421]	|■■■■■■■
  0.003 [445]	|■■
  0.004 [78]	|
  0.005 [53]	|
  0.006 [15]	|
  0.007 [24]	|
  0.008 [2]	|
  0.009 [3]	|
  0.010 [2]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0005 secs
  75% in 0.0009 secs
  90% in 0.0017 secs
  95% in 0.0021 secs
  99% in 0.0039 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0096 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0036 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0027 secs
  resp wait:	0.0005 secs, 0.0001 secs, 0.0043 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0044 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.8115 secs
  Slowest:	0.1699 secs
  Fastest:	0.0003 secs
  Average:	0.0037 secs
  Requests/sec:	12322.7934
  
  Total data:	317867 bytes
  Size/request:	31 bytes

Response time histogram:
  0.000 [1]	|
  0.017 [9311]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.034 [611]	|■■■
  0.051 [40]	|
  0.068 [16]	|
  0.085 [7]	|
  0.102 [5]	|
  0.119 [3]	|
  0.136 [2]	|
  0.153 [0]	|
  0.170 [4]	|


Latency distribution:
  10% in 0.0008 secs
  25% in 0.0016 secs
  50% in 0.0021 secs
  75% in 0.0028 secs
  90% in 0.0039 secs
  95% in 0.0193 secs
  99% in 0.0239 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.1699 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0036 secs, 0.0002 secs, 0.1674 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0011 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	3.2132 secs
  Slowest:	0.2772 secs
  Fastest:	0.0012 secs
  Average:	0.0156 secs
  Requests/sec:	3112.1298
  
  Total data:	3197833 bytes
  Size/request:	319 bytes

Response time histogram:
  0.001 [1]	|
  0.029 [9056]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.056 [812]	|■■■■
  0.084 [0]	|
  0.112 [2]	|
  0.139 [78]	|
  0.167 [36]	|
  0.194 [7]	|
  0.222 [0]	|
  0.250 [0]	|
  0.277 [8]	|


Latency distribution:
  10% in 0.0048 secs
  25% in 0.0082 secs
  50% in 0.0101 secs
  75% in 0.0250 secs
  90% in 0.0287 secs
  95% in 0.0302 secs
  99% in 0.1316 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0012 secs, 0.2772 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0015 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:	0.0155 secs, 0.0011 secs, 0.2772 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	1.2088 secs
  Slowest:	0.0523 secs
  Fastest:	0.0003 secs
  Average:	0.0058 secs
  Requests/sec:	8272.5450
  
  Total data:	19750000 bytes
  Size/request:	1975 bytes

Response time histogram:
  0.000 [1]	|
  0.006 [7075]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.011 [1502]	|■■■■■■■■
  0.016 [229]	|■
  0.021 [583]	|■■■
  0.026 [437]	|■■
  0.031 [114]	|■
  0.037 [34]	|
  0.042 [17]	|
  0.047 [4]	|
  0.052 [4]	|


Latency distribution:
  10% in 0.0012 secs
  25% in 0.0017 secs
  50% in 0.0032 secs
  75% in 0.0062 secs
  90% in 0.0182 secs
  95% in 0.0221 secs
  99% in 0.0287 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.0523 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0057 secs, 0.0003 secs, 0.0522 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0013 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	11.4410 secs
  Slowest:	0.1368 secs
  Fastest:	0.0082 secs
  Average:	0.0570 secs
  Requests/sec:	874.0487
  
  Total data:	3197846 bytes
  Size/request:	319 bytes

Response time histogram:
  0.008 [1]	|
  0.021 [10]	|
  0.034 [90]	|
  0.047 [259]	|■
  0.060 [7653]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.072 [920]	|■■■■■
  0.085 [901]	|■■■■■
  0.098 [102]	|■
  0.111 [40]	|
  0.124 [15]	|
  0.137 [9]	|


Latency distribution:
  10% in 0.0494 secs
  25% in 0.0519 secs
  50% in 0.0544 secs
  75% in 0.0579 secs
  90% in 0.0729 secs
  95% in 0.0765 secs
  99% in 0.0943 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0082 secs, 0.1368 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0019 secs
  resp wait:	0.0569 secs, 0.0081 secs, 0.1367 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1750 secs
  Slowest:	0.0124 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	57146.7797
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8358]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [1003]	|■■■■■
  0.004 [341]	|■■
  0.005 [151]	|■
  0.006 [94]	|
  0.007 [24]	|
  0.009 [12]	|
  0.010 [5]	|
  0.011 [8]	|
  0.012 [3]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0002 secs
  50% in 0.0004 secs
  75% in 0.0009 secs
  90% in 0.0019 secs
  95% in 0.0029 secs
  99% in 0.0054 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0124 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0024 secs
  resp wait:	0.0007 secs, 0.0001 secs, 0.0118 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0024 secs

Status code distribution:
  [200]	10000 responses



