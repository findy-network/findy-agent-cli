let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/fsm
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +1 b-fsm1.yaml
badd +133 f-fsm1.yaml
badd +5 f-fsm.yaml
badd +1 ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/eb3.yaml
badd +1 ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/select2.lua
badd +1 ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/helper.lua
badd +33 role.lua
badd +1 scratch.yaml
argglobal
%argdel
$argadd b-fsm1.yaml
$argadd f-fsm1.yaml
set stal=2
tabnew +setlocal\ bufhidden=wipe
tabnew +setlocal\ bufhidden=wipe
tabrewind
edit b-fsm1.yaml
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
wincmd _ | wincmd |
vsplit
wincmd _ | wincmd |
vsplit
3wincmd h
wincmd w
wincmd w
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 64 + 159) / 318)
exe 'vert 2resize ' . ((&columns * 94 + 159) / 318)
exe 'vert 3resize ' . ((&columns * 79 + 159) / 318)
exe 'vert 4resize ' . ((&columns * 78 + 159) / 318)
argglobal
setlocal fdm=manual
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 24 - ((23 * winheight(0) + 33) / 66)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 24
normal! 021|
wincmd w
argglobal
2argu
balt f-fsm.yaml
setlocal fdm=manual
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 133 - ((57 * winheight(0) + 33) / 66)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 133
normal! 021|
wincmd w
argglobal
2argu
if bufexists(fnamemodify("f-fsm.yaml", ":p")) | buffer f-fsm.yaml | else | edit f-fsm.yaml | endif
if &buftype ==# 'terminal'
  silent file f-fsm.yaml
endif
balt f-fsm1.yaml
setlocal fdm=manual
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 4 - ((3 * winheight(0) + 33) / 66)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 4
normal! 03|
wincmd w
argglobal
2argu
if bufexists(fnamemodify("role.lua", ":p")) | buffer role.lua | else | edit role.lua | endif
if &buftype ==# 'terminal'
  silent file role.lua
endif
balt f-fsm1.yaml
setlocal fdm=manual
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 33 - ((32 * winheight(0) + 33) / 66)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 33
normal! 037|
wincmd w
2wincmd w
exe 'vert 1resize ' . ((&columns * 64 + 159) / 318)
exe 'vert 2resize ' . ((&columns * 94 + 159) / 318)
exe 'vert 3resize ' . ((&columns * 79 + 159) / 318)
exe 'vert 4resize ' . ((&columns * 78 + 159) / 318)
tabnext
edit ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/eb3.yaml
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
wincmd _ | wincmd |
vsplit
2wincmd h
wincmd w
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 105 + 159) / 318)
exe 'vert 2resize ' . ((&columns * 105 + 159) / 318)
exe 'vert 3resize ' . ((&columns * 106 + 159) / 318)
argglobal
1argu
if bufexists(fnamemodify("~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/eb3.yaml", ":p")) | buffer ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/eb3.yaml | else | edit ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/eb3.yaml | endif
if &buftype ==# 'terminal'
  silent file ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/eb3.yaml
endif
balt ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/select2.lua
setlocal fdm=manual
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 50 - ((49 * winheight(0) + 33) / 66)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 50
normal! 07|
wincmd w
argglobal
1argu
if bufexists(fnamemodify("~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/select2.lua", ":p")) | buffer ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/select2.lua | else | edit ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/select2.lua | endif
if &buftype ==# 'terminal'
  silent file ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/select2.lua
endif
balt ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/helper.lua
setlocal fdm=manual
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 33) / 66)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 0
wincmd w
argglobal
1argu
if bufexists(fnamemodify("~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/helper.lua", ":p")) | buffer ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/helper.lua | else | edit ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/helper.lua | endif
if &buftype ==# 'terminal'
  silent file ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/helper.lua
endif
balt ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/select2.lua
setlocal fdm=manual
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 33) / 66)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 0
wincmd w
exe 'vert 1resize ' . ((&columns * 105 + 159) / 318)
exe 'vert 2resize ' . ((&columns * 105 + 159) / 318)
exe 'vert 3resize ' . ((&columns * 106 + 159) / 318)
tabnext
edit scratch.yaml
argglobal
if bufexists(fnamemodify("scratch.yaml", ":p")) | buffer scratch.yaml | else | edit scratch.yaml | endif
if &buftype ==# 'terminal'
  silent file scratch.yaml
endif
balt ~/go/src/github.com/findy-network/findy-agent-cli/scripts/fullstack/eb3.yaml
setlocal fdm=manual
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 17 - ((16 * winheight(0) + 33) / 66)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 17
normal! 015|
tabnext 1
set stal=1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
let g:this_session = v:this_session
let g:this_obsession = v:this_session
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
