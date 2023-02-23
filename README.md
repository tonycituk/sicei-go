# sicei-go

gin-gonic rest api coded with go, powered with docker.


### Build the app

```bash
docker build -t sicei-go:0.0.1 .
```

### Run it
```bash
docker run -d -p 80:80 sicei-go:0.0.1
```

### Test it
```bash
curl localhost:80/students
```


![http response](./jue%2023%20feb%202023%2016%3A48%3A09%20CST.png)