# Docker how to


## Backend

### Build
```bash
sudo docker build . -f Dockerfile.backend -t jonasbe25/simple-go-app-backend:latest
```

### Push
```bash
sudo docker push jonasbe25/simple-go-app-backend:latest
```

### Run 
```bash
sudo docker run -p 50051:50051 -e CONTENT=Hey jonasbe25/simple-go-app-backend
```


## Web

### Build
```bash
sudo docker build . -f Dockerfile.web -t jonasbe25/simple-go-app-web:latest
```
### Push
```bash
sudo docker push jonasbe25/simple-go-app-web:latest
```

### Run
```bash
sudo docker run -p 8080:8080 -e TARGET=<YOUR_HOST_IP>:50051 jonasbe25/simple-go-app-web
```