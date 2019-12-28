writ
====

Accepts http requests for a markdown file, responding with the html equivalent of the file.
Useful as part of a larger ecosystem that serves static files. 

## API

* `/list` - recursively lists all MD files that can be served. Trims the MD suffix.
* `/*` - Any other path is checked to see if it's an MD file. If not, 404. If so, converted to HTML and returned.

## Running

The executable itself serves markdown files anywhere in its current working directory, recursively downward.

However, usually this is run in a container, the cwd set to `/usr/share/writ`, and that location mounted to some other volume, or somewhere on the host.