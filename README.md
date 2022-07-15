# poc-client-side-caching

This is a CLI-based project that serves as poc for client-side caching implementation using Redis

There are two command that can be used:
Command Format | Note
-|-
`get <key>` | Get value of key `<key>`
`set <key> <value>` | Set value of `<value>` to key `<key>`

<br>

## Pre-requisites
- Go version >= 1.14.0
- Redis version >= 6.0
- Redis CLI

<br>

## Step-by-step
1. Run redis on your local machine

    ```
    redis-server
    ```

2. On another tab, run poc-CLI

    ```
    go run .
    ```

3. Try get value with key `key`. If key not exist, you will be notified as follows

    ```
    > get key
    2022/07/15 11:38:13 got no value from anywhere
    <nil>
    ```

4. On another tab, open Redis CLI. Try assign value to `key`

    ```
    ~ % redis-cli
    127.0.0.1:6379> set key value
    OK
    ```

5. Now `key` exist in Redis. Try get `key` using poc-CLI again. You will be notified that data retrieved from Redis

    ```
    > get key
    2022/07/15 11:43:43 got value from redis
    value
    ```

6. After retrieving data from Redis, data is stored on local memory. Get the value of `key` again. You will be notified that data retrieved from memory

    ```
    > get key
    2022/07/15 11:45:12 got value from memory cache
    value
    ```

7. Go back to Redis CLI tab. Try assigning new value to `key`

    ```
    127.0.0.1:6379> set key value2
    OK
    ```

8. Go back to poc-CLI tab. You will be notified that there are change in value for `key`, thus local memory cache for `key` is renewed

    ```
    > 
    refresh memory cache for key key
    key refreshed
    ```

9. If you try get `key` value again, you will notice that the value came from local memory cache. Local cache has been updated, so the data will be same as current data on Redis

    ```
    > get key
    2022/07/15 11:50:31 got value from memory cache
    value2
    ```

poc-CLI is also equipped with `set` command. You can try similar steps mentioned above by using two process of poc-CLI

<br>

## Related documents
- Redis client-side caching: https://redis.io/docs/manual/client-side-caching/
- Go package that support client-side caching: https://github.com/rueian/rueidis
- Practical example using redis-cli: https://www.linode.com/docs/guides/redis-client-side-caching/
