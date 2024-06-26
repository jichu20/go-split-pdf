# Usamos una imagen base de Go ligera
FROM golang:alpine as builder

ARG VERSION="development"
ARG BUILD="1999-01-01T00:00:00Z"
# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos los archivos del código fuente al contenedor
COPY . .

# Compilamos el binario
# RUN go build -o microservicio cmd/api/main.go -ldflags="-X 'main.Version=v1.0.0'"
# RUN go build -ldflags="-X 'main.Version=${VERSION}' -X main.Build=${BUILD}" -o ${BINARY} cmd/api/main.go 
RUN go build -ldflags="-X 'main.Version=${VERSION}' -X main.Build=${BUILD}" -o microservicio cmd/api/main.go 

# Usamos alpine para mantener nuestro contenedor lo más ligero posible
# FROM alpine:latest
FROM alpine:edge

WORKDIR /app/

# Copiamos el binario compilado desde el primer paso
COPY --from=builder /app/microservicio .
COPY static/ /app/static/

ENV STATIC_PATH=/app/static/
ENV TEMP_PATH=/tmp

# Exponemos el puerto en el que estamos escuchando
EXPOSE 8080

# Ejecutamos el microservicio
CMD ["./microservicio"]