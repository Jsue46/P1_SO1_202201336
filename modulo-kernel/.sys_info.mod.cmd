savedcmd_sys_info.mod := printf '%s\n'   sys_info.o | awk '!x[$$0]++ { print("./"$$0) }' > sys_info.mod
