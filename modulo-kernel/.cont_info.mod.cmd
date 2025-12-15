savedcmd_cont_info.mod := printf '%s\n'   cont_info.o | awk '!x[$$0]++ { print("./"$$0) }' > cont_info.mod
