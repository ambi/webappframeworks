##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.1335 secs
  Slowest:	0.0074 secs
  Fastest:	0.0001 secs
  Average:	0.0006 secs
  Requests/sec:	74891.2452
  
  Total data:	280000 bytes
  Size/request:	28 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [7728]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1240]	|■■■■■■
  0.002 [766]	|■■■■
  0.003 [188]	|■
  0.004 [27]	|
  0.004 [5]	|
  0.005 [30]	|
  0.006 [13]	|
  0.007 [1]	|
  0.007 [1]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0002 secs
  50% in 0.0004 secs
  75% in 0.0007 secs
  90% in 0.0016 secs
  95% in 0.0019 secs
  99% in 0.0028 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0074 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0029 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0022 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0027 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0028 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.3049 secs
  Slowest:	0.0382 secs
  Fastest:	0.0002 secs
  Average:	0.0015 secs
  Requests/sec:	32796.9585
  
  Total data:	317730 bytes
  Size/request:	31 bytes

Response time histogram:
  0.000 [1]	|
  0.004 [9947]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.008 [9]	|
  0.012 [1]	|
  0.015 [1]	|
  0.019 [12]	|
  0.023 [8]	|
  0.027 [12]	|
  0.031 [4]	|
  0.034 [3]	|
  0.038 [2]	|


Latency distribution:
  10% in 0.0009 secs
  25% in 0.0012 secs
  50% in 0.0014 secs
  75% in 0.0016 secs
  90% in 0.0018 secs
  95% in 0.0020 secs
  99% in 0.0025 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.0382 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:	0.0014 secs, 0.0002 secs, 0.0343 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.4178 secs
  Slowest:	0.2522 secs
  Fastest:	0.0008 secs
  Average:	0.0063 secs
  Requests/sec:	7053.0830
  
  Total data:	3197445 bytes
  Size/request:	319 bytes

Response time histogram:
  0.001 [1]	|
  0.026 [9725]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.051 [89]	|
  0.076 [67]	|
  0.101 [32]	|
  0.127 [32]	|
  0.152 [33]	|
  0.177 [11]	|
  0.202 [3]	|
  0.227 [0]	|
  0.252 [7]	|


Latency distribution:
  10% in 0.0020 secs
  25% in 0.0024 secs
  50% in 0.0036 secs
  75% in 0.0061 secs
  90% in 0.0071 secs
  95% in 0.0086 secs
  99% in 0.0920 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0008 secs, 0.2522 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0013 secs
  resp wait:	0.0062 secs, 0.0008 secs, 0.2522 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.5977 secs
  Slowest:	0.0248 secs
  Fastest:	0.0003 secs
  Average:	0.0029 secs
  Requests/sec:	16729.8865
  
  Total data:	19750000 bytes
  Size/request:	1975 bytes

Response time histogram:
  0.000 [1]	|
  0.003 [6391]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [1958]	|■■■■■■■■■■■■
  0.008 [1049]	|■■■■■■■
  0.010 [383]	|■■
  0.013 [141]	|■
  0.015 [46]	|
  0.017 [20]	|
  0.020 [5]	|
  0.022 [3]	|
  0.025 [3]	|


Latency distribution:
  10% in 0.0008 secs
  25% in 0.0011 secs
  50% in 0.0018 secs
  75% in 0.0040 secs
  90% in 0.0064 secs
  95% in 0.0081 secs
  99% in 0.0118 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.0248 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:	0.0028 secs, 0.0002 secs, 0.0248 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0011 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	5.3437 secs
  Slowest:	0.0785 secs
  Fastest:	0.0061 secs
  Average:	0.0265 secs
  Requests/sec:	1871.3756
  
  Total data:	3197920 bytes
  Size/request:	319 bytes

Response time histogram:
  0.006 [1]	|
  0.013 [12]	|
  0.021 [478]	|■■■
  0.028 [7008]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.035 [1600]	|■■■■■■■■■
  0.042 [703]	|■■■■
  0.050 [83]	|
  0.057 [29]	|
  0.064 [1]	|
  0.071 [43]	|
  0.078 [42]	|


Latency distribution:
  10% in 0.0216 secs
  25% in 0.0230 secs
  50% in 0.0248 secs
  75% in 0.0278 secs
  90% in 0.0341 secs
  95% in 0.0381 secs
  99% in 0.0527 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0061 secs, 0.0785 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0264 secs, 0.0060 secs, 0.0784 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1476 secs
  Slowest:	0.0120 secs
  Fastest:	0.0001 secs
  Average:	0.0007 secs
  Requests/sec:	67771.9297
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8714]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [771]	|■■■■
  0.004 [249]	|■
  0.005 [143]	|■
  0.006 [79]	|
  0.007 [19]	|
  0.008 [13]	|
  0.010 [7]	|
  0.011 [2]	|
  0.012 [2]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0007 secs
  90% in 0.0016 secs
  95% in 0.0025 secs
  99% in 0.0051 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0120 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0015 secs
  resp wait:	0.0006 secs, 0.0000 secs, 0.0115 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0013 secs

Status code distribution:
  [200]	10000 responses



