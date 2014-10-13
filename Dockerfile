#
# Go Dockerfile
#
# https://github.com/dockerfile/go
#

# Pull base image.
FROM dockerfile/ubuntu

MAINTAINER Aleksandr Guljajev <gulyayev.alex@gmail.com>
# Install Go
RUN \
  mkdir -p /goroot && \
  curl https://storage.googleapis.com/golang/go1.3.1.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1


# Set environment variables.
ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

# Define working directory.
WORKDIR /gopath

ENV APP_DIR $GOPATH/src/github.com/alex-glv/modelprovider
RUN go get github.com/alex-glv/modelprovider
RUN cd $APP_DIR && \
    go install

EXPOSE 8912
# Define default command.
CMD ["bash"]
