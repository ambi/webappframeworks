##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.2631 secs
  Slowest:	0.0116 secs
  Fastest:	0.0001 secs
  Average:	0.0013 secs
  Requests/sec:	38001.8224
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [5357]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [4433]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [159]	|■
  0.005 [0]	|
  0.006 [4]	|
  0.007 [13]	|
  0.008 [14]	|
  0.009 [7]	|
  0.010 [4]	|
  0.012 [8]	|


Latency distribution:
  10% in 0.0008 secs
  25% in 0.0010 secs
  50% in 0.0012 secs
  75% in 0.0015 secs
  90% in 0.0018 secs
  95% in 0.0020 secs
  99% in 0.0027 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0116 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0021 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:	0.0012 secs, 0.0001 secs, 0.0029 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0017 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	5.1157 secs
  Slowest:	0.0885 secs
  Fastest:	0.0048 secs
  Average:	0.0255 secs
  Requests/sec:	1954.7746
  
  Total data:	327801 bytes
  Size/request:	32 bytes

Response time histogram:
  0.005 [1]	|
  0.013 [393]	|■■■■
  0.022 [4432]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.030 [198]	|■■
  0.038 [4265]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.047 [661]	|■■■■■■
  0.055 [0]	|
  0.063 [0]	|
  0.072 [5]	|
  0.080 [0]	|
  0.088 [45]	|


Latency distribution:
  10% in 0.0141 secs
  25% in 0.0154 secs
  50% in 0.0294 secs
  75% in 0.0340 secs
  90% in 0.0373 secs
  95% in 0.0391 secs
  99% in 0.0434 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0048 secs, 0.0885 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:	0.0255 secs, 0.0031 secs, 0.0885 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	50.5227 secs
  Slowest:	0.4700 secs
  Fastest:	0.1832 secs
  Average:	0.2526 secs
  Requests/sec:	197.9309
  

Response time histogram:
  0.183 [1]	|
  0.212 [451]	|■■■■■
  0.241 [3828]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.269 [3695]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.298 [987]	|■■■■■■■■■■
  0.327 [355]	|■■■■
  0.355 [418]	|■■■■
  0.384 [175]	|■■
  0.413 [40]	|
  0.441 [0]	|
  0.470 [50]	|■


Latency distribution:
  10% in 0.2176 secs
  25% in 0.2310 secs
  50% in 0.2436 secs
  75% in 0.2580 secs
  90% in 0.2985 secs
  95% in 0.3471 secs
  99% in 0.3738 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.1832 secs, 0.4700 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0018 secs
  resp wait:	0.2525 secs, 0.1808 secs, 0.4699 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	10.1393 secs
  Slowest:	0.1294 secs
  Fastest:	0.0131 secs
  Average:	0.0506 secs
  Requests/sec:	986.2633
  

Response time histogram:
  0.013 [1]	|
  0.025 [18]	|
  0.036 [365]	|■■
  0.048 [2632]	|■■■■■■■■■■■■■■■■
  0.060 [6420]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.071 [287]	|■■
  0.083 [127]	|■
  0.094 [8]	|
  0.106 [96]	|■
  0.118 [24]	|
  0.129 [22]	|


Latency distribution:
  10% in 0.0458 secs
  25% in 0.0475 secs
  50% in 0.0496 secs
  75% in 0.0519 secs
  90% in 0.0549 secs
  95% in 0.0635 secs
  99% in 0.1033 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0131 secs, 0.1294 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0505 secs, 0.0114 secs, 0.1293 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	156.0334 secs
  Slowest:	1.1374 secs
  Fastest:	0.6728 secs
  Average:	0.7801 secs
  Requests/sec:	64.0888
  

Response time histogram:
  0.673 [1]	|
  0.719 [1063]	|■■■■■■■■■■
  0.766 [4312]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.812 [2142]	|■■■■■■■■■■■■■■■■■■■■
  0.859 [1546]	|■■■■■■■■■■■■■■
  0.905 [704]	|■■■■■■■
  0.952 [82]	|■
  0.998 [35]	|
  1.044 [15]	|
  1.091 [50]	|
  1.137 [50]	|


Latency distribution:
  10% in 0.7180 secs
  25% in 0.7418 secs
  50% in 0.7614 secs
  75% in 0.8115 secs
  90% in 0.8549 secs
  95% in 0.8879 secs
  99% in 1.0556 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.6728 secs, 1.1374 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:	0.7800 secs, 0.6727 secs, 1.1374 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.2074 secs
  Slowest:	0.0063 secs
  Fastest:	0.0001 secs
  Average:	0.0010 secs
  Requests/sec:	48226.8792
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [1238]	|■■■■■■■
  0.001 [7380]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1126]	|■■■■■■
  0.003 [194]	|■
  0.003 [11]	|
  0.004 [0]	|
  0.004 [0]	|
  0.005 [11]	|
  0.006 [28]	|
  0.006 [11]	|


Latency distribution:
  10% in 0.0007 secs
  25% in 0.0008 secs
  50% in 0.0009 secs
  75% in 0.0011 secs
  90% in 0.0015 secs
  95% in 0.0017 secs
  99% in 0.0023 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0063 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0020 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:	0.0009 secs, 0.0001 secs, 0.0020 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0017 secs

Status code distribution:
  [200]	10000 responses



