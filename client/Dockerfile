FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y curl bash dos2unix

WORKDIR /app

COPY request.sh .

RUN dos2unix request.sh
RUN chmod +x request.sh

ENTRYPOINT ["./request.sh"]