<img src="logo.png">

Hblend is a html, css and js preprocessor. Manage your code dependencies project into a nice package.

<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [Sample project](#sample-project)
- [Continuous preprocessing](#continuous-preprocessing)
- [Getting started](#getting-started)
    - [Components](#components)
    - [Www](#www)
- [Tags](#tags)
    - [path](#path)
    - [base64](#base64)
    - [content](#content)
    - [include](#include)
    - [link](#link)
- [How to build](#how-to-build)

<!-- /MarkdownTOC -->

# Sample project

There is a sample project with 4 components, blend it with the following command:

```bash
hblend page-1
```

And have a look the final result in `www/page-1.html`.

# Continuous preprocessing

The common way to work is refresh the browser or the mobile phone constantly. To avoid the boring task of launching the hblend command each time, you can use the inotify tools in this way:

```bash
watch "inotifywait components && hblend YOUR-COMPONENT"
```

# Getting started

Hblend works with two folders:

* `components` -> where all the sources, html, css, js, images, binaries, etc. are located
* `www` -> where the final result is generated

To blend the component `my-page` you should run:

```bash
hblend my-page
```

A `www/my-page.html` will be generated and all dependencies will be stored in `www/files/`.

## Components

TODO: write components organization

## Www

TODO: write www organization


# Tags

Hblend preprocessor works with tags with the following syntax:

```
[[TAGNAME flag1 flag 2 attribute1="value1" attribute2:'value2']]
```

NOTE (1): assignement operator can be `=` or `:`

NOTE (2): values can use simple quotes, double quotes or no quotes if it is no needed.

Some tags may not be available for all file types.

## path

```text
[[path {filename}]]
```

This tag will be replaced by the final filename, for example:

```html
<img src="[[path big-logo.png]]">
```

Will generate this final html code:

```html
<img src="files/d9832a84eab44a6ad4e9a6f50a84cf03.png">
```

## base64

```text
[[base64 {filename}]]
```

This tag will be replaced by the base64 representation of the filename content. It is extremely useful to embed small images into css and html code. Example:

```html
<img src="[[base64 small-icon.png]]">
```

Will generate this final html code:

```html
<img src="data:;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAICAYAAAAx8TU7AAAALElEQVQI13WOuQ0AIBDDbPbfOXQnBCGllc8k4dICUF9YnV+oTs3Ac6/GbZc2hNgLDevagOMAAAAASUVORK5CYII=">
```

NOTE: JavaScript files does not support this tag

## content

```text
[[content {filename} escape='{mode}']]
```

This tag will be replaced by the content of the file. It has several modes:

* `escape=none` -> [ default ] put the content as it is
* `escape=string` -> escape `'`, `"`, `\n`, `\t`, `\\`
* `escape=urlencode` -> use url encode to escape it
* `escape=html` -> escape html entities

NOTE (1): If the included file is html, it is not processed (TODO this feature in the future)

## include

```text
[[include {component-name}]]
```

This will include properly the corresponding css and js code.

NOTE (1): If this tag is inside an html file, will be replaced by the component html file.

## link

```text
[[link {component-name}]]
```

This tag will be replaced by the relative url to other hblended component. The component will be hblended too.

NOTE (1): This is only available to html files

# How to build

```bash
make
```

Will generate a binary file in `_vendor/bin/hblend`


```bash
make all
```

Will compile hblend for the main architectures and operative systems in `_vendor/bin/`