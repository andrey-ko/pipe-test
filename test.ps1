param(
)
$errorActionPreference = "stop"
write-host "building image...."
push-location $psScriptRoot
go build -o .out/pipe-test.exe .
docker build -t pipe-test -f ./Dockerfile.win .
pop-location
write-host "running container...."
docker run  --rm -ti -v "//./pipe/pipe-test://./pipe/pipe-test" pipe-test