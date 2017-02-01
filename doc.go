/*
This is a very simple package that will allow you to create gravatr images out of an email address.

GRAVATAR

Gravatar is an online service that helps put a phot to an email, or a photo to a user for that matter, with out making them upload it. If you haven't  registered there yet, your email will generate a funky looking photo. A lot of web apps use gravatar, it allows them with ease to display users photo, and even if user has never gone to gravatr to upload a photo, gravatr will automatically generate a photo that will reflect their inner being... or not.

PACKAGE

Currently there are two functions.

1. GetImg() takes in an email address and returns a link to gravatar image.

2. GetHash() This will return only a hash.

USAGE

run: go get github.com/ivanmrchk/gravatar and import it

  func main() {
    a := gravatar.GetImg("myemail@email.com")

    fmt.Println(a)
  }

You should use it in your next app.
*/
package gravatar
