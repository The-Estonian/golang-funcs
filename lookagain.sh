find . -name "*.sh" | cut -f 2 -d "/" | cut -f 1 -d "." | sort -rn
