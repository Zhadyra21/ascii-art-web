echo "Building image"
docker build -t ascii-art-app .
echo "Creating a container from the ascii-art-app image"
docker run -p 8080:8080 -d --rm --name ascii-art ascii-art-app
echo "Running container list"
docker ps
echo "Image list"
docker image ls
echo "Container list"
docker ps -a