##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0881 secs
  Slowest:	0.0049 secs
  Fastest:	0.0000 secs
  Average:	0.0004 secs
  Requests/sec:	113539.7581
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [7132]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [2138]	|■■■■■■■■■■■■
  0.001 [596]	|■■■
  0.002 [71]	|
  0.002 [12]	|
  0.003 [0]	|
  0.003 [0]	|
  0.004 [7]	|
  0.004 [22]	|
  0.005 [21]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0006 secs
  90% in 0.0009 secs
  95% in 0.0011 secs
  99% in 0.0016 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0049 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0011 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0013 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0023 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0014 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.2408 secs
  Slowest:	0.0150 secs
  Fastest:	0.0001 secs
  Average:	0.0012 secs
  Requests/sec:	41521.0475
  
  Total data:	307719 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [8799]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [1126]	|■■■■■
  0.005 [24]	|
  0.006 [0]	|
  0.008 [0]	|
  0.009 [0]	|
  0.011 [0]	|
  0.012 [0]	|
  0.014 [13]	|
  0.015 [37]	|


Latency distribution:
  10% in 0.0009 secs
  25% in 0.0009 secs
  50% in 0.0010 secs
  75% in 0.0013 secs
  90% in 0.0016 secs
  95% in 0.0019 secs
  99% in 0.0026 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0150 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0013 secs
  resp wait:	0.0012 secs, 0.0001 secs, 0.0138 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0014 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.0356 secs
  Slowest:	0.0150 secs
  Fastest:	0.0014 secs
  Average:	0.0052 secs
  Requests/sec:	9656.4291
  
  Total data:	3187895 bytes
  Size/request:	318 bytes

Response time histogram:
  0.001 [1]	|
  0.003 [5]	|
  0.004 [5]	|
  0.005 [7228]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.007 [2389]	|■■■■■■■■■■■■■
  0.008 [254]	|■
  0.010 [68]	|
  0.011 [11]	|
  0.012 [18]	|
  0.014 [2]	|
  0.015 [19]	|


Latency distribution:
  10% in 0.0045 secs
  25% in 0.0046 secs
  50% in 0.0049 secs
  75% in 0.0056 secs
  90% in 0.0060 secs
  95% in 0.0065 secs
  99% in 0.0089 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0014 secs, 0.0150 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:	0.0052 secs, 0.0014 secs, 0.0150 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.2586 secs
  Slowest:	0.0086 secs
  Fastest:	0.0001 secs
  Average:	0.0013 secs
  Requests/sec:	38668.9254
  
  Total data:	16050000 bytes
  Size/request:	1605 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [97]	|
  0.002 [8939]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [879]	|■■■■
  0.003 [34]	|
  0.004 [0]	|
  0.005 [0]	|
  0.006 [0]	|
  0.007 [9]	|
  0.008 [17]	|
  0.009 [24]	|


Latency distribution:
  10% in 0.0010 secs
  25% in 0.0011 secs
  50% in 0.0011 secs
  75% in 0.0013 secs
  90% in 0.0017 secs
  95% in 0.0021 secs
  99% in 0.0026 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0086 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0005 secs
  resp wait:	0.0013 secs, 0.0001 secs, 0.0072 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	3.5858 secs
  Slowest:	0.0499 secs
  Fastest:	0.0133 secs
  Average:	0.0179 secs
  Requests/sec:	2788.7763
  
  Total data:	3187977 bytes
  Size/request:	318 bytes

Response time histogram:
  0.013 [1]	|
  0.017 [1841]	|■■■■■■■■■■
  0.021 [7731]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.024 [231]	|■
  0.028 [146]	|■
  0.032 [0]	|
  0.035 [0]	|
  0.039 [0]	|
  0.043 [0]	|
  0.046 [10]	|
  0.050 [40]	|


Latency distribution:
  10% in 0.0166 secs
  25% in 0.0171 secs
  50% in 0.0175 secs
  75% in 0.0181 secs
  90% in 0.0188 secs
  95% in 0.0198 secs
  99% in 0.0262 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0133 secs, 0.0499 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:	0.0179 secs, 0.0133 secs, 0.0499 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0638 secs
  Slowest:	0.0038 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	156778.7191
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [8027]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1371]	|■■■■■■■
  0.001 [453]	|■■
  0.002 [76]	|
  0.002 [20]	|
  0.002 [1]	|
  0.003 [30]	|
  0.003 [20]	|
  0.003 [0]	|
  0.004 [1]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0002 secs
  75% in 0.0003 secs
  90% in 0.0006 secs
  95% in 0.0008 secs
  99% in 0.0014 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0038 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0006 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0015 secs
  resp wait:	0.0002 secs, 0.0000 secs, 0.0016 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0014 secs

Status code distribution:
  [200]	10000 responses



