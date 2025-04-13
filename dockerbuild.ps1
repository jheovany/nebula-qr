# Compila la imagen
docker build -t nebula-qr .

# Despliega la imagen
docker run -p 8080:8080 -e MONGODB_URI="mongodb+srv://<username>:<password>@<cluster-url>" nebula-qr