##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.2636 secs
  Slowest:	0.0085 secs
  Fastest:	0.0001 secs
  Average:	0.0013 secs
  Requests/sec:	37939.3722
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [2494]	|■■■■■■■■■■■■■■■■
  0.002 [6102]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [1195]	|■■■■■■■■
  0.003 [148]	|■
  0.004 [10]	|
  0.005 [19]	|
  0.006 [11]	|
  0.007 [7]	|
  0.008 [8]	|
  0.008 [5]	|


Latency distribution:
  10% in 0.0008 secs
  25% in 0.0009 secs
  50% in 0.0012 secs
  75% in 0.0015 secs
  90% in 0.0019 secs
  95% in 0.0022 secs
  99% in 0.0029 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0085 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0019 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0015 secs
  resp wait:	0.0012 secs, 0.0001 secs, 0.0029 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0017 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	1.7378 secs
  Slowest:	0.0444 secs
  Fastest:	0.0005 secs
  Average:	0.0087 secs
  Requests/sec:	5754.3666
  
  Total data:	307781 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.005 [3384]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.009 [4591]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.014 [195]	|■■
  0.018 [91]	|■
  0.022 [1202]	|■■■■■■■■■■
  0.027 [420]	|■■■■
  0.031 [47]	|
  0.036 [22]	|
  0.040 [22]	|
  0.044 [25]	|


Latency distribution:
  10% in 0.0041 secs
  25% in 0.0045 secs
  50% in 0.0058 secs
  75% in 0.0081 secs
  90% in 0.0211 secs
  95% in 0.0227 secs
  99% in 0.0291 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0005 secs, 0.0444 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0013 secs
  resp wait:	0.0086 secs, 0.0004 secs, 0.0404 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	11.0895 secs
  Slowest:	0.1882 secs
  Fastest:	0.0296 secs
  Average:	0.0554 secs
  Requests/sec:	901.7505
  

Response time histogram:
  0.030 [1]	|
  0.045 [90]	|
  0.061 [8693]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.077 [794]	|■■■■
  0.093 [209]	|■
  0.109 [63]	|
  0.125 [38]	|
  0.141 [21]	|
  0.156 [41]	|
  0.172 [27]	|
  0.188 [23]	|


Latency distribution:
  10% in 0.0483 secs
  25% in 0.0497 secs
  50% in 0.0519 secs
  75% in 0.0547 secs
  90% in 0.0683 secs
  95% in 0.0739 secs
  99% in 0.1273 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0296 secs, 0.1882 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0020 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0032 secs
  resp wait:	0.0553 secs, 0.0296 secs, 0.1881 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0014 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	2.8161 secs
  Slowest:	0.1536 secs
  Fastest:	0.0017 secs
  Average:	0.0140 secs
  Requests/sec:	3551.0423
  

Response time histogram:
  0.002 [1]	|
  0.017 [7164]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.032 [2649]	|■■■■■■■■■■■■■■■
  0.047 [110]	|■
  0.062 [16]	|
  0.078 [10]	|
  0.093 [0]	|
  0.108 [6]	|
  0.123 [10]	|
  0.138 [10]	|
  0.154 [24]	|


Latency distribution:
  10% in 0.0073 secs
  25% in 0.0079 secs
  50% in 0.0095 secs
  75% in 0.0219 secs
  90% in 0.0245 secs
  95% in 0.0266 secs
  99% in 0.0397 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0017 secs, 0.1536 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0001 secs, 0.0000 secs, 0.0021 secs
  resp wait:	0.0136 secs, 0.0016 secs, 0.1534 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0033 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	27.4593 secs
  Slowest:	0.2959 secs
  Fastest:	0.0976 secs
  Average:	0.1372 secs
  Requests/sec:	364.1754
  

Response time histogram:
  0.098 [1]	|
  0.117 [2728]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.137 [3137]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.157 [2871]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.177 [569]	|■■■■■■■
  0.197 [131]	|■■
  0.217 [203]	|■■■
  0.236 [121]	|■■
  0.256 [123]	|■■
  0.276 [63]	|■
  0.296 [53]	|■


Latency distribution:
  10% in 0.1100 secs
  25% in 0.1164 secs
  50% in 0.1328 secs
  75% in 0.1459 secs
  90% in 0.1623 secs
  95% in 0.2050 secs
  99% in 0.2631 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0976 secs, 0.2959 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:	0.1371 secs, 0.0975 secs, 0.2958 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0013 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.2326 secs
  Slowest:	0.0057 secs
  Fastest:	0.0001 secs
  Average:	0.0011 secs
  Requests/sec:	42988.7809
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [901]	|■■■■■■
  0.001 [6479]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1797]	|■■■■■■■■■■■
  0.002 [575]	|■■■■
  0.003 [161]	|■
  0.003 [34]	|
  0.004 [4]	|
  0.005 [11]	|
  0.005 [30]	|
  0.006 [7]	|


Latency distribution:
  10% in 0.0007 secs
  25% in 0.0009 secs
  50% in 0.0010 secs
  75% in 0.0013 secs
  90% in 0.0017 secs
  95% in 0.0021 secs
  99% in 0.0028 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0057 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0015 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0019 secs
  resp wait:	0.0010 secs, 0.0001 secs, 0.0030 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0023 secs

Status code distribution:
  [200]	10000 responses



