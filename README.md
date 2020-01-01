writ
====

A straightforward blog that uses markdown files for content.

Purpose
====

The author is not a big fan of large CMS's with lots of features. Generally, those CMS's tend to focus on SEO and plugins for ad networks and social media links. The author is not interested in making a living off a blog, and doesn't want the rest of those features.

Therefore, this project is meant to put as little friction as possible into the process of writing a post. All posts are just `*.md` files put into a directory, visible immediately after writing, editable anytime. The author mounts them with [webdav](https://github.com/Knetic/boji), and can write them from anywhere.

This eliminates the WYSIWYG editors, sidecar databases, network tracking, and other things that the author finds no value in.

Usage
====

To run everything from here, simply `make build run`. That'll compile the server, create the appropriate docker packages, and run them with the provided `docker-compose.yml`. It serves the `./sampledata` directory, so when you hit `localhost:1444` you'll see some basic posts, and a lorem ipsum about page.

If you want to just make a quick change to the static site (and note recompile the markdown server) you can use `make reload` to just rebuild/run that container.

To run in production, you'll probably want an SSL termination reverse proxy (the author uses [jwilder/nginx-proxy](https://github.com/jwilder/nginx-proxy)) instead of exposing ports directly. It would also be necessary to mount a data volume on your server to `/usr/share/writ`, which is where all the markdown files need to be located.

You'll probably want to change the `_about.md` (which populates the "about" page) to something about you, or the project you're running.