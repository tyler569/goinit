# vim: noet ts=8 sw=8 sts=8

root.img: init sub
	sudo mkdir -p root
	qemu-img create root.img 100M
	mkfs.ext2 root.img
	sudo mount -o loop root.img root
	sudo mkdir -p root/sbin
	sudo mkdir -p root/dev
	sudo mkdir -p root/proc
	sudo cp init root/sbin
	sudo cp sub root/sbin
	sudo umount root

sub: sub.go
	go build sub.go

init: init.go
	go build init.go

