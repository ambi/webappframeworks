##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0940 secs
  Slowest:	0.0090 secs
  Fastest:	0.0000 secs
  Average:	0.0005 secs
  Requests/sec:	106356.5774
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8996]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [935]	|■■■■
  0.003 [18]	|
  0.004 [0]	|
  0.005 [0]	|
  0.005 [0]	|
  0.006 [0]	|
  0.007 [0]	|
  0.008 [11]	|
  0.009 [39]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0005 secs
  90% in 0.0009 secs
  95% in 0.0012 secs
  99% in 0.0016 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0090 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0041 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0013 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.3022 secs
  Slowest:	0.0470 secs
  Fastest:	0.0001 secs
  Average:	0.0015 secs
  Requests/sec:	33090.9239
  
  Total data:	307839 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.005 [9875]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.009 [24]	|
  0.014 [0]	|
  0.019 [50]	|
  0.024 [0]	|
  0.028 [0]	|
  0.033 [0]	|
  0.038 [0]	|
  0.042 [3]	|
  0.047 [47]	|


Latency distribution:
  10% in 0.0009 secs
  25% in 0.0009 secs
  50% in 0.0011 secs
  75% in 0.0013 secs
  90% in 0.0017 secs
  95% in 0.0021 secs
  99% in 0.0142 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0470 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0015 secs, 0.0001 secs, 0.0470 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	0.9902 secs
  Slowest:	0.0095 secs
  Fastest:	0.0011 secs
  Average:	0.0049 secs
  Requests/sec:	10098.6537
  
  Total data:	3187810 bytes
  Size/request:	318 bytes

Response time histogram:
  0.001 [1]	|
  0.002 [3]	|
  0.003 [5]	|
  0.004 [9]	|
  0.004 [2135]	|■■■■■■■■■■■■■■■■
  0.005 [5213]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [2130]	|■■■■■■■■■■■■■■■■
  0.007 [290]	|■■
  0.008 [87]	|■
  0.009 [101]	|■
  0.010 [26]	|


Latency distribution:
  10% in 0.0044 secs
  25% in 0.0045 secs
  50% in 0.0047 secs
  75% in 0.0054 secs
  90% in 0.0058 secs
  95% in 0.0062 secs
  99% in 0.0080 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0011 secs, 0.0095 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0010 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:	0.0049 secs, 0.0011 secs, 0.0095 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.3378 secs
  Slowest:	0.0088 secs
  Fastest:	0.0004 secs
  Average:	0.0017 secs
  Requests/sec:	29605.6199
  
  Total data:	12340000 bytes
  Size/request:	1234 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [73]	|
  0.002 [8491]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [1185]	|■■■■■■
  0.004 [159]	|■
  0.005 [46]	|
  0.005 [8]	|
  0.006 [0]	|
  0.007 [10]	|
  0.008 [14]	|
  0.009 [13]	|


Latency distribution:
  10% in 0.0014 secs
  25% in 0.0014 secs
  50% in 0.0015 secs
  75% in 0.0017 secs
  90% in 0.0023 secs
  95% in 0.0025 secs
  99% in 0.0035 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.0088 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0008 secs
  resp wait:	0.0017 secs, 0.0004 secs, 0.0075 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	3.5631 secs
  Slowest:	0.0304 secs
  Fastest:	0.0116 secs
  Average:	0.0178 secs
  Requests/sec:	2806.5268
  
  Total data:	3187740 bytes
  Size/request:	318 bytes

Response time histogram:
  0.012 [1]	|
  0.013 [1]	|
  0.015 [12]	|
  0.017 [3218]	|■■■■■■■■■■■■■■■■■■■■■
  0.019 [6172]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.021 [337]	|■■
  0.023 [8]	|
  0.025 [16]	|
  0.027 [117]	|■
  0.029 [110]	|■
  0.030 [8]	|


Latency distribution:
  10% in 0.0168 secs
  25% in 0.0171 secs
  50% in 0.0175 secs
  75% in 0.0180 secs
  90% in 0.0188 secs
  95% in 0.0192 secs
  99% in 0.0268 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0116 secs, 0.0304 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0178 secs, 0.0116 secs, 0.0304 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0003 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0683 secs
  Slowest:	0.0028 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	146356.2790
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [6844]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1847]	|■■■■■■■■■■■
  0.001 [794]	|■■■■■
  0.001 [385]	|■■
  0.001 [54]	|
  0.002 [21]	|
  0.002 [4]	|
  0.002 [1]	|
  0.002 [17]	|
  0.003 [32]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0003 secs
  90% in 0.0007 secs
  95% in 0.0008 secs
  99% in 0.0012 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0028 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0014 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0011 secs

Status code distribution:
  [200]	10000 responses



