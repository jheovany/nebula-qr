# Imagen base
FROM golang:1.24.1-alpine

# Directorio de trabajo
WORKDIR /app

# Copiar go.mod y go.sum primero (para aprovechar cache)
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo el código
COPY . .

# Compilar la aplicación desde la carpeta cmd
RUN go build -o main ./cmd

# Puerto que expondrá el contenedor
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]