writstatic
====

The static content used as a reference consumption of the `writ` server.

## usage

The reference `docker-compose.yml` in the above folder shows how this is usually used. You can use the given `../reloadStatic.sh` as a oneshot to rebuild/relaunch the container when you make changes to the static content.

## nginx

While it's not required, this project assumes the given `nginx.conf` is used. It serves static content as expected, and forwards `/list` and all `/f/` traffic to `writ`.

## display

The client itself mostly exists to just show a navbar with the contents of `/list`, allow selection of each item (which in turn makes requests to `/f/` for that specific post), and style the markdown in a tasteful way.

It also supports direct linking to a specific post using `/a/` and the post name. This causes the client to load that content on-load.

## isn't this bad for SEO?

You'll probably notice that this makes things like the "About" page, and each post, a little light on SEO. It requires external content to be filled in later, which probably decreases the likelihood that any serious crawler will give it as much attention as it needs.

In the future, all this static content might get served directly by `writ` itself, and individual pages are full _pages_ (not a single-page application) that are served and crawl-able as usual.

Let's get real though, I wrote it this way, and I like it this way, and all major crawlers should be able to figure out how it works by the MIME type of the responses.