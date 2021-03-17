# watch

[![Go](https://github.com/li-zhixin/watch/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/li-zhixin/watch/actions/workflows/go.yml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/li-zhixin/watch)

print start time and end time when run command.

## Policy

- Time location is UNC.

- Time format is RFC3339.

## Example

```shell
watch ping baidu.com
```

output:
```shell

start time: 2021-03-17T14:10:22Z

Pinging baidu.com [220.181.38.148] with 32 bytes of data:
Reply from 220.181.38.148: bytes=32 time=29ms TTL=49
Reply from 220.181.38.148: bytes=32 time=28ms TTL=49
Reply from 220.181.38.148: bytes=32 time=29ms TTL=49
Reply from 220.181.38.148: bytes=32 time=29ms TTL=49

Ping statistics for 220.181.38.148:
    Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
    Minimum = 28ms, Maximum = 29ms, Average = 28ms
end time: 2021-03-17T14:10:26Z

```
