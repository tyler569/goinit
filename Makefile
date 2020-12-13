# vim: noet ts=8 sw=8 sts=8

root.img: init
	sudo mkdir -p root
	qemu-img create root.img 100M
	mkfs.ext2 root.img
	sudo mount -o loop root.img root
	sudo mkdir -p root/sbin
	sudo mkdir -p root/dev
	sudo mkdir -p root/proc
	sudo cp init root/sbin
	sudo umount root

init: $(shell find . -name '*.go')
	go build -o init

