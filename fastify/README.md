##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.5296 secs
  Slowest:	0.1511 secs
  Fastest:	0.0001 secs
  Average:	0.0023 secs
  Requests/sec:	18883.5973
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.015 [9868]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.030 [39]	|
  0.045 [54]	|
  0.060 [0]	|
  0.076 [1]	|
  0.091 [9]	|
  0.106 [1]	|
  0.121 [0]	|
  0.136 [15]	|
  0.151 [12]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0004 secs
  50% in 0.0007 secs
  75% in 0.0019 secs
  90% in 0.0047 secs
  95% in 0.0070 secs
  99% in 0.0295 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.1511 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0207 secs
  resp wait:	0.0020 secs, 0.0001 secs, 0.1510 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0065 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	2.2473 secs
  Slowest:	0.3609 secs
  Fastest:	0.0005 secs
  Average:	0.0100 secs
  Requests/sec:	4449.8055
  
  Total data:	307822 bytes
  Size/request:	30 bytes

Response time histogram:
  0.001 [1]	|
  0.037 [9319]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.073 [246]	|■
  0.109 [201]	|■
  0.145 [153]	|■
  0.181 [44]	|
  0.217 [25]	|
  0.253 [0]	|
  0.289 [0]	|
  0.325 [7]	|
  0.361 [4]	|


Latency distribution:
  10% in 0.0008 secs
  25% in 0.0010 secs
  50% in 0.0021 secs
  75% in 0.0053 secs
  90% in 0.0145 secs
  95% in 0.0599 secs
  99% in 0.1372 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0005 secs, 0.3609 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:	0.0099 secs, 0.0005 secs, 0.3555 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0020 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	7.9100 secs
  Slowest:	0.3831 secs
  Fastest:	0.0019 secs
  Average:	0.0322 secs
  Requests/sec:	1264.2268
  
  Total data:	3187749 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.040 [8165]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.078 [754]	|■■■■
  0.116 [275]	|■
  0.154 [310]	|■■
  0.192 [219]	|■
  0.231 [93]	|
  0.269 [91]	|
  0.307 [55]	|
  0.345 [24]	|
  0.383 [13]	|


Latency distribution:
  10% in 0.0037 secs
  25% in 0.0050 secs
  50% in 0.0131 secs
  75% in 0.0309 secs
  90% in 0.0886 secs
  95% in 0.1541 secs
  99% in 0.2641 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0019 secs, 0.3831 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0019 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0026 secs
  resp wait:	0.0321 secs, 0.0018 secs, 0.3830 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	5.3673 secs
  Slowest:	0.6511 secs
  Fastest:	0.0013 secs
  Average:	0.0242 secs
  Requests/sec:	1863.1220
  
  Total data:	14970000 bytes
  Size/request:	1497 bytes

Response time histogram:
  0.001 [1]	|
  0.066 [8952]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.131 [512]	|■■
  0.196 [383]	|■■
  0.261 [132]	|■
  0.326 [17]	|
  0.391 [0]	|
  0.456 [0]	|
  0.521 [0]	|
  0.586 [0]	|
  0.651 [3]	|


Latency distribution:
  10% in 0.0032 secs
  25% in 0.0043 secs
  50% in 0.0073 secs
  75% in 0.0165 secs
  90% in 0.0723 secs
  95% in 0.1353 secs
  99% in 0.2205 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0001 secs, 0.0013 secs, 0.6511 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0036 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0060 secs
  resp wait:	0.0241 secs, 0.0013 secs, 0.6319 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0011 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	15.4362 secs
  Slowest:	0.3992 secs
  Fastest:	0.0087 secs
  Average:	0.0737 secs
  Requests/sec:	647.8260
  
  Total data:	3187935 bytes
  Size/request:	318 bytes

Response time histogram:
  0.009 [1]	|
  0.048 [4556]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.087 [2624]	|■■■■■■■■■■■■■■■■■■■■■■■
  0.126 [918]	|■■■■■■■■
  0.165 [886]	|■■■■■■■■
  0.204 [753]	|■■■■■■■
  0.243 [99]	|■
  0.282 [102]	|■
  0.321 [53]	|
  0.360 [3]	|
  0.399 [5]	|


Latency distribution:
  10% in 0.0229 secs
  25% in 0.0324 secs
  50% in 0.0518 secs
  75% in 0.0958 secs
  90% in 0.1653 secs
  95% in 0.1815 secs
  99% in 0.2733 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0087 secs, 0.3992 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0109 secs
  resp wait:	0.0735 secs, 0.0086 secs, 0.3991 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.3873 secs
  Slowest:	0.1305 secs
  Fastest:	0.0001 secs
  Average:	0.0016 secs
  Requests/sec:	25821.1231
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.013 [9946]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.026 [14]	|
  0.039 [9]	|
  0.052 [0]	|
  0.065 [2]	|
  0.078 [0]	|
  0.091 [1]	|
  0.104 [5]	|
  0.117 [2]	|
  0.131 [20]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0005 secs
  75% in 0.0013 secs
  90% in 0.0031 secs
  95% in 0.0053 secs
  99% in 0.0099 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.1305 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0029 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0075 secs
  resp wait:	0.0013 secs, 0.0000 secs, 0.1297 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0073 secs

Status code distribution:
  [200]	10000 responses



