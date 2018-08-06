FROM scratch
ADD vue-go-stormdb /
EXPOSE 8080
CMD ["/vue-go-stormdb"]