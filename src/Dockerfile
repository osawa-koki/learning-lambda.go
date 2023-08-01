FROM public.ecr.aws/bitnami/golang:1.19 as build-image
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ../bin

FROM public.ecr.aws/lambda/go:1
COPY .env ./
COPY --from=build-image /go/bin/ /var/task/
CMD [ "program" ]
