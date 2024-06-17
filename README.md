# split-pdf

## Descripción 
Microservicio en golang creado para splitear un pdf en varios archivos.

### golang 
#### Dependencies
```sh 
go mod init
# Gestion de pdf
go get github.com/pdfcpu/pdfcpu/cmd/pdfcpu@latest 
# Gestión de healtcheck 
go get -u github.com/heptiolabs/healthcheck
``` 
### Docker 

#### Build

Para hacer la build con docker 

```sh
docker build --platform linux/amd64 --build-arg VERSION=any --build-arg BUILD=`date +%FT%T%z` -t mi-microservicio .
```

#### Run
Para hacer el run del contenedor

```sh
# Run with local build
docker run --rm -p 8080:8080 --name mi-servicio mi-microservicio

# Run with remote build
docker run --rm -p 8080:8080 --name mi-servicio jichu20/go-split-pdf


```
