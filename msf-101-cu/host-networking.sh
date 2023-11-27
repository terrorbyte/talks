if [ -z "$1" ]; 
	echo "Interface required as only option. Also run it as root"
	exit 1
fi
tunctl -u $(id -u -n) -t tap0
ifconfig tap0 192.168.100.1 up
# qemu-kvm -hda nixos-disc.img -m 1024 -net nic -net tap,ifname=tap0,script=no
echo 1 > /proc/sys/net/ipv4/ip_forward
iptables -t nat -A POSTROUTING -o "$1" -j MASQUERADE
iptables -I FORWARD 1 -i tap0 -j ACCEPT
iptables -I FORWARD 1 -o tap0 -m state --state RELATED,ESTABLISHED -j ACCEPT
