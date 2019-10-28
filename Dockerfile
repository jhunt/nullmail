FROM golang AS stage1

COPY . /go/src/github.com/jhunt/nullmail
WORKDIR /go/src/github.com/jhunt/nullmail
RUN CGO_ENABLED=0 go build

FROM scratch
COPY --from=stage1 /go/src/github.com/jhunt/nullmail/nullmail /nullmail
ENV PORT=25
CMD ["/nullmail"]
