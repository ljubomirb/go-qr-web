
# go-qr-web

[](https://github.com/ljubomirb/go-qr-web/edit/main/README.md#go-qr-web)

The simplest of them all, qr code generator and web server to serve it to on a single page.

Style index.html whatever you like, put it in the same folder as main app.

run on windows or linux as usual

docker setup has not been tested at all, but it should look something like this

1) git clone  [https://github.com/ljubomirb/go-qr-web](https://github.com/ljubomirb/go-qr-web)
2)  Build the Docker Image: In the same directory as the Dockerfile, run the following command:
> docker build -t go-qr-web .
3) Run the Docker Container: After building the image, you can run the container with the following command:
> docker run -d -p 8005:8005 go-qr-web

This command runs the container in detached mode (-d) and maps port 8005 of the container to port 8005 of your local machine.
