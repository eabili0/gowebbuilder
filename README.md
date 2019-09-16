![motive](https://raw.githubusercontent.com/abilioesteves/gowebbuilder/assets/motive.png)

# GoWebBuilder
[![Docker Pulls](https://img.shields.io/docker/pulls/abilioesteves/gowebbuilder.svg)](https://hub.docker.com/r/abilioesteves/gowebbuilder)

Have you ever faced problems with your internet connection and was unable to build your docker image because of `go get...` madness, like the one above?

Worry no more.

This is a public service repository that puts together in a single docker image all dependencies one might need to build his/her golang web projects.

# Example

The folder `example` holds a sample Dockerfile to show how one could write it without having to deal with the `go get...` madness.

Check it out!

# How to Contribute

Fork it and open a Pull Request with your dependency added to the `main.go` file. After doing this, run `go test` to update *go.mod*
dependencies.

For now, the only dependencies we will accept are those found in `github.com` and `golang.org`.

# Observations

This is an experimental initiative and it's open for suggestions. The questions I have for now are:

1. how do we certifiy that the added dependency is in a stable version? 
    *If you are using go mod this problem has been solved.*

2. does code introspection via `github.com` suffices to accept the PR and guarantee that the generated image is free of malwares?

3. How to enable new gowebbuilder releases for the update of only one dependecy? 
    *If you are using go mod this problem has solved.*
    You only need to change the version of your dependency at `go.mod` file

Thanks!


