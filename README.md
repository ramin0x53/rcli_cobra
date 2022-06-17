# RCLI

Easy to use Redis CLI that can be used from the command line mode.

## Build:
```
git clone https://github.com/ramin0x53/rcli.git
cd ./rcli
make build
```
## Docker:
```
docker build -t rcli-img .
docker run -d --name rcli-test rcli-img 
docker exec -it rcli-test bash
rcli config localhost:6379
```

## Supported commands
```
  config      Set ip address and port number for redis server
  set         This command sets the value at the specified key
  del         This command deletes the key, if it exists
  get         Gets the value of a key
  hdel        Deletes one or more hash fields
  help        Help about any command
  hget        Gets the value of a hash field stored at the specified key
  hgetall     Gets all the fields and values stored in a hash at the specified key
  hmset       Sets multiple hash fields to multiple values
  lpop        Removes and gets the first element in a list
  lpush       Prepends a value to a list
  lrange      Gets a range of elements from a list
```