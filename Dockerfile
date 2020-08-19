FROM alpine:3.7
RUN mkdir /app
COPY ./app-exe /app/
WORKDIR /app
ARG env
ARG config
ARG version
ENV env=${env}
ENV version=${version}
ENV config=${config}
CMD ["/app/app-exe"]