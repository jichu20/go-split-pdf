# split-pdf

## Descripción 
Microservicio en golang creado para splitear un pdf en varios archivos.

### Docker 

#### Build

Para hacer la build con docker 

```
docker build --platform linux/amd64 -t mi-microservicio .
```

#### Run
Para hacer el run del contenedor

```sh
# Run with local build
docker run --rm -p 8080:8080 --name mi-servicio mi-microservicio

# Run with remote build
docker run --rm -p 8080:8080 --name mi-servicio jichu20/go-split-pdf


```
