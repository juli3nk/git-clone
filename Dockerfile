FROM golang:alpine AS builder

RUN apk --update add \
		ca-certificates \
		gcc \
		git \
		musl-dev

RUN echo 'nobody:x:65534:65534:nobody:/:' > /tmp/passwd \
	&& echo 'nobody:x:65534:' > /tmp/group

WORKDIR /go/src/github.com/juli3nk/git-clone

COPY . .

ENV GO111MODULE=off
RUN go get
RUN go build -ldflags "-linkmode external -extldflags -static -s -w" -o /tmp/git-clone


FROM scratch

COPY --from=builder /tmp/group /tmp/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /tmp/git-clone /git-clone

USER nobody:nobody

ENTRYPOINT ["/git-clone"]
