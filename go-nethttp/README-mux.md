##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.2236 secs
  Slowest:	0.0109 secs
  Fastest:	0.0001 secs
  Average:	0.0011 secs
  Requests/sec:	44714.1689

  Total data:	280000 bytes
  Size/request:	28 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [7200]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1547]	|■■■■■■■■■
  0.003 [751]	|■■■■
  0.004 [291]	|■■
  0.005 [107]	|■
  0.007 [53]	|
  0.008 [34]	|
  0.009 [13]	|
  0.010 [1]	|
  0.011 [2]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0004 secs
  50% in 0.0006 secs
  75% in 0.0013 secs
  90% in 0.0026 secs
  95% in 0.0033 secs
  99% in 0.0055 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0109 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0015 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0047 secs
  resp wait:	0.0008 secs, 0.0001 secs, 0.0076 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0042 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.8278 secs
  Slowest:	0.1327 secs
  Fastest:	0.0003 secs
  Average:	0.0040 secs
  Requests/sec:	12079.8784

  Total data:	317829 bytes
  Size/request:	31 bytes

Response time histogram:
  0.000 [1]	|
  0.014 [9205]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.027 [786]	|■■■
  0.040 [0]	|
  0.053 [1]	|
  0.067 [0]	|
  0.080 [0]	|
  0.093 [0]	|
  0.106 [0]	|
  0.119 [0]	|
  0.133 [7]	|


Latency distribution:
  10% in 0.0014 secs
  25% in 0.0020 secs
  50% in 0.0026 secs
  75% in 0.0034 secs
  90% in 0.0055 secs
  95% in 0.0187 secs
  99% in 0.0205 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.1327 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0041 secs
  resp wait:	0.0039 secs, 0.0002 secs, 0.1327 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0015 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	4.7728 secs
  Slowest:	0.5051 secs
  Fastest:	0.0011 secs
  Average:	0.0229 secs
  Requests/sec:	2095.2059

  Total data:	3198066 bytes
  Size/request:	319 bytes

Response time histogram:
  0.001 [1]	|
  0.051 [9497]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.102 [239]	|■
  0.152 [155]	|■
  0.203 [74]	|
  0.253 [5]	|
  0.303 [13]	|
  0.354 [3]	|
  0.404 [5]	|
  0.455 [7]	|
  0.505 [1]	|


Latency distribution:
  10% in 0.0040 secs
  25% in 0.0091 secs
  50% in 0.0147 secs
  75% in 0.0291 secs
  90% in 0.0380 secs
  95% in 0.0516 secs
  99% in 0.1541 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0011 secs, 0.5051 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0015 secs
  resp wait:	0.0229 secs, 0.0011 secs, 0.5050 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	1.4782 secs
  Slowest:	0.1554 secs
  Fastest:	0.0004 secs
  Average:	0.0070 secs
  Requests/sec:	6764.8400

  Total data:	19750000 bytes
  Size/request:	1975 bytes

Response time histogram:
  0.000 [1]	|
  0.016 [8529]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.031 [1276]	|■■■■■■
  0.047 [104]	|
  0.062 [68]	|
  0.078 [5]	|
  0.093 [5]	|
  0.109 [6]	|
  0.124 [0]	|
  0.140 [2]	|
  0.155 [4]	|


Latency distribution:
  10% in 0.0015 secs
  25% in 0.0029 secs
  50% in 0.0041 secs
  75% in 0.0061 secs
  90% in 0.0200 secs
  95% in 0.0229 secs
  99% in 0.0411 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.1554 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0029 secs
  resp wait:	0.0069 secs, 0.0003 secs, 0.1553 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0171 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	17.3289 secs
  Slowest:	0.2593 secs
  Fastest:	0.0109 secs
  Average:	0.0860 secs
  Requests/sec:	577.0694

  Total data:	3197708 bytes
  Size/request:	319 bytes

Response time histogram:
  0.011 [1]	|
  0.036 [9]	|
  0.061 [1984]	|■■■■■■■■■■■■■■■■■■■■■
  0.085 [3863]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.110 [2628]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.135 [639]	|■■■■■■■
  0.160 [537]	|■■■■■■
  0.185 [172]	|■■
  0.210 [105]	|■
  0.234 [23]	|
  0.259 [39]	|


Latency distribution:
  10% in 0.0564 secs
  25% in 0.0626 secs
  50% in 0.0789 secs
  75% in 0.0971 secs
  90% in 0.1282 secs
  95% in 0.1517 secs
  99% in 0.1956 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0109 secs, 0.2593 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0032 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:	0.0859 secs, 0.0108 secs, 0.2592 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1781 secs
  Slowest:	0.0113 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	56150.1428

  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8098]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1359]	|■■■■■■■
  0.003 [374]	|■■
  0.005 [80]	|
  0.006 [27]	|
  0.007 [4]	|
  0.008 [7]	|
  0.009 [42]	|
  0.010 [6]	|
  0.011 [2]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0005 secs
  75% in 0.0010 secs
  90% in 0.0018 secs
  95% in 0.0024 secs
  99% in 0.0043 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0113 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0025 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0068 secs
  resp wait:	0.0006 secs, 0.0001 secs, 0.0069 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0064 secs

Status code distribution:
  [200]	10000 responses



