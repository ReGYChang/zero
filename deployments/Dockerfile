FROM golang:latest

WORKDIR $GOPATH/src/zero

# Instal package
RUN set -x \
  && apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y --no-install-recommends \
  unzip

# Set Timezone
ARG TZ=UTC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# User Aliases
COPY aliases.sh /root/aliases.sh
RUN sed -i 's/\r//g' /root/aliases.sh && \
  echo "" >> ~/.bashrc && \
  echo "# Load Custom Aliases" >> ~/.bashrc && \
  echo "source /root/aliases.sh" >> ~/.bashrc && \
  echo "" >> ~/.bashrc

# Install Staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

# Install Protocol Buffer Compiler
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-linux-x86_64.zip -o /tmp/protoc.zip && \
  unzip -o /tmp/protoc.zip -d /usr/local bin/protoc && \
  unzip -o /tmp/protoc.zip -d /usr/local include/*

# Install the Protocol Buffer Compiler plugins for Go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Clean up
RUN apt-get clean \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
