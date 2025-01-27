# Docker demo

This is a simple demo to allow people to put into practice the tips discussed in the workshop.

## Instructions

- Clone this repo
- Run `docker build -t workshop-start .`
- Open Docker and go to "Images" tab. Note the image size.
- Run `docker run --name workshop-start -p 8080:8080 -t workshop-start` to run the demo.
- Go to `http://localhost:8080` to see the demo.
- Refactor the Dockerfile to use the tips discussed:
  - Be mindfull of the commands used and the order
  - Use caching to your advantage (tip: the `COPY` command can change a lot!)
  - Use multi-layer builds to your advantage
  - Use the smallest base image possible Alpine > Debian-slim > Ubuntu/Debian
  - Make use of the `WORKDIR` command and comments to make the code readable
- Run `docker build -t workshop-refactored .` to build the refactored image.
- Run `docker run --name workshop-refactored -p 8080:8080 -t workshop-refactored` to run the refactored demo.
- Go to `http://localhost:8080` to see the refactored demo.
- Open Docker and go to "Images" tab. Note the difference in size.
- Optional: Remove the images from the workshop.
