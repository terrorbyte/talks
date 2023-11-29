cat "$1" | sed -e 's%\.png%\.ff%g' -e 's%\.jpg%\.ff%g' -e 's%\.jpeg%\.ff%g' > "$1"ff
