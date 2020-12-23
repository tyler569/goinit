# vim: noet ts=8 sw=8 sts=8

root.img: init Makefile
	qemu-img create root.img 100M
	mkfs.ext2 root.img
	e2mkdir root.img:/sbin
	e2mkdir root.img:/dev
	e2mkdir root.img:/proc
	e2cp -P755 init root.img:/sbin/init
	e2cp test_file root.img:/sbin

init: $(shell find . -name '*.go')
	go build -o init

