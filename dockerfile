FROM mcr.microsoft.com/playwright:bionic as base

FROM golang:1.15 as builder
WORKDIR /go/src/app
COPY . .
RUN chmod +x ./scripts/build.sh && ./scripts/build.sh

FROM base as production
COPY --from=builder /go/src/app/fay .
EXPOSE 8080
CMD ["./fay"]