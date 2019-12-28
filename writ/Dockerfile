FROM debian:9

WORKDIR /var/lib/writ
EXPOSE 8080

COPY ./.output/*.deb /etc/writ/packages/
RUN dpkg -i /etc/writ/packages/*.deb; \
    apt-get install -f;

VOLUME /usr/share/writ
WORKDIR /usr/share/writ

CMD ["/usr/local/bin/writ"]