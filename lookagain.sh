find . -type f -name "*.sh"| sort -nr | cut -f 2 -d "/" | cut -f 1 -d "."
