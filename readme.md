Shared container server pipe test
=================================


### build 
```$powershell
go build -o .out/pipe-test.exe .
docker build -t pipe-test -f ./Dockerfile.win .
```

### run 

 ```$powershell
docker run --rm -ti -v //./pipe/pipe-test://./pipe/pipe-test pipe-test
```
    
### test results
if pipe was created it should print `done`, we see this result when we run it
as host process: 
```$powershell
PS .../pipe-test> .\.out\pipe-test.exe
done
```
but when we run it inside container we get error:
```$powershell
PS .../pipe-test> docker run --rm -ti -v //./pipe/pipe-test://./pipe/pipe-test pipe-test
open \\.\pipe\pipe-test: The filename, directory name, or volume label syntax is incorrect.
```

