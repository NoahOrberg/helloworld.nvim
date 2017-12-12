if exists('g:loaded_helloworld')
  finish
endif
let g:loaded_helloworld = 1

function! s:RequireHelloworld(host) abort
  return jobstart(['helloworld.nvim'], { 'rpc': v:true })
endfunction

call remote#host#Register('helloworld.nvim', 'x', function('s:RequireHelloworld'))

call remote#host#RegisterPlugin('helloworld.nvim', '0', [
\ {'type': 'function', 'name': 'Hello', 'sync': 1, 'opts': {}},
\ ])
