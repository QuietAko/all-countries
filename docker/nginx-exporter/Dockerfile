FROM alpine:latest

# Устанавливаем bash и другие необходимые утилиты
RUN apk add --no-cache bash curl

# Скачиваем nginx-prometheus-exporter
ARG EXPORTER_VERSION=0.11.0  # Укажи нужную версию
RUN curl -L https://github.com/nginxinc/nginx-prometheus-exporter/releases/download/v${EXPORTER_VERSION}/nginx-prometheus-exporter_${EXPORTER_VERSION}_linux_amd64.tar.gz -o /tmp/exporter.tar.gz && \
    tar -xzf /tmp/exporter.tar.gz -C /usr/local/bin/ --strip-components=1 && \
    rm /tmp/exporter.tar.gz

# Указываем команду для запуска экспортера
CMD ["nginx-prometheus-exporter", "--nginx.scrape-uri=http://nginx:80/nginx-status"]