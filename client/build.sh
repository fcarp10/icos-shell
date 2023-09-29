docker build . -t shell-client
docker create --name builder shell-client
docker cp builder:/shell-client ./
docker rm -f builder
