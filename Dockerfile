# Give me a tiny computer with Go already set up.
FROM golang:alpine 
# Docker creates a folder called /app and makes it the current working directory.
WORKDIR /app
# All your project files get copied into /app in the container.
COPY . .
# Go compiles your code into a binary called app
# This dockerfile is now a real program (like an .exe on Windows).
RUN go build -o app .
# Docker starts the container and runs: /app/app
# This is the command that runs when you start the container.
CMD ["./app"]