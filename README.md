Doopla
======

Using a URL list for security testing can be painful as there are a lot of URLs that have uninteresting/duplicate content, doopla aims to make your life easier.

Doopla cleans up :

* Human Content
* Duplicate URLs based on parameters
* Removal of static content

Doopla is simple, It doesn't connect or verify your URLs.

installation :

```
$ go get -u github.com/affix/doopla
```

usage :

```
$ cat urls.txt | doopla
```

You can even chain doopla with other tools to create powerful one liners

```
$ waybackurls example.com | doopla | qsreplace '"><script>alert(1)</script>' | airixss -payload '"><script>alert(1)</script>'
```