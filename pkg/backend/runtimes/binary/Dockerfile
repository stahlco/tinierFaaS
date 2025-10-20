# scratch is an empty container without any OS, shell, etc.
FROM scratch
# Runtime Environment (shell) -> Created earlier in the justfile/makefile
ADD blob.tar.gz /

EXPOSE 8000

WORKDIR /usr/src/app

# Copy the function into the container
# fn
#  |_ fn.sh
COPY fn/* ./

# Makes sure the function is executable
RUN chmod +x fn.sh

# Executes the functionhandler.go
CMD ["./handler.bin"]

