FROM golang:1.18

# Set destination for COPY
WORKDIR /app

ENV API_PORT=9000
ENV SHORT_URL_MAX_LEN=6
# online mongo (mongo atlas)
#ENV DB_URI=mongodb+srv://userTest:123password123@cluster0.er1tcsx.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0
# local mongo
# ENV DB_URI=mongodb://localhost:27017/
# docker-compose version
ENV DB_URI=mongodb://mongodb:27017/


# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Build
RUN go build -o /url-shortener-api

EXPOSE 9000
EXPOSE 27017

# Run
CMD ["/url-shortener-api"]