#the app uses the version 1.17.3
FROM golang:1.17.3 

#About the app
LABEL project="ascii-art-web-dockerize"
LABEL authors="Zhadyra21, Seitimova"
LABEL description="My first project that uses Docker"

#Workspace in docker
WORKDIR /ascii-art

#Copying from local ascii-art-web-dockerize to /ascii-art in docker
COPY . .
RUN go build -o /server
#Chosen port in docker
EXPOSE 8080
#The command to execute when our image is used to start the container
CMD [ "/server" ]