# split-pdf

## Descripción 
Microservicio en golang creado para splitear un pdf en varios archivos.

### golang 
#### Dependencies
```sh 
go mod init
go get github.com/pdfcpu/pdfcpu/cmd/pdfcpu@latest
``` 
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
