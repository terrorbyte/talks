{ lib, python3Packages }:
with python3Packages;
buildPythonApplication {
  pname = "demo-vulnerable-app";
  version = "1.0";

  propagatedBuildInputs = [ flask pyyaml ];

  src = ./.;
   postInstall = ''
    cp -r static $out/bin/static
  '';
}
