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
      ensureDatabases = [ "${cfg.db}" ];
      authentication = pkgs.lib.mkOverride 10 ''
        #type database  DBuser  auth-method
        local ${cfg.db}       all     trust
      '';
    };
  };
}
