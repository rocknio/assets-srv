FROM alpine
ADD asset-srv /asset-srv
ENTRYPOINT [ "/asset-srv" ]
