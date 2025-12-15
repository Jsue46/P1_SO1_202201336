# Makefile simple y funcional para compilar módulos de kernel

# Lista de módulos que deseas compilar
obj-m := cont_info.o sys_info.o

# Ruta al build del kernel
KDIR := /lib/modules/$(shell uname -r)/build
PWD  := $(shell pwd)

all:
	make -C $(KDIR) M=$(PWD) modules

clean:
	make -C $(KDIR) M=$(PWD) clean
