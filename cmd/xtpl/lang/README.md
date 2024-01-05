
```bash
# in current dir; cmd/xtpl/lang
xgettext -C --from-code=UTF-8 -o messages.pot -kT -kN:1,2 -kN64:1,2 -kX:2,1c -kXN:2,3,1c -kXN64:2,3,1c ../*.go

# 初始化翻译
msginit -i messages.pot -l zh_CN -o zh_CN.po

# 更新模板
# 注意需要在 zh_CN.po 中设置 Content-Type: text/plain; charset=UTF-8\n
msgmerge -U zh_CN.po messages.pot
```
