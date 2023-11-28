{ config, lib, pkgs, ... }:
with lib;
let cfg = config.services.msfDemo;
in {
  options.services.msfDemo = {
    enable = mkOption {
      type = types.bool;
      default = true;
      description = "";
    };

    db = mkOption {
      type = types.str;
      default = "msf";
      example = "msf";
      description = "";
    };
  };

  config = {
    services.postgresql = {
  enable = true;
  package = pkgs.postgresql_15;
  ensureDatabases = [ "msf" ];
  enableTCPIP = true;
  port = 5432;
  settings.listen_addresses = lib.mkForce "127.0.0.1";
  #settings = pkgs.lib.mkDefault {
  #  "listen_addresses" = "127.0.0.1";
  #};
  authentication = pkgs.lib.mkOverride 10 ''
  local all all              trust
  host  all all 127.0.0.1/32 trust
  host  all all ::1/128      trust
'';
  
  initialScript = pkgs.writeText "backend-initScript" ''
    CREATE ROLE msfadmin WITH LOGIN PASSWORD 'metasploit-class-cu' SUPERUSER;
  '';
};
  };
}
