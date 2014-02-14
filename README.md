shovel
======

Append outputs of programs into SQL tables

###Usage

```
Usage of ./shovel:
  -database="shovel": <dbname>
  -host="localhost:3306": <hostname>:<port>
  -pass="": <dbpass>
  -tablename="": <tablename> else it will make a new one
  -user="root": <dbuser>

```

###Example

```
ben@daring:~/shovel$ ping 8.8.8.8 | shovel
[Shovel]14:10:07 Connecting to DB
[Shovel]14:10:07 Logging line by line into Shovel_1392387007.

ben@daring:~/shovel$ echo "SELECT * FROM Shovel_1392387007 LIMIT 5" | mysql shovel
id      line    logtime
1       PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.    2014-02-14 14:10:07
2       64 bytes from 8.8.8.8: icmp_req=1 ttl=49 time=5.92 ms   2014-02-14 14:10:07
3       64 bytes from 8.8.8.8: icmp_req=2 ttl=49 time=5.99 ms   2014-02-14 14:10:08
4       64 bytes from 8.8.8.8: icmp_req=3 ttl=49 time=5.92 ms   2014-02-14 14:10:09
5       64 bytes from 8.8.8.8: icmp_req=4 ttl=49 time=5.98 ms   2014-02-14 14:10:10

```
