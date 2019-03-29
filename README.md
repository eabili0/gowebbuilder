![motive](https://raw.githubusercontent.com/abilioesteves/gowebbuilder/assets/motive.png)

# GoWebBuilder

Have you ever faced problems with your internet connection and was unable to build your docker image because of `go get...` madness, like the one above?

Worry no more.

This is a public service repository that puts together in a single docker image all dependencies one might need to build his/her golang web projects.

# How to Contribute

Fork it and open a Pull Request with your dependency added to the `main.go` file.

For now, the only dependencies we will accept are those found in `github.com`.

# Observations

This is an experimental initiative and it's open for suggestions. The questions I have for now are:

1. how do we certifiy that the added dependency is in a stable version?

2. does code introspection via `github.com` suffices to accept the PR and guarantee that the generated image is free of malwares?

Thanks!


