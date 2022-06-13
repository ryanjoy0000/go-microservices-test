# base go image
FROM golang:1.18 AS builder

ENV APP_HOME "/go/src/app"
ENV APP_MAIN "${APP_HOME}/cmd/api"
ENV APP_OUT "${APP_HOME}/brokerApp"

RUN mkdir "$APP_HOME"

ADD . "$APP_HOME"

# Set working dir
WORKDIR "$APP_HOME"

# CGO_ENABLED environment variable 
# CGO_ENABLED=0: ideal for scratch docker image deployments
# as no host OS needs to be bundled and you get a staticaly-linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ${APP_OUT} ${APP_MAIN}

# enable the permission to execute
RUN chmod +x ${APP_OUT}

# ---------------------------------

# # Create another tiny docker image
FROM alpine:latest

RUN mkdir /app

# # Copy folder from previous image
COPY --from=builder /go/src/app/brokerApp /app

EXPOSE 80

# # Run executable from new image
CMD [ "/app/brokerApp" ]