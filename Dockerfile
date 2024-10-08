ARG GO_VERSION="1.22"
ARG USER_NAME="nonroot"
ARG AMSA_MARIWEB_USERNAME="set a build argument or .env file"
ARG AMSA_MARIWEB_UPASSWORD="set a build argument or .env file"
ARG AMSA_MARIWEB_URI="set a build argument or .env file"
ARG AISSTREAM_BOUNDARY_LAT1="set a build argument or .env file"
ARG AISSTREAM_BOUNDARY_LON1="set a build argument or .env file"
ARG AISSTREAM_BOUNDARY_LAT2="set a build argument or .env file"
ARG AISSTREAM_BOUNDARY_LON2="set a build argument or .env file"
ARG AISSTREAM_BOUNDARY_QLD_LAT1="set a build argument or .env file"
ARG AISSTREAM_BOUNDARY_QLD_LON1="set a build argument or .env file"
ARG AISSTREAM_BOUNDARY_QLD_LAT2="set a build argument or .env file"
ARG AISSTREAM_BOUNDARY_QLD_LON2="set a build argument or .env file"
ARG AISSTREAM_RETRY_SECS="set a build argument or .env file"
ARG AISSTREAM_URI="set a build argument or .env file"
ARG AISSTREAM_API_KEY="set a build argument or .env file"
ARG AISHUB_RETRY_SECS="set a build argument or .env file"
ARG AISHUB_TIMEOUT_SECS="set a build argument or .env file"
ARG AISHUB_URI="set a build argument or .env file"
ARG AMSA_RETRY_SECS="set a build argument or .env file"
ARG AMSA_TIMEOUT_SECS="set a build argument or .env file"
ARG AMSA_URI="set a build argument or .env file"
ARG AISLOCAL_RETRY_SECS="set a build argument or .env file"
ARG AISLOCAL_TIMEOUT_SECS="set a build argument or .env file"
ARG AISLOCAL_URI="set a build argument or .env file"
ARG UDP_RETRY_SECS="set a build argument or .env file"
ARG UDP_TIMEOUT_SECS="set a build argument or .env file"
ARG UDP_URI="set a build argument or .env file"
ARG LOCAL_MONGODB_CONNECTION="set a build argument or .env file"
ARG ATLAS_MONGODB_CONNECTION="set a build argument or .env file"

# build stage
FROM golang:${GO_VERSION}-alpine AS builder

# carry forward ARGS
ARG USER_NAME
ARG AMSA_MARIWEB_USERNAME
ARG AMSA_MARIWEB_UPASSWORD
ARG AMSA_MARIWEB_URI
ARG AISSTREAM_BOUNDARY_LAT1
ARG AISSTREAM_BOUNDARY_LON1
ARG AISSTREAM_BOUNDARY_LAT2
ARG AISSTREAM_BOUNDARY_LON2
ARG AISSTREAM_BOUNDARY_QLD_LAT1
ARG AISSTREAM_BOUNDARY_QLD_LON1
ARG AISSTREAM_BOUNDARY_QLD_LAT2
ARG AISSTREAM_BOUNDARY_QLD_LON2
ARG AISSTREAM_RETRY_SECS
ARG AISSTREAM_URI
ARG AISSTREAM_API_KEY
ARG AISHUB_RETRY_SECS
ARG AISHUB_TIMEOUT_SECS
ARG AISHUB_URI
ARG AMSA_RETRY_SECS
ARG AMSA_TIMEOUT_SECS
ARG AMSA_URI
ARG AISLOCAL_RETRY_SECS
ARG AISLOCAL_TIMEOUT_SECS
ARG AISLOCAL_URI
ARG UDP_RETRY_SECS
ARG UDP_TIMEOUT_SECS
ARG UDP_URI
ARG LOCAL_MONGODB_CONNECTION
ARG ATLAS_MONGODB_CONNECTION

# add dependencies and certs
RUN apk update 
RUN apk add --no-cache git
RUN apk add -U --no-cache ca-certificates

# add a user here because addgroup and adduser are not available in scratch
RUN addgroup -S ${USER_NAME} && adduser -S -u 10000 -g ${USER_NAME} ${USER_NAME}

# access for private repos
# ENV CGO_ENABLED=0 
# ENV GOPRIVATE=github.com/cetf/*
# RUN git config --global url."https://${GH_ACTION_TOKEN}@github.com/".insteadOf "https://github.com/"

# get the dependencies
WORKDIR /go/src/app
COPY ./go.mod ./go.sum ./
RUN go mod download -x
COPY . .

# run tests (if any)
#RUN CGO_ENABLED=0 go test -timeout 60s -v github.com/cetf/agent99
 
