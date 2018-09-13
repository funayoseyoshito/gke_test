FROM golang:1.9.1

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

# CMD go-wrapper run && tail -f /dev/null
CMD ["go-wrapper", "run"] # ["app"]
# ENTRYPOINT /go/bin/app