##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.1584 secs
  Slowest:	0.0102 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	63138.4290
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8086]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1290]	|■■■■■■
  0.003 [472]	|■■
  0.004 [81]	|
  0.005 [15]	|
  0.006 [21]	|
  0.007 [16]	|
  0.008 [6]	|
  0.009 [2]	|
  0.010 [10]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0005 secs
  75% in 0.0009 secs
  90% in 0.0017 secs
  95% in 0.0023 secs
  99% in 0.0035 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0102 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0015 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0039 secs
  resp wait:	0.0005 secs, 0.0001 secs, 0.0048 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0044 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.7148 secs
  Slowest:	0.2509 secs
  Fastest:	0.0002 secs
  Average:	0.0034 secs
  Requests/sec:	13990.5872
  
  Total data:	307792 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.025 [9995]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.050 [0]	|
  0.075 [0]	|
  0.100 [0]	|
  0.126 [0]	|
  0.151 [2]	|
  0.176 [0]	|
  0.201 [0]	|
  0.226 [0]	|
  0.251 [2]	|


Latency distribution:
  10% in 0.0012 secs
  25% in 0.0017 secs
  50% in 0.0023 secs
  75% in 0.0028 secs
  90% in 0.0038 secs
  95% in 0.0184 secs
  99% in 0.0200 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.2509 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0022 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:	0.0033 secs, 0.0002 secs, 0.2509 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	3.3043 secs
  Slowest:	0.2773 secs
  Fastest:	0.0011 secs
  Average:	0.0160 secs
  Requests/sec:	3026.3437
  
  Total data:	3187477 bytes
  Size/request:	318 bytes

Response time histogram:
  0.001 [1]	|
  0.029 [8744]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.056 [1172]	|■■■■■
  0.084 [48]	|
  0.112 [0]	|
  0.139 [6]	|
  0.167 [15]	|
  0.194 [1]	|
  0.222 [0]	|
  0.250 [0]	|
  0.277 [13]	|


Latency distribution:
  10% in 0.0070 secs
  25% in 0.0088 secs
  50% in 0.0107 secs
  75% in 0.0262 secs
  90% in 0.0295 secs
  95% in 0.0310 secs
  99% in 0.0520 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0011 secs, 0.2773 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0018 secs
  resp wait:	0.0159 secs, 0.0011 secs, 0.2773 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.9359 secs
  Slowest:	0.1949 secs
  Fastest:	0.0003 secs
  Average:	0.0044 secs
  Requests/sec:	10684.6839
  
  Total data:	13080000 bytes
  Size/request:	1308 bytes

Response time histogram:
  0.000 [1]	|
  0.020 [9521]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.039 [432]	|■■
  0.059 [13]	|
  0.078 [14]	|
  0.098 [5]	|
  0.117 [2]	|
  0.136 [5]	|
  0.156 [4]	|
  0.175 [0]	|
  0.195 [3]	|


Latency distribution:
  10% in 0.0011 secs
  25% in 0.0020 secs
  50% in 0.0025 secs
  75% in 0.0033 secs
  90% in 0.0054 secs
  95% in 0.0197 secs
  99% in 0.0239 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.1949 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0022 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0026 secs
  resp wait:	0.0042 secs, 0.0002 secs, 0.1944 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0016 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	12.1890 secs
  Slowest:	0.2936 secs
  Fastest:	0.0090 secs
  Average:	0.0606 secs
  Requests/sec:	820.4140
  
  Total data:	3187808 bytes
  Size/request:	318 bytes

Response time histogram:
  0.009 [1]	|
  0.037 [83]	|
  0.066 [7833]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.094 [1729]	|■■■■■■■■■
  0.123 [185]	|■
  0.151 [46]	|
  0.180 [69]	|
  0.208 [30]	|
  0.237 [4]	|
  0.265 [4]	|
  0.294 [16]	|


Latency distribution:
  10% in 0.0495 secs
  25% in 0.0521 secs
  50% in 0.0554 secs
  75% in 0.0611 secs
  90% in 0.0764 secs
  95% in 0.0837 secs
  99% in 0.1690 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0090 secs, 0.2936 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0605 secs, 0.0089 secs, 0.2935 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1708 secs
  Slowest:	0.0110 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	58536.0825
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8162]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1278]	|■■■■■■
  0.003 [291]	|■
  0.004 [114]	|■
  0.006 [89]	|
  0.007 [43]	|
  0.008 [15]	|
  0.009 [4]	|
  0.010 [2]	|
  0.011 [1]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0005 secs
  75% in 0.0009 secs
  90% in 0.0017 secs
  95% in 0.0024 secs
  99% in 0.0051 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0110 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0028 secs
  resp wait:	0.0006 secs, 0.0001 secs, 0.0103 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0029 secs

Status code distribution:
  [200]	10000 responses



