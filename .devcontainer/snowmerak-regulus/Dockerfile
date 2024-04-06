FROM fedora:rawhide

RUN dnf upgrade --refresh -y && \
    dnf update -y && \
    dnf install go git curl -y && \
    echo 'export PATH="$PATH:$HOME/go/bin"' >> ~/.bashrc && \
    go install github.com/wailsapp/wails/v2/cmd/wails@latest && \
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/master/install.sh | bash && \
    source ~/.bashrc && \
    nvm install --lts && \
    nvm use --lts

CMD ["/bin/sh", "-c", "while true; do sleep 30; done;"]
