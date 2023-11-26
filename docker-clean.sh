#!/bin/bash
echo "Is there container named 'forum'?"
result=$( sudo docker ps -q -a -f name=forum )
if [ "$result" = "" ]; then
echo "...container does not exist"
else
echo "...container exists, is it running?"
result=$( sudo docker ps -q -f name=forum )
if [ "$result" = "" ]; then
echo "   ...no, container is not running"
else
echo "   ...yes, let's stop it"
sudo docker stop forum
fi
echo "...deleting the container"
sudo docker container rm -f forum
fi
echo "Is there image named 'forum-docker'?"
result=$( sudo docker images -q forum-docker )
if [[ -n "$result" ]]; then
echo "...image exists, deleting the image"
sudo docker rmi -f forum-docker
else
echo "...image does not exist"
fi

