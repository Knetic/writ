writ
====

Accepts http requests for a markdown file, responding with the html equivalent of the file.
Useful as part of a larger ecosystem that serves static files. 

## API

* `/list` - recursively lists all MD files that can be served. Trims the MD suffix.
* `/*` - Any other path is checked to see if it's an MD file. If not, 404. If so, converted to HTML and returned.

## nginx

The given `nginx.conf` is an example of how to rewrite requests and proxy to `writ`.

```
location /writs/ {
  rewrite ^/writs/(.*) /$1  break;
  proxy_pass http://writ/$uri;
}
```

Assuming `writ` is a linked docker container running writ on 80.