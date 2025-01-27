# Docker demo
- Clone this repo
- Run `docker build -t workshop-start
- Refactor the Dockerfile to use the tips discussed:
  - Be mindfull of the commands used and the order
  - Use caching to your advantage (tip: the `COPY` command can change a lot!)
  - Use multi-layer builds to your advantage
  - Use the smallest base image possible Alpine > Debian-slim > Ubuntu/Debian
  - Make use of the `WORKDIR` command and comments to make the code readable