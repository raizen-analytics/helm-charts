#!/bin/bash
if [ -n "$USERNAME" ]; then
    sudo USERNAME=$USERNAME bash -c 'echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/nopasswd && usermod -l $USERNAME coder';
fi

sudo service ssh restart && dumb-init $@