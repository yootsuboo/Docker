FROM golang:1.19.11-alpine3.18
WORKDIR /root
RUN apk --update add \
    git \
    vim \
    tmux \
    nodejs \
    npm \
    curl
RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
COPY ./build/vimrc /root/.vimrc
RUN mkdir -p /go/work
WORKDIR /go/work
ADD . /go/work
RUN go install golang.org/x/tools/cmd/goimports@latest

