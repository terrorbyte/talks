{ pkgs ? import <nixpkgs> {} }:
let
in pkgs.mkShell {

  buildInputs = [
    pkgs.scdoc
    pkgs.nasm
    pkgs.expect
    pkgs.gcc
  ];

  shellHook = ''
  '';

}
