##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.9151 secs
  Slowest:	0.1916 secs
  Fastest:	0.0002 secs
  Average:	0.0042 secs
  Requests/sec:	10927.2150
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.019 [9608]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.038 [144]	|■
  0.058 [121]	|■
  0.077 [30]	|
  0.096 [43]	|
  0.115 [20]	|
  0.134 [15]	|
  0.153 [11]	|
  0.172 [0]	|
  0.192 [7]	|


Latency distribution:
  10% in 0.0004 secs
  25% in 0.0007 secs
  50% in 0.0012 secs
  75% in 0.0025 secs
  90% in 0.0058 secs
  95% in 0.0131 secs
  99% in 0.0724 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0001 secs, 0.0002 secs, 0.1916 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0166 secs
  resp wait:	0.0040 secs, 0.0002 secs, 0.1916 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0090 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	3.4546 secs
  Slowest:	0.2192 secs
  Fastest:	0.0006 secs
  Average:	0.0157 secs
  Requests/sec:	2894.7196
  
  Total data:	307802 bytes
  Size/request:	30 bytes

Response time histogram:
  0.001 [1]	|
  0.022 [7789]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.044 [1160]	|■■■■■■
  0.066 [831]	|■■■■
  0.088 [82]	|
  0.110 [60]	|
  0.132 [39]	|
  0.154 [18]	|
  0.175 [4]	|
  0.197 [4]	|
  0.219 [12]	|


Latency distribution:
  10% in 0.0014 secs
  25% in 0.0025 secs
  50% in 0.0079 secs
  75% in 0.0205 secs
  90% in 0.0455 secs
  95% in 0.0520 secs
  99% in 0.1036 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0006 secs, 0.2192 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:	0.0156 secs, 0.0005 secs, 0.2147 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0013 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	13.8112 secs
  Slowest:	0.8865 secs
  Fastest:	0.0020 secs
  Average:	0.0675 secs
  Requests/sec:	724.0516
  
  Total data:	3187136 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.090 [7795]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.179 [1603]	|■■■■■■■■
  0.267 [429]	|■■
  0.356 [143]	|■
  0.444 [13]	|
  0.533 [3]	|
  0.621 [0]	|
  0.710 [0]	|
  0.798 [1]	|
  0.887 [12]	|


Latency distribution:
  10% in 0.0079 secs
  25% in 0.0290 secs
  50% in 0.0528 secs
  75% in 0.0847 secs
  90% in 0.1481 secs
  95% in 0.1893 secs
  99% in 0.2981 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0020 secs, 0.8865 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0019 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0085 secs
  resp wait:	0.0674 secs, 0.0020 secs, 0.8746 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	3.5791 secs
  Slowest:	0.3195 secs
  Fastest:	0.0006 secs
  Average:	0.0157 secs
  Requests/sec:	2793.9624
  
  Total data:	14970000 bytes
  Size/request:	1497 bytes

Response time histogram:
  0.001 [1]	|
  0.033 [8501]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.064 [1217]	|■■■■■■
  0.096 [183]	|■
  0.128 [75]	|
  0.160 [8]	|
  0.192 [2]	|
  0.224 [0]	|
  0.256 [0]	|
  0.288 [5]	|
  0.319 [8]	|


Latency distribution:
  10% in 0.0011 secs
  25% in 0.0017 secs
  50% in 0.0068 secs
  75% in 0.0210 secs
  90% in 0.0488 secs
  95% in 0.0545 secs
  99% in 0.0961 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0006 secs, 0.3195 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0019 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:	0.0156 secs, 0.0006 secs, 0.3158 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0014 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	25.1412 secs
  Slowest:	0.2931 secs
  Fastest:	0.0094 secs
  Average:	0.1252 secs
  Requests/sec:	397.7532
  
  Total data:	3188179 bytes
  Size/request:	318 bytes

Response time histogram:
  0.009 [1]	|
  0.038 [10]	|
  0.066 [34]	|
  0.094 [266]	|■■
  0.123 [5054]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.151 [3501]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.180 [701]	|■■■■■■
  0.208 [305]	|■■
  0.236 [87]	|■
  0.265 [19]	|
  0.293 [22]	|


Latency distribution:
  10% in 0.1013 secs
  25% in 0.1068 secs
  50% in 0.1201 secs
  75% in 0.1386 secs
  90% in 0.1528 secs
  95% in 0.1748 secs
  99% in 0.2138 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0094 secs, 0.2931 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0021 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0043 secs
  resp wait:	0.1251 secs, 0.0093 secs, 0.2930 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.6540 secs
  Slowest:	0.1609 secs
  Fastest:	0.0002 secs
  Average:	0.0029 secs
  Requests/sec:	15290.8835
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.016 [9708]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.032 [111]	|
  0.048 [49]	|
  0.064 [32]	|
  0.081 [29]	|
  0.097 [44]	|
  0.113 [17]	|
  0.129 [4]	|
  0.145 [4]	|
  0.161 [1]	|


Latency distribution:
  10% in 0.0004 secs
  25% in 0.0006 secs
  50% in 0.0009 secs
  75% in 0.0015 secs
  90% in 0.0031 secs
  95% in 0.0080 secs
  99% in 0.0632 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.1609 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0020 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0072 secs
  resp wait:	0.0027 secs, 0.0001 secs, 0.1608 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0081 secs

Status code distribution:
  [200]	10000 responses



