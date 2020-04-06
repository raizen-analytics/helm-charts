#!/bin/bash
if [ -n "$USERNAME" ]; then
    sudo USERNAME=$USERNAME bash -c 'echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/nopasswd && usermod -l $USERNAME coder';
fi
["dumb-init", "fixuid", "-q", "/usr/local/bin/code-server", "--host", "0.0.0.0", "."]
sudo service ssh restart && dumb-init fixuid -q /usr/local/bin/code-server --host 0.0.0.0 . $@