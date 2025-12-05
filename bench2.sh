#!/usr/bin/env bash

echo -e "\033[1;36mAdvent of Code 2025 — Performance Dashboard (perf)\033[0m"
echo -e "\033[1;34m=================================================\033[0m"

# Require perf
if ! command -v perf >/dev/null 2>&1; then
    echo -e "\033[1;31mERROR:\033[0m perf not found. Install it (e.g. 'sudo pacman -S perf')."
    exit 1
fi

mkdir -p .bin

total_ms=0

for d in day??; do
    [[ ! -d "$d" ]] && continue

    printf "\033[1;33m%-8s\033[0m " "$d"

    # Build binary once
    (
        cd "$d" || exit 1
        go build -o "../.bin/$d"
    ) >/dev/null 2>&1

    # Run perf, discard program stdout, capture perf stderr
    perf_output=$(
      { perf stat -x, -e task-clock "./.bin/$d" "$d/input.txt" >/dev/null; } 2>&1
    )

    # Extract the first field on the task-clock line (nanoseconds)
    raw_ns=$(echo "$perf_output" | awk -F, '/task-clock/ {print $1; exit}')

    # If perf failed or gave something weird, fall back to 0
    if [[ ! "$raw_ns" =~ ^[0-9]+$ ]]; then
        raw_ns=0
    fi

    # Convert ns → ms (float)
    ms_float=$(awk -v ns="$raw_ns" 'BEGIN {printf "%.3f", ns / 1000000.0}')
    # Round to int ms for total
    ms_int=$(awk -v v="$ms_float" 'BEGIN {printf "%d", v + 0.5}')

    printf "\033[1;32m%8.3f ms\033[0m\n" "$ms_float"
    total_ms=$((total_ms + ms_int))
done

echo -e "\033[1;34m-------------------------------------------------\033[0m"
printf "\033[1;35mGRAND TOTAL (approx):%6d ms   (%.3f seconds)\033[0m\n" \
  "$total_ms" "$(awk -v ms="$total_ms" 'BEGIN {printf "%.3f", ms/1000.0}')"
