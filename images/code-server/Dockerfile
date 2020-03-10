FROM codercom/code-server:2.1698

USER root
ARG GO_VERSION=1.14
# Install go and kaniko to build images
RUN wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -xzf go1.14.linux-amd64.tar.gz -C /usr/local \
    && apt update \
    && apt install gcc -y
ENV PATH=$PATH:/usr/local/go/bin
RUN go get github.com/GoogleContainerTools/kaniko/cmd/executor/
RUN mv /root/go/bin/executor /usr/local/bin/kaniko \
    && rm -r /root/go/ 

# Install programming tools
RUN apt install python3-pip -y
USER coder
RUN pip3 install --user streamlit \
    && pip3 install --user kubernetes
ENV PATH=$PATH:/home/coder/.local/bin