{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-23.05") {} }:

pkgs.mkShell {
  packages = [
    (pkgs.python3.withPackages (ps: [
      ps.flask
      ps.pyyaml
    ]))

    pkgs.curl
    pkgs.jq
    pkgs.metasploit
    pkgs.nmap
    pkgs.masscan
    pkgs.tunctl
  ];
}

