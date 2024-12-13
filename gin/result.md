##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0927 secs
  Slowest:	0.0050 secs
  Fastest:	0.0000 secs
  Average:	0.0005 secs
  Requests/sec:	107906.1486
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [6745]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [2504]	|■■■■■■■■■■■■■■■
  0.002 [599]	|■■■■
  0.002 [83]	|
  0.002 [18]	|
  0.003 [0]	|
  0.003 [0]	|
  0.004 [0]	|
  0.004 [20]	|
  0.005 [30]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0006 secs
  90% in 0.0009 secs
  95% in 0.0012 secs
  99% in 0.0018 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0050 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0015 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0027 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0017 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.2289 secs
  Slowest:	0.0191 secs
  Fastest:	0.0001 secs
  Average:	0.0011 secs
  Requests/sec:	43677.6669
  
  Total data:	307762 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9619]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [333]	|■
  0.006 [17]	|
  0.008 [6]	|
  0.010 [1]	|
  0.011 [5]	|
  0.013 [3]	|
  0.015 [5]	|
  0.017 [5]	|
  0.019 [5]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0004 secs
  50% in 0.0013 secs
  75% in 0.0015 secs
  90% in 0.0016 secs
  95% in 0.0018 secs
  99% in 0.0026 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0191 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:	0.0011 secs, 0.0000 secs, 0.0179 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.1500 secs
  Slowest:	0.3757 secs
  Fastest:	0.0003 secs
  Average:	0.0053 secs
  Requests/sec:	8695.6144
  
  Total data:	3188158 bytes
  Size/request:	318 bytes

Response time histogram:
  0.000 [1]	|
  0.038 [9946]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.075 [0]	|
  0.113 [1]	|
  0.150 [46]	|
  0.188 [0]	|
  0.226 [1]	|
  0.263 [3]	|
  0.301 [0]	|
  0.338 [0]	|
  0.376 [2]	|


Latency distribution:
  10% in 0.0008 secs
  25% in 0.0038 secs
  50% in 0.0049 secs
  75% in 0.0058 secs
  90% in 0.0068 secs
  95% in 0.0080 secs
  99% in 0.0111 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.3757 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0053 secs, 0.0003 secs, 0.3757 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0021 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.2867 secs
  Slowest:	0.2633 secs
  Fastest:	0.0001 secs
  Average:	0.0010 secs
  Requests/sec:	34877.7947
  
  Total data:	13080000 bytes
  Size/request:	1308 bytes

Response time histogram:
  0.000 [1]	|
  0.026 [9976]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.053 [0]	|
  0.079 [0]	|
  0.105 [0]	|
  0.132 [1]	|
  0.158 [18]	|
  0.184 [0]	|
  0.211 [0]	|
  0.237 [0]	|
  0.263 [4]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0005 secs
  75% in 0.0010 secs
  90% in 0.0014 secs
  95% in 0.0018 secs
  99% in 0.0027 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.2633 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0010 secs, 0.0001 secs, 0.2623 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	3.1010 secs
  Slowest:	0.0472 secs
  Fastest:	0.0019 secs
  Average:	0.0153 secs
  Requests/sec:	3224.7537
  
  Total data:	3187859 bytes
  Size/request:	318 bytes

Response time histogram:
  0.002 [1]	|
  0.006 [40]	|
  0.011 [430]	|■■■
  0.015 [5306]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.020 [3882]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.025 [191]	|■
  0.029 [80]	|■
  0.034 [27]	|
  0.038 [16]	|
  0.043 [17]	|
  0.047 [10]	|


Latency distribution:
  10% in 0.0128 secs
  25% in 0.0144 secs
  50% in 0.0152 secs
  75% in 0.0161 secs
  90% in 0.0172 secs
  95% in 0.0186 secs
  99% in 0.0263 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0019 secs, 0.0472 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:	0.0153 secs, 0.0018 secs, 0.0472 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0704 secs
  Slowest:	0.0023 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	142139.2564
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [4871]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.000 [3043]	|■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [931]	|■■■■■■■■
  0.001 [605]	|■■■■■
  0.001 [278]	|■■
  0.001 [129]	|■
  0.002 [48]	|
  0.002 [51]	|
  0.002 [30]	|
  0.002 [13]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0003 secs
  75% in 0.0004 secs
  90% in 0.0007 secs
  95% in 0.0009 secs
  99% in 0.0016 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0023 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0018 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0013 secs

Status code distribution:
  [200]	10000 responses



