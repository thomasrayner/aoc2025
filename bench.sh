#!/usr/bin/env bash

echo -e "\033[1;36mAdvent of Code 2025 — Performance Dashboard\033[0m"
echo -e "\033[1;34m============================================\033[0m"

total=0

for d in day??; do
    [[ ! -d "$d" ]] && continue

    printf "\033[1;33m%-8s\033[0m " "$d"

    # Run from inside the day directory so its go.mod is respected
    timing_secs=$(
      (
        cd "$d" || exit 1
        # time -p prints:
        # real 0.028
        # user ...
        # sys  ...
        time -p go run . input.txt >/dev/null
      ) 2>&1 | awk '/^real / { print $2 }'
    )

    # Normalize decimal separator, default to 0 if empty
    secs=$(echo "${timing_secs:-0}" | tr ',' '.')

    ms=$(awk "BEGIN {printf \"%d\", $secs * 1000}")

    printf "\033[1;32m%5d ms\033[0m\n" "$ms"
    total=$((total + ms))
done

echo -e "\033[1;34m--------------------------------------------\033[0m"
printf "\033[1;35mGRAND TOTAL:%6d ms   (%.3f seconds)\033[0m\n" \
  "$total" "$(awk "BEGIN {printf \"%.3f\", $total/1000}")"
