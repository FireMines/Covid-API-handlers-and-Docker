FROM golang:1.17 as builder

LABEL maintainer="larsbho@stud.ntnu.no"
LABEL stage=builder

WORKDIR /go/src/app/cmd

# Copy relevant folders into container
COPY ./cases /go/src/app/cases
COPY ./cmd /go/src/app/cmd
COPY ./handler /go/src/app/handler
COPY ./notifications /go/src/app/notifications
COPY ./policy /go/src/app/policy
COPY ./status /go/src/app/status
COPY ./variables /go/src/app/variables
#COPY ./firebase-key.json /go/src/app/firebase-key.json
COPY ./go.sum /go/src/app/go.sum
COPY ./go.mod /go/src/app/go.mod

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server

# Instantiate binary
CMD ["./server"]
