FROM mcr.microsoft.com/devcontainers/base:ubuntu-22.04

# dependencies
RUN apt upgrade && apt update && apt install -y curl jq

RUN mkdir /home/vscode/.m2 && chown vscode:vscode -R /home/vscode/.m2
RUN mkdir /home/vscode/go && chown vscode:vscode -R /home/vscode/go

# non-root USER
USER vscode

# Install Go latest version
RUN VERSION=$(curl --silent https://go.dev/dl/\?mode\=json | jq -r '.[].files[].version' | sort | uniq | tail -n1) && \
    echo && \
    echo "Selected Version: $VERSION" && \
    echo && \
    GO="$VERSION.linux-amd64.tar.gz" && \
    sudo wget --quiet "https://go.dev/dl/${GO}" && \
    sudo rm -rf /usr/local/go && \
    sudo tar -C /usr/local -xzf "${GO}" && \
    sudo rm "${GO}"

# Install Java / Maven / Gradle
ARG JAVA_VERSION="21.0.2-oracle"
ARG MAVEN_VERSION="3.9.4"
ARG GRADLE_VERSION="8.3"
RUN curl -s "https://get.sdkman.io" | bash
RUN bash -c "source $HOME/.sdkman/bin/sdkman-init.sh && \
    yes | sdk install java $JAVA_VERSION && \
    yes | sdk install maven $MAVEN_VERSION && \
    yes | sdk install gradle $GRADLE_VERSION && \
    rm -rf $HOME/.sdkman/archives/* && \
    rm -rf $HOME/.sdkman/tmp/*"


# INSTALL lazygit latest version
RUN LAZYGIT_VERSION=$(curl "https://api.github.com/repos/jesseduffield/lazygit/releases/latest" | grep -Po '"tag_name": "v\K[^"]*') && \
    sudo curl -Lo lazygit.tar.gz "https://github.com/jesseduffield/lazygit/releases/latest/download/lazygit_${LAZYGIT_VERSION}_Linux_x86_64.tar.gz" && \
    sudo tar xf lazygit.tar.gz lazygit && \
    sudo install lazygit /usr/local/bin && \
    sudo rm lazygit.tar.gz

ENV JAVA_HOME="/home/vscode/.sdkman/candidates/java/current"
ENV MAVEN_HOME="/home/vscode/.sdkman/candidates/maven/current"
ENV GRADLE_HOME="/home/vscode/.sdkman/candidates/gradle/current"
ENV GO_HOME="/usr/local/go"
ENV PATH="$MAVEN_HOME/bin:GRADLE_VERSION:$JAVA_HOME/bin:$GO_HOME/bin:$PATH"

CMD [ "tail", "-f", "/dev/null"]