FROM golang:alpine

# Install dependency packages
RUN set -ex \
    && apk update \
    && apk add --no-cache ca-certificates git wget \
    && update-ca-certificates \
    && mkdir -p /go/src/openss/api \
    && git config --global http.https://gopkg.in.followRedirects true


#working directory in the container
WORKDIR /go/src/openss/api

# Install go libraries
RUN set -ex \
    && go get github.com/revel/cmd/revel \
    && go get gopkg.in/mgo.v2 \
    && go get github.com/kyawmyintthein/revel_mgo 

#Expose port to the host
EXPOSE 9000

#CMD ["revel", "run openss/api"]
ENTRYPOINT revel run openss/api prod 9000
