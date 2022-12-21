![motive](https://raw.githubusercontent.com/abilioesteves/gowebbuilder/assets/motive.png)

# GoWebBuilder
[![Docker Pulls](https://img.shields.io/docker/pulls/abilioesteves/gowebbuilder.svg)](https://hub.docker.com/r/abilioesteves/gowebbuilder)

Have you ever faced problems with your internet connection and was unable to build your docker image because `go get...` fails to resolve some dependencies, like showcased above?

This repository puts together in a single docker image all dependencies I might need to build my golang web projects. Please, use this project as an inspiration to setup your CI/CD pipeline with a base docker image with all relevant golang dependencies you might need for your projects.

# Example

The folder `example` holds a sample Dockerfile to show how one could define it without having to call `go get...` to resolve the needed dependencies.

Check it out!

# How to Contribute

Please don't. Just fork it and be happy :) 
