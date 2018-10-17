Pour lancer le serveur via le docker file : 
sudo docker build -t server . && 
sudo docker run --net=host -p 8000:8000 -t server
