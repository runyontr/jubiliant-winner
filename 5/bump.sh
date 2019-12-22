#!/bin/bash


# Bump the version of the chart and insert the value back into the chart

function bump() {
    # get version from chart file in folder
    version=`cat $1/Chart.yaml | grep version | awk '{print $2}'`
    parts=( ${version//./ } )

    # case docs: http://tldp.org/LDP/Bash-Beginners-Guide/html/sect_07_03.html
    case "$2" in
        m)
            ((parts[1]++))
            parts[2]=0
            ;;
        M)
            ((parts[0]++))
            parts[1]=0
            parts[2]=0
            ;;
        p)
            ((parts[2]++))
            ;;
        *)
            echo "Must provide which value to increment:  Major (M), Minor (m), or Patch (p)"
            exit 1
            ;;
    esac
        new_version="${parts[0]}.${parts[1]}.${parts[2]}"
        echo "Upgrading Chart version from ${version} to ${new_version}"
        # ((a[2]++))

    # Find and replace in the chart file
    # This is mac specific as outlined here: https://clubmate.fi/replace-strings-in-files-with-the-sed-bash-command/
    # the -i argument needs an empty value to prevent using a backup file
    sed -i "" "s/version: ${version}/version: ${new_version}/g" "$1/Chart.yaml"
}



bump $1 $2