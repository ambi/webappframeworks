##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.2393 secs
  Slowest:	0.0687 secs
  Fastest:	0.0000 secs
  Average:	0.0012 secs
  Requests/sec:	41785.1231
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.007 [9951]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.014 [5]	|
  0.021 [8]	|
  0.027 [5]	|
  0.034 [5]	|
  0.041 [5]	|
  0.048 [5]	|
  0.055 [6]	|
  0.062 [4]	|
  0.069 [5]	|


Latency distribution:
  10% in 0.0005 secs
  25% in 0.0007 secs
  50% in 0.0009 secs
  75% in 0.0013 secs
  90% in 0.0017 secs
  95% in 0.0022 secs
  99% in 0.0031 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0687 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0011 secs, 0.0000 secs, 0.0635 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0022 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.3976 secs
  Slowest:	0.0368 secs
  Fastest:	0.0001 secs
  Average:	0.0020 secs
  Requests/sec:	25148.4149
  
  Total data:	307883 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.004 [9745]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.007 [197]	|■
  0.011 [7]	|
  0.015 [0]	|
  0.018 [8]	|
  0.022 [10]	|
  0.026 [8]	|
  0.029 [6]	|
  0.033 [3]	|
  0.037 [15]	|


Latency distribution:
  10% in 0.0015 secs
  25% in 0.0016 secs
  50% in 0.0017 secs
  75% in 0.0020 secs
  90% in 0.0024 secs
  95% in 0.0030 secs
  99% in 0.0059 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0368 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0010 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:	0.0020 secs, 0.0001 secs, 0.0350 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.5826 secs
  Slowest:	0.0377 secs
  Fastest:	0.0023 secs
  Average:	0.0079 secs
  Requests/sec:	6318.8603
  
  Total data:	3188274 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.006 [22]	|
  0.009 [9867]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.013 [46]	|
  0.016 [14]	|
  0.020 [0]	|
  0.024 [0]	|
  0.027 [0]	|
  0.031 [0]	|
  0.034 [14]	|
  0.038 [36]	|


Latency distribution:
  10% in 0.0074 secs
  25% in 0.0076 secs
  50% in 0.0077 secs
  75% in 0.0079 secs
  90% in 0.0081 secs
  95% in 0.0083 secs
  99% in 0.0095 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0023 secs, 0.0377 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0011 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:	0.0079 secs, 0.0023 secs, 0.0377 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.5237 secs
  Slowest:	0.0173 secs
  Fastest:	0.0003 secs
  Average:	0.0026 secs
  Requests/sec:	19095.5625
  
  Total data:	12440000 bytes
  Size/request:	1244 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [192]	|■
  0.004 [9756]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [13]	|
  0.007 [3]	|
  0.009 [3]	|
  0.010 [1]	|
  0.012 [12]	|
  0.014 [14]	|
  0.016 [2]	|
  0.017 [3]	|


Latency distribution:
  10% in 0.0024 secs
  25% in 0.0025 secs
  50% in 0.0025 secs
  75% in 0.0027 secs
  90% in 0.0028 secs
  95% in 0.0030 secs
  99% in 0.0035 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.0173 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0005 secs
  resp wait:	0.0026 secs, 0.0002 secs, 0.0152 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0020 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	3.6762 secs
  Slowest:	0.0317 secs
  Fastest:	0.0156 secs
  Average:	0.0184 secs
  Requests/sec:	2720.1693
  
  Total data:	3187921 bytes
  Size/request:	318 bytes

Response time histogram:
  0.016 [1]	|
  0.017 [883]	|■■■■■
  0.019 [7744]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.020 [807]	|■■■■
  0.022 [119]	|■
  0.024 [61]	|
  0.025 [169]	|■
  0.027 [56]	|
  0.028 [12]	|
  0.030 [119]	|■
  0.032 [29]	|


Latency distribution:
  10% in 0.0172 secs
  25% in 0.0176 secs
  50% in 0.0180 secs
  75% in 0.0184 secs
  90% in 0.0191 secs
  95% in 0.0211 secs
  99% in 0.0291 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0156 secs, 0.0317 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0184 secs, 0.0156 secs, 0.0305 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1232 secs
  Slowest:	0.0150 secs
  Fastest:	0.0000 secs
  Average:	0.0006 secs
  Requests/sec:	81179.7581
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9925]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [38]	|
  0.005 [6]	|
  0.006 [5]	|
  0.007 [5]	|
  0.009 [9]	|
  0.010 [3]	|
  0.012 [2]	|
  0.013 [4]	|
  0.015 [2]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0005 secs
  50% in 0.0006 secs
  75% in 0.0006 secs
  90% in 0.0009 secs
  95% in 0.0011 secs
  99% in 0.0014 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0150 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0010 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:	0.0005 secs, 0.0000 secs, 0.0134 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



