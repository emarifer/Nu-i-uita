#!/bin/bash

# Script to update the production version with the corresponding tag

path="frontend/src/pages/About.svelte"
version=$(git describe --abbrev=0 --tags 2>/dev/null)

if [ "${version}" = "" ]; then
  echo "There are no tags set in this repository!!"
  exit 1
fi

OPTSTRING=":ur"

while getopts ${OPTSTRING} opt; do
  case ${opt} in
    u)
      sed -i "s/DEV_VERSION/$version/g" $path
      echo "Version tag added to repository."
      ;;
    r)
      sed -i "s/$version/DEV_VERSION/g" $path
      echo "Restored development version tag."
      ;;
    ?)
      echo "Invalid option: -${OPTARG}!!"
      exit 1
      ;;
  esac
done

if [ "${OPTIND}" -eq 1 ]; then
  echo "No flags provided!!"
  exit 1
fi
