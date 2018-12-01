# Read Me

During the second Python Web Dev class for Django 2.1, the default name
for the app created by Heroku was `infinite-tundra-77435`.

To ensure this app name is not subsequently abused, this application
unconditionally redirects https://infinite-tundra-77435.herokuapp.com/
to https://www.pywebdev.com/ , the official site for the classes.

# Installation

This app requires:

- Go 1.12.2+
- Dep 0.5+

While `dep ensure` would normally install all of the dependencies for
the application, this project is small enough that all of the
dependencies are vendored under the `vendor/` directory, following the
advice in [`dep`'s FAQ].

[`dep`'s FAQ]: https://github.com/golang/dep/blob/master/docs/FAQ.md#should-i-commit-my-vendor-directory

# Deployment

If you've just cloned the repo, you first need to know about the Heroku
remote.

```shell
$ heroku git:remote -a infinite-tundra-77435
```

You can then make changes and push.

```shell
$ git push heroku master
```

What do you mean, where are the tests? What about CI? Pshah, this is
life on the edge, my friend!
