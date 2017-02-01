# Gravatar



Gravatr is a golang package for gravatar.com

##### What is Gravatar.com
Your Gravatar is an image that follows you from site to site appearing beside your name when you do things like comment or post on a blog. Avatars help identify your posts on blogs and web forums, so why not on any site?

Developers use gravatar images to bring to life user database or contacts in their apps. All you need to have to get a gravatar image is an email. If an email address is not registered at gravatar.com, then the user with that email address will have a funky looking image.

## What does this package do?

This is a very small and timy package compared to some other ones, it has only two functions.
- GetHash() - will return only the hash  
- GetImg() - will return a complete url to the image.

## Usage
First run get the package
```sh
$ go get github.com/ivanmrchk/gravatar
```
Import it into your project
Example:
```Go
func main() {

	x := gravatar.GetImg("example@example.com")
}
```
*This will return such string:* http://www.gravatar.com/avatar/23463b99b62a72f26ed677cc556c44e8?d=identicon

#### Info
This package comes with `doc.go` file.

License
----

MIT


**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)


   [dill]: <https://github.com/joemccann/dillinger>
   [git-repo-url]: <https://github.com/joemccann/dillinger.git>
   [john gruber]: <http://daringfireball.net>
   [@thomasfuchs]: <http://twitter.com/thomasfuchs>
   [df1]: <http://daringfireball.net/projects/markdown/>
   [markdown-it]: <https://github.com/markdown-it/markdown-it>
   [Ace Editor]: <http://ace.ajax.org>
   [node.js]: <http://nodejs.org>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [keymaster.js]: <https://github.com/madrobby/keymaster>
   [jQuery]: <http://jquery.com>
   [@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
   [express]: <http://expressjs.com>
   [AngularJS]: <http://angularjs.org>
   [Gulp]: <http://gulpjs.com>

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]:  <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
