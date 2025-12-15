// cont_info.c
// Módulo del kernel: expone /proc/continfo_so1_202201336 con JSON de procesos "usuario"
// Diseñado para que un daemon en Go lea /proc/<pid>/cgroup y detecte Docker containers.

#include <linux/init.h>
#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/sched.h>
#include <linux/sched/signal.h>
#include <linux/mm.h>
#include <linux/mm_types.h>
#include <linux/sysinfo.h>
#include <linux/cred.h>
#include <linux/rcupdate.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Estuardo");
MODULE_DESCRIPTION("Modulo para exportar procesos de usuario via /proc para deteccion de contenedores (JSON)");
MODULE_VERSION("1.1");

#define PROC_NAME "continfo_so1_202201336"

static struct proc_dir_entry *proc_file;

/* Páginas -> KB */
static inline unsigned long pages_to_kb(unsigned long pages)
{
    return (pages << (PAGE_SHIFT - 10));
}

static int cont_proc_show(struct seq_file *m, void *v)
{
    struct task_struct *task;
    unsigned int total = 0;
    struct sysinfo si;
    unsigned long total_ram_kb;

    si_meminfo(&si);
    total_ram_kb = (si.totalram << (PAGE_SHIFT - 10));

    seq_puts(m, "{\n");
    seq_puts(m, "  \"procesos\": [\n");

    rcu_read_lock();
    for_each_process(task) {

        /* Filtrar: solo procesos de usuario con mm (excluir kernel threads) */
        if ((task->flags & PF_KTHREAD) || task->mm == NULL)
            continue;

        if (total > 0)
            seq_puts(m, ",\n");

        {
            pid_t pid = task->pid;
            const char *name = task->comm;
            kuid_t ku = task->real_cred ? task->real_cred->uid : INVALID_UID;
            unsigned int uid = from_kuid(&init_user_ns, ku);
            unsigned long vsz_kb = task->mm ? pages_to_kb(task->mm->total_vm) : 0;
            unsigned long rss_kb = task->mm ? pages_to_kb(get_mm_rss(task->mm)) : 0;
            unsigned long ram_percent = 0;

            if (total_ram_kb > 0)
                ram_percent = (rss_kb * 100) / total_ram_kb;

            /* JSON por proceso: el daemon usará pid para leer /proc/<pid>/cgroup */
            seq_printf(m,
                "    {\n"
                "      \"pid\": %d,\n"
                "      \"nombre\": \"%s\",\n"
                "      \"uid\": %u,\n"
                "      \"vsz_kb\": %lu,\n"
                "      \"rss_kb\": %lu,\n"
                "      \"cpu\": 0,\n"
                "      \"ram_percent\": %lu\n"
                "    }",
                pid, name, uid, vsz_kb, rss_kb, ram_percent
            );

            total++;
        }
    }
    rcu_read_unlock();

    seq_puts(m, "\n  ],\n");
    seq_printf(m, "  \"total_procesos\": %u\n", total);
    seq_puts(m, "}\n");

    return 0;
}

static int __init cont_init(void)
{
    pr_info("cont_info: cargando modulo.\n");

    proc_file = proc_create_single(PROC_NAME, 0444, NULL, cont_proc_show);
    if (!proc_file) {
        pr_err("cont_info: error creando /proc/%s\n", PROC_NAME);
        return -ENOMEM;
    }

    pr_info("cont_info: /proc/%s creado exitosamente.\n", PROC_NAME);
    return 0;
}

static void __exit cont_exit(void)
{
    proc_remove(proc_file);
    pr_info("cont_info: modulo removido.\n");
}

module_init(cont_init);
module_exit(cont_exit);
