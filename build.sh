docker stop todo-backend
docker rm todo-backend
docker rmi todo-backend
git pull
docker build -t todo-backend .
sh start.sh