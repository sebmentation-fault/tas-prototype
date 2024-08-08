# initial prototyping

just simple HTML files so that we can experiment with so we can figure out what
looks good and what might require changing around from the initial mock up

when changing an HTML element's tailwind classes, you need to re-compile
tailwind so that the necessary styles are added to the stylesheet. this can be
done in the background, so what seb found useful is to have a terminal window
open while coding with `tailwindcss -i index.css -o styles.css --watch`. this
monitors the HTML, and recompiles the stylesheet for you :)

see below for how to install tailwindcss

## tailwindcss

requires that [tailwindcss is installed on your machine](https://tailwindcss.com/blog/standalone-cli).

steps:

1. go to linked site and install the binary as described
2. add to zsh path
    * (e.g. what seb did)
    * move the binary to a folder in your `PATH` (e.g. run
      `mv ./tailwindcss /usr/local/bin/tailwindcss`)
