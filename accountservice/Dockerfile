FROM iron/base

ADD accountservice-linux-amd64 /
ADD healthchecker-linux-amd64 /

EXPOSE 3000
HEALTHCHECK --interval=5s --timeout=5s CMD ["./healthchecker-linux-amd64", "-port=3000"] || exit 1
ENTRYPOINT ["./accountservice-linux-amd64"]