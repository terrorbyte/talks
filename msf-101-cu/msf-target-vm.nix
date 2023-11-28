let
  #pkgs = import <nixpkgs> { config.permittedInsecurePackages = [ "python-2.7.18.6" "openssl-1.1.1w" ]; };
  pkgs = import (builtins.fetchTarball {
    name = "nixos-23.05-msf-cu";
    url =
      "https://github.com/nixos/nixpkgs/archive/d2e4de209881b38392933fabf303cde3454b0b4c.tar.gz";
    sha256 = "0nm6pxnji2jq7v3dx9v7cdj8rfzp64rgnxvhyphk39dhmlyyd9m0";
  }) {
    config.permittedInsecurePackages = [ "python-2.7.18.6" "openssl-1.1.1w" ];
  };
  vulnFlaskApp = pkgs.callPackage ./app/derivation.nix { };
  debugVm = { modulesPath, config, ... }: {
    imports = [ "${modulesPath}/virtualisation/qemu-vm.nix" ];
    boot.loader.systemd-boot.enable = true;
    boot.loader.efi.canTouchEfiVariables = true;
    services.xserver.enable = false;
    environment.systemPackages = with pkgs; [
      (python3.withPackages (ps: [ ps.flask ps.pyyaml ]))
      vulnFlaskApp
      nmap
      bash
      openssl
      curl
      coreutils-full
    ];
    users.users.alice = {
      isNormalUser = true;
      extraGroups = [ "wheel" ]; # ‘sudo’ for the user.
      #extraGroups = [ ]; # No ‘sudo’ for the user.
      initialPassword = "testpw";
    };
    virtualisation.qemu.networkingOptions =
      [ "-net nic -net tap,ifname=tap0,script=no" ];
    virtualisation.diskSize = 10240;
    networking.interfaces."eth1" = {
      ipv4.addresses = [{
        address = "192.168.100.2";
        prefixLength = 24;
      }];

      ipv4.routes = [{
        address = "192.168.100.1";
        prefixLength = 32;
      }];
    };
    networking.defaultGateway = {
      interface = "eth1";
      address = "192.168.100.1";
    };
    networking.firewall.enable = true;
    networking.firewall.interfaces."eth1".allowedTCPPorts = [ 8000 22 ];
    services.openssh.enable = true;
    systemd.services.vulnerableFlaskApp = {
      path =
        [ pkgs.bash pkgs.openssl pkgs.curl pkgs.coreutils-full pkgs.which ];
      wantedBy = [ "multi-user.target" ];
      after = [ "network.target" ];
      description = "Start the vulnerable flask application";
      path =
        [ pkgs.bash pkgs.openssl pkgs.curl pkgs.coreutils-full pkgs.which ];
      serviceConfig = {
        Type = "simple";
        User = "alice";
        ExecStart = "${vulnFlaskApp}/bin/vulnerable-app.py";
        WorkingDirectory = "/home/alice";
      };
    };
    services.cassandra = {
      enable = true;
      user = "root";
      group = "root";
    };
    services.zookeeper = { enable = true; };
    services.apache-kafka = {
      enable = true;
      jvmOptions = [
        "-Djava.net.preferIPv4Stack=true"
        "-Dcom.sun.management.jmxremote=true"
        "-Dcom.sun.management.jmxremote.port=8011"
        "-Dcom.sun.management.jmxremote.local.only=true"
        "-Dcom.sun.management.jmxremote.authenticate=false"
        "-Dcom.sun.management.jmxremote.ssl=false"
        "-Djava.rmi.server.hostname=192.168.100.2"
      ];

    };
    systemd.services.apache-kafka = {
      after = [ "network.target" "zookeper.service" ];
      path =
        [ pkgs.bash pkgs.openssl pkgs.curl pkgs.coreutils-full pkgs.which ];
      serviceConfig = {
        User = pkgs.lib.mkForce "root";
        Group = pkgs.lib.mkForce "root";
        Restart = "on-failure";
      };
    };
    system.stateVersion = "23.05";
  };
  nixosEvaluation = pkgs.nixos [ debugVm ];
in nixosEvaluation.config.system.build.vm