# build the executable
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /go/bin/app -v 
 
# final stage
FROM scratch AS final

# carry forward ARGS
ARG USER_NAME
ARG AMSA_MARIWEB_USERNAME
ARG AMSA_MARIWEB_UPASSWORD
ARG AMSA_MARIWEB_URI
ARG AISSTREAM_BOUNDARY_LAT1
ARG AISSTREAM_BOUNDARY_LON1
ARG AISSTREAM_BOUNDARY_LAT2
ARG AISSTREAM_BOUNDARY_LON2
ARG AISSTREAM_BOUNDARY_QLD_LAT1
ARG AISSTREAM_BOUNDARY_QLD_LON1
ARG AISSTREAM_BOUNDARY_QLD_LAT2
ARG AISSTREAM_BOUNDARY_QLD_LON2
ARG AISSTREAM_RETRY_SECS
ARG AISSTREAM_URI
ARG AISSTREAM_API_KEY
ARG AISHUB_RETRY_SECS
ARG AISHUB_TIMEOUT_SECS
ARG AISHUB_URI
ARG AMSA_RETRY_SECS
ARG AMSA_TIMEOUT_SECS
ARG AMSA_URI
ARG AISLOCAL_RETRY_SECS
ARG AISLOCAL_TIMEOUT_SECS
ARG AISLOCAL_URI
ARG UDP_RETRY_SECS
ARG UDP_TIMEOUT_SECS
ARG UDP_URI
ARG LOCAL_MONGODB_CONNECTION
ARG ATLAS_MONGODB_CONNECTION

ENV AMSA_MARIWEB_USERNAME=${AMSA_MARIWEB_USERNAME}
ENV AMSA_MARIWEB_UPASSWORD=${AMSA_MARIWEB_UPASSWORD}
ENV AMSA_MARIWEB_URI=${AMSA_MARIWEB_URI}
ENV AISSTREAM_BOUNDARY_LAT1=${AISSTREAM_BOUNDARY_LAT1}
ENV AISSTREAM_BOUNDARY_LON1=${AISSTREAM_BOUNDARY_LON1}
ENV AISSTREAM_BOUNDARY_LAT2=${AISSTREAM_BOUNDARY_LAT2}
ENV AISSTREAM_BOUNDARY_LON2=${AISSTREAM_BOUNDARY_LON2}
ENV AISSTREAM_BOUNDARY_QLD_LAT1=${AISSTREAM_BOUNDARY_QLD_LAT1}
ENV AISSTREAM_BOUNDARY_QLD_LON1=${AISSTREAM_BOUNDARY_QLD_LON1}
ENV AISSTREAM_BOUNDARY_QLD_LAT2=${AISSTREAM_BOUNDARY_QLD_LAT2}
ENV AISSTREAM_BOUNDARY_QLD_LON2=${AISSTREAM_BOUNDARY_QLD_LON2}
ENV AISSTREAM_RETRY_SECS=${AISSTREAM_RETRY_SECS}
ENV AISSTREAM_URI=${AISSTREAM_URI}
ENV AISSTREAM_API_KEY=${AISSTREAM_API_KEY}
ENV AISHUB_RETRY_SECS=${AISHUB_RETRY_SECS}
ENV AISHUB_TIMEOUT_SECS=${AISHUB_TIMEOUT_SECS}
ENV AISHUB_URI=${AISHUB_URI}
ENV AMSA_RETRY_SECS=${AMSA_RETRY_SECS}
ENV AMSA_TIMEOUT_SECS=${AMSA_TIMEOUT_SECS}
ENV AMSA_URI=${AMSA_URI}
ENV AISLOCAL_RETRY_SECS=${AISLOCAL_RETRY_SECS}
ENV AISLOCAL_TIMEOUT_SECS=${AISLOCAL_TIMEOUT_SECS}
ENV AISLOCAL_URI=${AISLOCAL_URI}
ENV UDP_RETRY_SECS=${UDP_RETRY_SECS}
ENV UDP_TIMEOUT_SECS=${UDP_TIMEOUT_SECS}
ENV UDP_URI=${UDP_URI}
ENV LOCAL_MONGODB_CONNECTION=${LOCAL_MONGODB_CONNECTION}
ENV ATLAS_MONGODB_CONNECTION=${ATLAS_MONGODB_CONNECTION}

COPY --from=builder /go/bin/app /app
# COPY --from=builder /go/src/app/config.json /config.json
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

LABEL maintainer="github.com/dhskinner"
USER ${USER_NAME}

# EXPOSE 8080/tcp
# EXPOSE 3500/tcp
ENTRYPOINT ["/app", "run"]
