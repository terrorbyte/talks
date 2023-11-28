set -x
ls --color=force -lah "$(which msfconsole | sed 's%bin/msfconsole%%g')/share/msf/modules/${1}"
