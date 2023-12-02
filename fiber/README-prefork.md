##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.1724 secs
  Slowest:	0.0115 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	58009.8979
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8306]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1601]	|■■■■■■■■
  0.003 [40]	|
  0.005 [2]	|
  0.006 [0]	|
  0.007 [3]	|
  0.008 [14]	|
  0.009 [18]	|
  0.010 [7]	|
  0.011 [8]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0004 secs
  50% in 0.0007 secs
  75% in 0.0010 secs
  90% in 0.0014 secs
  95% in 0.0017 secs
  99% in 0.0023 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0115 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0021 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:	0.0007 secs, 0.0000 secs, 0.0024 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0032 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.8526 secs
  Slowest:	0.0359 secs
  Fastest:	0.0003 secs
  Average:	0.0042 secs
  Requests/sec:	11729.0383
  
  Total data:	307792 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.004 [7864]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.007 [1241]	|■■■■■■
  0.011 [26]	|
  0.015 [27]	|
  0.018 [212]	|■
  0.022 [497]	|■■■
  0.025 [83]	|
  0.029 [3]	|
  0.032 [9]	|
  0.036 [37]	|


Latency distribution:
  10% in 0.0013 secs
  25% in 0.0017 secs
  50% in 0.0027 secs
  75% in 0.0036 secs
  90% in 0.0061 secs
  95% in 0.0187 secs
  99% in 0.0223 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.0359 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0013 secs
  resp wait:	0.0041 secs, 0.0002 secs, 0.0314 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0013 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	3.8989 secs
  Slowest:	0.1071 secs
  Fastest:	0.0016 secs
  Average:	0.0192 secs
  Requests/sec:	2564.8302
  
  Total data:	3187665 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.012 [3003]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.023 [3223]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.033 [3080]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.044 [633]	|■■■■■■■■
  0.054 [10]	|
  0.065 [0]	|
  0.075 [0]	|
  0.086 [1]	|
  0.097 [20]	|
  0.107 [29]	|


Latency distribution:
  10% in 0.0072 secs
  25% in 0.0108 secs
  50% in 0.0153 secs
  75% in 0.0292 secs
  90% in 0.0324 secs
  95% in 0.0338 secs
  99% in 0.0376 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0016 secs, 0.1071 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0015 secs
  resp wait:	0.0191 secs, 0.0016 secs, 0.1070 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0012 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	1.5576 secs
  Slowest:	0.0312 secs
  Fastest:	0.0003 secs
  Average:	0.0076 secs
  Requests/sec:	6420.1870
  
  Total data:	13080000 bytes
  Size/request:	1308 bytes

Response time histogram:
  0.000 [1]	|
  0.003 [2150]	|■■■■■■■■■■■■■■■■■■■■■
  0.006 [4108]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [1966]	|■■■■■■■■■■■■■■■■■■■
  0.013 [222]	|■■
  0.016 [58]	|■
  0.019 [295]	|■■■
  0.022 [722]	|■■■■■■■
  0.025 [391]	|■■■■
  0.028 [77]	|■
  0.031 [10]	|


Latency distribution:
  10% in 0.0028 secs
  25% in 0.0036 secs
  50% in 0.0058 secs
  75% in 0.0076 secs
  90% in 0.0199 secs
  95% in 0.0219 secs
  99% in 0.0248 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.0312 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:	0.0076 secs, 0.0003 secs, 0.0312 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0015 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	9.9782 secs
  Slowest:	0.1833 secs
  Fastest:	0.0070 secs
  Average:	0.0497 secs
  Requests/sec:	1002.1848
  
  Total data:	3187875 bytes
  Size/request:	318 bytes

Response time histogram:
  0.007 [1]	|
  0.025 [70]	|
  0.042 [1099]	|■■■■■
  0.060 [8010]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.077 [453]	|■■
  0.095 [117]	|■
  0.113 [144]	|■
  0.130 [7]	|
  0.148 [3]	|
  0.166 [56]	|
  0.183 [40]	|


Latency distribution:
  10% in 0.0415 secs
  25% in 0.0447 secs
  50% in 0.0476 secs
  75% in 0.0513 secs
  90% in 0.0576 secs
  95% in 0.0715 secs
  99% in 0.1265 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0070 secs, 0.1833 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:	0.0496 secs, 0.0069 secs, 0.1832 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1892 secs
  Slowest:	0.0106 secs
  Fastest:	0.0001 secs
  Average:	0.0009 secs
  Requests/sec:	52864.3650
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [7428]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [2255]	|■■■■■■■■■■■■
  0.003 [197]	|■
  0.004 [31]	|
  0.005 [20]	|
  0.006 [54]	|
  0.007 [9]	|
  0.009 [1]	|
  0.010 [1]	|
  0.011 [3]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0005 secs
  50% in 0.0008 secs
  75% in 0.0011 secs
  90% in 0.0015 secs
  95% in 0.0019 secs
  99% in 0.0037 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0106 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0044 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0043 secs
  resp wait:	0.0007 secs, 0.0000 secs, 0.0039 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0046 secs

Status code distribution:
  [200]	10000 responses



