##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0887 secs
  Slowest:	0.0036 secs
  Fastest:	0.0000 secs
  Average:	0.0004 secs
  Requests/sec:	112754.5602
  
  Total data:	280000 bytes
  Size/request:	28 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [6075]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [2224]	|■■■■■■■■■■■■■■■
  0.001 [1078]	|■■■■■■■
  0.001 [397]	|■■■
  0.002 [141]	|■
  0.002 [24]	|
  0.003 [8]	|
  0.003 [9]	|
  0.003 [30]	|
  0.004 [13]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0006 secs
  90% in 0.0010 secs
  95% in 0.0012 secs
  99% in 0.0017 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0036 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0010 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0028 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0021 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0024 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.2290 secs
  Slowest:	0.0216 secs
  Fastest:	0.0001 secs
  Average:	0.0011 secs
  Requests/sec:	43664.8015
  
  Total data:	317882 bytes
  Size/request:	31 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9905]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [49]	|
  0.007 [10]	|
  0.009 [2]	|
  0.011 [0]	|
  0.013 [0]	|
  0.015 [14]	|
  0.017 [4]	|
  0.019 [9]	|
  0.022 [6]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0004 secs
  50% in 0.0013 secs
  75% in 0.0015 secs
  90% in 0.0016 secs
  95% in 0.0017 secs
  99% in 0.0022 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0216 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:	0.0011 secs, 0.0000 secs, 0.0205 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.1083 secs
  Slowest:	0.5004 secs
  Fastest:	0.0003 secs
  Average:	0.0049 secs
  Requests/sec:	9022.6259
  
  Total data:	3197795 bytes
  Size/request:	319 bytes

Response time histogram:
  0.000 [1]	|
  0.050 [9896]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.100 [41]	|
  0.150 [49]	|
  0.200 [1]	|
  0.250 [0]	|
  0.300 [3]	|
  0.350 [0]	|
  0.400 [7]	|
  0.450 [0]	|
  0.500 [2]	|


Latency distribution:
  10% in 0.0004 secs
  25% in 0.0009 secs
  50% in 0.0037 secs
  75% in 0.0049 secs
  90% in 0.0059 secs
  95% in 0.0080 secs
  99% in 0.0678 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.5004 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0010 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0049 secs, 0.0002 secs, 0.5004 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.3013 secs
  Slowest:	0.2730 secs
  Fastest:	0.0001 secs
  Average:	0.0010 secs
  Requests/sec:	33184.2255
  
  Total data:	19750000 bytes
  Size/request:	1975 bytes

Response time histogram:
  0.000 [1]	|
  0.027 [9956]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.055 [4]	|
  0.082 [20]	|
  0.109 [0]	|
  0.137 [3]	|
  0.164 [11]	|
  0.191 [0]	|
  0.218 [0]	|
  0.246 [1]	|
  0.273 [4]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0005 secs
  75% in 0.0009 secs
  90% in 0.0011 secs
  95% in 0.0013 secs
  99% in 0.0020 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.2730 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0010 secs, 0.0001 secs, 0.2722 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	3.1324 secs
  Slowest:	0.0434 secs
  Fastest:	0.0018 secs
  Average:	0.0155 secs
  Requests/sec:	3192.4489
  
  Total data:	3197632 bytes
  Size/request:	319 bytes

Response time histogram:
  0.002 [1]	|
  0.006 [24]	|
  0.010 [76]	|
  0.014 [2379]	|■■■■■■■■■■■■■■
  0.018 [6839]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.023 [429]	|■■■
  0.027 [185]	|■
  0.031 [19]	|
  0.035 [10]	|
  0.039 [21]	|
  0.043 [17]	|


Latency distribution:
  10% in 0.0131 secs
  25% in 0.0143 secs
  50% in 0.0153 secs
  75% in 0.0162 secs
  90% in 0.0177 secs
  95% in 0.0193 secs
  99% in 0.0252 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0018 secs, 0.0434 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:	0.0155 secs, 0.0018 secs, 0.0434 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0689 secs
  Slowest:	0.0034 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	145190.5626
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [6882]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [2001]	|■■■■■■■■■■■■
  0.001 [814]	|■■■■■
  0.001 [219]	|■
  0.002 [48]	|
  0.002 [23]	|
  0.002 [2]	|
  0.003 [2]	|
  0.003 [3]	|
  0.003 [5]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0002 secs
  75% in 0.0004 secs
  90% in 0.0007 secs
  95% in 0.0009 secs
  99% in 0.0013 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0034 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0015 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0014 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0019 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0015 secs

Status code distribution:
  [200]	10000 responses



