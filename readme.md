

GoServe (CURRENTLY IN DEV)
=======

**GoServe! is a cross platform simple yet powerful static website server that is fully standalone, contained in a single binary, and enables quick website deployment and serving.**

**Platforms:**
OSX, Unix, Linux, Windows

**Fully Standalone:**
Doesn't require Apache, NGINX, or anything other than the single GoServe! binary

GoServe! was started as an open source version of enterprise features used in various large companies serving single or multiple websites internally and externally.

GoServe! was created for friends and the community as a whole to quickly spin up a single static website, or scale to entire webhosts.

GoServe! is perfect in the education field teaching HTML/CSS/JS with zero need for platform dependant server setup. Focus on the content your students develop, not the setup.

If you are using GoServe! in your company or organization, add to the list and any details you wish to provide!
https://github.com/pathaugen/GoServe/wiki/Those-Using-GoServe!

***

Quick Start
-----------

 1. Download latest release of GoServe! for your platform (OSX, Unix, Linux, Windows)
    * https://github.com/pathaugen/GoServe/releases
 2. Launch Binary/Executable
 3. In the GoServe! console, listen on a port by typing this example:
    * > listen 8080
 4. Open any web browser and visit:
    * http://localhost:8080/ (from the same machine)
    * http://[ip-address]:8080/ (from another machine)
      * If the machine IP is 10.0.1.15:
        * http://10.0.1.15:8080/

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
   * Source: https://github.com/shiena/ansicolor
   * Commit included: https://github.com/shiena/ansicolor/commit/a422bbe96644373c5753384a59d678f7d261ff10
   * License: [MIT](https://en.wikipedia.org/wiki/MIT_License)
   * Description: Coloring of command prompt, including Windows 10.

***

External Libraries (No Longer Used)
-----------------------------------

 * **Color**
   * Source: https://github.com/fatih/color
   * Commit included: https://github.com/fatih/color/commit/87d4004f2ab62d0d255e0a38f1680aa534549fe3
   * License: [MIT](https://en.wikipedia.org/wiki/MIT_License)
 * **go-colorable** (Dependency of Color)
   * Source: https://github.com/mattn/go-colorable
   * Commit included: https://github.com/mattn/go-colorable/commit/9056b7a9f2d1f2d96498d6d146acd1f9d5ed3d59
   * License: [MIT](https://en.wikipedia.org/wiki/MIT_License)
   * Description: Windows coloring of command line.
 * **go-isatty** (Dependency of Color)
   * Source: https://github.com/mattn/go-isatty
   * Commit included: https://github.com/mattn/go-isatty/commit/56b76bdf51f7708750eac80fa38b952bb9f32639
   * License: [MIT](https://en.wikipedia.org/wiki/MIT_License)

***

v1.0.0 (UNRELEASED)
------

 * xxx

***

License
-------

xxx
