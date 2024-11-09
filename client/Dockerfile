FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y curl

WORKDIR /app

COPY request.sh .

RUN chmod +x request.sh

ENTRYPOINT ["./request.sh"]