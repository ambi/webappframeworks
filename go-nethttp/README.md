##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.1615 secs
  Slowest:	0.0112 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	61920.9823

  Total data:	280000 bytes
  Size/request:	28 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8101]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1411]	|■■■■■■■
  0.003 [344]	|■■
  0.005 [73]	|
  0.006 [37]	|
  0.007 [25]	|
  0.008 [3]	|
  0.009 [2]	|
  0.010 [0]	|
  0.011 [3]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0005 secs
  75% in 0.0009 secs
  90% in 0.0018 secs
  95% in 0.0023 secs
  99% in 0.0042 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0112 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0037 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0031 secs
  resp wait:	0.0005 secs, 0.0001 secs, 0.0047 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0031 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.8417 secs
  Slowest:	0.1945 secs
  Fastest:	0.0002 secs
  Average:	0.0037 secs
  Requests/sec:	11880.5341

  Total data:	317809 bytes
  Size/request:	31 bytes

Response time histogram:
  0.000 [1]	|
  0.020 [9596]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.039 [326]	|■
  0.059 [43]	|
  0.078 [0]	|
  0.097 [3]	|
  0.117 [4]	|
  0.136 [2]	|
  0.156 [12]	|
  0.175 [2]	|
  0.195 [11]	|


Latency distribution:
  10% in 0.0005 secs
  25% in 0.0010 secs
  50% in 0.0018 secs
  75% in 0.0028 secs
  90% in 0.0041 secs
  95% in 0.0191 secs
  99% in 0.0271 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.1945 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0029 secs
  resp wait:	0.0036 secs, 0.0002 secs, 0.1884 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0012 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	3.6355 secs
  Slowest:	0.3496 secs
  Fastest:	0.0012 secs
  Average:	0.0172 secs
  Requests/sec:	2750.6481

  Total data:	3197633 bytes
  Size/request:	319 bytes

Response time histogram:
  0.001 [1]	|
  0.036 [9377]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.071 [258]	|■
  0.106 [84]	|
  0.141 [82]	|
  0.175 [109]	|
  0.210 [17]	|
  0.245 [6]	|
  0.280 [63]	|
  0.315 [0]	|
  0.350 [3]	|


Latency distribution:
  10% in 0.0029 secs
  25% in 0.0039 secs
  50% in 0.0085 secs
  75% in 0.0210 secs
  90% in 0.0298 secs
  95% in 0.0480 secs
  99% in 0.1612 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0012 secs, 0.3496 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0171 secs, 0.0012 secs, 0.3496 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	1.2153 secs
  Slowest:	0.1771 secs
  Fastest:	0.0004 secs
  Average:	0.0056 secs
  Requests/sec:	8228.7222

  Total data:	19750000 bytes
  Size/request:	1975 bytes

Response time histogram:
  0.000 [1]	|
  0.018 [8953]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.036 [973]	|■■■■
  0.053 [53]	|
  0.071 [2]	|
  0.089 [6]	|
  0.106 [8]	|
  0.124 [2]	|
  0.142 [0]	|
  0.159 [0]	|
  0.177 [2]	|


Latency distribution:
  10% in 0.0015 secs
  25% in 0.0025 secs
  50% in 0.0035 secs
  75% in 0.0046 secs
  90% in 0.0188 secs
  95% in 0.0211 secs
  99% in 0.0286 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.1771 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0013 secs
  resp wait:	0.0055 secs, 0.0003 secs, 0.1770 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0250 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	11.9561 secs
  Slowest:	0.1770 secs
  Fastest:	0.0072 secs
  Average:	0.0595 secs
  Requests/sec:	836.3914

  Total data:	3197907 bytes
  Size/request:	319 bytes

Response time histogram:
  0.007 [1]	|
  0.024 [22]	|
  0.041 [94]	|■
  0.058 [6448]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.075 [2406]	|■■■■■■■■■■■■■■■
  0.092 [860]	|■■■■■
  0.109 [77]	|
  0.126 [22]	|
  0.143 [33]	|
  0.160 [35]	|
  0.177 [2]	|


Latency distribution:
  10% in 0.0504 secs
  25% in 0.0530 secs
  50% in 0.0557 secs
  75% in 0.0613 secs
  90% in 0.0752 secs
  95% in 0.0795 secs
  99% in 0.1079 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0072 secs, 0.1770 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0594 secs, 0.0071 secs, 0.1769 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1676 secs
  Slowest:	0.0092 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	59673.3798

  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [7647]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1390]	|■■■■■■■
  0.003 [624]	|■■■
  0.004 [191]	|■
  0.005 [63]	|
  0.006 [29]	|
  0.006 [6]	|
  0.007 [34]	|
  0.008 [14]	|
  0.009 [1]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0005 secs
  75% in 0.0009 secs
  90% in 0.0019 secs
  95% in 0.0025 secs
  99% in 0.0044 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0092 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:	0.0005 secs, 0.0001 secs, 0.0049 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0040 secs

Status code distribution:
  [200]	10000 responses


