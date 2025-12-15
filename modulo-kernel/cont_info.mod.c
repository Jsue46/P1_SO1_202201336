#include <linux/module.h>
#include <linux/export-internal.h>
#include <linux/compiler.h>

MODULE_INFO(name, KBUILD_MODNAME);

__visible struct module __this_module
__section(".gnu.linkonce.this_module") = {
	.name = KBUILD_MODNAME,
	.init = init_module,
#ifdef CONFIG_MODULE_UNLOAD
	.exit = cleanup_module,
#endif
	.arch = MODULE_ARCH_INIT,
};



static const struct modversion_info ____versions[]
__used __section("__versions") = {
	{ 0xd272d446, "__x86_return_thunk" },
	{ 0xc7ffe1aa, "si_meminfo" },
	{ 0x08fc6844, "seq_write" },
	{ 0xd272d446, "__rcu_read_lock" },
	{ 0x1790826a, "init_task" },
	{ 0x93247f89, "init_user_ns" },
	{ 0x7883379a, "from_kuid" },
	{ 0xf2c4f3f1, "seq_printf" },
	{ 0xd272d446, "__rcu_read_unlock" },
	{ 0xd272d446, "__stack_chk_fail" },
	{ 0x9487f592, "proc_remove" },
	{ 0xd272d446, "__fentry__" },
	{ 0xe8213e80, "_printk" },
	{ 0xfd53b4b4, "proc_create_single_data" },
	{ 0xba157484, "module_layout" },
};

static const u32 ____version_ext_crcs[]
__used __section("__version_ext_crcs") = {
	0xd272d446,
	0xc7ffe1aa,
	0x08fc6844,
	0xd272d446,
	0x1790826a,
	0x93247f89,
	0x7883379a,
	0xf2c4f3f1,
	0xd272d446,
	0xd272d446,
	0x9487f592,
	0xd272d446,
	0xe8213e80,
	0xfd53b4b4,
	0xba157484,
};
static const char ____version_ext_names[]
__used __section("__version_ext_names") =
	"__x86_return_thunk\0"
	"si_meminfo\0"
	"seq_write\0"
	"__rcu_read_lock\0"
	"init_task\0"
	"init_user_ns\0"
	"from_kuid\0"
	"seq_printf\0"
	"__rcu_read_unlock\0"
	"__stack_chk_fail\0"
	"proc_remove\0"
	"__fentry__\0"
	"_printk\0"
	"proc_create_single_data\0"
	"module_layout\0"
;

MODULE_INFO(depends, "");


MODULE_INFO(srcversion, "9A1562220808FBCB241CC8C");
