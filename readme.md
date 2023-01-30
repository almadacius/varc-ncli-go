# Varc Go Ncli

`VARiable Cache GOlang Non-interactive CLI`

The `golang` implementation of `varc` project.

used to simplify bash extensions with variable get/set by scope.

---

## why

`golang` was used on this installment due to performance issues with nodejs.

baseline defined by single run of `(bash on macOs) time <command>` on `tornado` machine.

these were runs of a single console log command on different technologies.

- nodejs baseline: `0.08s`
- python baseline: `0.04s`
- golang baseline: `0.01s`

- nodejs implementation average baseline: `0.3s` (very slow)

---

Copyright 2023 Almadash
