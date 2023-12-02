##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.1287 secs
  Slowest:	0.0127 secs
  Fastest:	0.0001 secs
  Average:	0.0006 secs
  Requests/sec:	77716.9671
  
  Total data:	280000 bytes
  Size/request:	28 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8936]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [824]	|■■■■
  0.004 [146]	|■
  0.005 [37]	|
  0.006 [17]	|
  0.008 [34]	|
  0.009 [2]	|
  0.010 [1]	|
  0.011 [0]	|
  0.013 [2]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0002 secs
  50% in 0.0004 secs
  75% in 0.0007 secs
  90% in 0.0014 secs
  95% in 0.0020 secs
  99% in 0.0036 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0127 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0045 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0027 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0025 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0045 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.7231 secs
  Slowest:	0.1911 secs
  Fastest:	0.0002 secs
  Average:	0.0032 secs
  Requests/sec:	13829.0534
  
  Total data:	317851 bytes
  Size/request:	31 bytes

Response time histogram:
  0.000 [1]	|
  0.019 [9689]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.038 [253]	|■
  0.057 [22]	|
  0.077 [5]	|
  0.096 [7]	|
  0.115 [0]	|
  0.134 [10]	|
  0.153 [12]	|
  0.172 [0]	|
  0.191 [1]	|


Latency distribution:
  10% in 0.0005 secs
  25% in 0.0008 secs
  50% in 0.0014 secs
  75% in 0.0022 secs
  90% in 0.0038 secs
  95% in 0.0182 secs
  99% in 0.0272 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.1911 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0031 secs, 0.0002 secs, 0.1911 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0014 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	3.7423 secs
  Slowest:	0.2924 secs
  Fastest:	0.0011 secs
  Average:	0.0176 secs
  Requests/sec:	2672.1354
  
  Total data:	3198052 bytes
  Size/request:	319 bytes

Response time histogram:
  0.001 [1]	|
  0.030 [8804]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.059 [914]	|■■■■
  0.088 [54]	|
  0.118 [40]	|
  0.147 [126]	|■
  0.176 [5]	|
  0.205 [0]	|
  0.234 [1]	|
  0.263 [9]	|
  0.292 [46]	|


Latency distribution:
  10% in 0.0027 secs
  25% in 0.0057 secs
  50% in 0.0098 secs
  75% in 0.0247 secs
  90% in 0.0315 secs
  95% in 0.0413 secs
  99% in 0.1424 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0011 secs, 0.2924 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0019 secs
  resp wait:	0.0175 secs, 0.0010 secs, 0.2923 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	1.1895 secs
  Slowest:	0.3046 secs
  Fastest:	0.0003 secs
  Average:	0.0054 secs
  Requests/sec:	8406.6863
  
  Total data:	19750000 bytes
  Size/request:	1975 bytes

Response time histogram:
  0.000 [1]	|
  0.031 [9782]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.061 [134]	|■
  0.092 [32]	|
  0.122 [13]	|
  0.152 [25]	|
  0.183 [1]	|
  0.213 [0]	|
  0.244 [0]	|
  0.274 [11]	|
  0.305 [1]	|


Latency distribution:
  10% in 0.0007 secs
  25% in 0.0011 secs
  50% in 0.0022 secs
  75% in 0.0040 secs
  90% in 0.0141 secs
  95% in 0.0209 secs
  99% in 0.0497 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.3046 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0015 secs
  resp wait:	0.0053 secs, 0.0003 secs, 0.3037 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0016 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	12.8804 secs
  Slowest:	0.1859 secs
  Fastest:	0.0088 secs
  Average:	0.0641 secs
  Requests/sec:	776.3762
  
  Total data:	3197908 bytes
  Size/request:	319 bytes

Response time histogram:
  0.009 [1]	|
  0.027 [19]	|
  0.044 [76]	|
  0.062 [6253]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.080 [2406]	|■■■■■■■■■■■■■■■
  0.097 [813]	|■■■■■
  0.115 [229]	|■
  0.133 [102]	|■
  0.150 [96]	|■
  0.168 [3]	|
  0.186 [2]	|


Latency distribution:
  10% in 0.0515 secs
  25% in 0.0544 secs
  50% in 0.0584 secs
  75% in 0.0721 secs
  90% in 0.0822 secs
  95% in 0.0939 secs
  99% in 0.1329 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0088 secs, 0.1859 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0014 secs
  resp wait:	0.0640 secs, 0.0087 secs, 0.1859 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1311 secs
  Slowest:	0.0103 secs
  Fastest:	0.0001 secs
  Average:	0.0006 secs
  Requests/sec:	76269.8922
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8561]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [976]	|■■■■■
  0.003 [331]	|■■
  0.004 [62]	|
  0.005 [28]	|
  0.006 [25]	|
  0.007 [0]	|
  0.008 [0]	|
  0.009 [6]	|
  0.010 [10]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0002 secs
  50% in 0.0004 secs
  75% in 0.0007 secs
  90% in 0.0014 secs
  95% in 0.0020 secs
  99% in 0.0035 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0103 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0038 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0048 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0043 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0054 secs

Status code distribution:
  [200]	10000 responses



