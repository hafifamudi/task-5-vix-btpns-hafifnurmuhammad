FROM golang:alpine

# create working directory
# copy go.sum if exist
# and do cache if exist
WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . ./
CMD go run main.go