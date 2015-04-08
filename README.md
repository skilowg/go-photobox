# Go Photobox

Go Photobox is a self-contained photo gallery manager designed to run
on a Raspberry Pi as a front-end to an external hard drive.

Ok, maybe that was a mouthful. Let's start with the basics.

Who is this useful for? Go Photobox might be helpful if you:

* Know you could use Dropbox, Google Drive, or Flickr, but like the idea of having
your photos on a hard drive you own
* Don't really enjoy messing with cables, formats, and looking for your hard
drive every time you need to manage your photos
* Don't mind "some assembly required" projects for your home

# Why Go?

Go offers a couple really attractive advantages:

* The ability to create fast, self-contained programs
* The ability to compile for other targets

In addition, Go is really great for writing systems-level applications as well
as small web API servers.

I have been really excited about learning Go for a long time, so while there
are clear competitors and replacements in this space, this is a great learning
opportunity to explore:

* Building web applications in Go
* Go as a means to make more use of small hardware, like the Raspberry Pi

# Why Raspberry Pi?

Small hardware platforms like Raspberry Pi, Beaglebone, and Arduino offer
great opportunities to hobbyists and home hackers.

Because Raspberry Pi offers the power of a computer with low power consumption,
it seems like a great choice for home automation projects.

# Installing and Running

More to write later! This project is still in development.

The general overview will look something like this:

* Gather required materials
  * A Raspberry Pi
  * An external USB hard drive
  * A powered USB hub
* Connect Raspberry Pi and the external drive to the USB hub and power
everything up
* Connect to your Raspberry Pi and configure it to:
  * Access your home network via Wifi or Ethernet
  * Mount and connect to your external drive
* Deploy Go Photobox to the hard drive and start it up
* Configure your home computers to point to Go Photobox over the network and enjoy!

# Contributing

Fork this project and clone to your computer. The project is broken up into
the back-end and front-end concerns since the development approaches for each
are quite different.

## Web App

All front-end code lives in the `webapp` directory and is organized as a
self-contained JavaScript single-page application.

More information for working on the web app can be found in
the [README for the web app](webapp/README.md).

## Server and API

All API and file system handling is performed in `main.go`. Until routes
get more complicated, we basically have two main areas of concern:

* File server mounted at `/`
* Photos file server mounted at `/photos`
