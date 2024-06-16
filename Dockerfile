# Usamos una imagen base de Go ligera
FROM golang:alpine as builder

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos los archivos del código fuente al contenedor
COPY . .

# Compilamos el binario
RUN go build -o microservicio cmd/api/main.go

# Usamos alpine para mantener nuestro contenedor lo más ligero posible
FROM alpine:latest

WORKDIR /root/

# Copiamos el binario compilado desde el primer paso
COPY --from=builder /app/microservicio .

# Exponemos el puerto en el que estamos escuchando
EXPOSE 8080

# Ejecutamos el microservicio
CMD ["./microservicio"]