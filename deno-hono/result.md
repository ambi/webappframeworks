##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.1018 secs
  Slowest:	0.0121 secs
  Fastest:	0.0000 secs
  Average:	0.0005 secs
  Requests/sec:	98279.3739
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [9773]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [163]	|■
  0.004 [13]	|
  0.005 [0]	|
  0.006 [0]	|
  0.007 [0]	|
  0.008 [0]	|
  0.010 [0]	|
  0.011 [8]	|
  0.012 [42]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0004 secs
  75% in 0.0006 secs
  90% in 0.0008 secs
  95% in 0.0010 secs
  99% in 0.0018 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0121 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0067 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0011 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.3116 secs
  Slowest:	0.0269 secs
  Fastest:	0.0001 secs
  Average:	0.0016 secs
  Requests/sec:	32090.7526
  
  Total data:	307785 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.003 [9793]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [155]	|■
  0.008 [1]	|
  0.011 [0]	|
  0.013 [0]	|
  0.016 [0]	|
  0.019 [0]	|
  0.022 [0]	|
  0.024 [0]	|
  0.027 [50]	|


Latency distribution:
  10% in 0.0011 secs
  25% in 0.0011 secs
  50% in 0.0013 secs
  75% in 0.0016 secs
  90% in 0.0019 secs
  95% in 0.0022 secs
  99% in 0.0045 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0269 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:	0.0015 secs, 0.0001 secs, 0.0251 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.7316 secs
  Slowest:	0.0121 secs
  Fastest:	0.0057 secs
  Average:	0.0086 secs
  Requests/sec:	5774.9616
  

Response time histogram:
  0.006 [1]	|
  0.006 [2]	|
  0.007 [7]	|
  0.008 [10]	|
  0.008 [387]	|■■
  0.009 [8522]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [870]	|■■■■
  0.010 [34]	|
  0.011 [21]	|
  0.011 [128]	|■
  0.012 [18]	|


Latency distribution:
  10% in 0.0083 secs
  25% in 0.0085 secs
  50% in 0.0086 secs
  75% in 0.0087 secs
  90% in 0.0089 secs
  95% in 0.0090 secs
  99% in 0.0112 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0057 secs, 0.0121 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:	0.0086 secs, 0.0055 secs, 0.0117 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.4393 secs
  Slowest:	0.0060 secs
  Fastest:	0.0001 secs
  Average:	0.0022 secs
  Requests/sec:	22762.2891
  

Response time histogram:
  0.000 [1]	|
  0.001 [9]	|
  0.001 [16]	|
  0.002 [918]	|■■■■■
  0.002 [7695]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [1164]	|■■■■■■
  0.004 [111]	|■
  0.004 [54]	|
  0.005 [19]	|
  0.005 [7]	|
  0.006 [6]	|


Latency distribution:
  10% in 0.0019 secs
  25% in 0.0020 secs
  50% in 0.0022 secs
  75% in 0.0023 secs
  90% in 0.0025 secs
  95% in 0.0027 secs
  99% in 0.0035 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0060 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0008 secs
  resp wait:	0.0021 secs, 0.0001 secs, 0.0042 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	4.0181 secs
  Slowest:	0.0468 secs
  Fastest:	0.0165 secs
  Average:	0.0201 secs
  Requests/sec:	2488.7431
  

Response time histogram:
  0.016 [1]	|
  0.019 [4904]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.023 [4521]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.026 [235]	|■■
  0.029 [160]	|■
  0.032 [98]	|■
  0.035 [30]	|
  0.038 [1]	|
  0.041 [0]	|
  0.044 [7]	|
  0.047 [43]	|


Latency distribution:
  10% in 0.0187 secs
  25% in 0.0191 secs
  50% in 0.0195 secs
  75% in 0.0201 secs
  90% in 0.0214 secs
  95% in 0.0228 secs
  99% in 0.0311 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0165 secs, 0.0468 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0200 secs, 0.0164 secs, 0.0468 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1089 secs
  Slowest:	0.0242 secs
  Fastest:	0.0000 secs
  Average:	0.0005 secs
  Requests/sec:	91856.0786
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9848]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [53]	|
  0.007 [49]	|
  0.010 [1]	|
  0.012 [0]	|
  0.015 [0]	|
  0.017 [1]	|
  0.019 [1]	|
  0.022 [0]	|
  0.024 [46]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0004 secs
  90% in 0.0007 secs
  95% in 0.0010 secs
  99% in 0.0035 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0242 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0012 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0043 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0242 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0054 secs

Status code distribution:
  [200]	10000 responses



