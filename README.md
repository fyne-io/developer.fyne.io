# developer.fyne.io Website

> This repository has been archived as we are now using [docs.fyne.io](/fyne-io/docs.fyne.io).

This repository contains the source of the https://developer.fyne.io website.

## Running the website locally

The website uses [jekyll](https://jekyllrb.com/) to generate the website from markdown files.
With ruby installed it is as easy as running the following commands in a terminal (the first only needs to run once):

```bash
gem install bundler jekyll
bundle exec jekyll serve
```

## Run generator scripts

To generate new images etc you can use the following commands:

```bash
cd _gen
go run genwidgets.go # generate widget images
go run genlayouts.go # generate layout images
go run gendialogs.go # generate dialog images
```

