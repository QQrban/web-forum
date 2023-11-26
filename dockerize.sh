#!/bin/bash
./docker-clean.sh
echo ""
echo "Building new docker image 'forum-docker'..."
sudo docker image build -f Dockerfile -t forum-docker .
echo "...built docker image and proceeding to check if there exists container named 'forum'"
echo "Deploying the container..."
sudo docker container run -p 8080:8080 --detach --name forum forum-docker
echo ""
echo "Please open browser and go to http://localhost:8080"
echo ""
