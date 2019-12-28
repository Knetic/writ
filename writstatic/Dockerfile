FROM nginx

RUN mkdir -p /tmp/nginx/cache/list; \
	mkdir -p /tmp/nginx/cache/post;
	
COPY index.html /usr/share/writ/index.html
COPY about.html /usr/share/writ/about.html
COPY src/ /usr/share/writ/src
COPY css/ /usr/share/writ/css
COPY nginx.conf /etc/nginx/nginx.conf
