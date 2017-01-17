

GoServe!
========
v1.0.0-alpha-4 (CURRENTLY IN DEV)
v1.0.0-alpha-3 Launched for testing (New GUI)

**GoServe! is a powerful and simple to use cross platform web server that is fully standalone.**

The entire web stack is contained in a single binary, and enables quick web deployment and powerful web serving capabilities.

**Platforms:**
OSX, Unix, Linux, Windows

**Fully Standalone:**
Doesn't require Apache, NGINX, or anything other than this single GoServe! binary. No dependencies!

GoServe! was started as an open source version of enterprise features used in various large companies serving single or multiple web projects both internally and externally.

GoServe! was created for friends, and the community as a whole, to quickly spin up projects ranging from a single static website, to enterprise scale for entire webhosts.

GoServe! is perfect in the education field, mainly those teaching HTML/CSS/JS, with zero need for platform dependant server setup. Focus on the content your students are learning and developing, instead of them getting lost and frustrated in the setup process.

For more information on GoServe! in education, view the education specific wiki documentation:
<https://github.com/pathaugen/GoServe/wiki/GoServe!-in-Education>

If you are using GoServe! in your company or organization, please consider adding them to the list and any details you wish to provide!
<https://github.com/pathaugen/GoServe/wiki/Those-Using-GoServe!>

