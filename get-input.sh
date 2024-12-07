#!/bin/zsh
session=$(cat $HOME/sotion/aoc2020/session)
[[ -n $1 ]] || (echo "\$1 is required" && exit -1 )
bn=$(basename $1)
curl 'https://adventofcode.com/'"$1" -O \
  -H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7' \
  -H 'accept-language: en-US,en;q=0.9' \
  -H 'cache-control: no-cache' \
  -H "cookie: ${session}"  \
  -H 'pragma: no-cache' \
  -H 'priority: u=0, i' \
  -H 'referer: https://adventofcode.com/2024/day/7' \
  -H 'sec-ch-ua: "Not A(Brand";v="8", "Chromium";v="132", "Microsoft Edge";v="132"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Windows"' \
  -H 'sec-fetch-dest: document' \
  -H 'sec-fetch-mode: navigate' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-fetch-user: ?1' \
  -H 'upgrade-insecure-requests: 1' \
  -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0'
  mv $bn ${bn}.txt