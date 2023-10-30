docker build . -t shell-client
docker create --name builder shell-client
docker cp builder:/shell-client ./
docker rm -f builder
scp ./shell-client office:/opt/development/shell/client/
echo "testing specific deployment get"
ssh office "/opt/development/shell/client/shell-client --config /opt/development/shell/client/config.yml get deployment --id 123"
echo "testing all deployments get"
ssh office "/opt/development/shell/client/shell-client --config /opt/development/shell/client/config.yml get deployment"
