#!/bin/sh

qemu-system-x86_64 \
    -kernel vmlinuz \
    -hda root.img \
    -append "root=/dev/sda console=tty0 console=ttyS0,115200" \
    -serial stdio \
    -display none
