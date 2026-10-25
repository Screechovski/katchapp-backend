# --- build ---
FROM golang:1.25-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/app .

# --- runtime ---
FROM alpine:3.22
WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=build /out/app .
# сид-картинки: при первом создании named volume compose наполнит его этим содержимым
COPY --from=build /src/images ./images

EXPOSE 8088
CMD ["./app"]
