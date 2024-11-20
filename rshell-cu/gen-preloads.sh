set -x
export NIX_SHELL_PRESERVE_PROMPT=1
#nasm -f bin -o preloadlib_amd64 preloadlib_amd64.asm
#xxd -ps preloadlib_amd64 | tr -d '\n'
gcc preloadlib_amd64_alpinesafe.c -s -fPIC -shared -o preloadlib_amd64.so
xxd -ps preloadlib_amd64.so | tr -d '\n' | head -c 250
xxd -ps preloadlib_amd64.so | tr -d '\n' | xclip -selection clipboard
