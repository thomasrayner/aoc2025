#!/usr/bin/env bash

echo -e "\033[1;36mAdvent of Code 2025 — Performance Dashboard\033[0m"
echo -e "\033[1;34m============================================\033[0m"

total=0

for d in day??; do
    [[ ! -d "$d" ]] && continue

    printf "\033[1;33m%-8s\033[0m " "$d"

    # Capture the timing line cleanly
    timing_line=$( /usr/bin/time -f "%e" go run "$d"/*.go "$d/input.txt" 2>&1 >/dev/null | tail -1 )

    # Extract the number (handles both "0.123" and "0,123" locales)
    secs=$(echo "$timing_line" | grep -o -E '[0-9]*\.[0-9]+|[0-9]+,[0-9]+' | tr ',' '.' | head -1)

    # If no number found (should never happen), default to 0
    if [[ -z "$secs" ]]; then
        secs="0"
    fi

    ms=$(awk "BEGIN {printf \"%d\", $secs * 1000}")

    printf "\033[1;32m%5d ms\033[0m\n" "$ms"
    total=$((total + ms))
done

echo -e "\033[1;34m--------------------------------------------\033[0m"
printf "\033[1;35mGRAND TOTAL:%6d ms   (%.3f seconds)\033[0m\n" "$total" "$(awk "BEGIN {printf \"%.3f\", $total/1000}")"
