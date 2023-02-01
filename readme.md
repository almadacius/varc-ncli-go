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
  + nodejs operational: `0.3s`
- python baseline: `0.04s`
- golang baseline: `0.01s`
  + golang operational: `0.01s`

- nodejs implementation average baseline: `0.3s` (very slow)

---

## external API

```
  // variable management
  varc set "{{scope}}.{{name}}" "{{value}}"
  varc unset "{{scope}}.{{name}}"
  varc get "{{scope}}.{{name}}"
  varc scopes
  varc del "{{scope}}"
  varc keys "{{scope}}"

  // benchmark helper
  // > hft - human formatted time
  varc timercreate -> timestamp - new <timerkey>
  varc timerstep <timerkey> -> hft - time elapsed since start of timer
  // > @desc - close the timer session for this key
  varc timerend <timerkey> -> hft - total time elapsed
```

---

## structure

- deployment
  + linked repo
  + packaged app (@TODO)

- storage
  + local json

- features
  + scope separation
  + per-scope get/set
  + delete scope
  + list scopes

- output
  + output is the only value written to stdout (&1)
  + that's to simplify the interop, as the regular fd 3 (&3) strategy adds some boilerplate structures.
  + logging info can be bundled on stderr (&2)

---

Copyright 2023 Almadash
