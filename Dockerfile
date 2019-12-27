FROM debian:9

WORKDIR /var/lib/writ
EXPOSE 8080

COPY ./.output/*.deb /etc/writ/packages/
RUN dpkg -i /etc/writ/packages/*.deb; \
    apt-get install -f;

CMD ["/usr/local/bin/writ"]