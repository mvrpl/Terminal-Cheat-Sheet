#### GENERAL
|||
|-|-|
`docker rm $(docker ps -a -q)`|Remove all standing containers.
`docker stop $(docker ps -a -q)`|Stop all containers.
`docker ps`|List online containers.
`docker export <CONTAINER ID> > /home/export.tar`|Export container.
`docker cp <CONTAINER ID>:/<NOME ARQUIVO>.txt <NOME ARQUIVO>.txt`|Copy of container to place.
`docker load -i <IMAGEM>_image.docker`|Matter image docker.
`docker save -o <IMAGEM>_image.docker <IMAGEM>`|Export image docker.
`docker inspect --format '{{ .NetworkSettings.IPAddress }}' <CONTAINER ID>`|Get the internal ip container.
`docker-machine ip`|Get ip docker.
`docker ps -a`|List containers off-line.
`docker cp <NOME ARQUIVO>.txt <CONTAINER ID>:/<NOME ARQUIVO>.txt`|Local copy into the container.
`docker exec -u root -it <CONTAINER ID> bash`|Access container as root user.
`docker run --rm -it centos:centos6 /bin/bash -c "pwd;ls"`|Run command inside container.
`docker rmi -f $(docker images -a -q)`|Remove all images
`docker exec -it <CONTAINER ID> bash`|Open shell existing container.
`cat /home/export.tar | docker import - <NOVO NOME IMAGEM>:latest`|Import container (such as new image).
`docker run -i -t --rm <NOME IMAGEM> bash`|Run container and delete after output.
