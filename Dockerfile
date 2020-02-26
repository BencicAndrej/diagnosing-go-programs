FROM scratch

COPY app_linux /app

ENTRYPOINT [ "/app" ]
