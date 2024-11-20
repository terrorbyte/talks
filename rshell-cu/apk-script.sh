#!/bin/sh

SSH_KEY="ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILG8dpWe3Xb/1ivYCrUwXI9muU355HN/cb5T4N0plwti cblack"

#add community
apk add sudo go tmux vim bash findutils man-pages git less shadow

go get github.com/terrorbyte/talks/rshell-cu/trash@latest
cp ~/go/bin/trash /bin/trash

adduser -D user1
adduser -D user2 -s /bin/trash
adduser -D user3 -s /bin/bash
adduser -D user4 -s /bin/trash

echo -e 'TargetPassword\nTargetPassword' | passwd user1
echo -e 'TargetPassword\nTargetPassword' | passwd user2
echo -e 'TargetPassword\nTargetPassword' | passwd user3
echo -e 'TargetPassword\nTargetPassword' | passwd user4

mkdir -m 700 /root/.ssh
mkdir -m 700 /home/user1/.ssh
mkdir -m 700 /home/user2/.ssh
mkdir -m 700 /home/user3/.ssh
mkdir -m 700 /home/user4/.ssh

chown user1:user1 /home/user1/.ssh
chown user2:user2 /home/user2/.ssh
chown user3:user3 /home/user3/.ssh
chown user4:user4 /home/user4/.ssh

echo "$SSH_KEY" >/home/root/.ssh/authorized_keys
echo "$SSH_KEY" >/home/user1/.ssh/authorized_keys
echo "$SSH_KEY" >/home/user2/.ssh/authorized_keys
echo "$SSH_KEY" >/home/user3/.ssh/authorized_keys
echo "$SSH_KEY" >/home/user4/.ssh/authorized_keys

cat <<- "EOF" > /home/user3/.profile
if [ -f "$HOME"/.bashrc ]; then
	. "$HOME"/.bashrc
fi
EOF

cat <<- "EOF" > /home/user3/.bashrc
trash && exit
EOF
