##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0932 secs
  Slowest:	0.0076 secs
  Fastest:	0.0000 secs
  Average:	0.0005 secs
  Requests/sec:	107275.1784
  
  Total data:	280000 bytes
  Size/request:	28 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8437]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1275]	|■■■■■■
  0.002 [228]	|■
  0.003 [9]	|
  0.004 [0]	|
  0.005 [0]	|
  0.005 [8]	|
  0.006 [27]	|
  0.007 [12]	|
  0.008 [3]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0006 secs
  90% in 0.0010 secs
  95% in 0.0013 secs
  99% in 0.0018 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0076 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0020 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0024 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0024 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.2135 secs
  Slowest:	0.0196 secs
  Fastest:	0.0001 secs
  Average:	0.0010 secs
  Requests/sec:	46831.2512
  
  Total data:	317786 bytes
  Size/request:	31 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9658]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [300]	|■
  0.006 [2]	|
  0.008 [2]	|
  0.010 [13]	|
  0.012 [7]	|
  0.014 [3]	|
  0.016 [3]	|
  0.018 [5]	|
  0.020 [6]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0003 secs
  50% in 0.0013 secs
  75% in 0.0015 secs
  90% in 0.0016 secs
  95% in 0.0018 secs
  99% in 0.0025 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0196 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0006 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0005 secs
  resp wait:	0.0010 secs, 0.0000 secs, 0.0181 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.1472 secs
  Slowest:	0.5023 secs
  Fastest:	0.0003 secs
  Average:	0.0054 secs
  Requests/sec:	8717.2169
  
  Total data:	3197587 bytes
  Size/request:	319 bytes

Response time histogram:
  0.000 [1]	|
  0.050 [9937]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.101 [10]	|
  0.151 [45]	|
  0.201 [3]	|
  0.251 [0]	|
  0.301 [1]	|
  0.352 [0]	|
  0.402 [2]	|
  0.452 [0]	|
  0.502 [1]	|


Latency distribution:
  10% in 0.0009 secs
  25% in 0.0026 secs
  50% in 0.0046 secs
  75% in 0.0059 secs
  90% in 0.0077 secs
  95% in 0.0093 secs
  99% in 0.0380 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.5023 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:	0.0054 secs, 0.0003 secs, 0.5023 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.3101 secs
  Slowest:	0.1965 secs
  Fastest:	0.0001 secs
  Average:	0.0013 secs
  Requests/sec:	32250.7691
  
  Total data:	19750000 bytes
  Size/request:	1975 bytes

Response time histogram:
  0.000 [1]	|
  0.020 [9912]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.039 [68]	|
  0.059 [17]	|
  0.079 [0]	|
  0.098 [0]	|
  0.118 [0]	|
  0.138 [0]	|
  0.157 [0]	|
  0.177 [1]	|
  0.197 [1]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0004 secs
  75% in 0.0017 secs
  90% in 0.0020 secs
  95% in 0.0031 secs
  99% in 0.0177 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.1965 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0006 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0013 secs, 0.0001 secs, 0.1965 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	3.1024 secs
  Slowest:	0.0517 secs
  Fastest:	0.0017 secs
  Average:	0.0154 secs
  Requests/sec:	3223.2687
  
  Total data:	3198203 bytes
  Size/request:	319 bytes

Response time histogram:
  0.002 [1]	|
  0.007 [27]	|
  0.012 [322]	|■■
  0.017 [8323]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.022 [1038]	|■■■■■
  0.027 [189]	|■
  0.032 [43]	|
  0.037 [6]	|
  0.042 [9]	|
  0.047 [28]	|
  0.052 [14]	|


Latency distribution:
  10% in 0.0133 secs
  25% in 0.0142 secs
  50% in 0.0151 secs
  75% in 0.0160 secs
  90% in 0.0171 secs
  95% in 0.0186 secs
  99% in 0.0268 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0017 secs, 0.0517 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0154 secs, 0.0017 secs, 0.0517 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0688 secs
  Slowest:	0.0021 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	145368.0287
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [4342]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.000 [3313]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1162]	|■■■■■■■■■■■
  0.001 [663]	|■■■■■■
  0.001 [325]	|■■■
  0.001 [96]	|■
  0.001 [49]	|
  0.002 [23]	|
  0.002 [14]	|
  0.002 [12]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0004 secs
  90% in 0.0007 secs
  95% in 0.0009 secs
  99% in 0.0013 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0021 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0014 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0015 secs

Status code distribution:
  [200]	10000 responses



