# BUILD STAGE
FROM golang:1.22-alpine AS build

# Adds certificates to make HTTP requests. It is essential if you'll need services like AWS Secrets Manager + STS.
RUN apk update && apk add --no-cache ca-certificates procps openssl tzdata

# Sets the working directory
WORKDIR /app

# Copy the project files to generate the application binary
COPY . .

# Generates the application binary
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main main.go

# Sets a non-root user for the image
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# RUN STAGE
FROM scratch

# Uses non-root user
COPY --from=build /etc/passwd /etc/passwd
USER appuser

# Sets the working directory
WORKDIR /app

# Copies the certificates installed on the BUILD STAGE
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copies the BUILD STAGE binary
COPY --from=build /app/main .

# Copies the relational database migrations
COPY --from=build /app/external/rdb/migration ./external/rdb/migration

# Copies the environment file. This is not required if environment is loaded from Cloud (e.g. AWS Secrets Manager)
COPY --from=build /app/.env .

# Copies the VERSION file, as the application uses it internally for tracking
COPY --from=build /app/VERSION .

# Sets the ENV variable to production
ENV ENV=production

# Starts the application
ENTRYPOINT [ "/app/main" ]