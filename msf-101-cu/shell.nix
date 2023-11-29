{ pkgs ?
  import (fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-23.05")
  { } }:

pkgs.mkShell {
  packages = [
    (pkgs.python3.withPackages (ps: [ ps.flask ps.pyyaml ]))

    pkgs.curl
    pkgs.jq
    pkgs.metasploit
    pkgs.nmap
    pkgs.masscan
    pkgs.tunctl
    pkgs.farbfeld
    #(pkgs.sent.overrideAttrs (oldAttrs: rec {
    #  patches = [
    #    (lib.fetchpatch {
    #      url =
    #        "https://tools.suckless.org/sent/patches/bilinear_scaling/sent-bilinearscaling-1.0.diff";
    #      sha256 = "1xhyhdl88jc2g6m77amw278waw1ahwg02y2c21sg77m41ksfzwb5";
    #    })
    #  ];
    #}))
  ];
}

