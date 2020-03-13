FROM codercom/code-server:2.1698

USER root
ARG GO_VERSION=1.14
# Install go and kaniko to build images
RUN wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -xzf go1.14.linux-amd64.tar.gz -C /usr/local \
    && apt update \
    && apt install gcc -y \
    && rm -r go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

# RUN go get github.com/GoogleContainerTools/kaniko/cmd/executor/
# RUN mv /root/go/bin/executor /usr/local/bin/kaniko \
#     && rm -r /root/go/

# Install programming tools
RUN apt install -y python3-pip \
    && apt install -y python3-venv \
    && pip3 install virtualenv \
    && pip3 install streamlit
ENV PATH=$PATH:/home/coder/.local/bin

# Install kaniko-build tool

COPY kaniko_build /home/coder/kaniko_build
RUN pip3 install /home/coder/kaniko_build \
    && rm -r /home/coder/kaniko_build
    
USER coder