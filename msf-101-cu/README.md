# Metasploit Framework Introduction

An introduction and exploration of the Metasploit framework.

## Info

Where: CU
When: 2023-11-30

## Documentation

* `slides/` - HTML and [sent](https://tools.suckless.org/sent/) slides
    * A copy will be hosted on [hosakacorp.net](https://hosakacorp.net/t/msf-101-cu.html)
    * The slides are supposed to be thin, this is demo heavy
* `payloads/` - All the modules used or customized for the demo and curl ones too
* `msf-target-vm.nix` -

## Setting up test environment

### Host setup

1. Setup networking - `host-networking.sh` - `./host-networking.sh poptart enp0s31f6`
2. Install the postgresql NixOS module - `imports = [ ./msf-host-nixos-module.nix ];`
3. Build the demo VM (see below)
4. Drop into a configured shell - `nix-shell`
5. Init the msf database `msfdb init --connection-string=postgresql://msfadmin:metasploit-class-cu@127.0.0.1:5432/postgres`
6. Launch `msfconsole` [1] 

### Building demo VM

```
$ nix-build msf-target-vm.nix
[...snip...]
/nix/store/142x8wqxxyids009q5g4a3i7gz7z26dy-cassandra-3.11.12/bin/.nodetool-wrapped: interpreter directive changed from "#!/bin/sh" to "/nix/store/0rwyq0j954a7143p0wzd4rhycny8i967-bash-5.2-p15/bin/sh"
/nix/store/142x8wqxxyids009q5g4a3i7gz7z26dy-cassandra-3.11.12/bin/.cassandra-wrapped: interpreter directive changed from "#!/bin/sh" to "/nix/store/0rwyq0j954a7143p0wzd4rhycny8i967-bash-5.2-p15/bin/sh"
stripping (with command strip and flags -S -p) in  /nix/store/142x8wqxxyids009q5g4a3i7gz7z26dy-cassandra-3.11.12/lib /nix/store/142x8wqxxyids009q5g4a3i7gz7z26dy-cassandra-3.11.12/bin
building '/nix/store/y04qm45ipli9xjyg0pd8l7gzj6x7183c-cassandra-etc.drv'...
building '/nix/store/p4cmfvsvi16nmi6plpz7x6xrgdjyn9fb-unit-cassandra-full-repair.service.drv'...
building '/nix/store/01mdlrv797hkmgzya48g7vdrdpssy74f-unit-cassandra-incremental-repair.service.drv'...
building '/nix/store/jah7pb167fjz065m3azy4hi2qg1jbhzf-unit-cassandra.service.drv'...
building '/nix/store/zvbr4hlddc5s40clhwww4yr2zzbzb92s-system-units.drv'...
building '/nix/store/1a6v6fws85a7qk91h04n3x5x62l0jgv7-etc.drv'...
building '/nix/store/rx4hdbp2ch18vr7fd6ybixa50iz7yrix-nixos-system-nixos-23.05pre-git.drv'...
building '/nix/store/vjgl60n6xla5l7k6s6n14lcb5ijijmng-closure-info.drv'...
building '/nix/store/aqc371p9xw0rhd50x4h2jg27wgsqyhgc-run-nixos-vm.drv'...
building '/nix/store/lvbfx1glgcqyycxks9s0s565n3cij9cj-nixos-vm.drv'...
/nix/store/vijanak0b01awvqz2ifsbz9l55zvhdqs-nixos-vm
$ ./result/bin/run-nixos-vm
Disk image do not exist, creating the virtualisation disk image...
Formatting '/home/poptart/src/git/talks/msf-101-cu/nixos.qcow2', fmt=qcow2 cluster_size=65536 extended_l2=off compression_type=zlib size=10737418240 lazy_refcounts=off refcount_bits=16
Virtualisation disk image created.
```

### Starting Testing Shell

```
$ nix-shell
```

## References

- [https://docs.rapid7.com/metasploit/](https://docs.rapid7.com/metasploit/)* this has some MSF Pro content
- [https://www.offsec.com/metasploit-unleashed/](https://www.offsec.com/metasploit-unleashed/)
- [https://docs.metasploit.com/api/](https://docs.metasploit.com/api/)

PyYAML:

- [https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/Insecure%20Deserialization/YAML.md#pyyaml](https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/Insecure%20Deserialization/YAML.md#pyyaml)

## Notes

[1] - weird issue with the generated config from msfdb, may require `msfconsole -y whatever` or modifying `~/.msf4/database.yml` to utilize the msfadmin cred vs the generated one that is triggering a stack trace

