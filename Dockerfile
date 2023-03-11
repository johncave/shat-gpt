FROM docker.io/golang:latest as go

# Build Go stuffs
WORKDIR /shatgpt/
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o ./shatgpt

FROM docker.io/node:latest as node
# Build the frontend
WORKDIR /shatgpt-frontend/
COPY ./shatgpt-frontend/ .
#RUN apt-get -y update && apt-get -y install npm
RUN npm install 
RUN npm run build  

FROM ubuntu:latest 
WORKDIR /shatgpt/

COPY --from=go /shatgpt/shatgpt ./shatgpt-bin
COPY --from=node /shatgpt-frontend/dist ./shatgpt-frontend/dist



EXPOSE 8080
CMD ["./shatgpt-bin"]