Thank you from the GoServe! [team](https://github.com/pathaugen/GoServe/wiki/Team)

***

Quick Start
-----------

 1. Download latest release of GoServe! for your platform (OSX, Unix, Linux, Windows)
    * <https://github.com/pathaugen/GoServe/releases>
 2. Launch Binary/Executable
 3. In the GoServe! console, listen on a port by typing this example:
    * $ listen 8080
 4. Open any web browser and visit:
    * [http://localhost:8080/](http://localhost:8080/) (from the same machine)
    * http://[ip-address]:8080/ (from another machine)
      * If the machine IP is 10.0.1.15:
        * http://10.0.1.15:8080/

**For more information and full documentation, view the wiki:**
<https://github.com/pathaugen/GoServe/wiki>

***

Features
--------

 * **Multiple Server and Domain Support**
   * Power of a large webhost in a single binary
   * Multiple domains pointed to same application instance supported
   * Spin up one webserver on a port, or multiple webservers on different ports
     * Outside the application can route differing bandwidth, static or dynamic IPs to these various servers/ports
   * Each webserver on each port customizable with one or more domains
   * Perfect for one website or hundreds
 * **Content and Pages**
   * Static website pages support and switchable between text (.txt), HTML (.html), and Markdown (.md)
     * MarkDown pages translated and output to HTML
     * Text pages translated to HTML with simple whitespace translations
     * Pure HTML pages support linked CSS and JS per-page
 * **Templates**
   * Templates following Golang standard template format
   * Templates selectable per domain in application: domain.com/templates
   * Template include HTML, CSS, and JS resources
 * **HTTPS**
   * HTTPS SSL with pre-generated placeholder certificate (make sure to replace with your own!)
 * **Speed and Caching**
   * File based caching for low server load and speed
   * Memory based caching for high performance and high number of users
 * **Administration**
   * **Console Controls**
 	 * Full console controlling on all platforms (Unix/Linux/Windows)
 	 * Full monitoring via console command: **monitor**
 	 * Full monitoring of active users via console command: **active**
   * Form based administration and editing of all templates: /templates
     * Full form editing of HTML, CSS, JS for each template
   * Form based editing of all page data via extension: **.edit**
     * HTML content also includes administration of CSS and JS per page
 * **API**
   * API endpoints baked into the website URLs themselves via dot notation
     * .create
     * .edit
     * .save
     * .text
     * .md
     * .html
     * .css
     * .js
     * .template

***

External Libraries
------------------

 * **ansicolor**
   * Source: <https://github.com/shiena/ansicolor>
   * Commit included: <https://github.com/shiena/ansicolor/commit/a422bbe96644373c5753384a59d678f7d261ff10>
   * License: [MIT](https://en.wikipedia.org/wiki/MIT_License)
   * Description: Coloring of command prompt, including Windows 10.
 * **GoThrust**
   * Source: <https://github.com/miketheprogrammer/go-thrust>
   * Commit included: <https://github.com/miketheprogrammer/go-thrust/commit/92e66a5b9de59b0c0e07ced3a964135ddcdeb883>
   * License: [MIT](https://en.wikipedia.org/wiki/MIT_License)
   * Description: Chromium-based cross-platform / cross-language application framework.
   * **profile** (Dependency of GoThrust)
     * Source: <https://github.com/pkg/profile>
     * Commit included: <https://github.com/pkg/profile/commit/8a808a6967b79da66deacfe508b26d398a69518f>
     * License: <https://github.com/pkg/profile/blob/master/LICENSE>
     * Description: Simple profiling for Go.
   * **pb** (Dependency of GoThrust)
     * Source: <https://github.com/cheggaaa/pb>
     * Commit included: <https://github.com/cheggaaa/pb/commit/04b234c80d661c663dbcebd52fc7218fdacc6d0c>
     * License: <https://github.com/cheggaaa/pb/blob/master/LICENSE>
     * Description: Console progress bar for Golang
   * **websocket** (Dependency of GoThrust)
     * Source: <https://github.com/gorilla/websocket>
     * Commit included: <https://github.com/gorilla/websocket/commit/54f9decdbfb3fe254c5df3d5aa234480989f4af0>
     * License: <https://github.com/gorilla/websocket/blob/master/LICENSE>
     * Description: A WebSocket implementation for Go.
   * **cli** (Dependency of GoThrust)
     * Source: <https://github.com/urfave/cli>
     * Commit included: <https://github.com/urfave/cli/commit/1efa31f08b9333f1bd4882d61f9d668a70cd902e>
     * License: [MIT](https://en.wikipedia.org/wiki/MIT_License)
     * Description: A simple, fast, and fun package for building command line apps in Go.
     * **yaml** (Dependency of cli)
       * Source: <https://github.com/go-yaml/yaml/tree/v2>
       * Commit included: <https://github.com/go-yaml/yaml/commit/a83829b6f1293c91addabc89d0571c246397bbf4>
       * License: <https://github.com/go-yaml/yaml/blob/v2/LICENSE>
       * Description: YAML support for the Go language.
   * **birpc** (Dependency of GoThrust)
     * Source: <https://github.com/tv42/birpc>
     * Commit included: <https://github.com/tv42/birpc/commit/22dcbfff0024e83a23cb42c800c1523a3ba08a4c>
     * License: <https://github.com/tv42/birpc/blob/master/LICENSE>
     * Description: Bi-directional RPC library for Go, including JSON-over-WebSocket.
     * **topic** (Dependency of birpc)
       * Source: <https://github.com/tv42/topic>
       * Commit included: <https://github.com/tv42/topic/commit/aa72cbe81b4823f349da47a4d749cdda61677c09>
       * License: <https://github.com/tv42/topic/blob/master/LICENSE>
       * Description: Go library for in-process single-topic pub-sub.

***

External Libraries (No Longer Used)
-----------------------------------

 * **Color**
   * Source: <https://github.com/fatih/color>
   * Commit included: <https://github.com/fatih/color/commit/87d4004f2ab62d0d255e0a38f1680aa534549fe3>
   * License: [MIT](https://en.wikipedia.org/wiki/MIT_License)
   * **go-colorable** (Dependency of Color)
     * Source: <https://github.com/mattn/go-colorable>
     * Commit included: <https://github.com/mattn/go-colorable/commit/9056b7a9f2d1f2d96498d6d146acd1f9d5ed3d59>
     * License: [MIT](https://en.wikipedia.org/wiki/MIT_License)
     * Description: Windows coloring of command line.
   * **go-isatty** (Dependency of Color)
     * Source: <https://github.com/mattn/go-isatty>
     * Commit included: <https://github.com/mattn/go-isatty/commit/56b76bdf51f7708750eac80fa38b952bb9f32639>
     * License: [MIT](https://en.wikipedia.org/wiki/MIT_License)

***

v1.0.0 (UNRELEASED)
------

 * Major Changes:
 * Minor Changes:
 * Bugs Fixed:

***

License
-------

TBA
