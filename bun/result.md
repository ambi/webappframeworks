##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0913 secs
  Slowest:	0.0079 secs
  Fastest:	0.0000 secs
  Average:	0.0004 secs
  Requests/sec:	109531.3742
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8679]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1261]	|■■■■■■
  0.002 [9]	|
  0.003 [0]	|
  0.004 [0]	|
  0.005 [0]	|
  0.006 [0]	|
  0.006 [0]	|
  0.007 [17]	|
  0.008 [33]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0006 secs
  90% in 0.0009 secs
  95% in 0.0012 secs
  99% in 0.0015 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0079 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0006 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0029 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.2488 secs
  Slowest:	0.0193 secs
  Fastest:	0.0004 secs
  Average:	0.0012 secs
  Requests/sec:	40196.1572
  
  Total data:	307825 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9809]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [140]	|■
  0.006 [0]	|
  0.008 [0]	|
  0.010 [0]	|
  0.012 [0]	|
  0.014 [0]	|
  0.016 [2]	|
  0.017 [11]	|
  0.019 [37]	|


Latency distribution:
  10% in 0.0009 secs
  25% in 0.0009 secs
  50% in 0.0011 secs
  75% in 0.0013 secs
  90% in 0.0016 secs
  95% in 0.0019 secs
  99% in 0.0027 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.0193 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:	0.0012 secs, 0.0004 secs, 0.0179 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.0063 secs
  Slowest:	0.0097 secs
  Fastest:	0.0024 secs
  Average:	0.0050 secs
  Requests/sec:	9937.1253
  
  Total data:	3188048 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.003 [2]	|
  0.004 [5]	|
  0.005 [2685]	|■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [4460]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [2458]	|■■■■■■■■■■■■■■■■■■■■■■
  0.007 [214]	|■■
  0.008 [100]	|■
  0.008 [36]	|
  0.009 [22]	|
  0.010 [17]	|


Latency distribution:
  10% in 0.0045 secs
  25% in 0.0046 secs
  50% in 0.0048 secs
  75% in 0.0055 secs
  90% in 0.0058 secs
  95% in 0.0060 secs
  99% in 0.0073 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0024 secs, 0.0097 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0050 secs, 0.0024 secs, 0.0097 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.2600 secs
  Slowest:	0.0100 secs
  Fastest:	0.0002 secs
  Average:	0.0013 secs
  Requests/sec:	38461.2796
  
  Total data:	16050000 bytes
  Size/request:	1605 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [6031]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [3709]	|■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [198]	|■
  0.004 [11]	|
  0.005 [0]	|
  0.006 [0]	|
  0.007 [5]	|
  0.008 [12]	|
  0.009 [8]	|
  0.010 [25]	|


Latency distribution:
  10% in 0.0010 secs
  25% in 0.0011 secs
  50% in 0.0011 secs
  75% in 0.0013 secs
  90% in 0.0017 secs
  95% in 0.0020 secs
  99% in 0.0025 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.0100 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0006 secs
  resp wait:	0.0013 secs, 0.0002 secs, 0.0087 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	3.5788 secs
  Slowest:	0.0269 secs
  Fastest:	0.0136 secs
  Average:	0.0179 secs
  Requests/sec:	2794.2008
  
  Total data:	3188042 bytes
  Size/request:	318 bytes

Response time histogram:
  0.014 [1]	|
  0.015 [18]	|
  0.016 [472]	|■■■■
  0.018 [4530]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.019 [3992]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.020 [549]	|■■■■■
  0.022 [57]	|■
  0.023 [96]	|■
  0.024 [114]	|■
  0.026 [79]	|■
  0.027 [92]	|■


Latency distribution:
  10% in 0.0168 secs
  25% in 0.0172 secs
  50% in 0.0176 secs
  75% in 0.0182 secs
  90% in 0.0189 secs
  95% in 0.0196 secs
  99% in 0.0255 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0136 secs, 0.0269 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:	0.0179 secs, 0.0136 secs, 0.0269 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0642 secs
  Slowest:	0.0030 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	155652.6223
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [7200]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1697]	|■■■■■■■■■
  0.001 [759]	|■■■■
  0.001 [206]	|■
  0.002 [51]	|
  0.002 [42]	|
  0.002 [31]	|
  0.002 [6]	|
  0.003 [4]	|
  0.003 [3]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0002 secs
  75% in 0.0003 secs
  90% in 0.0007 secs
  95% in 0.0008 secs
  99% in 0.0014 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0030 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:	0.0002 secs, 0.0000 secs, 0.0018 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0024 secs

Status code distribution:
  [200]	10000 responses



